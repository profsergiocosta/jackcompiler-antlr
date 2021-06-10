// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJackListener is a complete listener for a parse tree produced by JackParser.
type BaseJackListener struct{}

var _ JackListener = &BaseJackListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJackListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJackListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJackListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJackListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseJackListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseJackListener) ExitStart(ctx *StartContext) {}

// EnterClassvardec is called when production classvardec is entered.
func (s *BaseJackListener) EnterClassvardec(ctx *ClassvardecContext) {}

// ExitClassvardec is called when production classvardec is exited.
func (s *BaseJackListener) ExitClassvardec(ctx *ClassvardecContext) {}
