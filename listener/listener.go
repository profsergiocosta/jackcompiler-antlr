package listener

import (
	"fmt"
	"path"
	"strings"

	"github.com/profsergiocosta/ccompiler/token"
	"github.com/profsergiocosta/jackcompiler-antlr/parser"
	"github.com/profsergiocosta/jackcompiler-antlr/symboltable"
	"github.com/profsergiocosta/jackcompiler-antlr/vmwriter"
)

type JackListener struct {
	*parser.BaseJackListener

	st *symboltable.SymbolTable

	vm *vmwriter.VMWriter
}

func New() *JackListener {

	s := &JackListener{}
	s.st = symboltable.NewSymbolTable()

	pathName := "/home/sergio/go/src/github.com/profsergiocosta/jackcompiler-antlr/teste.jack"
	s.vm = vmwriter.New(filenameWithoutExtension(pathName) + ".vm")

	return s
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
		fmt.Println(element.GetText(), ttype, scope)
	}

}

func (s *JackListener) EnterVardec(ctx *parser.VardecContext) {

	var scope symboltable.SymbolScope = symboltable.VAR

	ttype := ctx.Atype().GetText()

	varnames := ctx.AllVarname()
	for _, element := range varnames {
		s.st.Define(element.GetText(), ttype, scope)
		fmt.Println(element.GetText(), ttype, scope)
	}
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}
