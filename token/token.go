package token

// não era para precisar desse pacote

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//  Symbols
	// Operators
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	AND      = "&"
	OR       = "|"
	NOT      = "~"
	DOT      = "."
	LT       = "<"
	GT       = ">"
	EQ       = "="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Keywords
	METHOD      = "METHOD"
	STATIC      = "STATIC"
	INT         = "INT"
	BOOLEAN     = "BOOLEAN"
	TRUE        = "TRUE"
	NULL        = "NULL"
	LET         = "LET"
	IF          = "IF"
	WHILE       = "WHILE"
	CONSTRUCTOR = "CONSTRUCTOR"
	FIELD       = "FIELD"
	VAR         = "VAR"
	CHAR        = "CHAR"
	VOID        = "VOID"
	CLASS       = "CLASS"
	FALSE       = "FALSE"
	DO          = "DO"
	ELSE        = "ELSE"
	RETURN      = "RETURN"
	FUNCTION    = "FUNCTION"
	THIS        = "THIS"
)
