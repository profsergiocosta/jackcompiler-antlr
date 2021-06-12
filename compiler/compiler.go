package compiler

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"

	"github.com/profsergiocosta/jackcompiler-antlr/listener"
)

func compile(s string) {

	// Setup the input
	is := antlr.NewInputStream(s)

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
