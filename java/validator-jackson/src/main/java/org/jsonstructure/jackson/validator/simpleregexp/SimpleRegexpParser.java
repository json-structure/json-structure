package org.jsonstructure.jackson.validator.simpleregexp;

// Generated from SimpleRegexp.g4 by ANTLR 4.6
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SimpleRegexpParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.6", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Dot=1, Plus=2, Star=3, Question=4, LeftParen=5, RightParen=6, Pipe=7, 
		LeftBrace=8, RightBrace=9, Dollar=10, Slash=11, LeftBracket=12, RightBracket=13, 
		Caret=14, Dash=15, Comma=16, Digit=17, Letter=18, Space=19;
	public static final int
		RULE_start = 0, RULE_characterClass = 1, RULE_characterClassExpression = 2, 
		RULE_characterClassRange = 3, RULE_characterClassAtom = 4, RULE_characterClassEscaped = 5, 
		RULE_quantifier = 6, RULE_simpleQuantifier = 7, RULE_rangeQuantifier = 8, 
		RULE_rangeQuantifierExpression = 9, RULE_rangeQuantifierExact = 10, RULE_rangeQuantifierMinMax = 11, 
		RULE_rangeQuantifierMin = 12, RULE_number = 13, RULE_term = 14, RULE_parenExpression = 15, 
		RULE_orExpression = 16, RULE_expression = 17, RULE_atomEscaped = 18, RULE_atom = 19, 
		RULE_characterClassSimple = 20;
	public static final String[] ruleNames = {
		"start", "characterClass", "characterClassExpression", "characterClassRange", 
		"characterClassAtom", "characterClassEscaped", "quantifier", "simpleQuantifier", 
		"rangeQuantifier", "rangeQuantifierExpression", "rangeQuantifierExact", 
		"rangeQuantifierMinMax", "rangeQuantifierMin", "number", "term", "parenExpression", 
		"orExpression", "expression", "atomEscaped", "atom", "characterClassSimple"
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

	@Override
	public String getGrammarFileName() { return "SimpleRegexp.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SimpleRegexpParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class StartContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(SimpleRegexpParser.EOF, 0); }
		public TerminalNode Caret() { return getToken(SimpleRegexpParser.Caret, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode Dollar() { return getToken(SimpleRegexpParser.Dollar, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterStart(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitStart(this);
		}
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(43);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Caret) {
				{
				setState(42);
				match(Caret);
				}
			}

			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Dot) | (1L << LeftParen) | (1L << Pipe) | (1L << Slash) | (1L << LeftBracket) | (1L << Dash) | (1L << Comma) | (1L << Digit) | (1L << Letter) | (1L << Space))) != 0)) {
				{
				{
				setState(45);
				expression();
				}
				}
				setState(50);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Dollar) {
				{
				setState(51);
				match(Dollar);
				}
			}

			setState(54);
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

	public static class CharacterClassContext extends ParserRuleContext {
		public TerminalNode LeftBracket() { return getToken(SimpleRegexpParser.LeftBracket, 0); }
		public TerminalNode RightBracket() { return getToken(SimpleRegexpParser.RightBracket, 0); }
		public TerminalNode Caret() { return getToken(SimpleRegexpParser.Caret, 0); }
		public List<CharacterClassExpressionContext> characterClassExpression() {
			return getRuleContexts(CharacterClassExpressionContext.class);
		}
		public CharacterClassExpressionContext characterClassExpression(int i) {
			return getRuleContext(CharacterClassExpressionContext.class,i);
		}
		public CharacterClassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClass; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClass(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClass(this);
		}
	}

	public final CharacterClassContext characterClass() throws RecognitionException {
		CharacterClassContext _localctx = new CharacterClassContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_characterClass);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(LeftBracket);
			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Caret) {
				{
				setState(57);
				match(Caret);
				}
			}

			setState(61); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(60);
				characterClassExpression();
				}
				}
				setState(63); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Dot) | (1L << Plus) | (1L << Star) | (1L << Question) | (1L << LeftParen) | (1L << RightParen) | (1L << Pipe) | (1L << LeftBrace) | (1L << RightBrace) | (1L << Dollar) | (1L << Slash) | (1L << Comma) | (1L << Digit) | (1L << Letter) | (1L << Space))) != 0) );
			setState(65);
			match(RightBracket);
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

	public static class CharacterClassExpressionContext extends ParserRuleContext {
		public CharacterClassAtomContext characterClassAtom() {
			return getRuleContext(CharacterClassAtomContext.class,0);
		}
		public CharacterClassRangeContext characterClassRange() {
			return getRuleContext(CharacterClassRangeContext.class,0);
		}
		public CharacterClassExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClassExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClassExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClassExpression(this);
		}
	}

	public final CharacterClassExpressionContext characterClassExpression() throws RecognitionException {
		CharacterClassExpressionContext _localctx = new CharacterClassExpressionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_characterClassExpression);
		try {
			setState(69);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(67);
				characterClassAtom();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				characterClassRange();
				}
				break;
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

	public static class CharacterClassRangeContext extends ParserRuleContext {
		public List<CharacterClassAtomContext> characterClassAtom() {
			return getRuleContexts(CharacterClassAtomContext.class);
		}
		public CharacterClassAtomContext characterClassAtom(int i) {
			return getRuleContext(CharacterClassAtomContext.class,i);
		}
		public TerminalNode Dash() { return getToken(SimpleRegexpParser.Dash, 0); }
		public CharacterClassRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClassRange; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClassRange(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClassRange(this);
		}
	}

	public final CharacterClassRangeContext characterClassRange() throws RecognitionException {
		CharacterClassRangeContext _localctx = new CharacterClassRangeContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_characterClassRange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			characterClassAtom();
			setState(72);
			match(Dash);
			setState(73);
			characterClassAtom();
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

	public static class CharacterClassAtomContext extends ParserRuleContext {
		public CharacterClassSimpleContext characterClassSimple() {
			return getRuleContext(CharacterClassSimpleContext.class,0);
		}
		public CharacterClassEscapedContext characterClassEscaped() {
			return getRuleContext(CharacterClassEscapedContext.class,0);
		}
		public CharacterClassAtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClassAtom; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClassAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClassAtom(this);
		}
	}

	public final CharacterClassAtomContext characterClassAtom() throws RecognitionException {
		CharacterClassAtomContext _localctx = new CharacterClassAtomContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_characterClassAtom);
		try {
			setState(77);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				characterClassSimple();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(76);
				characterClassEscaped();
				}
				break;
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

	public static class CharacterClassEscapedContext extends ParserRuleContext {
		public List<TerminalNode> Slash() { return getTokens(SimpleRegexpParser.Slash); }
		public TerminalNode Slash(int i) {
			return getToken(SimpleRegexpParser.Slash, i);
		}
		public TerminalNode LeftBracket() { return getToken(SimpleRegexpParser.LeftBracket, 0); }
		public TerminalNode RightBracket() { return getToken(SimpleRegexpParser.RightBracket, 0); }
		public TerminalNode Caret() { return getToken(SimpleRegexpParser.Caret, 0); }
		public TerminalNode Dash() { return getToken(SimpleRegexpParser.Dash, 0); }
		public CharacterClassEscapedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClassEscaped; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClassEscaped(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClassEscaped(this);
		}
	}

	public final CharacterClassEscapedContext characterClassEscaped() throws RecognitionException {
		CharacterClassEscapedContext _localctx = new CharacterClassEscapedContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_characterClassEscaped);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(Slash);
			setState(80);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Slash) | (1L << LeftBracket) | (1L << RightBracket) | (1L << Caret) | (1L << Dash))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	public static class QuantifierContext extends ParserRuleContext {
		public SimpleQuantifierContext simpleQuantifier() {
			return getRuleContext(SimpleQuantifierContext.class,0);
		}
		public RangeQuantifierContext rangeQuantifier() {
			return getRuleContext(RangeQuantifierContext.class,0);
		}
		public TerminalNode Question() { return getToken(SimpleRegexpParser.Question, 0); }
		public QuantifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_quantifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterQuantifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitQuantifier(this);
		}
	}

	public final QuantifierContext quantifier() throws RecognitionException {
		QuantifierContext _localctx = new QuantifierContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_quantifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Plus:
			case Star:
			case Question:
				{
				setState(82);
				simpleQuantifier();
				}
				break;
			case LeftBrace:
				{
				setState(83);
				rangeQuantifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(87);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(86);
				match(Question);
				}
				break;
			}
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

	public static class SimpleQuantifierContext extends ParserRuleContext {
		public TerminalNode Plus() { return getToken(SimpleRegexpParser.Plus, 0); }
		public TerminalNode Star() { return getToken(SimpleRegexpParser.Star, 0); }
		public TerminalNode Question() { return getToken(SimpleRegexpParser.Question, 0); }
		public SimpleQuantifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleQuantifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterSimpleQuantifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitSimpleQuantifier(this);
		}
	}

	public final SimpleQuantifierContext simpleQuantifier() throws RecognitionException {
		SimpleQuantifierContext _localctx = new SimpleQuantifierContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_simpleQuantifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Plus) | (1L << Star) | (1L << Question))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	public static class RangeQuantifierContext extends ParserRuleContext {
		public TerminalNode LeftBrace() { return getToken(SimpleRegexpParser.LeftBrace, 0); }
		public RangeQuantifierExpressionContext rangeQuantifierExpression() {
			return getRuleContext(RangeQuantifierExpressionContext.class,0);
		}
		public TerminalNode RightBrace() { return getToken(SimpleRegexpParser.RightBrace, 0); }
		public RangeQuantifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeQuantifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterRangeQuantifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitRangeQuantifier(this);
		}
	}

	public final RangeQuantifierContext rangeQuantifier() throws RecognitionException {
		RangeQuantifierContext _localctx = new RangeQuantifierContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_rangeQuantifier);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			match(LeftBrace);
			setState(92);
			rangeQuantifierExpression();
			setState(93);
			match(RightBrace);
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

	public static class RangeQuantifierExpressionContext extends ParserRuleContext {
		public RangeQuantifierExactContext rangeQuantifierExact() {
			return getRuleContext(RangeQuantifierExactContext.class,0);
		}
		public RangeQuantifierMinMaxContext rangeQuantifierMinMax() {
			return getRuleContext(RangeQuantifierMinMaxContext.class,0);
		}
		public RangeQuantifierMinContext rangeQuantifierMin() {
			return getRuleContext(RangeQuantifierMinContext.class,0);
		}
		public RangeQuantifierExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeQuantifierExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterRangeQuantifierExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitRangeQuantifierExpression(this);
		}
	}

	public final RangeQuantifierExpressionContext rangeQuantifierExpression() throws RecognitionException {
		RangeQuantifierExpressionContext _localctx = new RangeQuantifierExpressionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_rangeQuantifierExpression);
		try {
			setState(98);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(95);
				rangeQuantifierExact();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				rangeQuantifierMinMax();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(97);
				rangeQuantifierMin();
				}
				break;
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

	public static class RangeQuantifierExactContext extends ParserRuleContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public RangeQuantifierExactContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeQuantifierExact; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterRangeQuantifierExact(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitRangeQuantifierExact(this);
		}
	}

	public final RangeQuantifierExactContext rangeQuantifierExact() throws RecognitionException {
		RangeQuantifierExactContext _localctx = new RangeQuantifierExactContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_rangeQuantifierExact);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			number();
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

	public static class RangeQuantifierMinMaxContext extends ParserRuleContext {
		public List<NumberContext> number() {
			return getRuleContexts(NumberContext.class);
		}
		public NumberContext number(int i) {
			return getRuleContext(NumberContext.class,i);
		}
		public TerminalNode Comma() { return getToken(SimpleRegexpParser.Comma, 0); }
		public RangeQuantifierMinMaxContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeQuantifierMinMax; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterRangeQuantifierMinMax(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitRangeQuantifierMinMax(this);
		}
	}

	public final RangeQuantifierMinMaxContext rangeQuantifierMinMax() throws RecognitionException {
		RangeQuantifierMinMaxContext _localctx = new RangeQuantifierMinMaxContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_rangeQuantifierMinMax);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			number();
			setState(103);
			match(Comma);
			setState(104);
			number();
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

	public static class RangeQuantifierMinContext extends ParserRuleContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public TerminalNode Comma() { return getToken(SimpleRegexpParser.Comma, 0); }
		public RangeQuantifierMinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rangeQuantifierMin; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterRangeQuantifierMin(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitRangeQuantifierMin(this);
		}
	}

	public final RangeQuantifierMinContext rangeQuantifierMin() throws RecognitionException {
		RangeQuantifierMinContext _localctx = new RangeQuantifierMinContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_rangeQuantifierMin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			number();
			setState(107);
			match(Comma);
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

	public static class NumberContext extends ParserRuleContext {
		public List<TerminalNode> Digit() { return getTokens(SimpleRegexpParser.Digit); }
		public TerminalNode Digit(int i) {
			return getToken(SimpleRegexpParser.Digit, i);
		}
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitNumber(this);
		}
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(109);
				match(Digit);
				}
				}
				setState(112); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Digit );
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

	public static class TermContext extends ParserRuleContext {
		public List<AtomContext> atom() {
			return getRuleContexts(AtomContext.class);
		}
		public AtomContext atom(int i) {
			return getRuleContext(AtomContext.class,i);
		}
		public List<AtomEscapedContext> atomEscaped() {
			return getRuleContexts(AtomEscapedContext.class);
		}
		public AtomEscapedContext atomEscaped(int i) {
			return getRuleContext(AtomEscapedContext.class,i);
		}
		public List<CharacterClassContext> characterClass() {
			return getRuleContexts(CharacterClassContext.class);
		}
		public CharacterClassContext characterClass(int i) {
			return getRuleContext(CharacterClassContext.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterTerm(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitTerm(this);
		}
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_term);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(117); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					setState(117);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case Dot:
					case Dash:
					case Comma:
					case Digit:
					case Letter:
					case Space:
						{
						setState(114);
						atom();
						}
						break;
					case Slash:
						{
						setState(115);
						atomEscaped();
						}
						break;
					case LeftBracket:
						{
						setState(116);
						characterClass();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(119); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
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

	public static class ParenExpressionContext extends ParserRuleContext {
		public TerminalNode LeftParen() { return getToken(SimpleRegexpParser.LeftParen, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RightParen() { return getToken(SimpleRegexpParser.RightParen, 0); }
		public ParenExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parenExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterParenExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitParenExpression(this);
		}
	}

	public final ParenExpressionContext parenExpression() throws RecognitionException {
		ParenExpressionContext _localctx = new ParenExpressionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_parenExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			match(LeftParen);
			setState(122);
			expression();
			setState(123);
			match(RightParen);
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

	public static class OrExpressionContext extends ParserRuleContext {
		public TerminalNode Pipe() { return getToken(SimpleRegexpParser.Pipe, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public OrExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_orExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterOrExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitOrExpression(this);
		}
	}

	public final OrExpressionContext orExpression() throws RecognitionException {
		OrExpressionContext _localctx = new OrExpressionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_orExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(Pipe);
			setState(126);
			expression();
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

	public static class ExpressionContext extends ParserRuleContext {
		public ParenExpressionContext parenExpression() {
			return getRuleContext(ParenExpressionContext.class,0);
		}
		public OrExpressionContext orExpression() {
			return getRuleContext(OrExpressionContext.class,0);
		}
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public QuantifierContext quantifier() {
			return getRuleContext(QuantifierContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LeftParen:
				{
				setState(128);
				parenExpression();
				}
				break;
			case Pipe:
				{
				setState(129);
				orExpression();
				}
				break;
			case Dot:
			case Slash:
			case LeftBracket:
			case Dash:
			case Comma:
			case Digit:
			case Letter:
			case Space:
				{
				setState(130);
				term();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(134);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(133);
				quantifier();
				}
				break;
			}
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

	public static class AtomEscapedContext extends ParserRuleContext {
		public List<TerminalNode> Slash() { return getTokens(SimpleRegexpParser.Slash); }
		public TerminalNode Slash(int i) {
			return getToken(SimpleRegexpParser.Slash, i);
		}
		public TerminalNode Dot() { return getToken(SimpleRegexpParser.Dot, 0); }
		public TerminalNode Plus() { return getToken(SimpleRegexpParser.Plus, 0); }
		public TerminalNode Star() { return getToken(SimpleRegexpParser.Star, 0); }
		public TerminalNode Question() { return getToken(SimpleRegexpParser.Question, 0); }
		public TerminalNode LeftParen() { return getToken(SimpleRegexpParser.LeftParen, 0); }
		public TerminalNode RightParen() { return getToken(SimpleRegexpParser.RightParen, 0); }
		public TerminalNode Pipe() { return getToken(SimpleRegexpParser.Pipe, 0); }
		public TerminalNode LeftBracket() { return getToken(SimpleRegexpParser.LeftBracket, 0); }
		public TerminalNode RightBracket() { return getToken(SimpleRegexpParser.RightBracket, 0); }
		public TerminalNode LeftBrace() { return getToken(SimpleRegexpParser.LeftBrace, 0); }
		public TerminalNode RightBrace() { return getToken(SimpleRegexpParser.RightBrace, 0); }
		public TerminalNode Caret() { return getToken(SimpleRegexpParser.Caret, 0); }
		public TerminalNode Dollar() { return getToken(SimpleRegexpParser.Dollar, 0); }
		public AtomEscapedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atomEscaped; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterAtomEscaped(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitAtomEscaped(this);
		}
	}

	public final AtomEscapedContext atomEscaped() throws RecognitionException {
		AtomEscapedContext _localctx = new AtomEscapedContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_atomEscaped);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(Slash);
			setState(137);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Dot) | (1L << Plus) | (1L << Star) | (1L << Question) | (1L << LeftParen) | (1L << RightParen) | (1L << Pipe) | (1L << LeftBrace) | (1L << RightBrace) | (1L << Dollar) | (1L << Slash) | (1L << LeftBracket) | (1L << RightBracket) | (1L << Caret))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	public static class AtomContext extends ParserRuleContext {
		public TerminalNode Letter() { return getToken(SimpleRegexpParser.Letter, 0); }
		public TerminalNode Digit() { return getToken(SimpleRegexpParser.Digit, 0); }
		public TerminalNode Space() { return getToken(SimpleRegexpParser.Space, 0); }
		public TerminalNode Dot() { return getToken(SimpleRegexpParser.Dot, 0); }
		public TerminalNode Dash() { return getToken(SimpleRegexpParser.Dash, 0); }
		public TerminalNode Comma() { return getToken(SimpleRegexpParser.Comma, 0); }
		public AtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atom; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterAtom(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitAtom(this);
		}
	}

	public final AtomContext atom() throws RecognitionException {
		AtomContext _localctx = new AtomContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_atom);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Dot) | (1L << Dash) | (1L << Comma) | (1L << Digit) | (1L << Letter) | (1L << Space))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	public static class CharacterClassSimpleContext extends ParserRuleContext {
		public TerminalNode Letter() { return getToken(SimpleRegexpParser.Letter, 0); }
		public TerminalNode Digit() { return getToken(SimpleRegexpParser.Digit, 0); }
		public TerminalNode Space() { return getToken(SimpleRegexpParser.Space, 0); }
		public TerminalNode Dot() { return getToken(SimpleRegexpParser.Dot, 0); }
		public TerminalNode Plus() { return getToken(SimpleRegexpParser.Plus, 0); }
		public TerminalNode Star() { return getToken(SimpleRegexpParser.Star, 0); }
		public TerminalNode Question() { return getToken(SimpleRegexpParser.Question, 0); }
		public TerminalNode LeftParen() { return getToken(SimpleRegexpParser.LeftParen, 0); }
		public TerminalNode RightParen() { return getToken(SimpleRegexpParser.RightParen, 0); }
		public TerminalNode Pipe() { return getToken(SimpleRegexpParser.Pipe, 0); }
		public TerminalNode LeftBrace() { return getToken(SimpleRegexpParser.LeftBrace, 0); }
		public TerminalNode RightBrace() { return getToken(SimpleRegexpParser.RightBrace, 0); }
		public TerminalNode Dollar() { return getToken(SimpleRegexpParser.Dollar, 0); }
		public TerminalNode Slash() { return getToken(SimpleRegexpParser.Slash, 0); }
		public TerminalNode Comma() { return getToken(SimpleRegexpParser.Comma, 0); }
		public CharacterClassSimpleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_characterClassSimple; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).enterCharacterClassSimple(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SimpleRegexpListener ) ((SimpleRegexpListener)listener).exitCharacterClassSimple(this);
		}
	}

	public final CharacterClassSimpleContext characterClassSimple() throws RecognitionException {
		CharacterClassSimpleContext _localctx = new CharacterClassSimpleContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_characterClassSimple);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Dot) | (1L << Plus) | (1L << Star) | (1L << Question) | (1L << LeftParen) | (1L << RightParen) | (1L << Pipe) | (1L << LeftBrace) | (1L << RightBrace) | (1L << Dollar) | (1L << Slash) | (1L << Comma) | (1L << Digit) | (1L << Letter) | (1L << Space))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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
		"\3\u0430\ud6d1\u8206\uad2d\u4417\uaef1\u8d80\uaadd\3\25\u0092\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\5\2.\n\2\3\2\7\2\61\n\2\f"+
		"\2\16\2\64\13\2\3\2\5\2\67\n\2\3\2\3\2\3\3\3\3\5\3=\n\3\3\3\6\3@\n\3\r"+
		"\3\16\3A\3\3\3\3\3\4\3\4\5\4H\n\4\3\5\3\5\3\5\3\5\3\6\3\6\5\6P\n\6\3\7"+
		"\3\7\3\7\3\b\3\b\5\bW\n\b\3\b\5\bZ\n\b\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3"+
		"\13\3\13\5\13e\n\13\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\17\6\17q"+
		"\n\17\r\17\16\17r\3\20\3\20\3\20\6\20x\n\20\r\20\16\20y\3\21\3\21\3\21"+
		"\3\21\3\22\3\22\3\22\3\23\3\23\3\23\5\23\u0086\n\23\3\23\5\23\u0089\n"+
		"\23\3\24\3\24\3\24\3\25\3\25\3\26\3\26\3\26\2\2\27\2\4\6\b\n\f\16\20\22"+
		"\24\26\30\32\34\36 \"$&(*\2\7\3\2\r\21\3\2\4\6\3\2\3\20\4\2\3\3\21\25"+
		"\4\2\3\r\22\25\u008e\2-\3\2\2\2\4:\3\2\2\2\6G\3\2\2\2\bI\3\2\2\2\nO\3"+
		"\2\2\2\fQ\3\2\2\2\16V\3\2\2\2\20[\3\2\2\2\22]\3\2\2\2\24d\3\2\2\2\26f"+
		"\3\2\2\2\30h\3\2\2\2\32l\3\2\2\2\34p\3\2\2\2\36w\3\2\2\2 {\3\2\2\2\"\177"+
		"\3\2\2\2$\u0085\3\2\2\2&\u008a\3\2\2\2(\u008d\3\2\2\2*\u008f\3\2\2\2,"+
		".\7\20\2\2-,\3\2\2\2-.\3\2\2\2.\62\3\2\2\2/\61\5$\23\2\60/\3\2\2\2\61"+
		"\64\3\2\2\2\62\60\3\2\2\2\62\63\3\2\2\2\63\66\3\2\2\2\64\62\3\2\2\2\65"+
		"\67\7\f\2\2\66\65\3\2\2\2\66\67\3\2\2\2\678\3\2\2\289\7\2\2\39\3\3\2\2"+
		"\2:<\7\16\2\2;=\7\20\2\2<;\3\2\2\2<=\3\2\2\2=?\3\2\2\2>@\5\6\4\2?>\3\2"+
		"\2\2@A\3\2\2\2A?\3\2\2\2AB\3\2\2\2BC\3\2\2\2CD\7\17\2\2D\5\3\2\2\2EH\5"+
		"\n\6\2FH\5\b\5\2GE\3\2\2\2GF\3\2\2\2H\7\3\2\2\2IJ\5\n\6\2JK\7\21\2\2K"+
		"L\5\n\6\2L\t\3\2\2\2MP\5*\26\2NP\5\f\7\2OM\3\2\2\2ON\3\2\2\2P\13\3\2\2"+
		"\2QR\7\r\2\2RS\t\2\2\2S\r\3\2\2\2TW\5\20\t\2UW\5\22\n\2VT\3\2\2\2VU\3"+
		"\2\2\2WY\3\2\2\2XZ\7\6\2\2YX\3\2\2\2YZ\3\2\2\2Z\17\3\2\2\2[\\\t\3\2\2"+
		"\\\21\3\2\2\2]^\7\n\2\2^_\5\24\13\2_`\7\13\2\2`\23\3\2\2\2ae\5\26\f\2"+
		"be\5\30\r\2ce\5\32\16\2da\3\2\2\2db\3\2\2\2dc\3\2\2\2e\25\3\2\2\2fg\5"+
		"\34\17\2g\27\3\2\2\2hi\5\34\17\2ij\7\22\2\2jk\5\34\17\2k\31\3\2\2\2lm"+
		"\5\34\17\2mn\7\22\2\2n\33\3\2\2\2oq\7\23\2\2po\3\2\2\2qr\3\2\2\2rp\3\2"+
		"\2\2rs\3\2\2\2s\35\3\2\2\2tx\5(\25\2ux\5&\24\2vx\5\4\3\2wt\3\2\2\2wu\3"+
		"\2\2\2wv\3\2\2\2xy\3\2\2\2yw\3\2\2\2yz\3\2\2\2z\37\3\2\2\2{|\7\7\2\2|"+
		"}\5$\23\2}~\7\b\2\2~!\3\2\2\2\177\u0080\7\t\2\2\u0080\u0081\5$\23\2\u0081"+
		"#\3\2\2\2\u0082\u0086\5 \21\2\u0083\u0086\5\"\22\2\u0084\u0086\5\36\20"+
		"\2\u0085\u0082\3\2\2\2\u0085\u0083\3\2\2\2\u0085\u0084\3\2\2\2\u0086\u0088"+
		"\3\2\2\2\u0087\u0089\5\16\b\2\u0088\u0087\3\2\2\2\u0088\u0089\3\2\2\2"+
		"\u0089%\3\2\2\2\u008a\u008b\7\r\2\2\u008b\u008c\t\4\2\2\u008c\'\3\2\2"+
		"\2\u008d\u008e\t\5\2\2\u008e)\3\2\2\2\u008f\u0090\t\6\2\2\u0090+\3\2\2"+
		"\2\21-\62\66<AGOVYdrwy\u0085\u0088";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}