package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"
)

type jackListener struct {
	*parser.BaseJackListener
}

func main() {
	// Setup the input
	is := antlr.NewInputStream(`
	class Main {
		static a;
	}
				`)

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
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewJackParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&jackListener{}, p.Start())

	fmt.Println("Fim")
	
}