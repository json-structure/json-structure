package org.jsonstructure.jackson.validator.simpleregexp;

// Generated from SimpleRegexp.g4 by ANTLR 4.6
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SimpleRegexpLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.6", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Dot=1, Plus=2, Star=3, Question=4, LeftParen=5, RightParen=6, Pipe=7, 
		LeftBrace=8, RightBrace=9, Dollar=10, Slash=11, LeftBracket=12, RightBracket=13, 
		Caret=14, Dash=15, Comma=16, Digit=17, Letter=18, Space=19;
	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"Dot", "Plus", "Star", "Question", "LeftParen", "RightParen", "Pipe", 
		"LeftBrace", "RightBrace", "Dollar", "Slash", "LeftBracket", "RightBracket", 
		"Caret", "Dash", "Comma", "Digit", "Letter", "Space"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'.'", "'+'", "'*'", "'?'", "'('", "')'", "'|'", "'{'", "'}'", "'$'", 
		"'\\'", "'['", "']'", "'^'", "'-'", "','", null, null, "' '"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "Dot", "Plus", "Star", "Question", "LeftParen", "RightParen", "Pipe", 
		"LeftBrace", "RightBrace", "Dollar", "Slash", "LeftBracket", "RightBracket", 
		"Caret", "Dash", "Comma", "Digit", "Letter", "Space"
	};
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


	public SimpleRegexpLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SimpleRegexp.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u0430\ud6d1\u8206\uad2d\u4417\uaef1\u8d80\uaadd\2\25O\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3"+
		"\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3"+
		"\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\2\2\25\3\3\5\4\7\5\t"+
		"\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25\3\2\4\3\2\62;\4\2C\\c|N\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2"+
		"\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3"+
		"\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2"+
		"\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\3)\3\2\2"+
		"\2\5+\3\2\2\2\7-\3\2\2\2\t/\3\2\2\2\13\61\3\2\2\2\r\63\3\2\2\2\17\65\3"+
		"\2\2\2\21\67\3\2\2\2\239\3\2\2\2\25;\3\2\2\2\27=\3\2\2\2\31?\3\2\2\2\33"+
		"A\3\2\2\2\35C\3\2\2\2\37E\3\2\2\2!G\3\2\2\2#I\3\2\2\2%K\3\2\2\2\'M\3\2"+
		"\2\2)*\7\60\2\2*\4\3\2\2\2+,\7-\2\2,\6\3\2\2\2-.\7,\2\2.\b\3\2\2\2/\60"+
		"\7A\2\2\60\n\3\2\2\2\61\62\7*\2\2\62\f\3\2\2\2\63\64\7+\2\2\64\16\3\2"+
		"\2\2\65\66\7~\2\2\66\20\3\2\2\2\678\7}\2\28\22\3\2\2\29:\7\177\2\2:\24"+
		"\3\2\2\2;<\7&\2\2<\26\3\2\2\2=>\7^\2\2>\30\3\2\2\2?@\7]\2\2@\32\3\2\2"+
		"\2AB\7_\2\2B\34\3\2\2\2CD\7`\2\2D\36\3\2\2\2EF\7/\2\2F \3\2\2\2GH\7.\2"+
		"\2H\"\3\2\2\2IJ\t\2\2\2J$\3\2\2\2KL\t\3\2\2L&\3\2\2\2MN\7\"\2\2N(\3\2"+
		"\2\2\3\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}