package tree_sitter_spectra

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

// AdapterContractVersion identifies the semantics documented in adapter-contract.md.
const AdapterContractVersion = 1

// ByteRange is a half-open range of UTF-8 source byte offsets.
type ByteRange struct {
	Start uint
	End   uint
}

// Analysis owns its strings and contains no pointers into the syntax tree.
type Analysis struct {
	Statements   []Statement
	Declarations []Declaration
	References   []Reference
	Conditionals []Conditional
	Includes     []Include
	Diagnostics  []Diagnostic
}

// Statement groups a statement node with its continuation nodes. Parts excludes
// intervening comments and whitespace; Range is their enclosing source span.
type Statement struct {
	Kind          string
	Name          string
	CanonicalName string
	NameRange     ByteRange
	Range         ByteRange
	Parts         []ByteRange
	Parameters    []Parameter
	Active        bool
	Malformed     bool
}

// Parameter retains an assignment, including assignments nested in mask groups.
// Group numbers are zero-based DEFINE NAME / EXTRACT PARAMETER groups.
type Parameter struct {
	Name          string
	CanonicalName string
	NameRange     ByteRange
	Range         ByteRange
	Value         string
	ValueKind     string
	ValueRange    ByteRange
	Group         int
}

type DeclarationKind string

const (
	DefineDeclaration  DeclarationKind = "define"
	ExtractDeclaration DeclarationKind = "extract"
)

// Declaration names are case sensitive; Spelling includes any source quotes.
type Declaration struct {
	Name      string
	Spelling  string
	Range     ByteRange
	Kind      DeclarationKind
	Statement int
	Parameter int
	Group     int
	Active    bool
}

type ReferenceKind string

const (
	NumericReference   ReferenceKind = "numeric"
	CharacterReference ReferenceKind = "character"
	ExtractedReference ReferenceKind = "extracted"
)

// Reference is a syntactic use, not a claim that a declaration was resolved.
// Parameter is -1 for references outside assignments (for example IF).
type Reference struct {
	Name      string
	Spelling  string
	Range     ByteRange
	NameRange ByteRange
	Kind      ReferenceKind
	Linked    bool
	Statement int
	Parameter int
	Active    bool
}

// Conditional pairs flat statement indices. Absent ELSE/ENDIF indices are -1.
type Conditional struct {
	IfStatement    int
	ElseStatement  int
	EndIfStatement int
}

type IncludeStatus string

const (
	IncludeLiteral   IncludeStatus = "literal"
	IncludeSymbolic  IncludeStatus = "symbolic"
	IncludeInvalid   IncludeStatus = "invalid"
	IncludeInactive  IncludeStatus = "inactive"
	IncludeResolved  IncludeStatus = "resolved"
	IncludeFileError IncludeStatus = "file-error"
)

// Include describes INSERT FILE or DEFINE FILE, not simulator data/output files.
type Include struct {
	Spelling  string
	Range     ByteRange
	Path      string
	Status    IncludeStatus
	Statement int
	Parameter int
	Active    bool
}

// IncludeResolution holds a filesystem result without loading or executing input.
// Path is absolute for filesystem attempts and empty for unresolved/inactive input.
type IncludeResolution struct {
	Include Include
	Path    string
	Status  IncludeStatus
	Err     error
}

// Diagnostic codes are stable; messages are explanatory rather than API keys.
// Statement is -1 when a syntax range has no logical statement owner.
type Diagnostic struct {
	Code      string
	Message   string
	Severity  string
	Range     ByteRange
	Statement int
}

type analysisBuilder struct {
	result       Analysis
	source       string
	nameField    uint16
	valueField   uint16
	owner        int
	active       bool
	groups       []int
	groupStarted []bool
}

