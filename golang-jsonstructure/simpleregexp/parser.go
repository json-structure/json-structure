package simpleregexp // SimpleRegexp

// Generated from SimpleRegexp.g4 by ANTLR 4.6.

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 21, 146, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 18,
	4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2, 5, 2, 46,
	10, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3, 2, 5, 2, 55,
	10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 61, 10, 3, 3, 3, 6, 3, 64, 10, 3,
	13, 3, 14, 3, 65, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 72, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 80, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	5, 8, 87, 10, 8, 3, 8, 5, 8, 90, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 101, 10, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 6, 15, 113, 10, 15, 13, 15,
	14, 15, 114, 3, 16, 3, 16, 3, 16, 6, 16, 120, 10, 16, 13, 16, 14, 16, 121,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5,
	19, 134, 10, 19, 3, 19, 5, 19, 137, 10, 19, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 2, 2, 23, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 2, 7, 3, 2, 13, 17, 3, 2, 4,
	6, 3, 2, 3, 16, 4, 2, 3, 3, 17, 21, 4, 2, 3, 13, 18, 21, 142, 2, 45, 3,
	2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2, 10, 79,
	3, 2, 2, 2, 12, 81, 3, 2, 2, 2, 14, 86, 3, 2, 2, 2, 16, 91, 3, 2, 2, 2,
	18, 93, 3, 2, 2, 2, 20, 100, 3, 2, 2, 2, 22, 102, 3, 2, 2, 2, 24, 104,
	3, 2, 2, 2, 26, 108, 3, 2, 2, 2, 28, 112, 3, 2, 2, 2, 30, 119, 3, 2, 2,
	2, 32, 123, 3, 2, 2, 2, 34, 127, 3, 2, 2, 2, 36, 133, 3, 2, 2, 2, 38, 138,
	3, 2, 2, 2, 40, 141, 3, 2, 2, 2, 42, 143, 3, 2, 2, 2, 44, 46, 7, 16, 2,
	2, 45, 44, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 50, 3, 2, 2, 2, 47, 49,
	5, 36, 19, 2, 48, 47, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2,
	2, 50, 51, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 55,
	7, 12, 2, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 57, 7, 2, 2, 3, 57, 3, 3, 2, 2, 2, 58, 60, 7, 14, 2, 2, 59, 61, 7,
	16, 2, 2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 63, 3, 2, 2, 2, 62,
	64, 5, 6, 4, 2, 63, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2,
	2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 7, 15, 2, 2, 68, 5,
	3, 2, 2, 2, 69, 72, 5, 10, 6, 2, 70, 72, 5, 8, 5, 2, 71, 69, 3, 2, 2, 2,
	71, 70, 3, 2, 2, 2, 72, 7, 3, 2, 2, 2, 73, 74, 5, 10, 6, 2, 74, 75, 7,
	17, 2, 2, 75, 76, 5, 10, 6, 2, 76, 9, 3, 2, 2, 2, 77, 80, 5, 42, 22, 2,
	78, 80, 5, 12, 7, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 11, 3,
	2, 2, 2, 81, 82, 7, 13, 2, 2, 82, 83, 9, 2, 2, 2, 83, 13, 3, 2, 2, 2, 84,
	87, 5, 16, 9, 2, 85, 87, 5, 18, 10, 2, 86, 84, 3, 2, 2, 2, 86, 85, 3, 2,
	2, 2, 87, 89, 3, 2, 2, 2, 88, 90, 7, 6, 2, 2, 89, 88, 3, 2, 2, 2, 89, 90,
	3, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 92, 9, 3, 2, 2, 92, 17, 3, 2, 2, 2,
	93, 94, 7, 10, 2, 2, 94, 95, 5, 20, 11, 2, 95, 96, 7, 11, 2, 2, 96, 19,
	3, 2, 2, 2, 97, 101, 5, 22, 12, 2, 98, 101, 5, 24, 13, 2, 99, 101, 5, 26,
	14, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2, 2, 2, 101,
	21, 3, 2, 2, 2, 102, 103, 5, 28, 15, 2, 103, 23, 3, 2, 2, 2, 104, 105,
	5, 28, 15, 2, 105, 106, 7, 18, 2, 2, 106, 107, 5, 28, 15, 2, 107, 25, 3,
	2, 2, 2, 108, 109, 5, 28, 15, 2, 109, 110, 7, 18, 2, 2, 110, 27, 3, 2,
	2, 2, 111, 113, 7, 19, 2, 2, 112, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2,
	114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 29, 3, 2, 2, 2, 116, 120,
	5, 40, 21, 2, 117, 120, 5, 38, 20, 2, 118, 120, 5, 4, 3, 2, 119, 116, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2,
	2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 31, 3, 2, 2, 2, 123,
	124, 7, 7, 2, 2, 124, 125, 5, 36, 19, 2, 125, 126, 7, 8, 2, 2, 126, 33,
	3, 2, 2, 2, 127, 128, 7, 9, 2, 2, 128, 129, 5, 36, 19, 2, 129, 35, 3, 2,
	2, 2, 130, 134, 5, 32, 17, 2, 131, 134, 5, 34, 18, 2, 132, 134, 5, 30,
	16, 2, 133, 130, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2,
	134, 136, 3, 2, 2, 2, 135, 137, 5, 14, 8, 2, 136, 135, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 37, 3, 2, 2, 2, 138, 139, 7, 13, 2, 2, 139, 140,
	9, 4, 2, 2, 140, 39, 3, 2, 2, 2, 141, 142, 9, 5, 2, 2, 142, 41, 3, 2, 2,
	2, 143, 144, 9, 6, 2, 2, 144, 43, 3, 2, 2, 2, 17, 45, 50, 54, 60, 65, 71,
	79, 86, 89, 100, 114, 119, 121, 133, 136,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'+'", "'*'", "'?'", "'('", "')'", "'|'", "'{'", "'}'", "'$'",
	"'\\'", "'['", "']'", "'^'", "'-'", "','", "", "", "' '",
}

