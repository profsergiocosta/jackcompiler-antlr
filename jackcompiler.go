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

		
		function void main() {
			var int a, b  ,c ;


			let a = 10 + 50 * 60 - 80;

			let b = false;
			let c = "ola";
		

			return ;
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