// Analyze reads the fields of a SPECTRA source_file and the exact bytes used to
// parse it. It does not evaluate expressions, bind references, or read files.
func Analyze(root *tree_sitter.Node, source []byte) Analysis {
	b := analysisBuilder{source: string(source), owner: -1, active: true}
	if root == nil || root.Kind() != "source_file" || root.EndByte() > uint(len(source)) {
		b.diagnostic("invalid-root", "expected a source_file tree for the supplied source", "error", ByteRange{}, -1)
		return b.result
	}
	language := root.Language()
	b.nameField = language.FieldIdForName("name")
	b.valueField = language.FieldIdForName("value")
	cursor := root.Walk()
	defer cursor.Close()
	b.syntax(cursor)
	cursor.Reset(*root)
	if cursor.GotoFirstChild() {
		for {
			if node := cursor.Node(); node.IsNamed() {
				b.line(node)
			}
			if !cursor.GotoNextSibling() {
				break
			}
		}
	}
	b.matchConditionals()
	// Assign syntax diagnostics to enclosing physical parts, not comment gaps.
	for i := range b.result.Diagnostics {
		d := &b.result.Diagnostics[i]
		if d.Statement >= 0 {
			continue
		}
		j := sort.Search(len(b.result.Statements), func(j int) bool { return b.result.Statements[j].Range.End >= d.Range.End })
		if j < len(b.result.Statements) {
			for _, part := range b.result.Statements[j].Parts {
				if part.Start <= d.Range.Start && d.Range.End <= part.End {
					d.Statement = j
					break
				}
			}
		}
	}
	sort.SliceStable(b.result.Diagnostics, func(i, j int) bool {
		return b.result.Diagnostics[i].Range.Start < b.result.Diagnostics[j].Range.Start
	})
	return b.result
}

func nodeRange(n *tree_sitter.Node) ByteRange {
	if n == nil {
		return ByteRange{}
	}
	return ByteRange{Start: n.StartByte(), End: n.EndByte()}
}

func (b *analysisBuilder) text(n *tree_sitter.Node) string {
	if n == nil {
		return ""
	}
	return b.source[n.StartByte():n.EndByte()]
}

func (b *analysisBuilder) diagnostic(code, message, severity string, r ByteRange, statement int) {
	b.result.Diagnostics = append(b.result.Diagnostics, Diagnostic{Code: code, Message: message, Severity: severity, Range: r, Statement: statement})
}

func (b *analysisBuilder) syntax(cursor *tree_sitter.TreeCursor) {
	n := cursor.Node()
	switch {
	case n.IsMissing():
		b.diagnostic("syntax-missing", "missing "+n.Kind(), "error", nodeRange(n), -1)
	case n.IsError():
		b.diagnostic("syntax-error", "unrecognized syntax", "error", nodeRange(n), -1)
	case n.Kind() == "unparsed_line":
		b.diagnostic("unparsed-line", "line could not be parsed", "error", nodeRange(n), -1)
	}
	if cursor.GotoFirstChild() {
		for {
			b.syntax(cursor)
			if !cursor.GotoNextSibling() {
				break
			}
		}
		cursor.GotoParent()
	}
}

func (b *analysisBuilder) line(n *tree_sitter.Node) {
	kind := n.Kind()
	if kind == "comment" || strings.TrimSpace(b.text(n)) == "" {
		return
	}
	if kind == "unparsed_line" || n.IsError() {
		b.owner = -1
		// Recovery may wrap intact subsequent statements inside ERROR. Preserve
		// those tree-recognized statements without interpreting opaque text.
		if n.IsError() {
			for i, count := uint(0), n.NamedChildCount(); i < count; i++ {
				b.line(n.NamedChild(i))
			}
		}
		b.owner = -1
		return
	}
	continuation := kind == "continuation_line"
	nameNode := n.ChildByFieldId(b.nameField)
	if !continuation && (nameNode == nil || nameNode.Kind() != "statement_name") {
		b.owner = -1
		return
	}
	index := b.owner
	if !continuation || index < 0 {
		name := b.text(nameNode)
		canonical := canonicalStatement(name)
		index = len(b.result.Statements)
		b.result.Statements = append(b.result.Statements, Statement{Kind: kind, Name: name, CanonicalName: canonical, NameRange: nodeRange(nameNode), Range: nodeRange(n), Active: b.active})
		b.groups = append(b.groups, 0)
		b.groupStarted = append(b.groupStarted, false)
		if continuation {
			b.diagnostic("orphan-continuation", "continuation has no preceding parameter statement", "error", nodeRange(n), index)
		} else if _, known := parameterCatalogue[canonical]; !known {
			b.diagnostic("unknown-statement", "statement is not in the checked-in reference catalogue", "warning", nodeRange(nameNode), index)
		}
	}
	s := &b.result.Statements[index]
	s.Range.End = n.EndByte()
	s.Parts = append(s.Parts, nodeRange(n))
	s.Malformed = s.Malformed || n.HasError()
	cursor := n.Walk()
	b.collect(cursor, index, -1, false, false)
	cursor.Close()
	b.owner = -1
	if !s.Malformed && (kind == "statement" || (continuation && s.Name != "")) {
		b.owner = index
	}
	if kind == "end_statement" && !s.Malformed {
		b.active = false
	}
}

