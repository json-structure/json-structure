package simpleregexp // SimpleRegexp

// Generated from SimpleRegexp.g4 by ANTLR 4.6.

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 2, 21, 79, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18,
	9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 2, 2,
	21, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21,
	3, 2, 4, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 78, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 3, 41, 3, 2, 2, 2, 5, 43, 3,
	2, 2, 2, 7, 45, 3, 2, 2, 2, 9, 47, 3, 2, 2, 2, 11, 49, 3, 2, 2, 2, 13,
	51, 3, 2, 2, 2, 15, 53, 3, 2, 2, 2, 17, 55, 3, 2, 2, 2, 19, 57, 3, 2, 2,
	2, 21, 59, 3, 2, 2, 2, 23, 61, 3, 2, 2, 2, 25, 63, 3, 2, 2, 2, 27, 65,
	3, 2, 2, 2, 29, 67, 3, 2, 2, 2, 31, 69, 3, 2, 2, 2, 33, 71, 3, 2, 2, 2,
	35, 73, 3, 2, 2, 2, 37, 75, 3, 2, 2, 2, 39, 77, 3, 2, 2, 2, 41, 42, 7,
	48, 2, 2, 42, 4, 3, 2, 2, 2, 43, 44, 7, 45, 2, 2, 44, 6, 3, 2, 2, 2, 45,
	46, 7, 44, 2, 2, 46, 8, 3, 2, 2, 2, 47, 48, 7, 65, 2, 2, 48, 10, 3, 2,
	2, 2, 49, 50, 7, 42, 2, 2, 50, 12, 3, 2, 2, 2, 51, 52, 7, 43, 2, 2, 52,
	14, 3, 2, 2, 2, 53, 54, 7, 126, 2, 2, 54, 16, 3, 2, 2, 2, 55, 56, 7, 125,
	2, 2, 56, 18, 3, 2, 2, 2, 57, 58, 7, 127, 2, 2, 58, 20, 3, 2, 2, 2, 59,
	60, 7, 38, 2, 2, 60, 22, 3, 2, 2, 2, 61, 62, 7, 94, 2, 2, 62, 24, 3, 2,
	2, 2, 63, 64, 7, 93, 2, 2, 64, 26, 3, 2, 2, 2, 65, 66, 7, 95, 2, 2, 66,
	28, 3, 2, 2, 2, 67, 68, 7, 96, 2, 2, 68, 30, 3, 2, 2, 2, 69, 70, 7, 47,
	2, 2, 70, 32, 3, 2, 2, 2, 71, 72, 7, 46, 2, 2, 72, 34, 3, 2, 2, 2, 73,
	74, 9, 2, 2, 2, 74, 36, 3, 2, 2, 2, 75, 76, 9, 3, 2, 2, 76, 38, 3, 2, 2,
	2, 77, 78, 7, 34, 2, 2, 78, 40, 3, 2, 2, 2, 3, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "'+'", "'*'", "'?'", "'('", "')'", "'|'", "'{'", "'}'", "'$'",
	"'\\'", "'['", "']'", "'^'", "'-'", "','", "", "", "' '",
}

var lexerSymbolicNames = []string{
	"", "Dot", "Plus", "Star", "Question", "LeftParen", "RightParen", "Pipe",
	"LeftBrace", "RightBrace", "Dollar", "Slash", "LeftBracket", "RightBracket",
	"Caret", "Dash", "Comma", "Digit", "Letter", "Space",
}

var lexerRuleNames = []string{
	"Dot", "Plus", "Star", "Question", "LeftParen", "RightParen", "Pipe", "LeftBrace",
	"RightBrace", "Dollar", "Slash", "LeftBracket", "RightBracket", "Caret",
	"Dash", "Comma", "Digit", "Letter", "Space",
}

type SimpleRegexpLexer struct {
	*antlr.BaseLexer
	modeNames []string
	// TODO: EOF string
}

func NewSimpleRegexpLexer(input antlr.CharStream) *SimpleRegexpLexer {
	var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	l := new(SimpleRegexpLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SimpleRegexp.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleRegexpLexer tokens.
const (
	SimpleRegexpLexerDot          = 1
	SimpleRegexpLexerPlus         = 2
	SimpleRegexpLexerStar         = 3
	SimpleRegexpLexerQuestion     = 4
	SimpleRegexpLexerLeftParen    = 5
	SimpleRegexpLexerRightParen   = 6
	SimpleRegexpLexerPipe         = 7
	SimpleRegexpLexerLeftBrace    = 8
	SimpleRegexpLexerRightBrace   = 9
	SimpleRegexpLexerDollar       = 10
	SimpleRegexpLexerSlash        = 11
	SimpleRegexpLexerLeftBracket  = 12
	SimpleRegexpLexerRightBracket = 13
	SimpleRegexpLexerCaret        = 14
	SimpleRegexpLexerDash         = 15
	SimpleRegexpLexerComma        = 16
	SimpleRegexpLexerDigit        = 17
	SimpleRegexpLexerLetter       = 18
	SimpleRegexpLexerSpace        = 19
)
