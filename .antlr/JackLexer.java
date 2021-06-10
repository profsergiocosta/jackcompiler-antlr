// Generated from /home/sergio/go/src/github.com/profsergiocosta/jackcompiler-antlr/Jack.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JackLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ASSIGN=1, PLUS=2, MINUS=3, ASTERISK=4, SLASH=5, AND=6, OR=7, NOT=8, LT=9, 
		GT=10, EQ=11, DOT=12, COMMA=13, SEMICOLON=14, LPAREN=15, RPAREN=16, LBRACE=17, 
		RBRACE=18, LBRACKET=19, RBRACKET=20, METHOD=21, STATIC=22, INT=23, BOOLEAN=24, 
		TRUE=25, NULL=26, LET=27, IF=28, WHILE=29, CONSTRUCTOR=30, FIELD=31, VAR=32, 
		CHAR=33, VOID=34, CLASS=35, FALSE=36, DO=37, ELSE=38, RETURN=39, FUNCTION=40, 
		THIS=41, ID=42, STRING=43, INTEGER=44, WHITESPACE=45;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT", "LT", 
			"GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE", 
			"RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN", 
			"TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", 
			"CHAR", "VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", "FUNCTION", 
			"THIS", "ID", "STRING", "INTEGER", "WHITESPACE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'~'", "'<'", 
			"'>'", "'=='", "'.'", "','", "';'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'method'", "'static'", "'int'", "'boolean'", "'true'", "'null'", 
			"'let'", "'if'", "'while'", "'constructor'", "'field'", "'var'", "'char'", 
			"'void'", "'class'", "'false'", "'do'", "'else'", "'return'", "'function'", 
			"'this'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", "AND", "OR", "NOT", 
			"LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", "RPAREN", "LBRACE", 
			"RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", "INT", "BOOLEAN", 
			"TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", "FIELD", "VAR", 
			"CHAR", "VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", "FUNCTION", 
			"THIS", "ID", "STRING", "INTEGER", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public JackLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Jack.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2/\u011e\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3"+
		"\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3"+
		"\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3"+
		"\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3"+
		"\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\36\3"+
		"\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#"+
		"\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3\'"+
		"\3\'\3(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3"+
		"+\5+\u0102\n+\3+\7+\u0105\n+\f+\16+\u0108\13+\3,\3,\7,\u010c\n,\f,\16"+
		",\u010f\13,\3,\3,\3-\6-\u0114\n-\r-\16-\u0115\3.\6.\u0119\n.\r.\16.\u011a"+
		"\3.\3.\3\u010d\2/\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/\3\2\6\4\2C\\c|\5\2"+
		"\62;C\\c|\3\2\62;\5\2\13\f\17\17\"\"\2\u0121\2\3\3\2\2\2\2\5\3\2\2\2\2"+
		"\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2"+
		"\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2"+
		"\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2"+
		"\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2"+
		"\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2"+
		"\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2"+
		"M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3"+
		"\2\2\2\2[\3\2\2\2\3]\3\2\2\2\5_\3\2\2\2\7a\3\2\2\2\tc\3\2\2\2\13e\3\2"+
		"\2\2\rg\3\2\2\2\17i\3\2\2\2\21k\3\2\2\2\23m\3\2\2\2\25o\3\2\2\2\27q\3"+
		"\2\2\2\31t\3\2\2\2\33v\3\2\2\2\35x\3\2\2\2\37z\3\2\2\2!|\3\2\2\2#~\3\2"+
		"\2\2%\u0080\3\2\2\2\'\u0082\3\2\2\2)\u0084\3\2\2\2+\u0086\3\2\2\2-\u008d"+
		"\3\2\2\2/\u0094\3\2\2\2\61\u0098\3\2\2\2\63\u00a0\3\2\2\2\65\u00a5\3\2"+
		"\2\2\67\u00aa\3\2\2\29\u00ae\3\2\2\2;\u00b1\3\2\2\2=\u00b7\3\2\2\2?\u00c3"+
		"\3\2\2\2A\u00c9\3\2\2\2C\u00cd\3\2\2\2E\u00d2\3\2\2\2G\u00d7\3\2\2\2I"+
		"\u00dd\3\2\2\2K\u00e3\3\2\2\2M\u00e6\3\2\2\2O\u00eb\3\2\2\2Q\u00f2\3\2"+
		"\2\2S\u00fb\3\2\2\2U\u0101\3\2\2\2W\u0109\3\2\2\2Y\u0113\3\2\2\2[\u0118"+
		"\3\2\2\2]^\7?\2\2^\4\3\2\2\2_`\7-\2\2`\6\3\2\2\2ab\7/\2\2b\b\3\2\2\2c"+
		"d\7,\2\2d\n\3\2\2\2ef\7\61\2\2f\f\3\2\2\2gh\7(\2\2h\16\3\2\2\2ij\7~\2"+
		"\2j\20\3\2\2\2kl\7\u0080\2\2l\22\3\2\2\2mn\7>\2\2n\24\3\2\2\2op\7@\2\2"+
		"p\26\3\2\2\2qr\7?\2\2rs\7?\2\2s\30\3\2\2\2tu\7\60\2\2u\32\3\2\2\2vw\7"+
		".\2\2w\34\3\2\2\2xy\7=\2\2y\36\3\2\2\2z{\7*\2\2{ \3\2\2\2|}\7+\2\2}\""+
		"\3\2\2\2~\177\7}\2\2\177$\3\2\2\2\u0080\u0081\7\177\2\2\u0081&\3\2\2\2"+
		"\u0082\u0083\7]\2\2\u0083(\3\2\2\2\u0084\u0085\7_\2\2\u0085*\3\2\2\2\u0086"+
		"\u0087\7o\2\2\u0087\u0088\7g\2\2\u0088\u0089\7v\2\2\u0089\u008a\7j\2\2"+
		"\u008a\u008b\7q\2\2\u008b\u008c\7f\2\2\u008c,\3\2\2\2\u008d\u008e\7u\2"+
		"\2\u008e\u008f\7v\2\2\u008f\u0090\7c\2\2\u0090\u0091\7v\2\2\u0091\u0092"+
		"\7k\2\2\u0092\u0093\7e\2\2\u0093.\3\2\2\2\u0094\u0095\7k\2\2\u0095\u0096"+
		"\7p\2\2\u0096\u0097\7v\2\2\u0097\60\3\2\2\2\u0098\u0099\7d\2\2\u0099\u009a"+
		"\7q\2\2\u009a\u009b\7q\2\2\u009b\u009c\7n\2\2\u009c\u009d\7g\2\2\u009d"+
		"\u009e\7c\2\2\u009e\u009f\7p\2\2\u009f\62\3\2\2\2\u00a0\u00a1\7v\2\2\u00a1"+
		"\u00a2\7t\2\2\u00a2\u00a3\7w\2\2\u00a3\u00a4\7g\2\2\u00a4\64\3\2\2\2\u00a5"+
		"\u00a6\7p\2\2\u00a6\u00a7\7w\2\2\u00a7\u00a8\7n\2\2\u00a8\u00a9\7n\2\2"+
		"\u00a9\66\3\2\2\2\u00aa\u00ab\7n\2\2\u00ab\u00ac\7g\2\2\u00ac\u00ad\7"+
		"v\2\2\u00ad8\3\2\2\2\u00ae\u00af\7k\2\2\u00af\u00b0\7h\2\2\u00b0:\3\2"+
		"\2\2\u00b1\u00b2\7y\2\2\u00b2\u00b3\7j\2\2\u00b3\u00b4\7k\2\2\u00b4\u00b5"+
		"\7n\2\2\u00b5\u00b6\7g\2\2\u00b6<\3\2\2\2\u00b7\u00b8\7e\2\2\u00b8\u00b9"+
		"\7q\2\2\u00b9\u00ba\7p\2\2\u00ba\u00bb\7u\2\2\u00bb\u00bc\7v\2\2\u00bc"+
		"\u00bd\7t\2\2\u00bd\u00be\7w\2\2\u00be\u00bf\7e\2\2\u00bf\u00c0\7v\2\2"+
		"\u00c0\u00c1\7q\2\2\u00c1\u00c2\7t\2\2\u00c2>\3\2\2\2\u00c3\u00c4\7h\2"+
		"\2\u00c4\u00c5\7k\2\2\u00c5\u00c6\7g\2\2\u00c6\u00c7\7n\2\2\u00c7\u00c8"+
		"\7f\2\2\u00c8@\3\2\2\2\u00c9\u00ca\7x\2\2\u00ca\u00cb\7c\2\2\u00cb\u00cc"+
		"\7t\2\2\u00ccB\3\2\2\2\u00cd\u00ce\7e\2\2\u00ce\u00cf\7j\2\2\u00cf\u00d0"+
		"\7c\2\2\u00d0\u00d1\7t\2\2\u00d1D\3\2\2\2\u00d2\u00d3\7x\2\2\u00d3\u00d4"+
		"\7q\2\2\u00d4\u00d5\7k\2\2\u00d5\u00d6\7f\2\2\u00d6F\3\2\2\2\u00d7\u00d8"+
		"\7e\2\2\u00d8\u00d9\7n\2\2\u00d9\u00da\7c\2\2\u00da\u00db\7u\2\2\u00db"+
		"\u00dc\7u\2\2\u00dcH\3\2\2\2\u00dd\u00de\7h\2\2\u00de\u00df\7c\2\2\u00df"+
		"\u00e0\7n\2\2\u00e0\u00e1\7u\2\2\u00e1\u00e2\7g\2\2\u00e2J\3\2\2\2\u00e3"+
		"\u00e4\7f\2\2\u00e4\u00e5\7q\2\2\u00e5L\3\2\2\2\u00e6\u00e7\7g\2\2\u00e7"+
		"\u00e8\7n\2\2\u00e8\u00e9\7u\2\2\u00e9\u00ea\7g\2\2\u00eaN\3\2\2\2\u00eb"+
		"\u00ec\7t\2\2\u00ec\u00ed\7g\2\2\u00ed\u00ee\7v\2\2\u00ee\u00ef\7w\2\2"+
		"\u00ef\u00f0\7t\2\2\u00f0\u00f1\7p\2\2\u00f1P\3\2\2\2\u00f2\u00f3\7h\2"+
		"\2\u00f3\u00f4\7w\2\2\u00f4\u00f5\7p\2\2\u00f5\u00f6\7e\2\2\u00f6\u00f7"+
		"\7v\2\2\u00f7\u00f8\7k\2\2\u00f8\u00f9\7q\2\2\u00f9\u00fa\7p\2\2\u00fa"+
		"R\3\2\2\2\u00fb\u00fc\7v\2\2\u00fc\u00fd\7j\2\2\u00fd\u00fe\7k\2\2\u00fe"+
		"\u00ff\7u\2\2\u00ffT\3\2\2\2\u0100\u0102\t\2\2\2\u0101\u0100\3\2\2\2\u0102"+
		"\u0106\3\2\2\2\u0103\u0105\t\3\2\2\u0104\u0103\3\2\2\2\u0105\u0108\3\2"+
		"\2\2\u0106\u0104\3\2\2\2\u0106\u0107\3\2\2\2\u0107V\3\2\2\2\u0108\u0106"+
		"\3\2\2\2\u0109\u010d\7$\2\2\u010a\u010c\13\2\2\2\u010b\u010a\3\2\2\2\u010c"+
		"\u010f\3\2\2\2\u010d\u010e\3\2\2\2\u010d\u010b\3\2\2\2\u010e\u0110\3\2"+
		"\2\2\u010f\u010d\3\2\2\2\u0110\u0111\7$\2\2\u0111X\3\2\2\2\u0112\u0114"+
		"\t\4\2\2\u0113\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115\u0113\3\2\2\2\u0115"+
		"\u0116\3\2\2\2\u0116Z\3\2\2\2\u0117\u0119\t\5\2\2\u0118\u0117\3\2\2\2"+
		"\u0119\u011a\3\2\2\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u011c"+
		"\3\2\2\2\u011c\u011d\b.\2\2\u011d\\\3\2\2\2\t\2\u0101\u0104\u0106\u010d"+
		"\u0115\u011a\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}