grammar Jack;

// Rules
classdef:
	CLASS classname LBRACE classvardec* subrotinedec* RBRACE EOF;

classvardec:
	kind = (STATIC | FIELD) atype varname (COMMA varname)* SEMICOLON;

atype: (INT | CHAR | BOOLEAN | ID);

subrotinedec:
	kind = (CONSTRUCTOR | FUNCTION | METHOD) subroutinetype subroutinename LPAREN parameterList
		RPAREN subroutinebody;

subroutinetype: (VOID | atype);

parameterList: (atype varname (COMMA atype varname)*)?;

subroutinebody: LBRACE vardecs statement* RBRACE;

vardecs: vardec*;

vardec: VAR atype varname (COMMA varname)* SEMICOLON;

statement:
	letStatement
	| ifStatement
	| whileStatement
	| doStatement
	| returnStatement;

letStatement:
	LET varname (LBRACKET expression RBRACKET)? EQ expression SEMICOLON;
ifStatement:
	IF LPAREN expression RPAREN LBRACE statement* RBRACE (
		(ELSE LBRACE statement* RBRACE)?
	);

whileStatement:
	WHILE LPAREN expression RPAREN LBRACE statement* RBRACE;

doStatement: DO subroutinecall SEMICOLON;

returnStatement: RETURN expression? SEMICOLON;

subroutinecall: ((classname | varname) DOT)? subroutinename LPAREN expressionlist RPAREN;

expressionlist: (expression (COMMA expression)*)?;

expression: term binaryoperation*;

binaryoperation: (
		operator = (
			PLUS
			| MINUS
			| ASTERISK
			| SLASH
			| AND
			| OR
			| LT
			| GT
			| EQ
		) term
	);

term:
	INTEGER											# IntegerTerm
	| STRING										# StringTerm
	| keywordConst = (TRUE | FALSE | NULL | THIS)	# KeywordTerm
	| varname										# VarnameTerm
	| varname LBRACKET expression RBRACKET			# ArrayTerm
	| subroutinecall								# SubroutineTerm
	| LPAREN expression RPAREN						# ParsTerm
	| unaryop term									# unaryopTerm;

binop:
	PLUS
	| MINUS
	| ASTERISK
	| SLASH
	| AND
	| OR
	| LT
	| GT
	| EQ;

unaryop: MINUS | NOT;

classname: ID;
varname: ID;
subroutinename: ID;

// Symbols

// operators

PLUS: '+';
MINUS: '-';
ASTERISK: '*';
SLASH: '/';
AND: '&';
OR: '|';
NOT: '~';
LT: '<';
GT: '>';
EQ: '=';

// Delimiters
DOT: '.';
COMMA: ',';
SEMICOLON: ';';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';

// Keywords
METHOD: 'method';
STATIC: 'static';
INT: 'int';
BOOLEAN: 'boolean';
TRUE: 'true';
NULL: 'null';
LET: 'let';
IF: 'if';
WHILE: 'while';
CONSTRUCTOR: 'constructor';
FIELD: 'field';
VAR: 'var';
CHAR: 'char';
VOID: 'void';
CLASS: 'class';
FALSE: 'false';
DO: 'do';
ELSE: 'else';
RETURN: 'return';
FUNCTION: 'function';
THIS: 'this';

ID: ([a-z] | [A-Z]) ([a-z] | [A-Z] | [0-9])*;

STRING: '"' .*? '"';

INTEGER: [0-9]+;

WHITESPACE: [ \r\n\t]+ -> skip;