package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type TokenType int

const (
	TokenUnkown TokenType = iota
	OpenBracket
	ClosedBracket
	FnBlock
	Function
	Variable
	NewLine
)

type Token struct {
	Type  TokenType
	Name  string
	Value string
}

func Tokenize(input string) []Token {
	var tokens []Token

	for idx, ch := range input {

		switch ch {
		case '[':
			tokens = append(tokens, Token{OpenBracket, "O_BRACKET", string(ch)})
		case ']':
			tokens = append(tokens, Token{ClosedBracket, "C_BRACKET", string(ch)})
		case '-':
			if idx < len(input)-1 {
				// fmt.Println("Checking next character\n")
				// fmt.Println(string(input[idx+1]))

				nextChar := input[idx+1]

				if nextChar == '>' {
					tokens = append(tokens, Token{FnBlock, "FN_BLOCK", string("->")})
				}
			}
		case '\n':
			tokens = append(tokens, Token{NewLine, "NEW_LINE", string(ch)})

		case '>':
			continue
		case '@':
			tokens = append(tokens, Token{Variable, "VARIABLE", string(ch)})
		default:
			// log.Fatalf("Invalid token: %s", string(ch))
			// just ignore other characters for now
			continue
		}
	}

	return tokens
}

func main() {

	source, err := ioutil.ReadFile("program.snk")

	if err != nil {
		log.Fatal(err)
	}

	// input := " @x=1 []->[]"

	input := string(source)

	tokens := Tokenize(input)

	for _, token := range tokens {
		fmt.Printf("%v -- %v: %v\n", token.Type, token.Name, token.Value)
	}
}
