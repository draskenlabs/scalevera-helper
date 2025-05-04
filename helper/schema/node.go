package schema

import (
	"fmt"
	"strings"

	"github.com/draskenlabs/drasken-go-lexer/lexer"
)

type ScaleveraNode struct {
	Key      string
	Value    string
	Token    lexer.Token
	Children []*ScaleveraNode
}

var keywords = []string{
	"global",
	"platform",

	"plan",
	"step",
	"task",
	"module",
	"resource",

	"requirements",
	"require",
	"provider",
	"source",
	"version",
}

func GenerateNewNode(token *lexer.Token) *ScaleveraNode {
	return &ScaleveraNode{
		Key:   strings.TrimSpace(token.Literal),
		Value: strings.TrimSpace(token.Literal),
		Token: *token,
	}
}

func (n *ScaleveraNode) IsKeyword() bool {
	for _, keyword := range keywords {
		if strings.TrimSpace(n.Token.Literal) == keyword {
			return true
		}
	}
	return false
}

func (n *ScaleveraNode) IsIdentifier() bool {
	return n.Token.Type == lexer.IDENT && !n.IsKeyword()
}

func (n *ScaleveraNode) IsBraceBlockStart() bool {
	return strings.TrimSpace(n.Token.Literal) == "{"
}

func (n *ScaleveraNode) IsBraceBlockEnd() bool {
	return strings.TrimSpace(n.Token.Literal) == "}"
}

func (n *ScaleveraNode) IsSingleLinerString() bool {
	return n.Token.Type == lexer.STRING && (strings.TrimSpace(n.Token.Literal)[0] == '"' || strings.TrimSpace(n.Token.Literal)[0] == '\'')
}

func (n *ScaleveraNode) IsAssignmentOperator() bool {
	return strings.TrimSpace(n.Token.Literal) == "="
}

func (n *ScaleveraNode) IsTerminatorOperator() bool {
	return strings.TrimSpace(n.Token.Literal) == ";"
}

func (n *ScaleveraNode) IsValue() bool {
	return n.Token.Type == lexer.STRING || n.Token.Type == lexer.INT || n.Token.Type == lexer.BOOL
}

func (n *ScaleveraNode) IsEOF() bool {
	return n.Token.Type == lexer.EOF
}

func (n *ScaleveraNode) String() string {
	return fmt.Sprintf("ScaleveraNode{Key: %s, Value: %s, Token: %s, Children: %s}", n.Key, n.Value, n.Token, n.Children)
}

// Helpers for more core and semantic functions
func (n *ScaleveraNode) IsRequirementsNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "requirements"
}

func (n *ScaleveraNode) IsRequireNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "require"
}

func (n *ScaleveraNode) IsProviderNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "provider"
}

func (n *ScaleveraNode) IsGlobalNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "global"
}

func (n *ScaleveraNode) IsPlanNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "plan"
}

func (n *ScaleveraNode) IsStepNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "step"
}

func (n *ScaleveraNode) IsTaskNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "task"
}

func (n *ScaleveraNode) IsModuleNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "module"
}

func (n *ScaleveraNode) IsResourceNode() bool {
	return n.IsKeyword() && strings.TrimSpace(n.Token.Literal) == "resource"
}