var symbolicNames = []string{
	"", "Dot", "Plus", "Star", "Question", "LeftParen", "RightParen", "Pipe",
	"LeftBrace", "RightBrace", "Dollar", "Slash", "LeftBracket", "RightBracket",
	"Caret", "Dash", "Comma", "Digit", "Letter", "Space",
}

var ruleNames = []string{
	"start", "characterClass", "characterClassExpression", "characterClassRange",
	"characterClassAtom", "characterClassEscaped", "quantifier", "simpleQuantifier",
	"rangeQuantifier", "rangeQuantifierExpression", "rangeQuantifierExact",
	"rangeQuantifierMinMax", "rangeQuantifierMin", "number", "term", "parenExpression",
	"orExpression", "expression", "atomEscaped", "atom", "characterClassSimple",
}

type SimpleRegexpParser struct {
	*antlr.BaseParser
}

func NewSimpleRegexpParser(input antlr.TokenStream) *SimpleRegexpParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(SimpleRegexpParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SimpleRegexp.g4"

	return this
}

// SimpleRegexpParser tokens.
const (
	SimpleRegexpParserEOF          = antlr.TokenEOF
	SimpleRegexpParserDot          = 1
	SimpleRegexpParserPlus         = 2
	SimpleRegexpParserStar         = 3
	SimpleRegexpParserQuestion     = 4
	SimpleRegexpParserLeftParen    = 5
	SimpleRegexpParserRightParen   = 6
	SimpleRegexpParserPipe         = 7
	SimpleRegexpParserLeftBrace    = 8
	SimpleRegexpParserRightBrace   = 9
	SimpleRegexpParserDollar       = 10
	SimpleRegexpParserSlash        = 11
	SimpleRegexpParserLeftBracket  = 12
	SimpleRegexpParserRightBracket = 13
	SimpleRegexpParserCaret        = 14
	SimpleRegexpParserDash         = 15
	SimpleRegexpParserComma        = 16
	SimpleRegexpParserDigit        = 17
	SimpleRegexpParserLetter       = 18
	SimpleRegexpParserSpace        = 19
)

// SimpleRegexpParser rules.
const (
	SimpleRegexpParserRULE_start                     = 0
	SimpleRegexpParserRULE_characterClass            = 1
	SimpleRegexpParserRULE_characterClassExpression  = 2
	SimpleRegexpParserRULE_characterClassRange       = 3
	SimpleRegexpParserRULE_characterClassAtom        = 4
	SimpleRegexpParserRULE_characterClassEscaped     = 5
	SimpleRegexpParserRULE_quantifier                = 6
	SimpleRegexpParserRULE_simpleQuantifier          = 7
	SimpleRegexpParserRULE_rangeQuantifier           = 8
	SimpleRegexpParserRULE_rangeQuantifierExpression = 9
	SimpleRegexpParserRULE_rangeQuantifierExact      = 10
	SimpleRegexpParserRULE_rangeQuantifierMinMax     = 11
	SimpleRegexpParserRULE_rangeQuantifierMin        = 12
	SimpleRegexpParserRULE_number                    = 13
	SimpleRegexpParserRULE_term                      = 14
	SimpleRegexpParserRULE_parenExpression           = 15
	SimpleRegexpParserRULE_orExpression              = 16
	SimpleRegexpParserRULE_expression                = 17
	SimpleRegexpParserRULE_atomEscaped               = 18
	SimpleRegexpParserRULE_atom                      = 19
	SimpleRegexpParserRULE_characterClassSimple      = 20
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserEOF, 0)
}

func (s *StartContext) Caret() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserCaret, 0)
}