func (b *analysisBuilder) collect(cursor *tree_sitter.TreeCursor, statement, parameter int, linked, nested bool) {
	n := cursor.Node()
	kind := n.Kind()
	s := &b.result.Statements[statement]
	if kind == "unparsed_line" || n.IsError() || n.IsMissing() {
		s.Malformed = true
		return
	}
	if kind == "mask_group" {
		nested = true
	}
	if kind == "linked_series" {
		linked = true
	}
	if kind == "parameter" {
		nameNode := n.ChildByFieldId(b.nameField)
		valueNode := n.ChildByFieldId(b.valueField)
		name := b.text(nameNode)
		canonical := canonicalParameter(s.CanonicalName, name)
		if !nested && ((s.CanonicalName == "DEFINE" && canonical == "NAME") || (s.CanonicalName == "EXTRACT" && canonical == "PARAMETER")) {
			if b.groupStarted[statement] {
				b.groups[statement]++
			}
			b.groupStarted[statement] = true
		}
		p := Parameter{Name: name, CanonicalName: canonical, NameRange: nodeRange(nameNode), Range: nodeRange(n), Value: b.text(valueNode), ValueRange: nodeRange(valueNode), Group: b.groups[statement]}
		if valueNode != nil {
			p.ValueKind = valueNode.Kind()
		}
		parameter = len(s.Parameters)
		s.Parameters = append(s.Parameters, p)
		if catalogue, known := parameterCatalogue[s.CanonicalName]; known {
			if _, known := catalogue[canonical]; !known {
				b.diagnostic("unknown-parameter", "parameter is not in this statement's checked-in reference catalogue", "warning", p.NameRange, statement)
			}
		}
		if !nested && canonical == "NAME" && (s.CanonicalName == "DEFINE" || s.CanonicalName == "EXTRACT") && valueNode != nil && !n.HasError() {
			if name, literal := literalText(p.ValueKind, p.Value); literal && name != "" {
				declarationKind := DefineDeclaration
				if s.CanonicalName == "EXTRACT" {
					declarationKind = ExtractDeclaration
				}
				b.result.Declarations = append(b.result.Declarations, Declaration{Name: name, Spelling: p.Value, Range: p.ValueRange, Kind: declarationKind, Statement: statement, Parameter: parameter, Group: p.Group, Active: s.Active})
			}
		}
		if !nested && canonical == "FILE" && (s.CanonicalName == "INSERT" || s.CanonicalName == "DEFINE") {
			include := Include{Spelling: p.Value, Range: p.ValueRange, Status: IncludeSymbolic, Statement: statement, Parameter: parameter, Active: s.Active}
			if n.HasError() || valueNode == nil || p.Value == "" {
				include.Status = IncludeInvalid
			} else if path, literal := literalText(p.ValueKind, p.Value); literal {
				include.Path = path
				include.Status = IncludeLiteral
				if path == "" {
					include.Status = IncludeInvalid
				}
			}
			b.result.Includes = append(b.result.Includes, include)
		}
	}
	if kind == "variable_reference" {
		spelling := b.text(n)
		prefix := 0
		for prefix < len(spelling) && (spelling[prefix] == '@' || spelling[prefix] == '$' || spelling[prefix] == '#') {
			prefix++
		}
		if prefix > 0 && prefix < len(spelling) {
			referenceKind := NumericReference
			if spelling[0] == '$' {
				referenceKind = CharacterReference
			} else if spelling[0] == '#' {
				referenceKind = ExtractedReference
			}
			b.result.References = append(b.result.References, Reference{Name: spelling[prefix:], Spelling: spelling, Range: nodeRange(n), NameRange: ByteRange{Start: n.StartByte() + uint(prefix), End: n.EndByte()}, Kind: referenceKind, Linked: linked && spelling[0] == '#', Statement: statement, Parameter: parameter, Active: s.Active})
		}
	}
	if cursor.GotoFirstChild() {
		for {
			if cursor.Node().IsNamed() {
				b.collect(cursor, statement, parameter, linked, nested)
			}
			if !cursor.GotoNextSibling() {
				break
			}
		}
		cursor.GotoParent()
	}
}

