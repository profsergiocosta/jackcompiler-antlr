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
		function void main () {
		  var String nome;
		  let nome = Keyboard.readLine("Qual o seu nome? ");
		  do Output.printString("Ola "); 
		  do Output.printString(nome); 
		  do Output.println(); 
		  return;
		 }
   }
				`)

	// Create the Lexer
	lexer := parser.NewJackLexer(is)


	
	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}


	/*
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewJackParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&jackListener{}, p.Start())
	*/
}