func (s *StartContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *StartContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) Dollar() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDollar, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *SimpleRegexpParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleRegexpParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleRegexpParserCaret {
		{
			p.SetState(42)
			p.Match(SimpleRegexpParserCaret)
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserDot)|(1<<SimpleRegexpParserLeftParen)|(1<<SimpleRegexpParserPipe)|(1<<SimpleRegexpParserSlash)|(1<<SimpleRegexpParserLeftBracket)|(1<<SimpleRegexpParserDash)|(1<<SimpleRegexpParserComma)|(1<<SimpleRegexpParserDigit)|(1<<SimpleRegexpParserLetter)|(1<<SimpleRegexpParserSpace))) != 0 {
		{
			p.SetState(45)
			p.Expression()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleRegexpParserDollar {
		{
			p.SetState(51)
			p.Match(SimpleRegexpParserDollar)
		}

	}
	{
		p.SetState(54)
		p.Match(SimpleRegexpParserEOF)
	}

	return localctx
}

// ICharacterClassContext is an interface to support dynamic dispatch.
type ICharacterClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassContext differentiates from other interfaces.
	IsCharacterClassContext()
}

type CharacterClassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassContext() *CharacterClassContext {
	var p = new(CharacterClassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClass
	return p
}

func (*CharacterClassContext) IsCharacterClassContext() {}

func NewCharacterClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassContext {
	var p = new(CharacterClassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClass

	return p
}

func (s *CharacterClassContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBracket, 0)
}

func (s *CharacterClassContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBracket, 0)
}

func (s *CharacterClassContext) Caret() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserCaret, 0)
}

func (s *CharacterClassContext) AllCharacterClassExpression() []ICharacterClassExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterClassExpressionContext)(nil)).Elem())
	var tst = make([]ICharacterClassExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterClassExpressionContext)
		}
	}

	return tst
}

func (s *CharacterClassContext) CharacterClassExpression(i int) ICharacterClassExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassExpressionContext)
}

func (s *CharacterClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClass(s)
	}
}

func (s *CharacterClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClass(s)
	}
}

func (p *SimpleRegexpParser) CharacterClass() (localctx ICharacterClassContext) {
	localctx = NewCharacterClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleRegexpParserRULE_characterClass)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(SimpleRegexpParserLeftBracket)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleRegexpParserCaret {
		{
			p.SetState(57)
			p.Match(SimpleRegexpParserCaret)
		}

	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserDot)|(1<<SimpleRegexpParserPlus)|(1<<SimpleRegexpParserStar)|(1<<SimpleRegexpParserQuestion)|(1<<SimpleRegexpParserLeftParen)|(1<<SimpleRegexpParserRightParen)|(1<<SimpleRegexpParserPipe)|(1<<SimpleRegexpParserLeftBrace)|(1<<SimpleRegexpParserRightBrace)|(1<<SimpleRegexpParserDollar)|(1<<SimpleRegexpParserSlash)|(1<<SimpleRegexpParserComma)|(1<<SimpleRegexpParserDigit)|(1<<SimpleRegexpParserLetter)|(1<<SimpleRegexpParserSpace))) != 0) {
		{
			p.SetState(60)
			p.CharacterClassExpression()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(SimpleRegexpParserRightBracket)
	}

	return localctx
}

// ICharacterClassExpressionContext is an interface to support dynamic dispatch.
type ICharacterClassExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassExpressionContext differentiates from other interfaces.
	IsCharacterClassExpressionContext()
}

type CharacterClassExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassExpressionContext() *CharacterClassExpressionContext {
	var p = new(CharacterClassExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClassExpression
	return p
}

func (*CharacterClassExpressionContext) IsCharacterClassExpressionContext() {}

func NewCharacterClassExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassExpressionContext {
	var p = new(CharacterClassExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClassExpression

	return p
}

func (s *CharacterClassExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassExpressionContext) CharacterClassAtom() ICharacterClassAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassAtomContext)
}

func (s *CharacterClassExpressionContext) CharacterClassRange() ICharacterClassRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassRangeContext)
}

func (s *CharacterClassExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClassExpression(s)
	}
}

func (s *CharacterClassExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClassExpression(s)
	}
}

func (p *SimpleRegexpParser) CharacterClassExpression() (localctx ICharacterClassExpressionContext) {
	localctx = NewCharacterClassExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleRegexpParserRULE_characterClassExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.CharacterClassAtom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.CharacterClassRange()
		}

	}

	return localctx
}

// ICharacterClassRangeContext is an interface to support dynamic dispatch.
type ICharacterClassRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassRangeContext differentiates from other interfaces.
	IsCharacterClassRangeContext()
}

type CharacterClassRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassRangeContext() *CharacterClassRangeContext {
	var p = new(CharacterClassRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClassRange
	return p
}

func (*CharacterClassRangeContext) IsCharacterClassRangeContext() {}

func NewCharacterClassRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassRangeContext {
	var p = new(CharacterClassRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClassRange

	return p
}

func (s *CharacterClassRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassRangeContext) AllCharacterClassAtom() []ICharacterClassAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterClassAtomContext)(nil)).Elem())
	var tst = make([]ICharacterClassAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterClassAtomContext)
		}
	}

	return tst
}