func (b *analysisBuilder) matchConditionals() {
	var stack []int
	closeOpen := func() {
		for _, index := range stack {
			statement := b.result.Conditionals[index].IfStatement
			b.diagnostic("unclosed-if", "IF has no matching ENDIF before END or end of file", "error", b.result.Statements[statement].NameRange, statement)
		}
		stack = stack[:0]
	}
	for i, s := range b.result.Statements {
		switch s.Kind {
		case "if_statement":
			stack = append(stack, len(b.result.Conditionals))
			b.result.Conditionals = append(b.result.Conditionals, Conditional{IfStatement: i, ElseStatement: -1, EndIfStatement: -1})
		case "else_statement":
			if len(stack) == 0 {
				b.diagnostic("unmatched-else", "ELSE has no matching IF", "error", s.NameRange, i)
			} else {
				c := &b.result.Conditionals[stack[len(stack)-1]]
				if c.ElseStatement >= 0 {
					b.diagnostic("duplicate-else", "IF already has an ELSE", "error", s.NameRange, i)
				} else {
					c.ElseStatement = i
				}
			}
		case "endif_statement":
			if len(stack) == 0 {
				b.diagnostic("unmatched-endif", "ENDIF has no matching IF", "error", s.NameRange, i)
			} else {
				b.result.Conditionals[stack[len(stack)-1]].EndIfStatement = i
				stack = stack[:len(stack)-1]
			}
		case "end_statement":
			if !s.Malformed {
				closeOpen()
			}
		}
	}
	closeOpen()
}

func literalText(kind, value string) (string, bool) {
	if kind == "quoted_string" && len(value) >= 2 {
		value = value[1 : len(value)-1]
		// Escape interpretation is not specified by the references. Do not
		// guess filesystem paths or declaration identities from escaped text.
		if strings.Contains(value, "\\") {
			return "", false
		}
	} else if kind != "character_value" && kind != "numeric_value" {
		return "", false
	}
	if strings.ContainsAny(value, "@$#{}") {
		return "", false
	}
	return value, true
}

// ResolveIncludes resolves literal active paths relative to the containing file.
// os.Stat errors are preserved in Err (including errors.Is support). Directories
// are rejected. Symlinks follow host filesystem semantics. No file is executed.
func ResolveIncludes(includes []Include, containingPath string) []IncludeResolution {
	results := make([]IncludeResolution, len(includes))
	for i, include := range includes {
		r := IncludeResolution{Include: include, Status: include.Status}
		if !include.Active {
			r.Status = IncludeInactive
		} else if include.Status == IncludeLiteral {
			path := include.Path
			if !filepath.IsAbs(path) {
				path = filepath.Join(filepath.Dir(containingPath), path)
			}
			r.Path, r.Err = filepath.Abs(path)
			if r.Err == nil {
				var info os.FileInfo
				info, r.Err = os.Stat(r.Path)
				if r.Err == nil && !info.Mode().IsRegular() {
					r.Err = &os.PathError{Op: "include", Path: r.Path, Err: errors.New("include target is not a regular file")}
				}
			}
			r.Status = IncludeResolved
			if r.Err != nil {
				r.Status = IncludeFileError
			}
		}
		results[i] = r
	}
	return results
}

