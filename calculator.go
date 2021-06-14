package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	EOF
)

type Token struct {
	Type  TokenType
	Value float64
}

func (t Token) String() string {
	switch t.Type {
	case NUMBER:
		return "NUMBER:" + fmt.Sprint(t.Value)
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case MULTIPLY:
		return "MULTIPLY"
	case DIVIDE:
		return "DIVIDE"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

type Lexer struct {
	text    string
	current byte
	pos     int
}

func (l *Lexer) advance() {
	if len(l.text) > l.pos {
		l.current = l.text[l.pos]
		l.pos++
	} else {
		l.current = byte(0)
	}
}

func LexerInit(text string) *Lexer {
	lexer := &Lexer{text: text, pos: 0}
	lexer.advance()
	return lexer
}

func (l *Lexer) GetTokens() ([]Token, error) {
	tokens := []Token{}

	for l.current != byte(0) {
		if unicode.IsSpace(rune(l.current)) {
			l.advance()
		} else if unicode.IsNumber(rune(l.current)) || l.current == '.' {
			result, err := l.getNumber()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, result)
		} else if l.current == '+' {
			l.advance()
			tokens = append(tokens, Token{Type: PLUS})
		} else if l.current == '-' {
			l.advance()
			tokens = append(tokens, Token{Type: MINUS})
		} else if l.current == '*' {
			l.advance()
			tokens = append(tokens, Token{Type: MULTIPLY})
		} else if l.current == '/' {
			l.advance()
			tokens = append(tokens, Token{Type: DIVIDE})
		} else if l.current == '(' {
			l.advance()
			tokens = append(tokens, Token{Type: LPAREN})
		} else if l.current == ')' {
			l.advance()
			tokens = append(tokens, Token{Type: RPAREN})
		} else {
			return nil, fmt.Errorf("illegal character '%c'", l.current)
		}
	}

	return tokens, nil
}

func (l *Lexer) getNumber() (Token, error) {
	decimal_point_count := 0
	number_str := string(l.current)
	l.advance()

	for l.current != byte(0) && (unicode.IsNumber(rune(l.current)) || l.current == '.') {
		if l.current == '.' {
			decimal_point_count += 1
			if decimal_point_count > 1 {
				break
			}
		}

		number_str += string(l.current)
		l.advance()
	}

	if number_str[0] == '.' {
		number_str = "0" + number_str
	}

	if number_str[len(number_str)-1] == '.' {
		number_str += "0"
	}

	number, err := strconv.ParseFloat(number_str, 64)

	if err != nil {
		return Token{}, err
	}

	return Token{Type: NUMBER, Value: number}, nil
}

type NodeType int

const (
	NumberNodeType NodeType = iota
	AddNodeType
	SubtractNodeType
	MultiplyNodeType
	DivideNodeType
	PlusNodeType
	MinusNodeType
)

type Node interface {
	Type() NodeType
	String() string
}

type NumberNode struct {
	Value float64
}

func (n *NumberNode) Type() NodeType {
	return NumberNodeType
}

func (n *NumberNode) String() string {
	return fmt.Sprint(n.Value)
}

type AddNode struct {
	NodeA Node
	NodeB Node
}

func (n *AddNode) Type() NodeType {
	return AddNodeType
}

func (n *AddNode) String() string {
	return fmt.Sprintf("(%s+%s)", n.NodeA, n.NodeB)
}

type SubtractNode struct {
	NodeA Node
	NodeB Node
}

func (n *SubtractNode) Type() NodeType {
	return SubtractNodeType
}

func (n *SubtractNode) String() string {
	return fmt.Sprintf("(%s-%s)", n.NodeA, n.NodeB)
}

type MultiplyNode struct {
	NodeA Node
	NodeB Node
}

func (n *MultiplyNode) Type() NodeType {
	return MultiplyNodeType
}

func (n *MultiplyNode) String() string {
	return fmt.Sprintf("(%s*%s)", n.NodeA, n.NodeB)
}

type DivideNode struct {
	NodeA Node
	NodeB Node
}

func (n *DivideNode) Type() NodeType {
	return DivideNodeType
}

func (n *DivideNode) String() string {
	return fmt.Sprintf("(%s/%s)", n.NodeA, n.NodeB)
}

type PlusNode struct {
	Node Node
}

func (n *PlusNode) Type() NodeType {
	return PlusNodeType
}

func (n *PlusNode) String() string {
	return fmt.Sprintf("(+%s)", n.Node)
}

type MinusNode struct {
	Node Node
}

func (n *MinusNode) Type() NodeType {
	return MinusNodeType
}

func (n *MinusNode) String() string {
	return fmt.Sprintf("(-%s)", n.Node)
}

type Parser struct {
	tokens  []Token
	current *Token
	pos     int
}

func (p *Parser) advance() {
	if len(p.tokens) > p.pos {
		p.current = &p.tokens[p.pos]
		p.pos++
	} else {
		p.current = nil
	}
}

func ParserInit(tokens []Token) *Parser {
	parser := &Parser{tokens: tokens}
	parser.advance()
	return parser
}

func (p *Parser) raiseError() error {
	return fmt.Errorf("invalid syntax")
}

func (p *Parser) Parse() (Node, error) {
	if p.current == nil {
		return nil, nil
	}

	result, err := p.expr()

	if err != nil {
		return nil, err
	}

	if p.current != nil {
		return nil, p.raiseError()
	}

	return result, nil
}

func (p *Parser) expr() (Node, error) {
	result, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.current != nil && (p.current.Type == PLUS || p.current.Type == MINUS) {
		switch p.current.Type {
		case PLUS:
			p.advance()
			res, err := p.term()

			if err != nil {
				return nil, err
			}

			result = &AddNode{NodeA: result, NodeB: res}
		case MINUS:
			p.advance()
			res, err := p.term()

			if err != nil {
				return nil, err
			}

			result = &SubtractNode{NodeA: result, NodeB: res}
		}
	}

	return result, nil
}

func (p *Parser) term() (Node, error) {
	result, err := p.factor()

	if err != nil {
		return nil, err
	}

	for p.current != nil && (p.current.Type == MULTIPLY || p.current.Type == DIVIDE) {
		switch p.current.Type {
		case MULTIPLY:
			p.advance()
			res, err := p.factor()

			if err != nil {
				return nil, err
			}

			result = &MultiplyNode{NodeA: result, NodeB: res}
		case DIVIDE:
			p.advance()
			res, err := p.factor()

			if err != nil {
				return nil, err
			}

			result = &DivideNode{NodeA: result, NodeB: res}
		}
	}

	return result, nil
}

func (p *Parser) factor() (Node, error) {
	token := p.current

	if token == nil {
		return nil, p.raiseError()
	}

	switch token.Type {
	case LPAREN:
		p.advance()
		result, err := p.expr()

		if err != nil {
			return nil, err
		}

		if p.current.Type != RPAREN {
			return nil, p.raiseError()
		}

		p.advance()
		return result, nil
	case NUMBER:
		p.advance()
		return &NumberNode{Value: token.Value}, nil
	case PLUS:
		p.advance()
		result, err := p.factor()

		if err != nil {
			return nil, err
		}

		return &PlusNode{Node: result}, nil
	case MINUS:
		p.advance()
		result, err := p.factor()

		if err != nil {
			return nil, err
		}

		return &MinusNode{Node: result}, nil
	}

	return nil, p.raiseError()
}

type Number struct {
	Value float64
}

func (n Number) String() string {
	return fmt.Sprint(n.Value)
}

func visit(node Node) (Number, error) {
	switch node.Type() {
	case NumberNodeType:
		return Number{Value: node.(*NumberNode).Value}, nil
	case AddNodeType:
		node := node.(*AddNode)
		value1, err := visit(node.NodeA)

		if err != nil {
			return Number{}, err
		}

		value2, err := visit(node.NodeB)

		if err != nil {
			return Number{}, err
		}

		return Number{Value: value1.Value + value2.Value}, nil
	case SubtractNodeType:
		node := node.(*SubtractNode)
		value1, err := visit(node.NodeA)

		if err != nil {
			return Number{}, err
		}

		value2, err := visit(node.NodeB)

		if err != nil {
			return Number{}, err
		}

		return Number{Value: value1.Value - value2.Value}, nil
	case MultiplyNodeType:
		node := node.(*MultiplyNode)
		value1, err := visit(node.NodeA)

		if err != nil {
			return Number{}, err
		}

		value2, err := visit(node.NodeB)

		if err != nil {
			return Number{}, err
		}

		return Number{Value: value1.Value * value2.Value}, nil
	case DivideNodeType:
		node := node.(*DivideNode)
		value1, err := visit(node.NodeA)

		if err != nil {
			return Number{}, err
		}

		value2, err := visit(node.NodeB)

		if err != nil {
			return Number{}, err
		}

		if value2.Value == 0 {
			return Number{}, fmt.Errorf("runtime math error")
		}

		return Number{Value: value1.Value / value2.Value}, nil
	case PlusNodeType:
		node := node.(*PlusNode)
		value, err := visit(node.Node)

		if err != nil {
			return Number{}, err
		}

		return value, nil
	case MinusNodeType:
		node := node.(*MinusNode)
		value, err := visit(node.Node)

		if err != nil {
			return Number{}, err
		}

		return Number{Value: -value.Value}, nil
	default:
		return Number{}, fmt.Errorf("unknown node")
	}
}

func CalculatorRun(text string) (string, error) {
	if text == "" {
		return "", fmt.Errorf("empty text/string")
	}

	lexer := LexerInit(text)
	tokens, err := lexer.GetTokens()

	if err != nil {
		return "", err
	}

	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }

	parser := ParserInit(tokens)
	tree, err := parser.Parse()

	if err != nil {
		return "", err
	}

	result, err := visit(tree)

	if err != nil {
		return "", err
	}

	return result.String(), nil
}