func (s *CharacterClassRangeContext) CharacterClassAtom(i int) ICharacterClassAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassAtomContext)
}

func (s *CharacterClassRangeContext) Dash() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDash, 0)
}

func (s *CharacterClassRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClassRange(s)
	}
}

func (s *CharacterClassRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClassRange(s)
	}
}

func (p *SimpleRegexpParser) CharacterClassRange() (localctx ICharacterClassRangeContext) {
	localctx = NewCharacterClassRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleRegexpParserRULE_characterClassRange)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.CharacterClassAtom()
	}
	{
		p.SetState(72)
		p.Match(SimpleRegexpParserDash)
	}
	{
		p.SetState(73)
		p.CharacterClassAtom()
	}

	return localctx
}

// ICharacterClassAtomContext is an interface to support dynamic dispatch.
type ICharacterClassAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassAtomContext differentiates from other interfaces.
	IsCharacterClassAtomContext()
}

type CharacterClassAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassAtomContext() *CharacterClassAtomContext {
	var p = new(CharacterClassAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClassAtom
	return p
}

func (*CharacterClassAtomContext) IsCharacterClassAtomContext() {}

func NewCharacterClassAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassAtomContext {
	var p = new(CharacterClassAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClassAtom

	return p
}

func (s *CharacterClassAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassAtomContext) CharacterClassSimple() ICharacterClassSimpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassSimpleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassSimpleContext)
}

func (s *CharacterClassAtomContext) CharacterClassEscaped() ICharacterClassEscapedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassEscapedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassEscapedContext)
}

func (s *CharacterClassAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClassAtom(s)
	}
}

func (s *CharacterClassAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClassAtom(s)
	}
}

func (p *SimpleRegexpParser) CharacterClassAtom() (localctx ICharacterClassAtomContext) {
	localctx = NewCharacterClassAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleRegexpParserRULE_characterClassAtom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.CharacterClassSimple()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.CharacterClassEscaped()
		}

	}

	return localctx
}

// ICharacterClassEscapedContext is an interface to support dynamic dispatch.
type ICharacterClassEscapedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassEscapedContext differentiates from other interfaces.
	IsCharacterClassEscapedContext()
}

type CharacterClassEscapedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassEscapedContext() *CharacterClassEscapedContext {
	var p = new(CharacterClassEscapedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClassEscaped
	return p
}

func (*CharacterClassEscapedContext) IsCharacterClassEscapedContext() {}

func NewCharacterClassEscapedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassEscapedContext {
	var p = new(CharacterClassEscapedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClassEscaped

	return p
}

func (s *CharacterClassEscapedContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassEscapedContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(SimpleRegexpParserSlash)
}

func (s *CharacterClassEscapedContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserSlash, i)
}

func (s *CharacterClassEscapedContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBracket, 0)
}

func (s *CharacterClassEscapedContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBracket, 0)
}

func (s *CharacterClassEscapedContext) Caret() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserCaret, 0)
}

func (s *CharacterClassEscapedContext) Dash() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDash, 0)
}

func (s *CharacterClassEscapedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassEscapedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassEscapedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClassEscaped(s)
	}
}

func (s *CharacterClassEscapedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClassEscaped(s)
	}
}

func (p *SimpleRegexpParser) CharacterClassEscaped() (localctx ICharacterClassEscapedContext) {
	localctx = NewCharacterClassEscapedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SimpleRegexpParserRULE_characterClassEscaped)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(SimpleRegexpParserSlash)
	}
	p.SetState(80)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserSlash)|(1<<SimpleRegexpParserLeftBracket)|(1<<SimpleRegexpParserRightBracket)|(1<<SimpleRegexpParserCaret)|(1<<SimpleRegexpParserDash))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IQuantifierContext is an interface to support dynamic dispatch.
type IQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantifierContext differentiates from other interfaces.
	IsQuantifierContext()
}

type QuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantifierContext() *QuantifierContext {
	var p = new(QuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_quantifier
	return p
}

func (*QuantifierContext) IsQuantifierContext() {}

func NewQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantifierContext {
	var p = new(QuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_quantifier

	return p
}

func (s *QuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantifierContext) SimpleQuantifier() ISimpleQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleQuantifierContext)
}

func (s *QuantifierContext) RangeQuantifier() IRangeQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeQuantifierContext)
}

func (s *QuantifierContext) Question() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserQuestion, 0)
}

func (s *QuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterQuantifier(s)
	}
}

func (s *QuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitQuantifier(s)
	}
}

