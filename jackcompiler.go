package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"

	"github.com/profsergiocosta/jackcompiler-antlr/listener"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream(
		`
	class Main {

		
		function void main(int x) {
			var int a, b  ,c ;
			let a[10] = 42;
			let c = x;

		
		

			return 10;
		}

		function void x() {
			var int a, b  ,c ;
			let a[10] = 42;

		
		}
	}
	
	
	`)

	// Create the Lexer
	lexer := parser.NewJackLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewJackParser(stream)

	// generate AST
	tree := p.Classdef()

	// visit tree
	antlr.ParseTreeWalkerDefault.Walk(listener.New(), tree)

	fmt.Println("compiled")

}
