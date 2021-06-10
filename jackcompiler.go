package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/profsergiocosta/jackcompiler-go/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	
	// Create the Lexer
	lexer := parser.NewJackLexer(is)
/*
	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	*/
}