// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 10, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 5, 2, 8, 10, 2, 3, 2, 2, 2, 3, 2, 2, 2, 2, 9,
	2, 7, 3, 2, 2, 2, 4, 8, 7, 5, 2, 2, 5, 6, 7, 4, 2, 2, 6, 8, 7, 2, 2, 3,
	7, 4, 3, 2, 2, 2, 7, 5, 3, 2, 2, 2, 8, 3, 3, 2, 2, 2, 3, 7,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'='", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'~'",
	"'<'", "'>'", "'=='", "'.'", "','", "';'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'METHOD'", "'STATIC'", "'INT'", "'BOOLEAN'", "'TRUE'", "'NULL'",
	"'LET'", "'IF'", "'WHILE'", "'CONSTRUCTOR'", "'FIELD'", "'VAR'", "'CHAR'",
	"'VOID'", "'CLASS'", "'FALSE'", "'DO'", "'ELSE'", "'RETURN'", "'FUNCTION'",
	"'THIS'",
}
var symbolicNames = []string{
	"", "ID", "STRING", "INTEGER", "ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH",
	"AND", "OR", "NOT", "LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC",
	"INT", "BOOLEAN", "TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR",
	"FIELD", "VAR", "CHAR", "VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN",
	"FUNCTION", "THIS", "WHITESPACE",
}

var ruleNames = []string{
	"start",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JackParser struct {
	*antlr.BaseParser
}

func NewJackParser(input antlr.TokenStream) *JackParser {
	this := new(JackParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Jack.g4"

	return this
}

// JackParser tokens.
const (
	JackParserEOF         = antlr.TokenEOF
	JackParserID          = 1
	JackParserSTRING      = 2
	JackParserINTEGER     = 3
	JackParserASSIGN      = 4
	JackParserPLUS        = 5
	JackParserMINUS       = 6
	JackParserASTERISK    = 7
	JackParserSLASH       = 8
	JackParserAND         = 9
	JackParserOR          = 10
	JackParserNOT         = 11
	JackParserLT          = 12
	JackParserGT          = 13
	JackParserEQ          = 14
	JackParserDOT         = 15
	JackParserCOMMA       = 16
	JackParserSEMICOLON   = 17
	JackParserLPAREN      = 18
	JackParserRPAREN      = 19
	JackParserLBRACE      = 20
	JackParserRBRACE      = 21
	JackParserLBRACKET    = 22
	JackParserRBRACKET    = 23
	JackParserMETHOD      = 24
	JackParserSTATIC      = 25
	JackParserINT         = 26
	JackParserBOOLEAN     = 27
	JackParserTRUE        = 28
	JackParserNULL        = 29
	JackParserLET         = 30
	JackParserIF          = 31
	JackParserWHILE       = 32
	JackParserCONSTRUCTOR = 33
	JackParserFIELD       = 34
	JackParserVAR         = 35
	JackParserCHAR        = 36
	JackParserVOID        = 37
	JackParserCLASS       = 38
	JackParserFALSE       = 39
	JackParserDO          = 40
	JackParserELSE        = 41
	JackParserRETURN      = 42
	JackParserFUNCTION    = 43
	JackParserTHIS        = 44
	JackParserWHITESPACE  = 45
)

// JackParserRULE_start is the JackParser rule.
const JackParserRULE_start = 0

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(JackParserINTEGER, 0)
}

func (s *StartContext) STRING() antlr.TerminalNode {
	return s.GetToken(JackParserSTRING, 0)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(JackParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *JackParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JackParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(5)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JackParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(2)
			p.Match(JackParserINTEGER)
		}

	case JackParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(3)
			p.Match(JackParserSTRING)
		}
		{
			p.SetState(4)
			p.Match(JackParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
