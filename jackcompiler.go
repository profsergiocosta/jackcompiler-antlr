package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"

	
)

type JackListener struct {
	*parser.BaseJackListener
}


func NewJackListener() *JackListener {
	return new(JackListener)
}

func (this *JackListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
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

	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//p.BuildParseTrees = true

	// Finally parse the expression

	tree := p.Start()
	
	//fmt.Println (tree.ToStringTree(p) )

	antlr.ParseTreeWalkerDefault.Walk(NewJackListener(), tree)
	

	fmt.Println("Fim")
	
}