func (p *SimpleRegexpParser) Quantifier() (localctx IQuantifierContext) {
	localctx = NewQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleRegexpParserRULE_quantifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleRegexpParserPlus, SimpleRegexpParserStar, SimpleRegexpParserQuestion:
		{
			p.SetState(82)
			p.SimpleQuantifier()
		}

	case SimpleRegexpParserLeftBrace:
		{
			p.SetState(83)
			p.RangeQuantifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(86)
			p.Match(SimpleRegexpParserQuestion)
		}

	}

	return localctx
}

// ISimpleQuantifierContext is an interface to support dynamic dispatch.
type ISimpleQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleQuantifierContext differentiates from other interfaces.
	IsSimpleQuantifierContext()
}

type SimpleQuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleQuantifierContext() *SimpleQuantifierContext {
	var p = new(SimpleQuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_simpleQuantifier
	return p
}

func (*SimpleQuantifierContext) IsSimpleQuantifierContext() {}

func NewSimpleQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleQuantifierContext {
	var p = new(SimpleQuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_simpleQuantifier

	return p
}

func (s *SimpleQuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleQuantifierContext) Plus() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPlus, 0)
}

func (s *SimpleQuantifierContext) Star() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserStar, 0)
}

func (s *SimpleQuantifierContext) Question() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserQuestion, 0)
}

func (s *SimpleQuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleQuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleQuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterSimpleQuantifier(s)
	}
}

func (s *SimpleQuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitSimpleQuantifier(s)
	}
}

func (p *SimpleRegexpParser) SimpleQuantifier() (localctx ISimpleQuantifierContext) {
	localctx = NewSimpleQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SimpleRegexpParserRULE_simpleQuantifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserPlus)|(1<<SimpleRegexpParserStar)|(1<<SimpleRegexpParserQuestion))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IRangeQuantifierContext is an interface to support dynamic dispatch.
type IRangeQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeQuantifierContext differentiates from other interfaces.
	IsRangeQuantifierContext()
}

type RangeQuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQuantifierContext() *RangeQuantifierContext {
	var p = new(RangeQuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifier
	return p
}

func (*RangeQuantifierContext) IsRangeQuantifierContext() {}

func NewRangeQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQuantifierContext {
	var p = new(RangeQuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifier

	return p
}

func (s *RangeQuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQuantifierContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBrace, 0)
}

func (s *RangeQuantifierContext) RangeQuantifierExpression() IRangeQuantifierExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeQuantifierExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeQuantifierExpressionContext)
}

func (s *RangeQuantifierContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBrace, 0)
}

func (s *RangeQuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterRangeQuantifier(s)
	}
}

func (s *RangeQuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitRangeQuantifier(s)
	}
}

func (p *SimpleRegexpParser) RangeQuantifier() (localctx IRangeQuantifierContext) {
	localctx = NewRangeQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SimpleRegexpParserRULE_rangeQuantifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(SimpleRegexpParserLeftBrace)
	}
	{
		p.SetState(92)
		p.RangeQuantifierExpression()
	}
	{
		p.SetState(93)
		p.Match(SimpleRegexpParserRightBrace)
	}

	return localctx
}

// IRangeQuantifierExpressionContext is an interface to support dynamic dispatch.
type IRangeQuantifierExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeQuantifierExpressionContext differentiates from other interfaces.
	IsRangeQuantifierExpressionContext()
}

type RangeQuantifierExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQuantifierExpressionContext() *RangeQuantifierExpressionContext {
	var p = new(RangeQuantifierExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierExpression
	return p
}

func (*RangeQuantifierExpressionContext) IsRangeQuantifierExpressionContext() {}

func NewRangeQuantifierExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQuantifierExpressionContext {
	var p = new(RangeQuantifierExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierExpression

	return p
}

func (s *RangeQuantifierExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQuantifierExpressionContext) RangeQuantifierExact() IRangeQuantifierExactContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeQuantifierExactContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeQuantifierExactContext)
}

func (s *RangeQuantifierExpressionContext) RangeQuantifierMinMax() IRangeQuantifierMinMaxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeQuantifierMinMaxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeQuantifierMinMaxContext)
}

func (s *RangeQuantifierExpressionContext) RangeQuantifierMin() IRangeQuantifierMinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeQuantifierMinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeQuantifierMinContext)
}

func (s *RangeQuantifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQuantifierExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQuantifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterRangeQuantifierExpression(s)
	}
}

func (s *RangeQuantifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitRangeQuantifierExpression(s)
	}
}

func (p *SimpleRegexpParser) RangeQuantifierExpression() (localctx IRangeQuantifierExpressionContext) {
	localctx = NewRangeQuantifierExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SimpleRegexpParserRULE_rangeQuantifierExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.RangeQuantifierExact()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.RangeQuantifierMinMax()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.RangeQuantifierMin()
		}

	}

	return localctx
}

// IRangeQuantifierExactContext is an interface to support dynamic dispatch.
type IRangeQuantifierExactContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeQuantifierExactContext differentiates from other interfaces.
	IsRangeQuantifierExactContext()
}

type RangeQuantifierExactContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQuantifierExactContext() *RangeQuantifierExactContext {
	var p = new(RangeQuantifierExactContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierExact
	return p
}

func (*RangeQuantifierExactContext) IsRangeQuantifierExactContext() {}

func NewRangeQuantifierExactContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQuantifierExactContext {
	var p = new(RangeQuantifierExactContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierExact

	return p
}

func (s *RangeQuantifierExactContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQuantifierExactContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RangeQuantifierExactContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQuantifierExactContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQuantifierExactContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterRangeQuantifierExact(s)
	}
}

func (s *RangeQuantifierExactContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitRangeQuantifierExact(s)
	}
}

func (p *SimpleRegexpParser) RangeQuantifierExact() (localctx IRangeQuantifierExactContext) {
	localctx = NewRangeQuantifierExactContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SimpleRegexpParserRULE_rangeQuantifierExact)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Number()
	}

	return localctx
}

// IRangeQuantifierMinMaxContext is an interface to support dynamic dispatch.
type IRangeQuantifierMinMaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeQuantifierMinMaxContext differentiates from other interfaces.
	IsRangeQuantifierMinMaxContext()
}

type RangeQuantifierMinMaxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQuantifierMinMaxContext() *RangeQuantifierMinMaxContext {
	var p = new(RangeQuantifierMinMaxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierMinMax
	return p
}

func (*RangeQuantifierMinMaxContext) IsRangeQuantifierMinMaxContext() {}

func NewRangeQuantifierMinMaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQuantifierMinMaxContext {
	var p = new(RangeQuantifierMinMaxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierMinMax

	return p
}

func (s *RangeQuantifierMinMaxContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQuantifierMinMaxContext) AllNumber() []INumberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberContext)(nil)).Elem())
	var tst = make([]INumberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberContext)
		}
	}

	return tst
}

func (s *RangeQuantifierMinMaxContext) Number(i int) INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RangeQuantifierMinMaxContext) Comma() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserComma, 0)
}

func (s *RangeQuantifierMinMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQuantifierMinMaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQuantifierMinMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterRangeQuantifierMinMax(s)
	}
}

func (s *RangeQuantifierMinMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitRangeQuantifierMinMax(s)
	}
}

func (p *SimpleRegexpParser) RangeQuantifierMinMax() (localctx IRangeQuantifierMinMaxContext) {
	localctx = NewRangeQuantifierMinMaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SimpleRegexpParserRULE_rangeQuantifierMinMax)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Number()
	}
	{
		p.SetState(103)
		p.Match(SimpleRegexpParserComma)
	}
	{
		p.SetState(104)
		p.Number()
	}

	return localctx
}

// IRangeQuantifierMinContext is an interface to support dynamic dispatch.
type IRangeQuantifierMinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeQuantifierMinContext differentiates from other interfaces.
	IsRangeQuantifierMinContext()
}

type RangeQuantifierMinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeQuantifierMinContext() *RangeQuantifierMinContext {
	var p = new(RangeQuantifierMinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierMin
	return p
}

func (*RangeQuantifierMinContext) IsRangeQuantifierMinContext() {}

func NewRangeQuantifierMinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeQuantifierMinContext {
	var p = new(RangeQuantifierMinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_rangeQuantifierMin

	return p
}

func (s *RangeQuantifierMinContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeQuantifierMinContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RangeQuantifierMinContext) Comma() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserComma, 0)
}

func (s *RangeQuantifierMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeQuantifierMinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeQuantifierMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterRangeQuantifierMin(s)
	}
}

func (s *RangeQuantifierMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitRangeQuantifierMin(s)
	}
}

func (p *SimpleRegexpParser) RangeQuantifierMin() (localctx IRangeQuantifierMinContext) {
	localctx = NewRangeQuantifierMinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SimpleRegexpParserRULE_rangeQuantifierMin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Number()
	}
	{
		p.SetState(107)
		p.Match(SimpleRegexpParserComma)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(SimpleRegexpParserDigit)
}

func (s *NumberContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDigit, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SimpleRegexpParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SimpleRegexpParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SimpleRegexpParserDigit {
		{
			p.SetState(109)
			p.Match(SimpleRegexpParserDigit)
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllAtom() []IAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomContext)(nil)).Elem())
	var tst = make([]IAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomContext)
		}
	}

	return tst
}

func (s *TermContext) Atom(i int) IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TermContext) AllAtomEscaped() []IAtomEscapedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtomEscapedContext)(nil)).Elem())
	var tst = make([]IAtomEscapedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtomEscapedContext)
		}
	}

	return tst
}

func (s *TermContext) AtomEscaped(i int) IAtomEscapedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomEscapedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtomEscapedContext)
}

