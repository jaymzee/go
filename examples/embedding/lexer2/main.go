package main

// embedding example
//
// say we are creating a Java Lexer
//
// public class Hello {}
//
// <public keyword>, "public"
// <class keyword>, "class"
// <identifier>, "Hello"
// <left-bracket>, "{"
// <right-bracket>, "}"

import "fmt"

type TokenType uint16

const (
	KEYWORD TokenType = iota
	IDENTIFIER
	LBRACKET
	RBRACKET
	INT
)

func (t TokenType) String() string {
	return [...]string{
		"KEYWORD",
		"IDENTIFIER",
		"LBRACKET",
		"RBRACKET",
		"INT",
	}[t]
}

type Token interface {
	Type() TokenType
	Lexeme() string
}

// Match implements Token interface
type Match struct {
	toktype TokenType // type is a reserved word
	lexeme string
}

func (m *Match) Type() TokenType {
	return m.toktype
}

func (m *Match) Lexeme() string {
	return m.lexeme
}

// IntegerConstant also implements Token interface
// by not giving Token field a name, it "inherits" the methods and
// fields of Token.
type IntegerConstant struct {
	Token
	value uint64
}

func (i *IntegerConstant) Value() uint64 {
	return i.value
}

func display(tokens []Token) {
	for _, t := range tokens {
		switch t := t.(type) {
		case *IntegerConstant:
			fmt.Printf("{%s %q} = %d\n", t.Type(), t.Lexeme(), t.Value())
		default:
			fmt.Printf("{%s %q}\n", t.Type(), t.Lexeme())
		}
	}
}

func main() {
	c := Match{KEYWORD, "class"}
	p := Match{KEYWORD, "public"}
	h := Match{IDENTIFIER, "Hello"}
	l := Match{LBRACKET, "{"}
	r := Match{RBRACKET, "}"}

	i := IntegerConstant{&Match{INT, "42"}, 42}
	j := IntegerConstant{&Match{INT, "88"}, 88}

	tokens := []Token{&p, &c, &h, &l, &r}
	display(tokens)

	ints := []Token{&i, &j}
	display(ints)
}
