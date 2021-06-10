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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 18, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 4, 2, 4, 2, 3, 4, 2, 24, 24, 33, 33, 2, 15,
	2, 6, 3, 2, 2, 2, 4, 13, 3, 2, 2, 2, 6, 7, 7, 37, 2, 2, 7, 8, 7, 44, 2,
	2, 8, 9, 7, 19, 2, 2, 9, 10, 5, 4, 3, 2, 10, 11, 7, 20, 2, 2, 11, 12, 7,
	2, 2, 3, 12, 3, 3, 2, 2, 2, 13, 14, 9, 2, 2, 2, 14, 15, 7, 44, 2, 2, 15,
	16, 7, 16, 2, 2, 16, 5, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'~'", "'<'", "'>'",
	"'=='", "'.'", "','", "';'", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'method'", "'static'", "'int'", "'boolean'", "'true'", "'null'", "'let'",
	"'if'", "'while'", "'constructor'", "'field'", "'var'", "'char'", "'void'",
	"'class'", "'false'", "'do'", "'else'", "'return'", "'function'", "'this'",
}
var symbolicNames = []string{
	"", "ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT",
	"LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN",
	"TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", "CHAR",
	"VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", "FUNCTION", "THIS", "ID",
	"STRING", "INTEGER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "classvardec",
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
	JackParserASSIGN      = 1
	JackParserPLUS        = 2
	JackParserMINUS       = 3
	JackParserASTERISK    = 4
	JackParserSLASH       = 5
	JackParserAND         = 6
	JackParserOR          = 7
	JackParserNOT         = 8
	JackParserLT          = 9
	JackParserGT          = 10
	JackParserEQ          = 11
	JackParserDOT         = 12
	JackParserCOMMA       = 13
	JackParserSEMICOLON   = 14
	JackParserLPAREN      = 15
	JackParserRPAREN      = 16
	JackParserLBRACE      = 17
	JackParserRBRACE      = 18
	JackParserLBRACKET    = 19
	JackParserRBRACKET    = 20
	JackParserMETHOD      = 21
	JackParserSTATIC      = 22
	JackParserINT         = 23
	JackParserBOOLEAN     = 24
	JackParserTRUE        = 25
	JackParserNULL        = 26
	JackParserLET         = 27
	JackParserIF          = 28
	JackParserWHILE       = 29
	JackParserCONSTRUCTOR = 30
	JackParserFIELD       = 31
	JackParserVAR         = 32
	JackParserCHAR        = 33
	JackParserVOID        = 34
	JackParserCLASS       = 35
	JackParserFALSE       = 36
	JackParserDO          = 37
	JackParserELSE        = 38
	JackParserRETURN      = 39
	JackParserFUNCTION    = 40
	JackParserTHIS        = 41
	JackParserID          = 42
	JackParserSTRING      = 43
	JackParserINTEGER     = 44
	JackParserWHITESPACE  = 45
)

// JackParser rules.
const (
	JackParserRULE_start       = 0
	JackParserRULE_classvardec = 1
)

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

func (s *StartContext) CLASS() antlr.TerminalNode {
	return s.GetToken(JackParserCLASS, 0)
}

func (s *StartContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *StartContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserLBRACE, 0)
}

func (s *StartContext) Classvardec() IClassvardecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassvardecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassvardecContext)
}

func (s *StartContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JackParserRBRACE, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.Match(JackParserCLASS)
	}
	{
		p.SetState(5)
		p.Match(JackParserID)
	}
	{
		p.SetState(6)
		p.Match(JackParserLBRACE)
	}
	{
		p.SetState(7)
		p.Classvardec()
	}
	{
		p.SetState(8)
		p.Match(JackParserRBRACE)
	}
	{
		p.SetState(9)
		p.Match(JackParserEOF)
	}

	return localctx
}

// IClassvardecContext is an interface to support dynamic dispatch.
type IClassvardecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassvardecContext differentiates from other interfaces.
	IsClassvardecContext()
}

type ClassvardecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassvardecContext() *ClassvardecContext {
	var p = new(ClassvardecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JackParserRULE_classvardec
	return p
}

func (*ClassvardecContext) IsClassvardecContext() {}

func NewClassvardecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassvardecContext {
	var p = new(ClassvardecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JackParserRULE_classvardec

	return p
}

func (s *ClassvardecContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassvardecContext) ID() antlr.TerminalNode {
	return s.GetToken(JackParserID, 0)
}

func (s *ClassvardecContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(JackParserSEMICOLON, 0)
}

func (s *ClassvardecContext) STATIC() antlr.TerminalNode {
	return s.GetToken(JackParserSTATIC, 0)
}

func (s *ClassvardecContext) FIELD() antlr.TerminalNode {
	return s.GetToken(JackParserFIELD, 0)
}

func (s *ClassvardecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassvardecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassvardecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.EnterClassvardec(s)
	}
}

func (s *ClassvardecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JackListener); ok {
		listenerT.ExitClassvardec(s)
	}
}

func (p *JackParser) Classvardec() (localctx IClassvardecContext) {
	localctx = NewClassvardecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JackParserRULE_classvardec)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	_la = p.GetTokenStream().LA(1)

	if !(_la == JackParserSTATIC || _la == JackParserFIELD) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(12)
		p.Match(JackParserID)
	}
	{
		p.SetState(13)
		p.Match(JackParserSEMICOLON)
	}

	return localctx
}
