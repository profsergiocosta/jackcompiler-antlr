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
	is := antlr.NewInputStream(
	`
	class Main {
		function void main() {
			var Array a;
			var int length;
			var int i, sum;
		
		let length = Keyboard.readInt("HOW MANY NUMBERS? ");
		let a = Array.new(length);
		let i = 0;
		
		while (i < length) {
			let a[i] = Keyboard.readInt("ENTER THE NEXT NUMBER: ");
			let i = i + 1;
		}
		
		let i = 0;
		let sum = 0;
		
		while (i < length) {
			let sum = sum + a[i];
			let i = i + 1;
		}
		
		do Output.printString("THE AVERAGE IS: ");
		do Output.printInt(sum / length);
		do Output.println();
		
		return 0;
		}
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

	tree := p.Classdef()
	
	//fmt.Println (tree.ToStringTree(p) )

	antlr.ParseTreeWalkerDefault.Walk(NewJackListener(), tree)
	

	fmt.Println("Fim")
	
}