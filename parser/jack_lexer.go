// Generated from Jack.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 47, 286,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 5, 43,
	258, 10, 43, 3, 43, 7, 43, 261, 10, 43, 12, 43, 14, 43, 264, 11, 43, 3,
	44, 3, 44, 7, 44, 268, 10, 44, 12, 44, 14, 44, 271, 11, 44, 3, 44, 3, 44,
	3, 45, 6, 45, 276, 10, 45, 13, 45, 14, 45, 277, 3, 46, 6, 46, 281, 10,
	46, 13, 46, 14, 46, 282, 3, 46, 3, 46, 3, 269, 2, 47, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 3, 2, 6, 4, 2, 67, 92, 99,
	124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 289, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2,
	2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3,
	2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 3, 93,
	3, 2, 2, 2, 5, 95, 3, 2, 2, 2, 7, 97, 3, 2, 2, 2, 9, 99, 3, 2, 2, 2, 11,
	101, 3, 2, 2, 2, 13, 103, 3, 2, 2, 2, 15, 105, 3, 2, 2, 2, 17, 107, 3,
	2, 2, 2, 19, 109, 3, 2, 2, 2, 21, 111, 3, 2, 2, 2, 23, 113, 3, 2, 2, 2,
	25, 116, 3, 2, 2, 2, 27, 118, 3, 2, 2, 2, 29, 120, 3, 2, 2, 2, 31, 122,
	3, 2, 2, 2, 33, 124, 3, 2, 2, 2, 35, 126, 3, 2, 2, 2, 37, 128, 3, 2, 2,
	2, 39, 130, 3, 2, 2, 2, 41, 132, 3, 2, 2, 2, 43, 134, 3, 2, 2, 2, 45, 141,
	3, 2, 2, 2, 47, 148, 3, 2, 2, 2, 49, 152, 3, 2, 2, 2, 51, 160, 3, 2, 2,
	2, 53, 165, 3, 2, 2, 2, 55, 170, 3, 2, 2, 2, 57, 174, 3, 2, 2, 2, 59, 177,
	3, 2, 2, 2, 61, 183, 3, 2, 2, 2, 63, 195, 3, 2, 2, 2, 65, 201, 3, 2, 2,
	2, 67, 205, 3, 2, 2, 2, 69, 210, 3, 2, 2, 2, 71, 215, 3, 2, 2, 2, 73, 221,
	3, 2, 2, 2, 75, 227, 3, 2, 2, 2, 77, 230, 3, 2, 2, 2, 79, 235, 3, 2, 2,
	2, 81, 242, 3, 2, 2, 2, 83, 251, 3, 2, 2, 2, 85, 257, 3, 2, 2, 2, 87, 265,
	3, 2, 2, 2, 89, 275, 3, 2, 2, 2, 91, 280, 3, 2, 2, 2, 93, 94, 7, 63, 2,
	2, 94, 4, 3, 2, 2, 2, 95, 96, 7, 45, 2, 2, 96, 6, 3, 2, 2, 2, 97, 98, 7,
	47, 2, 2, 98, 8, 3, 2, 2, 2, 99, 100, 7, 44, 2, 2, 100, 10, 3, 2, 2, 2,
	101, 102, 7, 49, 2, 2, 102, 12, 3, 2, 2, 2, 103, 104, 7, 40, 2, 2, 104,
	14, 3, 2, 2, 2, 105, 106, 7, 126, 2, 2, 106, 16, 3, 2, 2, 2, 107, 108,
	7, 128, 2, 2, 108, 18, 3, 2, 2, 2, 109, 110, 7, 62, 2, 2, 110, 20, 3, 2,
	2, 2, 111, 112, 7, 64, 2, 2, 112, 22, 3, 2, 2, 2, 113, 114, 7, 63, 2, 2,
	114, 115, 7, 63, 2, 2, 115, 24, 3, 2, 2, 2, 116, 117, 7, 48, 2, 2, 117,
	26, 3, 2, 2, 2, 118, 119, 7, 46, 2, 2, 119, 28, 3, 2, 2, 2, 120, 121, 7,
	61, 2, 2, 121, 30, 3, 2, 2, 2, 122, 123, 7, 42, 2, 2, 123, 32, 3, 2, 2,
	2, 124, 125, 7, 43, 2, 2, 125, 34, 3, 2, 2, 2, 126, 127, 7, 125, 2, 2,
	127, 36, 3, 2, 2, 2, 128, 129, 7, 127, 2, 2, 129, 38, 3, 2, 2, 2, 130,
	131, 7, 93, 2, 2, 131, 40, 3, 2, 2, 2, 132, 133, 7, 95, 2, 2, 133, 42,
	3, 2, 2, 2, 134, 135, 7, 111, 2, 2, 135, 136, 7, 103, 2, 2, 136, 137, 7,
	118, 2, 2, 137, 138, 7, 106, 2, 2, 138, 139, 7, 113, 2, 2, 139, 140, 7,
	102, 2, 2, 140, 44, 3, 2, 2, 2, 141, 142, 7, 117, 2, 2, 142, 143, 7, 118,
	2, 2, 143, 144, 7, 99, 2, 2, 144, 145, 7, 118, 2, 2, 145, 146, 7, 107,
	2, 2, 146, 147, 7, 101, 2, 2, 147, 46, 3, 2, 2, 2, 148, 149, 7, 107, 2,
	2, 149, 150, 7, 112, 2, 2, 150, 151, 7, 118, 2, 2, 151, 48, 3, 2, 2, 2,
	152, 153, 7, 100, 2, 2, 153, 154, 7, 113, 2, 2, 154, 155, 7, 113, 2, 2,
	155, 156, 7, 110, 2, 2, 156, 157, 7, 103, 2, 2, 157, 158, 7, 99, 2, 2,
	158, 159, 7, 112, 2, 2, 159, 50, 3, 2, 2, 2, 160, 161, 7, 118, 2, 2, 161,
	162, 7, 116, 2, 2, 162, 163, 7, 119, 2, 2, 163, 164, 7, 103, 2, 2, 164,
	52, 3, 2, 2, 2, 165, 166, 7, 112, 2, 2, 166, 167, 7, 119, 2, 2, 167, 168,
	7, 110, 2, 2, 168, 169, 7, 110, 2, 2, 169, 54, 3, 2, 2, 2, 170, 171, 7,
	110, 2, 2, 171, 172, 7, 103, 2, 2, 172, 173, 7, 118, 2, 2, 173, 56, 3,
	2, 2, 2, 174, 175, 7, 107, 2, 2, 175, 176, 7, 104, 2, 2, 176, 58, 3, 2,
	2, 2, 177, 178, 7, 121, 2, 2, 178, 179, 7, 106, 2, 2, 179, 180, 7, 107,
	2, 2, 180, 181, 7, 110, 2, 2, 181, 182, 7, 103, 2, 2, 182, 60, 3, 2, 2,
	2, 183, 184, 7, 101, 2, 2, 184, 185, 7, 113, 2, 2, 185, 186, 7, 112, 2,
	2, 186, 187, 7, 117, 2, 2, 187, 188, 7, 118, 2, 2, 188, 189, 7, 116, 2,
	2, 189, 190, 7, 119, 2, 2, 190, 191, 7, 101, 2, 2, 191, 192, 7, 118, 2,
	2, 192, 193, 7, 113, 2, 2, 193, 194, 7, 116, 2, 2, 194, 62, 3, 2, 2, 2,
	195, 196, 7, 104, 2, 2, 196, 197, 7, 107, 2, 2, 197, 198, 7, 103, 2, 2,
	198, 199, 7, 110, 2, 2, 199, 200, 7, 102, 2, 2, 200, 64, 3, 2, 2, 2, 201,
	202, 7, 120, 2, 2, 202, 203, 7, 99, 2, 2, 203, 204, 7, 116, 2, 2, 204,
	66, 3, 2, 2, 2, 205, 206, 7, 101, 2, 2, 206, 207, 7, 106, 2, 2, 207, 208,
	7, 99, 2, 2, 208, 209, 7, 116, 2, 2, 209, 68, 3, 2, 2, 2, 210, 211, 7,
	120, 2, 2, 211, 212, 7, 113, 2, 2, 212, 213, 7, 107, 2, 2, 213, 214, 7,
	102, 2, 2, 214, 70, 3, 2, 2, 2, 215, 216, 7, 101, 2, 2, 216, 217, 7, 110,
	2, 2, 217, 218, 7, 99, 2, 2, 218, 219, 7, 117, 2, 2, 219, 220, 7, 117,
	2, 2, 220, 72, 3, 2, 2, 2, 221, 222, 7, 104, 2, 2, 222, 223, 7, 99, 2,
	2, 223, 224, 7, 110, 2, 2, 224, 225, 7, 117, 2, 2, 225, 226, 7, 103, 2,
	2, 226, 74, 3, 2, 2, 2, 227, 228, 7, 102, 2, 2, 228, 229, 7, 113, 2, 2,
	229, 76, 3, 2, 2, 2, 230, 231, 7, 103, 2, 2, 231, 232, 7, 110, 2, 2, 232,
	233, 7, 117, 2, 2, 233, 234, 7, 103, 2, 2, 234, 78, 3, 2, 2, 2, 235, 236,
	7, 116, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7, 118, 2, 2, 238, 239,
	7, 119, 2, 2, 239, 240, 7, 116, 2, 2, 240, 241, 7, 112, 2, 2, 241, 80,
	3, 2, 2, 2, 242, 243, 7, 104, 2, 2, 243, 244, 7, 119, 2, 2, 244, 245, 7,
	112, 2, 2, 245, 246, 7, 101, 2, 2, 246, 247, 7, 118, 2, 2, 247, 248, 7,
	107, 2, 2, 248, 249, 7, 113, 2, 2, 249, 250, 7, 112, 2, 2, 250, 82, 3,
	2, 2, 2, 251, 252, 7, 118, 2, 2, 252, 253, 7, 106, 2, 2, 253, 254, 7, 107,
	2, 2, 254, 255, 7, 117, 2, 2, 255, 84, 3, 2, 2, 2, 256, 258, 9, 2, 2, 2,
	257, 256, 3, 2, 2, 2, 258, 262, 3, 2, 2, 2, 259, 261, 9, 3, 2, 2, 260,
	259, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 262, 263,
	3, 2, 2, 2, 263, 86, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 269, 7, 36,
	2, 2, 266, 268, 11, 2, 2, 2, 267, 266, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2,
	269, 270, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 270, 272, 3, 2, 2, 2, 271,
	269, 3, 2, 2, 2, 272, 273, 7, 36, 2, 2, 273, 88, 3, 2, 2, 2, 274, 276,
	9, 4, 2, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275, 3, 2,
	2, 2, 277, 278, 3, 2, 2, 2, 278, 90, 3, 2, 2, 2, 279, 281, 9, 5, 2, 2,
	280, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 8, 46, 2, 2, 285, 92,
	3, 2, 2, 2, 9, 2, 257, 260, 262, 269, 277, 282, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'~'", "'<'", "'>'",
	"'=='", "'.'", "','", "';'", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'method'", "'static'", "'int'", "'boolean'", "'true'", "'null'", "'let'",
	"'if'", "'while'", "'constructor'", "'field'", "'var'", "'char'", "'void'",
	"'class'", "'false'", "'do'", "'else'", "'return'", "'function'", "'this'",
}

