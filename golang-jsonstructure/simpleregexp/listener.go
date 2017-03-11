package simpleregexp // SimpleRegexp

// Generated from SimpleRegexp.g4 by ANTLR 4.6.

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleRegexpListener is a complete listener for a parse tree produced by SimpleRegexpParser.
type SimpleRegexpListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCharacterClass is called when entering the characterClass production.
	EnterCharacterClass(c *CharacterClassContext)

	// EnterCharacterClassExpression is called when entering the characterClassExpression production.
	EnterCharacterClassExpression(c *CharacterClassExpressionContext)

	// EnterCharacterClassRange is called when entering the characterClassRange production.
	EnterCharacterClassRange(c *CharacterClassRangeContext)

	// EnterCharacterClassAtom is called when entering the characterClassAtom production.
	EnterCharacterClassAtom(c *CharacterClassAtomContext)

	// EnterCharacterClassEscaped is called when entering the characterClassEscaped production.
	EnterCharacterClassEscaped(c *CharacterClassEscapedContext)

	// EnterQuantifier is called when entering the quantifier production.
	EnterQuantifier(c *QuantifierContext)

	// EnterSimpleQuantifier is called when entering the simpleQuantifier production.
	EnterSimpleQuantifier(c *SimpleQuantifierContext)

	// EnterRangeQuantifier is called when entering the rangeQuantifier production.
	EnterRangeQuantifier(c *RangeQuantifierContext)

	// EnterRangeQuantifierExpression is called when entering the rangeQuantifierExpression production.
	EnterRangeQuantifierExpression(c *RangeQuantifierExpressionContext)

	// EnterRangeQuantifierExact is called when entering the rangeQuantifierExact production.
	EnterRangeQuantifierExact(c *RangeQuantifierExactContext)

	// EnterRangeQuantifierMinMax is called when entering the rangeQuantifierMinMax production.
	EnterRangeQuantifierMinMax(c *RangeQuantifierMinMaxContext)

	// EnterRangeQuantifierMin is called when entering the rangeQuantifierMin production.
	EnterRangeQuantifierMin(c *RangeQuantifierMinContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterOrExpression is called when entering the orExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAtomEscaped is called when entering the atomEscaped production.
	EnterAtomEscaped(c *AtomEscapedContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterCharacterClassSimple is called when entering the characterClassSimple production.
	EnterCharacterClassSimple(c *CharacterClassSimpleContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCharacterClass is called when exiting the characterClass production.
	ExitCharacterClass(c *CharacterClassContext)

	// ExitCharacterClassExpression is called when exiting the characterClassExpression production.
	ExitCharacterClassExpression(c *CharacterClassExpressionContext)

	// ExitCharacterClassRange is called when exiting the characterClassRange production.
	ExitCharacterClassRange(c *CharacterClassRangeContext)

	// ExitCharacterClassAtom is called when exiting the characterClassAtom production.
	ExitCharacterClassAtom(c *CharacterClassAtomContext)

	// ExitCharacterClassEscaped is called when exiting the characterClassEscaped production.
	ExitCharacterClassEscaped(c *CharacterClassEscapedContext)

	// ExitQuantifier is called when exiting the quantifier production.
	ExitQuantifier(c *QuantifierContext)

	// ExitSimpleQuantifier is called when exiting the simpleQuantifier production.
	ExitSimpleQuantifier(c *SimpleQuantifierContext)

	// ExitRangeQuantifier is called when exiting the rangeQuantifier production.
	ExitRangeQuantifier(c *RangeQuantifierContext)

	// ExitRangeQuantifierExpression is called when exiting the rangeQuantifierExpression production.
	ExitRangeQuantifierExpression(c *RangeQuantifierExpressionContext)

	// ExitRangeQuantifierExact is called when exiting the rangeQuantifierExact production.
	ExitRangeQuantifierExact(c *RangeQuantifierExactContext)

	// ExitRangeQuantifierMinMax is called when exiting the rangeQuantifierMinMax production.
	ExitRangeQuantifierMinMax(c *RangeQuantifierMinMaxContext)

	// ExitRangeQuantifierMin is called when exiting the rangeQuantifierMin production.
	ExitRangeQuantifierMin(c *RangeQuantifierMinContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitOrExpression is called when exiting the orExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAtomEscaped is called when exiting the atomEscaped production.
	ExitAtomEscaped(c *AtomEscapedContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitCharacterClassSimple is called when exiting the characterClassSimple production.
	ExitCharacterClassSimple(c *CharacterClassSimpleContext)
}
