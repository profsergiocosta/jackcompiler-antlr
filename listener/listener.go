package listener

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"
	"github.com/profsergiocosta/jackcompiler-antlr/symboltable"
	"github.com/profsergiocosta/jackcompiler-antlr/vmwriter"
)

type JackListener struct {
	*parser.BaseJackListener

	st *symboltable.SymbolTable

	vm *vmwriter.VMWriter

	className    string
	functionName string

	subroutineKind int
}

func New() *JackListener {

	s := &JackListener{}
	s.st = symboltable.NewSymbolTable()

	pathName := "/home/sergio/go/src/github.com/profsergiocosta/jackcompiler-antlr/teste.jack"
	s.vm = vmwriter.New(filenameWithoutExtension(pathName) + ".vm")

	return s
}

func (s *JackListener) EnterClassdef(ctx *parser.ClassdefContext) {
	s.className = ctx.Classname().GetText()
}

func (s *JackListener) EnterClassvardec(ctx *parser.ClassvardecContext) {

	var scope symboltable.SymbolScope

	scope = symboltable.STATIC

	if ctx.GetKind().GetTokenType() == parser.JackLexerFIELD {
		scope = symboltable.FIELD
	}

	ttype := ctx.Atype().GetText()

	varnames := ctx.AllVarname()
	for _, element := range varnames {
		s.st.Define(element.GetText(), ttype, scope)
	}

}

func (s *JackListener) EnterVardec(ctx *parser.VardecContext) {

	var scope symboltable.SymbolScope = symboltable.VAR

	ttype := ctx.Atype().GetText()

	varnames := ctx.AllVarname()
	for _, element := range varnames {
		s.st.Define(element.GetText(), ttype, scope)
	}
}

func (s *JackListener) EnterParameter(ctx *parser.ParameterContext) {
	var scope symboltable.SymbolScope = symboltable.ARG
	ttype := ctx.Atype().GetText()
	name := ctx.Varname().GetText()
	s.st.Define(name, ttype, scope)
}

func (s *JackListener) ExitVardecs(ctx *parser.VardecsContext) {
	nLocals := s.st.VarCount(symboltable.VAR)
	s.vm.WriteFunction(s.functionName, nLocals)

	switch s.subroutineKind {
	case parser.JackLexerMETHOD:
		s.st.Define("this", s.className, symboltable.ARG)
		s.vm.WritePush(vmwriter.ARG, 0)
		s.vm.WritePop(vmwriter.POINTER, 0)
	case parser.JackLexerCONSTRUCTOR:
		s.vm.WritePush(vmwriter.CONST, s.st.VarCount(symboltable.FIELD))
		s.vm.WriteCall("Memory.alloc", 1)
		s.vm.WritePop(vmwriter.POINTER, 0)
	}

}

func (s *JackListener) EnterSubrotinedec(ctx *parser.SubrotinedecContext) {

	s.functionName = s.className + "." + ctx.Subroutinename().GetText()
	s.subroutineKind = ctx.GetKind().GetTokenType()
}

// ExitLvalue is called when production lvalue is exited.
func (s *JackListener) ExitLvalue(ctx *parser.LvalueContext) {
	sym := s.st.Resolve(ctx.GetIdent().GetText())
	if ctx.LBRACKET() != nil {
		s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
		s.vm.WriteArithmetic(vmwriter.ADD)
	}
}

func (s *JackListener) ExitLetStatement(ctx *parser.LetStatementContext) {
	sym := s.st.Resolve(ctx.Lvalue().GetIdent().GetText())
	if ctx.Lvalue().GetChildCount() > 1 { // is array ?
		s.vm.WritePop(vmwriter.TEMP, 0)
		s.vm.WritePop(vmwriter.POINTER, 1)
		s.vm.WritePush(vmwriter.TEMP, 0)
		s.vm.WritePop(vmwriter.THAT, 0)
	} else {
		s.vm.WritePop(scopeToSegment(sym.Scope), sym.Index)
	}

}

func (s *JackListener) ExitReturnStatement(ctx *parser.ReturnStatementContext) {
	if ctx.GetChildCount() > 2 { // has expression ?
		s.vm.WriteReturn()
	} else {
		s.vm.WritePush(vmwriter.CONST, 0)
		s.vm.WriteReturn()
	}
}