func canonicalStatement(name string) string {
	name = strings.ToUpper(name)
	switch name {
	case "CONST":
		return "CONSTANT"
	case "CONV":
		return "CONVERGENCE"
	case "DEPO":
		return "DEPOSIT"
	case "STRUCT":
		return "STRUCTURE"
	case "SUBS":
		return "SUBSTRATE"
	}
	return name
}

func canonicalParameter(statement, name string) string {
	name = strings.ToUpper(name)
	if alias := parameterAliases[statement][name]; alias != "" {
		return alias
	}
	if alias := coordinateAliases[name]; alias != "" {
		if _, known := parameterCatalogue[statement][alias]; known {
			return alias
		}
	}
	return name
}

// Only attested abbreviations are listed: no arbitrary prefix expansion.
// See adapter-contract.md for source links and ambiguous source spellings.
var coordinateAliases = map[string]string{"XMI": "XMIN", "YMI": "YMIN", "ZMI": "ZMIN", "XMA": "XMAX", "YMA": "YMAX", "ZMA": "ZMAX"}

var parameterAliases = map[string]map[string]string{
	"CONSTANT":    {"TEMP": "TEMPERATURE"},
	"CONVERGENCE": {"VE": "VERROR", "FE": "FERROR", "CA": "CARRIER", "NE": "NEWTON", "FC": "FCYCLE", "RANGEC": "RANGECHECK"},
	"DEFINE":      {"CHAR": "CHARACTER"},
	"DOPE":        {"ELEM": "ELEMENT", "PROF": "PROFILE", "CM": "CMAX", "CB": "CBACK", "XO": "XORIGIN", "YO": "YORIGIN", "MUL": "MULTIPLY", "NEGA": "NEGATIVE", "FXMUL": "FXMULTIPLE"},
	"ELECTRODE":   {"V": "VOLTAGE", "PAT": "PATTERN"},
	"EXTRACT":     {"PARA": "PARAMETER", "MATE": "MATERIAL", "DELIMIT": "DELIMITER"},
	"GRID":        {"RATE": "RATIO", "STEP": "SPACE"},
	"INTERFACE":   {"MATE": "MATERIAL"},
	"LIGHT":       {"LA": "LAMBDA", "POW": "POWER", "ABSORPTION": "ASI"},
	"MODEL":       {"BGN": "BGNARROW", "FNC": "FNCURRENT"},
	"NFERMI":      {"V": "VOLTAGE", "PAT": "PATTERN"},
	"PFERMI":      {"V": "VOLTAGE", "PAT": "PATTERN"},
	"REGION":      {"XO": "XORIGIN", "YO": "YORIGIN", "MULTI": "MULTIPLY", "NC": "NCURRENT", "FXMUL": "FXMULTIPLE", "FYMUL": "FYMULTIPLE"},
	"RESTART":     {"ADD_IMP": "ADD_IMPURITY", "OLD_ELE": "OLD_ELECTRODE"},
	"STRUCTURE":   {"XB": "XBOUNDARY", "YB": "YBOUNDARY"},
	"SUBSTRATE":   {"C": "CONCENTRATION", "CONC": "CONCENTRATION", "V": "VOLTAGE"},
}

// Formats and parameter tables in docs/<STATEMENT>.md are the authority.
// Warnings deliberately distinguish "not catalogued" from "invalid simulator input".
var parameterCatalogue = makeParameterCatalogue()

