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

		function void f1() {
			return;
		}

		function void main(int x) {
			var Point p;
			var int a, b  ,c ;
			let b = Main.f1();
			let a = p.getX();
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
