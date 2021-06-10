// Generated from /home/sergio/go/src/github.com/profsergiocosta/jackcompiler-antlr/Jack.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JackParser extends Parser {
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
	public static final int
		RULE_start = 0;
	private static String[] makeRuleNames() {
		return new String[] {
			"start"
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

	@Override
	public String getGrammarFileName() { return "Jack.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public JackParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public TerminalNode INTEGER() { return getToken(JackParser.INTEGER, 0); }
		public TerminalNode STRING() { return getToken(JackParser.STRING, 0); }
		public TerminalNode EOF() { return getToken(JackParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			setState(5);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTEGER:
				enterOuterAlt(_localctx, 1);
				{
				setState(2);
				match(INTEGER);
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(3);
				match(STRING);
				setState(4);
				match(EOF);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3/\n\4\2\t\2\3\2\3"+
		"\2\3\2\5\2\b\n\2\3\2\2\2\3\2\2\2\2\t\2\7\3\2\2\2\4\b\7\5\2\2\5\6\7\4\2"+
		"\2\6\b\7\2\2\3\7\4\3\2\2\2\7\5\3\2\2\2\b\3\3\2\2\2\3\7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}