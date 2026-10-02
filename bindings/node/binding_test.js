import assert from "node:assert";
import { test } from "node:test";
import Parser from "tree-sitter";
import language from "./index.js";

test("parses declaration fields and retains statements after END", () => {
  const parser = new Parser();
  parser.setLanguage(language);
  const root = parser.parse("DEFINE NAME=Width VALUE=1\nEND\nSAVE FILE=after.dat\n").rootNode;
  assert.strictEqual(root.hasError, false);
  const [define, end, save] = root.namedChildren;
  assert.strictEqual(define.type, "statement");
  assert.strictEqual(define.childForFieldName("name").text, "DEFINE");
  assert.deepStrictEqual(
    define.namedChildren.filter(node => node.type === "parameter").map(node => [
      node.childForFieldName("name").text,
      node.childForFieldName("value").type,
      node.childForFieldName("value").text,
    ]),
    [["NAME", "character_value", "Width"], ["VALUE", "numeric_value", "1"]],
  );
  assert.strictEqual(end.type, "end_statement");
  assert.strictEqual(save.childForFieldName("name").text, "SAVE");
  assert.strictEqual(save.startPosition.row, 2);
});

test("shipped highlight query captures numeric and character references", () => {
  const parser = new Parser();
  parser.setLanguage(language);
  const root = parser.parse("SAVE FILE=$Width V={@Width+1}\n").rootNode;
  assert.strictEqual(root.hasError, false);
  const query = new Parser.Query(language, language.HIGHLIGHTS_QUERY);
  assert.deepStrictEqual(
    query.captures(root).filter(capture => capture.name === "variable").map(capture => capture.node.text),
    ["$Width", "@Width"],
  );
});
