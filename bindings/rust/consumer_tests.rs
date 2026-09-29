use tree_sitter::{InputEdit, Node, Parser, Point};

fn parser() -> Parser {
    let mut parser = Parser::new();
    parser.set_language(&super::LANGUAGE.into()).unwrap();
    parser
}

fn collect<'tree>(node: Node<'tree>, kind: &str, nodes: &mut Vec<Node<'tree>>) {
    if node.kind() == kind {
        nodes.push(node);
    }
    let mut cursor = node.walk();
    for child in node.named_children(&mut cursor) {
        collect(child, kind, nodes);
    }
}

#[test]
fn navigation_and_symbols_have_independent_name_value_and_reference_ranges() {
    // Mirrors the fields consumed by spectrals::analysis and document symbols.
    let source = "TITLE μ😀\r\nDEFINE NAME=Width VALUE=1\r\n  CHARACTER=wide\r\nSAVE FILE={$Width}.dat V={@Width+1}\r\n";
    let tree = parser().parse(source, None).unwrap();
    let root = tree.root_node();
    assert!(!root.has_error(), "{}", root.to_sexp());
    let mut statements = Vec::new();
    collect(root, "statement", &mut statements);
    assert_eq!(statements.len(), 2);
    let define = statements[0];
    let name = define.child_by_field_name("name").unwrap();
    assert_eq!(&source[name.byte_range()], "DEFINE");
    assert_eq!(name.start_position(), Point::new(1, 0));
    let declaration = define.named_child(1).unwrap();
    assert_eq!(declaration.kind(), "parameter");
    assert_eq!(&source[declaration.child_by_field_name("name").unwrap().byte_range()], "NAME");
    let value = declaration.child_by_field_name("value").unwrap();
    assert_eq!(&source[value.byte_range()], "Width");
    assert_eq!(value.start_position(), Point::new(1, 12));
    let mut continuations = Vec::new();
    collect(root, "continuation_line", &mut continuations);
    assert_eq!(continuations.len(), 1);
    assert_eq!(continuations[0].start_position().row, 2);
    let mut references = Vec::new();
    collect(root, "variable_reference", &mut references);
    let values: Vec<_> = references.iter().map(|node| &source[node.byte_range()]).collect();
    assert_eq!(values, ["$Width", "@Width"]);
    assert!(references.iter().all(|node| node.start_position().row == 3));
    // Tree-sitter positions are byte-based, not the UTF-16 columns the LSP needs.
    let title = root.named_child(0).unwrap().child_by_field_name("title").unwrap();
    assert_eq!(&source[title.byte_range()], "μ😀");
    assert_eq!(title.end_position().column, 12);
    assert_eq!(source[..title.end_byte()].encode_utf16().count(), 9);
}

#[test]
fn malformed_edit_preserves_following_statement_and_fresh_tree_ranges() {
    let before = "DEFINE NAME=x VALUE=1\nSAVE FILE=after.dat\n";
    let after = "DEFINE NAME=x VALUE=\nSAVE FILE=after.dat\n";
    let mut parser = parser();
    let mut old = parser.parse(before, None).unwrap();
    let offset = before.find("1\n").unwrap();
    old.edit(&InputEdit {
        start_byte: offset,
        old_end_byte: offset + 1,
        new_end_byte: offset,
        start_position: Point::new(0, offset),
        old_end_position: Point::new(0, offset + 1),
        new_end_position: Point::new(0, offset),
    });
    let incremental = parser.parse(after, Some(&old)).unwrap();
    let fresh = parser.parse(after, None).unwrap();
    assert_eq!(incremental.root_node().to_sexp(), fresh.root_node().to_sexp());
    let mut statements = Vec::new();
    collect(incremental.root_node(), "statement", &mut statements);
    let save = statements.iter().find(|node| node.child_by_field_name("name").is_some_and(|name| &after[name.byte_range()] == "SAVE")).expect("SAVE must survive an incomplete previous parameter");
    assert_eq!(save.start_position(), Point::new(1, 0));
    assert_eq!(&after[save.byte_range()], "SAVE FILE=after.dat");
    fn equal_ranges(a: Node<'_>, b: Node<'_>) {
        assert_eq!(a.kind(), b.kind());
        assert_eq!(a.range(), b.range());
        assert_eq!(a.is_missing(), b.is_missing());
        assert_eq!(a.child_count(), b.child_count());
        for i in 0..a.child_count() {
            equal_ranges(a.child(i).unwrap(), b.child(i).unwrap());
        }
    }
    equal_ranges(incremental.root_node(), fresh.root_node());
}

#[test]
fn keyword_prefix_remains_available_to_consumer_unknown_name_diagnostics() {
    let source = "ENDPOINT X=1\n";
    let tree = parser().parse(source, None).unwrap();
    let statement = tree.root_node().named_child(0).unwrap();
    assert_eq!(statement.kind(), "statement");
    let name = statement.child_by_field_name("name").unwrap();
    assert_eq!(&source[name.byte_range()], "ENDPOINT");
    assert!(!tree.root_node().has_error());
}