func (s *JackListener) EnterTerm(ctx *parser.TermContext) {

}

func (s *JackListener) EnterIntegerTerm(ctx *parser.IntegerTermContext) {
	s.vm.WritePush(vmwriter.CONST, asInt(ctx.GetText()))
}

// EnterKeywordTerm is called when production KeywordTerm is entered.
func (s *JackListener) EnterKeywordTerm(ctx *parser.KeywordTermContext) {
	switch ctx.GetKeywordConst().GetTokenType() {

	case parser.JackLexerTRUE:
		s.vm.WritePush(vmwriter.CONST, 0)
		s.vm.WriteArithmetic(vmwriter.NOT)

	case parser.JackLexerFALSE, parser.JackLexerNULL:
		s.vm.WritePush(vmwriter.CONST, 0)

	case parser.JackLexerTHIS:
		s.vm.WritePush(vmwriter.POINTER, 0)

	}

}

// EnterStringTerm is called when production StringTerm is entered.
func (s *JackListener) EnterStringTerm(ctx *parser.StringTermContext) {
	str := ctx.GetText()
	s.vm.WritePush(vmwriter.CONST, len(str))
	s.vm.WriteCall("String.new", 1)
	for i := 0; i < len(str); i++ {
		s.vm.WritePush(vmwriter.CONST, int(str[i]))
		s.vm.WriteCall("String.appendChar", 2)
	}
}

// EnterVarnameTerm is called when production VarnameTerm is entered.
func (s *JackListener) EnterVarnameTerm(ctx *parser.VarnameTermContext) {
	varname := ctx.GetText()
	sym := s.st.Resolve(varname)
	s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
}

func (s *JackListener) ExitArrayTerm(ctx *parser.ArrayTermContext) {

	varname := ctx.GetIdent().GetText()
	fmt.Println("v=", varname)

	sym := s.st.Resolve(varname)

	s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
	s.vm.WriteArithmetic(vmwriter.ADD)

	s.vm.WritePop(vmwriter.POINTER, 1)
	s.vm.WritePush(vmwriter.THAT, 0)

}

// ExitUnaryopTerm is called when production unaryopTerm is exited.
func (s *JackListener) ExitUnaryopTerm(ctx *parser.UnaryopTermContext) {
	switch ctx.GetUnaryop().GetTokenType() {
	case parser.JackLexerMINUS:
		s.vm.WriteArithmetic(vmwriter.NEG)
	default:
		s.vm.WriteArithmetic(vmwriter.NOT)

	}
}

func (s *JackListener) ExitBinaryoperation(ctx *parser.BinaryoperationContext) {
	switch ctx.GetOperator().GetTokenType() {

	case parser.JackLexerPLUS:
		s.vm.WriteArithmetic(vmwriter.ADD)
	case parser.JackLexerMINUS:
		s.vm.WriteArithmetic(vmwriter.SUB)
	case parser.JackLexerASTERISK:
		s.vm.WriteCall("Math.multiply", 2)
	case parser.JackLexerSLASH:
		s.vm.WriteCall("Math.divide", 2)
	case parser.JackLexerAND:
		s.vm.WriteArithmetic(vmwriter.AND)
	case parser.JackLexerOR:
		s.vm.WriteArithmetic(vmwriter.OR)
	case parser.JackLexerLT:
		s.vm.WriteArithmetic(vmwriter.LT)
	case parser.JackLexerGT:
		s.vm.WriteArithmetic(vmwriter.GT)
	case parser.JackLexerEQ:
		s.vm.WriteArithmetic(vmwriter.EQ)
	case parser.JackLexerNOT:
		s.vm.WriteArithmetic(vmwriter.NOT)
	}

}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func asInt(s string) int {
	i1, err := strconv.Atoi(s)
	check(err)
	return i1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scopeToSegment(scope symboltable.SymbolScope) vmwriter.Segment {
	switch scope {
	case symboltable.STATIC:
		return vmwriter.STATIC
	case symboltable.FIELD:
		return vmwriter.THIS
	case symboltable.VAR:
		return vmwriter.LOCAL
	case symboltable.ARG:
		return vmwriter.ARG
	default:

	}
	panic("scope undefined")
}
