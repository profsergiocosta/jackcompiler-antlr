package listener

import (
	"path"
	"strconv"
	"strings"

	"github.com/profsergiocosta/jackcompiler-antlr/parser"
	"github.com/profsergiocosta/jackcompiler-antlr/symboltable"
	"github.com/profsergiocosta/jackcompiler-antlr/token"
	"github.com/profsergiocosta/jackcompiler-antlr/vmwriter"
)

type JackListener struct {
	*parser.BaseJackListener

	st *symboltable.SymbolTable

	vm *vmwriter.VMWriter

	className    string
	functionName string
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

	scope = token.STATIC

	if ctx.GetKind().GetTokenType() == parser.JackLexerFIELD {
		scope = token.FIELD
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

func (s *JackListener) ExitVardecs(ctx *parser.VardecsContext) {
	nLocals := s.st.VarCount(symboltable.VAR)
	s.vm.WriteFunction(s.functionName, nLocals)
}

func (s *JackListener) EnterSubrotinedec(ctx *parser.SubrotinedecContext) {

	if ctx.GetKind().GetTokenType() == parser.JackLexerMETHOD {
		s.st.Define("this", s.className, symboltable.ARG)
	}

	s.functionName = s.className + "." + ctx.Subroutinename().GetText()
}

func (s *JackListener) ExitLetStatement(ctx *parser.LetStatementContext) {

	sym := s.st.Resolve(ctx.Varname().GetText())
	s.vm.WritePop(scopeToSegment(sym.Scope), sym.Index)
}

func (s *JackListener) EnterTerm(ctx *parser.TermContext) {

}

func (s *JackListener) EnterIntegerTerm(ctx *parser.IntegerTermContext) {
	s.vm.WritePush(vmwriter.CONST, asInt(ctx.GetText()))
}

// ExitBinopterm is called when production binopterm is exited.
func (s *JackListener) ExitBinopterm(ctx *parser.BinoptermContext) {
	op := ctx.GetOperator().GetText() // não conseguir pegar pelo tipo de token

	switch op {
	case token.PLUS:
		s.vm.WriteArithmetic(vmwriter.ADD)
	case token.MINUS:
		s.vm.WriteArithmetic(vmwriter.SUB)
	case token.ASTERISK:
		s.vm.WriteCall("Math.multiply", 2)
	case token.SLASH:
		s.vm.WriteCall("Math.divide", 2)
	case token.AND:
		s.vm.WriteArithmetic(vmwriter.AND)
	case token.OR:
		s.vm.WriteArithmetic(vmwriter.OR)
	case token.LT:
		s.vm.WriteArithmetic(vmwriter.LT)
	case token.GT:
		s.vm.WriteArithmetic(vmwriter.GT)
	case token.EQ:
		s.vm.WriteArithmetic(vmwriter.EQ)
	case token.NOT:
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
