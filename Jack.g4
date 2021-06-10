grammar Jack;

// Rules
classdef: CLASS ID LBRACE classvardec* RBRACE EOF;

classvardec: (STATIC | FIELD) ID SEMICOLON;

// Symbols

// operators
ASSIGN: '=';

PLUS: '+';
MINUS: '-';
ASTERISK: '*';
SLASH: '/';
AND: '&';
OR: '|';
NOT: '~';
LT: '<';
GT: '>';
EQ: '==';

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