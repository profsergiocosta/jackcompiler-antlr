grammar Jack;

// Rules
start: INTEGER | STRING EOF;

ID: ([a-z] | [A-Z]) ([a-z] | [A-Z] | [0-9])*;

STRING: '"' .*? '"';

INTEGER: [0-9]+;

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
METHOD: 'METHOD';
STATIC: 'STATIC';
INT: 'INT';
BOOLEAN: 'BOOLEAN';
TRUE: 'TRUE';
NULL: 'NULL';
LET: 'LET';
IF: 'IF';
WHILE: 'WHILE';
CONSTRUCTOR: 'CONSTRUCTOR';
FIELD: 'FIELD';
VAR: 'VAR';
CHAR: 'CHAR';
VOID: 'VOID';
CLASS: 'CLASS';
FALSE: 'FALSE';
DO: 'DO';
ELSE: 'ELSE';
RETURN: 'RETURN';
FUNCTION: 'FUNCTION';
THIS: 'THIS';

WHITESPACE: [ \r\n\t]+ -> skip;