func (s *TermContext) AllCharacterClass() []ICharacterClassContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterClassContext)(nil)).Elem())
	var tst = make([]ICharacterClassContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterClassContext)
		}
	}

	return tst
}

func (s *TermContext) CharacterClass(i int) ICharacterClassContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterClassContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterClassContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *SimpleRegexpParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SimpleRegexpParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(117)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case SimpleRegexpParserDot, SimpleRegexpParserDash, SimpleRegexpParserComma, SimpleRegexpParserDigit, SimpleRegexpParserLetter, SimpleRegexpParserSpace:
				{
					p.SetState(114)
					p.Atom()
				}

			case SimpleRegexpParserSlash:
				{
					p.SetState(115)
					p.AtomEscaped()
				}

			case SimpleRegexpParserLeftBracket:
				{
					p.SetState(116)
					p.CharacterClass()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IParenExpressionContext is an interface to support dynamic dispatch.
type IParenExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParenExpressionContext differentiates from other interfaces.
	IsParenExpressionContext()
}

type ParenExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenExpressionContext() *ParenExpressionContext {
	var p = new(ParenExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_parenExpression
	return p
}

func (*ParenExpressionContext) IsParenExpressionContext() {}

func NewParenExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_parenExpression

	return p
}

func (s *ParenExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftParen, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightParen, 0)
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (p *SimpleRegexpParser) ParenExpression() (localctx IParenExpressionContext) {
	localctx = NewParenExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SimpleRegexpParserRULE_parenExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(SimpleRegexpParserLeftParen)
	}
	{
		p.SetState(122)
		p.Expression()
	}
	{
		p.SetState(123)
		p.Match(SimpleRegexpParserRightParen)
	}

	return localctx
}

// IOrExpressionContext is an interface to support dynamic dispatch.
type IOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrExpressionContext differentiates from other interfaces.
	IsOrExpressionContext()
}

type OrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExpressionContext() *OrExpressionContext {
	var p = new(OrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_orExpression
	return p
}

func (*OrExpressionContext) IsOrExpressionContext() {}

func NewOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_orExpression

	return p
}

func (s *OrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExpressionContext) Pipe() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPipe, 0)
}

func (s *OrExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

func (p *SimpleRegexpParser) OrExpression() (localctx IOrExpressionContext) {
	localctx = NewOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SimpleRegexpParserRULE_orExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(SimpleRegexpParserPipe)
	}
	{
		p.SetState(126)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ParenExpression() IParenExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenExpressionContext)
}

func (s *ExpressionContext) OrExpression() IOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrExpressionContext)
}

func (s *ExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExpressionContext) Quantifier() IQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantifierContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SimpleRegexpParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SimpleRegexpParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleRegexpParserLeftParen:
		{
			p.SetState(128)
			p.ParenExpression()
		}

	case SimpleRegexpParserPipe:
		{
			p.SetState(129)
			p.OrExpression()
		}

	case SimpleRegexpParserDot, SimpleRegexpParserSlash, SimpleRegexpParserLeftBracket, SimpleRegexpParserDash, SimpleRegexpParserComma, SimpleRegexpParserDigit, SimpleRegexpParserLetter, SimpleRegexpParserSpace:
		{
			p.SetState(130)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(133)
			p.Quantifier()
		}

	}

	return localctx
}

// IAtomEscapedContext is an interface to support dynamic dispatch.
type IAtomEscapedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomEscapedContext differentiates from other interfaces.
	IsAtomEscapedContext()
}

type AtomEscapedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomEscapedContext() *AtomEscapedContext {
	var p = new(AtomEscapedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_atomEscaped
	return p
}

func (*AtomEscapedContext) IsAtomEscapedContext() {}

func NewAtomEscapedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomEscapedContext {
	var p = new(AtomEscapedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_atomEscaped

	return p
}

func (s *AtomEscapedContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomEscapedContext) AllSlash() []antlr.TerminalNode {
	return s.GetTokens(SimpleRegexpParserSlash)
}

func (s *AtomEscapedContext) Slash(i int) antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserSlash, i)
}

func (s *AtomEscapedContext) Dot() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDot, 0)
}

func (s *AtomEscapedContext) Plus() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPlus, 0)
}

func (s *AtomEscapedContext) Star() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserStar, 0)
}

func (s *AtomEscapedContext) Question() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserQuestion, 0)
}

func (s *AtomEscapedContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftParen, 0)
}

func (s *AtomEscapedContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightParen, 0)
}

func (s *AtomEscapedContext) Pipe() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPipe, 0)
}

func (s *AtomEscapedContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBracket, 0)
}

func (s *AtomEscapedContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBracket, 0)
}

func (s *AtomEscapedContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBrace, 0)
}

func (s *AtomEscapedContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBrace, 0)
}

func (s *AtomEscapedContext) Caret() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserCaret, 0)
}

func (s *AtomEscapedContext) Dollar() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDollar, 0)
}

