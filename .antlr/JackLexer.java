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
		ID=1, STRING=2, INTEGER=3, ASSIGN=4, PLUS=5, MINUS=6, ASTERISK=7, SLASH=8, 
		AND=9, OR=10, NOT=11, LT=12, GT=13, EQ=14, DOT=15, COMMA=16, SEMICOLON=17, 
		LPAREN=18, RPAREN=19, LBRACE=20, RBRACE=21, LBRACKET=22, RBRACKET=23, 
		METHOD=24, STATIC=25, INT=26, BOOLEAN=27, TRUE=28, NULL=29, LET=30, IF=31, 
		WHILE=32, CONSTRUCTOR=33, FIELD=34, VAR=35, CHAR=36, VOID=37, CLASS=38, 
		FALSE=39, DO=40, ELSE=41, RETURN=42, FUNCTION=43, THIS=44, WHITESPACE=45;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ID", "STRING", "INTEGER", "ASSIGN", "PLUS", "MINUS", "ASTERISK", "SLASH", 
			"AND", "OR", "NOT", "LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", "LPAREN", 
			"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "METHOD", "STATIC", 
			"INT", "BOOLEAN", "TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", 
			"FIELD", "VAR", "CHAR", "VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", 
			"FUNCTION", "THIS", "WHITESPACE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'='", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", 
			"'~'", "'<'", "'>'", "'=='", "'.'", "','", "';'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "'METHOD'", "'STATIC'", "'INT'", "'BOOLEAN'", "'TRUE'", 
			"'NULL'", "'LET'", "'IF'", "'WHILE'", "'CONSTRUCTOR'", "'FIELD'", "'VAR'", 
			"'CHAR'", "'VOID'", "'CLASS'", "'FALSE'", "'DO'", "'ELSE'", "'RETURN'", 
			"'FUNCTION'", "'THIS'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ID", "STRING", "INTEGER", "ASSIGN", "PLUS", "MINUS", "ASTERISK", 
			"SLASH", "AND", "OR", "NOT", "LT", "GT", "EQ", "DOT", "COMMA", "SEMICOLON", 
			"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "METHOD", 
			"STATIC", "INT", "BOOLEAN", "TRUE", "NULL", "LET", "IF", "WHILE", "CONSTRUCTOR", 
			"FIELD", "VAR", "CHAR", "VOID", "CLASS", "FALSE", "DO", "ELSE", "RETURN", 
			"FUNCTION", "THIS", "WHITESPACE"
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
		",\t,\4-\t-\4.\t.\3\2\5\2_\n\2\3\2\7\2b\n\2\f\2\16\2e\13\2\3\3\3\3\7\3"+
		"i\n\3\f\3\16\3l\13\3\3\3\3\3\3\4\6\4q\n\4\r\4\16\4r\3\5\3\5\3\6\3\6\3"+
		"\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17"+
		"\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25"+
		"\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3 \3 \3 \3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3"+
		"\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3"+
		"%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3(\3)\3)\3)\3"+
		"*\3*\3*\3*\3*\3+\3+\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3,\3-\3-\3"+
		"-\3-\3-\3.\6.\u0119\n.\r.\16.\u011a\3.\3.\3j\2/\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/\3\2\6\4\2C\\c|\5\2\62;C\\c|\3\2\62;\5\2\13\f\17\17\"\"\2"+
		"\u0121\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2"+
		"\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3"+
		"\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2"+
		"\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2"+
		"/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2"+
		"\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2"+
		"G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3"+
		"\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\3^\3\2\2\2\5f\3\2\2"+
		"\2\7p\3\2\2\2\tt\3\2\2\2\13v\3\2\2\2\rx\3\2\2\2\17z\3\2\2\2\21|\3\2\2"+
		"\2\23~\3\2\2\2\25\u0080\3\2\2\2\27\u0082\3\2\2\2\31\u0084\3\2\2\2\33\u0086"+
		"\3\2\2\2\35\u0088\3\2\2\2\37\u008b\3\2\2\2!\u008d\3\2\2\2#\u008f\3\2\2"+
		"\2%\u0091\3\2\2\2\'\u0093\3\2\2\2)\u0095\3\2\2\2+\u0097\3\2\2\2-\u0099"+
		"\3\2\2\2/\u009b\3\2\2\2\61\u009d\3\2\2\2\63\u00a4\3\2\2\2\65\u00ab\3\2"+
		"\2\2\67\u00af\3\2\2\29\u00b7\3\2\2\2;\u00bc\3\2\2\2=\u00c1\3\2\2\2?\u00c5"+
		"\3\2\2\2A\u00c8\3\2\2\2C\u00ce\3\2\2\2E\u00da\3\2\2\2G\u00e0\3\2\2\2I"+
		"\u00e4\3\2\2\2K\u00e9\3\2\2\2M\u00ee\3\2\2\2O\u00f4\3\2\2\2Q\u00fa\3\2"+
		"\2\2S\u00fd\3\2\2\2U\u0102\3\2\2\2W\u0109\3\2\2\2Y\u0112\3\2\2\2[\u0118"+
		"\3\2\2\2]_\t\2\2\2^]\3\2\2\2_c\3\2\2\2`b\t\3\2\2a`\3\2\2\2be\3\2\2\2c"+
		"a\3\2\2\2cd\3\2\2\2d\4\3\2\2\2ec\3\2\2\2fj\7$\2\2gi\13\2\2\2hg\3\2\2\2"+
		"il\3\2\2\2jk\3\2\2\2jh\3\2\2\2km\3\2\2\2lj\3\2\2\2mn\7$\2\2n\6\3\2\2\2"+
		"oq\t\4\2\2po\3\2\2\2qr\3\2\2\2rp\3\2\2\2rs\3\2\2\2s\b\3\2\2\2tu\7?\2\2"+
		"u\n\3\2\2\2vw\7-\2\2w\f\3\2\2\2xy\7/\2\2y\16\3\2\2\2z{\7,\2\2{\20\3\2"+
		"\2\2|}\7\61\2\2}\22\3\2\2\2~\177\7(\2\2\177\24\3\2\2\2\u0080\u0081\7~"+
		"\2\2\u0081\26\3\2\2\2\u0082\u0083\7\u0080\2\2\u0083\30\3\2\2\2\u0084\u0085"+
		"\7>\2\2\u0085\32\3\2\2\2\u0086\u0087\7@\2\2\u0087\34\3\2\2\2\u0088\u0089"+
		"\7?\2\2\u0089\u008a\7?\2\2\u008a\36\3\2\2\2\u008b\u008c\7\60\2\2\u008c"+
		" \3\2\2\2\u008d\u008e\7.\2\2\u008e\"\3\2\2\2\u008f\u0090\7=\2\2\u0090"+
		"$\3\2\2\2\u0091\u0092\7*\2\2\u0092&\3\2\2\2\u0093\u0094\7+\2\2\u0094("+
		"\3\2\2\2\u0095\u0096\7}\2\2\u0096*\3\2\2\2\u0097\u0098\7\177\2\2\u0098"+
		",\3\2\2\2\u0099\u009a\7]\2\2\u009a.\3\2\2\2\u009b\u009c\7_\2\2\u009c\60"+
		"\3\2\2\2\u009d\u009e\7O\2\2\u009e\u009f\7G\2\2\u009f\u00a0\7V\2\2\u00a0"+
		"\u00a1\7J\2\2\u00a1\u00a2\7Q\2\2\u00a2\u00a3\7F\2\2\u00a3\62\3\2\2\2\u00a4"+
		"\u00a5\7U\2\2\u00a5\u00a6\7V\2\2\u00a6\u00a7\7C\2\2\u00a7\u00a8\7V\2\2"+
		"\u00a8\u00a9\7K\2\2\u00a9\u00aa\7E\2\2\u00aa\64\3\2\2\2\u00ab\u00ac\7"+
		"K\2\2\u00ac\u00ad\7P\2\2\u00ad\u00ae\7V\2\2\u00ae\66\3\2\2\2\u00af\u00b0"+
		"\7D\2\2\u00b0\u00b1\7Q\2\2\u00b1\u00b2\7Q\2\2\u00b2\u00b3\7N\2\2\u00b3"+
		"\u00b4\7G\2\2\u00b4\u00b5\7C\2\2\u00b5\u00b6\7P\2\2\u00b68\3\2\2\2\u00b7"+
		"\u00b8\7V\2\2\u00b8\u00b9\7T\2\2\u00b9\u00ba\7W\2\2\u00ba\u00bb\7G\2\2"+
		"\u00bb:\3\2\2\2\u00bc\u00bd\7P\2\2\u00bd\u00be\7W\2\2\u00be\u00bf\7N\2"+
		"\2\u00bf\u00c0\7N\2\2\u00c0<\3\2\2\2\u00c1\u00c2\7N\2\2\u00c2\u00c3\7"+
		"G\2\2\u00c3\u00c4\7V\2\2\u00c4>\3\2\2\2\u00c5\u00c6\7K\2\2\u00c6\u00c7"+
		"\7H\2\2\u00c7@\3\2\2\2\u00c8\u00c9\7Y\2\2\u00c9\u00ca\7J\2\2\u00ca\u00cb"+
		"\7K\2\2\u00cb\u00cc\7N\2\2\u00cc\u00cd\7G\2\2\u00cdB\3\2\2\2\u00ce\u00cf"+
		"\7E\2\2\u00cf\u00d0\7Q\2\2\u00d0\u00d1\7P\2\2\u00d1\u00d2\7U\2\2\u00d2"+
		"\u00d3\7V\2\2\u00d3\u00d4\7T\2\2\u00d4\u00d5\7W\2\2\u00d5\u00d6\7E\2\2"+
		"\u00d6\u00d7\7V\2\2\u00d7\u00d8\7Q\2\2\u00d8\u00d9\7T\2\2\u00d9D\3\2\2"+
		"\2\u00da\u00db\7H\2\2\u00db\u00dc\7K\2\2\u00dc\u00dd\7G\2\2\u00dd\u00de"+
		"\7N\2\2\u00de\u00df\7F\2\2\u00dfF\3\2\2\2\u00e0\u00e1\7X\2\2\u00e1\u00e2"+
		"\7C\2\2\u00e2\u00e3\7T\2\2\u00e3H\3\2\2\2\u00e4\u00e5\7E\2\2\u00e5\u00e6"+
		"\7J\2\2\u00e6\u00e7\7C\2\2\u00e7\u00e8\7T\2\2\u00e8J\3\2\2\2\u00e9\u00ea"+
		"\7X\2\2\u00ea\u00eb\7Q\2\2\u00eb\u00ec\7K\2\2\u00ec\u00ed\7F\2\2\u00ed"+
		"L\3\2\2\2\u00ee\u00ef\7E\2\2\u00ef\u00f0\7N\2\2\u00f0\u00f1\7C\2\2\u00f1"+
		"\u00f2\7U\2\2\u00f2\u00f3\7U\2\2\u00f3N\3\2\2\2\u00f4\u00f5\7H\2\2\u00f5"+
		"\u00f6\7C\2\2\u00f6\u00f7\7N\2\2\u00f7\u00f8\7U\2\2\u00f8\u00f9\7G\2\2"+
		"\u00f9P\3\2\2\2\u00fa\u00fb\7F\2\2\u00fb\u00fc\7Q\2\2\u00fcR\3\2\2\2\u00fd"+
		"\u00fe\7G\2\2\u00fe\u00ff\7N\2\2\u00ff\u0100\7U\2\2\u0100\u0101\7G\2\2"+
		"\u0101T\3\2\2\2\u0102\u0103\7T\2\2\u0103\u0104\7G\2\2\u0104\u0105\7V\2"+
		"\2\u0105\u0106\7W\2\2\u0106\u0107\7T\2\2\u0107\u0108\7P\2\2\u0108V\3\2"+
		"\2\2\u0109\u010a\7H\2\2\u010a\u010b\7W\2\2\u010b\u010c\7P\2\2\u010c\u010d"+
		"\7E\2\2\u010d\u010e\7V\2\2\u010e\u010f\7K\2\2\u010f\u0110\7Q\2\2\u0110"+
		"\u0111\7P\2\2\u0111X\3\2\2\2\u0112\u0113\7V\2\2\u0113\u0114\7J\2\2\u0114"+
		"\u0115\7K\2\2\u0115\u0116\7U\2\2\u0116Z\3\2\2\2\u0117\u0119\t\5\2\2\u0118"+
		"\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a\u0118\3\2\2\2\u011a\u011b\3\2"+
		"\2\2\u011b\u011c\3\2\2\2\u011c\u011d\b.\2\2\u011d\\\3\2\2\2\t\2^acjr\u011a"+
		"\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}