func makeParameterCatalogue() map[string]map[string]struct{} {
	const box = " XMIN YMIN ZMIN XMAX YMAX ZMAX"
	const masks = " MASK ANDMASK ANDNEGA XPLUS YPLUS XMULTIPLE YMULTIPLE TRIM"
	lists := map[string]string{
		"TITLE": "", "IF": "", "ELSE": "", "ENDIF": "", "END": "",
		"BIAS":        "ENAME VSTEP NSTEP ERROR NFNAME PFNAME FLOAT BMODEL",
		"CONSTANT":    "TEMPERATURE EOX ESI ENI EIN EME EAL EPS EM1 EM2 EM3 EM4 EM5 EM6 EM7 EM8 EMOB HMOB NI TNO TP0 CN CP CSRH AN AP BN BP VBG CBG ALPHAN ALPHAP EMMIN HMMIN CREFN CREFP BETAN BETAP XREF ESREF HSREF SREFN SREFP VCN VCP VSN VSP FACTN FACTP NSI ASI ASG AGE EGGE QE CTRAP DTRAP TRAPMOB TMASS DRANGE ET FNMN FNMP FNBN FNBP",
		"CONVERGENCE": "VERROR FERROR CARRIER PHASE1 PHASE2 PHASE3 TIME STEP TRATE SKIP PS0R CYCLE NEWTON FSOR FCYCLE ABORT METHOD POISSON RANGECHECK FPRINT",
		"DEFINE":      "NAME VALUE CHARACTER VSTEP NSTEP FILE OPTIMIZE TARGET ERROR",
		"DEPOSIT":     "TYPE CHARGE" + box,
		"DOPE":        "TYPE ELEMENT PROFILE FILE GFILE CMAX CBACK XJ RP LD THETA PHI MULTIPLY PATTERN STEP XORIGIN YORIGIN AXIS NAME ZPLUS ZMULTIPLE XSHIFT YSHIFT ZSHIFT FXMULTIPLE FYMULTIPLE CRATE DETECT APPEND NEGATIVE XREPEAT YREPEAT LOCOS MATERIAL SKIP EXPAND" + box + masks,
		"ELECTRODE":   "VOLTAGE PATTERN NAME WORK RESISTANCE SUPREM4 LOCOS" + box + masks,
		"EXTRACT":     "PARAMETER FILE ENAME AXIS PATTERN OPTION NAME VALUE PEAK MATERIAL RADIUS STEP COORDINATE WIDTH DELIMITER" + box + masks,
		"GRID":        "RATIO SPACE TMIN TMAX XPLUS YPLUS ZPLUS TPLUS" + box,
		"INSERT":      "FILE",
		"INTERFACE":   "MATERIAL QF SN SP NEGATIVE CTRAP DTRAP TRAPMOB NAME" + box + masks,
		"LIGHT":       "LAMBDA POWER THETA PHI ASI NSI XREPEAT YREPEAT TYPE FILE SKIP NSTEP LMAX ZPLUS REVERSE NAME BACK EXPAND" + box + masks,
		"MODEL":       "MOBILITY BGNARROW SRH TUNNELING AUGER IMPACT TRAP FNCURRENT",
		"NFERMI":      "VOLTAGE AXIS DV PATTERN NAME EXPAND" + box + masks,
		"PFERMI":      "VOLTAGE AXIS DV PATTERN NAME EXPAND" + box + masks,
		"REGION":      "MATERIAL PATTERN XORIGIN YORIGIN AXIS FILE TN0 TP0 RN RP GR CURRENT NCURRENT PCURRENT BEAK ZBEAK EDGE MULTIPLY ZMULTIPLE FXMULTIPLE FYMULTIPLE XSHIFT YSHIFT ZSHIFT EMOB HMOB SHAPE NAME EXPAND ZPLUS" + box + masks,
		"RESTART":     "FILE NEW_FERMI ADD_IMPURITY OLD_ELECTRODE",
		"SAVE":        "FILE TYPE TSTEP OPTION",
		"STRUCTURE":   "XBOUNDARY YBOUNDARY XMAX YMAX ZMAX ZMIN CELL",
		"SUBSTRATE":   "TYPE CONCENTRATION VOLTAGE",
	}
	catalogue := make(map[string]map[string]struct{}, len(lists))
	for statement, list := range lists {
		parameters := make(map[string]struct{})
		for _, name := range strings.Fields(list) {
			parameters[name] = struct{}{}
		}
		catalogue[statement] = parameters
	}
	return catalogue
}
