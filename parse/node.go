package parse

import "fmt"

// A node is an item in the AST
type Node interface {
	Pos() pos
	String() string
}

type pos struct {
	Line   int
	Offset int
}

func newPos(line, offset int) pos {
	return pos{line, offset}
}

func (p pos) Pos() pos {
	return p
}

func (p pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Offset)
}

// ModuleNode represents a set of nodes.
type ModuleNode struct {
	*BodyNode
	parent *ExtendsNode
}

func newModuleNode(nodes ...Node) *ModuleNode {
	return &ModuleNode{newBodyNode(newPos(1, 0), nodes...), nil}
}

func (l *ModuleNode) Parent() *ExtendsNode {
	return l.parent
}

func (l *ModuleNode) String() string {
	return fmt.Sprintf("Module%s", l.nodes)
}

type BodyNode struct {
	pos
	nodes []Node
}

func newBodyNode(pos pos, nodes ...Node) *BodyNode {
	return &BodyNode{pos, nodes}
}

func (l *BodyNode) append(n Node) {
	l.nodes = append(l.nodes, n)
}

func (l *BodyNode) String() string {
	return fmt.Sprintf("Body%s", l.nodes)
}

func (l *BodyNode) Children() []Node {
	return l.nodes
}

// TextNode represents raw, non Stick source code, like plain HTML.
type TextNode struct {
	pos
	data string
}

func newTextNode(data string, p pos) *TextNode {
	return &TextNode{p, data}
}

func (t *TextNode) String() string {
	return fmt.Sprintf("Text(%s)", t.data)
}

func (t *TextNode) Text() string {
	return t.data
}

// PrintNode represents a print statement
type PrintNode struct {
	pos
	exp Expr
}

func newPrintNode(exp Expr, p pos) *PrintNode {
	return &PrintNode{p, exp}
}

func (t *PrintNode) Expr() Expr {
	return t.exp
}

func (t *PrintNode) String() string {
	return fmt.Sprintf("Print(%s)", t.exp)
}

// BlockNode represents a block statement
type BlockNode struct {
	pos
	name string
	body Node
}

func newBlockNode(name string, body Node, p pos) *BlockNode {
	return &BlockNode{p, name, body}
}

func (t *BlockNode) Name() string {
	return t.name
}

func (t *BlockNode) Body() Node {
	return t.body
}

func (t *BlockNode) String() string {
	return fmt.Sprintf("Block(%s: %s)", t.name, t.body)
}

// IfNode represents an if statement
type IfNode struct {
	pos
	cond Expr
	body Node
	els  Node
}

func newIfNode(cond Expr, body Node, els Node, p pos) *IfNode {
	return &IfNode{p, cond, body, els}
}

func (t *IfNode) Cond() Expr {
	return t.cond
}

func (t *IfNode) Body() Node {
	return t.body
}

func (t *IfNode) Else() Node {
	return t.els
}

func (t *IfNode) String() string {
	return fmt.Sprintf("If(%s: %s Else: %s)", t.cond, t.body, t.els)
}

// ExtendsNode represents an extends statement
type ExtendsNode struct {
	pos
	tplRef Expr
}

func newExtendsNode(tplRef Expr, p pos) *ExtendsNode {
	return &ExtendsNode{p, tplRef}
}

func (t *ExtendsNode) TemplateRef() Expr {
	return t.tplRef
}

func (t *ExtendsNode) String() string {
	return fmt.Sprintf("Extends(%s)", t.tplRef)
}

// ForNode represents a for loop construct.
type ForNode struct {
	pos
	key  *NameExpr
	val  *NameExpr
	expr Expr
	body Node
	els  Node
}

func newForNode(k, v *NameExpr, expr, body, els Node, p pos) *ForNode {
	return &ForNode{p, k, v, expr, body, els}
}

func (t *ForNode) String() string {
	return fmt.Sprintf("For(%s, %s in %s: %s else %s)", t.key, t.val, t.expr, t.body, t.els)
}

// IncludeNode is an include statement.
type IncludeNode struct {
	pos
	tmpl Expr
	with Expr
}

func newIncludeNode(tmpl Expr, with Expr, pos pos) *IncludeNode {
	return &IncludeNode{pos, tmpl, with}
}

func (t *IncludeNode) String() string {
	return fmt.Sprintf("Include(%s with %s)", t.tmpl, t.with)
}
