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
		ASSIGN=1, PLUS=2, MINUS=3, ASTERISK=4, SLASH=5, AND=6, OR=7, NOT=8, LT=9, 
		GT=10, EQ=11, DOT=12, COMMA=13, SEMICOLON=14, LPAREN=15, RPAREN=16, LBRACE=17, 
		RBRACE=18, LBRACKET=19, RBRACKET=20, METHOD=21, STATIC=22, INT=23, BOOLEAN=24, 
		TRUE=25, NULL=26, LET=27, IF=28, WHILE=29, CONSTRUCTOR=30, FIELD=31, VAR=32, 
		CHAR=33, VOID=34, CLASS=35, FALSE=36, DO=37, ELSE=38, RETURN=39, FUNCTION=40, 
		THIS=41, ID=42, STRING=43, INTEGER=44, WHITESPACE=45;
	public static final int
		RULE_start = 0, RULE_classvardec = 1;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "classvardec"
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
		public TerminalNode CLASS() { return getToken(JackParser.CLASS, 0); }
		public TerminalNode ID() { return getToken(JackParser.ID, 0); }
		public TerminalNode LBRACE() { return getToken(JackParser.LBRACE, 0); }
		public ClassvardecContext classvardec() {
			return getRuleContext(ClassvardecContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(JackParser.RBRACE, 0); }
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
			enterOuterAlt(_localctx, 1);
			{
			setState(4);
			match(CLASS);
			setState(5);
			match(ID);
			setState(6);
			match(LBRACE);
			setState(7);
			classvardec();
			setState(8);
			match(RBRACE);
			setState(9);
			match(EOF);
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

	public static class ClassvardecContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(JackParser.ID, 0); }
		public TerminalNode SEMICOLON() { return getToken(JackParser.SEMICOLON, 0); }
		public TerminalNode STATIC() { return getToken(JackParser.STATIC, 0); }
		public TerminalNode FIELD() { return getToken(JackParser.FIELD, 0); }
		public ClassvardecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classvardec; }
	}

	public final ClassvardecContext classvardec() throws RecognitionException {
		ClassvardecContext _localctx = new ClassvardecContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_classvardec);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(11);
			_la = _input.LA(1);
			if ( !(_la==STATIC || _la==FIELD) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(12);
			match(ID);
			setState(13);
			match(SEMICOLON);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3/\22\4\2\t\2\4\3\t"+
		"\3\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\2\2\4\2\4\2\3\4\2\30"+
		"\30!!\2\17\2\6\3\2\2\2\4\r\3\2\2\2\6\7\7%\2\2\7\b\7,\2\2\b\t\7\23\2\2"+
		"\t\n\5\4\3\2\n\13\7\24\2\2\13\f\7\2\2\3\f\3\3\2\2\2\r\16\t\2\2\2\16\17"+
		"\7,\2\2\17\20\7\20\2\2\20\5\3\2\2\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}