func (s *AtomEscapedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomEscapedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomEscapedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterAtomEscaped(s)
	}
}

func (s *AtomEscapedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitAtomEscaped(s)
	}
}

func (p *SimpleRegexpParser) AtomEscaped() (localctx IAtomEscapedContext) {
	localctx = NewAtomEscapedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SimpleRegexpParserRULE_atomEscaped)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(SimpleRegexpParserSlash)
	}
	p.SetState(137)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserDot)|(1<<SimpleRegexpParserPlus)|(1<<SimpleRegexpParserStar)|(1<<SimpleRegexpParserQuestion)|(1<<SimpleRegexpParserLeftParen)|(1<<SimpleRegexpParserRightParen)|(1<<SimpleRegexpParserPipe)|(1<<SimpleRegexpParserLeftBrace)|(1<<SimpleRegexpParserRightBrace)|(1<<SimpleRegexpParserDollar)|(1<<SimpleRegexpParserSlash)|(1<<SimpleRegexpParserLeftBracket)|(1<<SimpleRegexpParserRightBracket)|(1<<SimpleRegexpParserCaret))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Letter() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLetter, 0)
}

func (s *AtomContext) Digit() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDigit, 0)
}

func (s *AtomContext) Space() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserSpace, 0)
}

func (s *AtomContext) Dot() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDot, 0)
}

func (s *AtomContext) Dash() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDash, 0)
}

func (s *AtomContext) Comma() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserComma, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *SimpleRegexpParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SimpleRegexpParserRULE_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserDot)|(1<<SimpleRegexpParserDash)|(1<<SimpleRegexpParserComma)|(1<<SimpleRegexpParserDigit)|(1<<SimpleRegexpParserLetter)|(1<<SimpleRegexpParserSpace))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ICharacterClassSimpleContext is an interface to support dynamic dispatch.
type ICharacterClassSimpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterClassSimpleContext differentiates from other interfaces.
	IsCharacterClassSimpleContext()
}

type CharacterClassSimpleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterClassSimpleContext() *CharacterClassSimpleContext {
	var p = new(CharacterClassSimpleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleRegexpParserRULE_characterClassSimple
	return p
}

func (*CharacterClassSimpleContext) IsCharacterClassSimpleContext() {}

func NewCharacterClassSimpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterClassSimpleContext {
	var p = new(CharacterClassSimpleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleRegexpParserRULE_characterClassSimple

	return p
}

func (s *CharacterClassSimpleContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterClassSimpleContext) Letter() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLetter, 0)
}

func (s *CharacterClassSimpleContext) Digit() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDigit, 0)
}

func (s *CharacterClassSimpleContext) Space() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserSpace, 0)
}

func (s *CharacterClassSimpleContext) Dot() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDot, 0)
}

func (s *CharacterClassSimpleContext) Plus() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPlus, 0)
}

func (s *CharacterClassSimpleContext) Star() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserStar, 0)
}

func (s *CharacterClassSimpleContext) Question() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserQuestion, 0)
}

func (s *CharacterClassSimpleContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftParen, 0)
}

func (s *CharacterClassSimpleContext) RightParen() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightParen, 0)
}

func (s *CharacterClassSimpleContext) Pipe() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserPipe, 0)
}

func (s *CharacterClassSimpleContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserLeftBrace, 0)
}

func (s *CharacterClassSimpleContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserRightBrace, 0)
}

func (s *CharacterClassSimpleContext) Dollar() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserDollar, 0)
}

func (s *CharacterClassSimpleContext) Slash() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserSlash, 0)
}

func (s *CharacterClassSimpleContext) Comma() antlr.TerminalNode {
	return s.GetToken(SimpleRegexpParserComma, 0)
}

func (s *CharacterClassSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterClassSimpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterClassSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.EnterCharacterClassSimple(s)
	}
}

func (s *CharacterClassSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleRegexpListener); ok {
		listenerT.ExitCharacterClassSimple(s)
	}
}

func (p *SimpleRegexpParser) CharacterClassSimple() (localctx ICharacterClassSimpleContext) {
	localctx = NewCharacterClassSimpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SimpleRegexpParserRULE_characterClassSimple)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SimpleRegexpParserDot)|(1<<SimpleRegexpParserPlus)|(1<<SimpleRegexpParserStar)|(1<<SimpleRegexpParserQuestion)|(1<<SimpleRegexpParserLeftParen)|(1<<SimpleRegexpParserRightParen)|(1<<SimpleRegexpParserPipe)|(1<<SimpleRegexpParserLeftBrace)|(1<<SimpleRegexpParserRightBrace)|(1<<SimpleRegexpParserDollar)|(1<<SimpleRegexpParserSlash)|(1<<SimpleRegexpParserComma)|(1<<SimpleRegexpParserDigit)|(1<<SimpleRegexpParserLetter)|(1<<SimpleRegexpParserSpace))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}