var lexerSymbolicNames = []string{
	"", "ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT",
	"LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN",
	"TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", "CHAR",
	"VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", "FUNCTION", "THIS", "ID",
	"STRING", "INTEGER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT", "LT",
	"GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE",
	"RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN",
	"TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", "CHAR",
	"VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", "FUNCTION", "THIS", "ID",
	"STRING", "INTEGER", "WHITESPACE",
}

type JackLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewJackLexer(input antlr.CharStream) *JackLexer {

	l := new(JackLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Jack.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JackLexer tokens.
const (
	JackLexerASSIGN      = 1
	JackLexerPLUS        = 2
	JackLexerMINUS       = 3
	JackLexerASTERISK    = 4
	JackLexerSLASH       = 5
	JackLexerAND         = 6
	JackLexerOR          = 7
	JackLexerNOT         = 8
	JackLexerLT          = 9
	JackLexerGT          = 10
	JackLexerEQ          = 11
	JackLexerDOT         = 12
	JackLexerCOMMA       = 13
	JackLexerSEMICOLON   = 14
	JackLexerLPAREN      = 15
	JackLexerRPAREN      = 16
	JackLexerLBRACE      = 17
	JackLexerRBRACE      = 18
	JackLexerLBRACKET    = 19
	JackLexerRBRACKET    = 20
	JackLexerMETHOD      = 21
	JackLexerSTATIC      = 22
	JackLexerINT         = 23
	JackLexerBOOLEAN     = 24
	JackLexerTRUE        = 25
	JackLexerNULL        = 26
	JackLexerLET         = 27
	JackLexerIF          = 28
	JackLexerWHILE       = 29
	JackLexerCONSTRUCTOR = 30
	JackLexerFIELD       = 31
	JackLexerVAR         = 32
	JackLexerCHAR        = 33
	JackLexerVOID        = 34
	JackLexerCLASS       = 35
	JackLexerFALSE       = 36
	JackLexerDO          = 37
	JackLexerELSE        = 38
	JackLexerRETURN      = 39
	JackLexerFUNCTION    = 40
	JackLexerTHIS        = 41
	JackLexerID          = 42
	JackLexerSTRING      = 43
	JackLexerINTEGER     = 44
	JackLexerWHITESPACE  = 45
)
