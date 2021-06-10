// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JackListener is a complete listener for a parse tree produced by JackParser.
type JackListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterClassvardec is called when entering the classvardec production.
	EnterClassvardec(c *ClassvardecContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitClassvardec is called when exiting the classvardec production.
	ExitClassvardec(c *ClassvardecContext)
}
