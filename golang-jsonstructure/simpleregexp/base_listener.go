package simpleregexp // SimpleRegexp

// Generated from SimpleRegexp.g4 by ANTLR 4.6.

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleRegexpListener is a complete listener for a parse tree produced by SimpleRegexpParser.
type BaseSimpleRegexpListener struct{}

var _ SimpleRegexpListener = &BaseSimpleRegexpListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleRegexpListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleRegexpListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleRegexpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleRegexpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSimpleRegexpListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSimpleRegexpListener) ExitStart(ctx *StartContext) {}

// EnterCharacterClass is called when production characterClass is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClass(ctx *CharacterClassContext) {}

// ExitCharacterClass is called when production characterClass is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClass(ctx *CharacterClassContext) {}

// EnterCharacterClassExpression is called when production characterClassExpression is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClassExpression(ctx *CharacterClassExpressionContext) {
}

// ExitCharacterClassExpression is called when production characterClassExpression is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClassExpression(ctx *CharacterClassExpressionContext) {
}

// EnterCharacterClassRange is called when production characterClassRange is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClassRange(ctx *CharacterClassRangeContext) {}

// ExitCharacterClassRange is called when production characterClassRange is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClassRange(ctx *CharacterClassRangeContext) {}

// EnterCharacterClassAtom is called when production characterClassAtom is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClassAtom(ctx *CharacterClassAtomContext) {}

// ExitCharacterClassAtom is called when production characterClassAtom is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClassAtom(ctx *CharacterClassAtomContext) {}

// EnterCharacterClassEscaped is called when production characterClassEscaped is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClassEscaped(ctx *CharacterClassEscapedContext) {}

// ExitCharacterClassEscaped is called when production characterClassEscaped is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClassEscaped(ctx *CharacterClassEscapedContext) {}

// EnterQuantifier is called when production quantifier is entered.
func (s *BaseSimpleRegexpListener) EnterQuantifier(ctx *QuantifierContext) {}

// ExitQuantifier is called when production quantifier is exited.
func (s *BaseSimpleRegexpListener) ExitQuantifier(ctx *QuantifierContext) {}

// EnterSimpleQuantifier is called when production simpleQuantifier is entered.
func (s *BaseSimpleRegexpListener) EnterSimpleQuantifier(ctx *SimpleQuantifierContext) {}

// ExitSimpleQuantifier is called when production simpleQuantifier is exited.
func (s *BaseSimpleRegexpListener) ExitSimpleQuantifier(ctx *SimpleQuantifierContext) {}

// EnterRangeQuantifier is called when production rangeQuantifier is entered.
func (s *BaseSimpleRegexpListener) EnterRangeQuantifier(ctx *RangeQuantifierContext) {}

// ExitRangeQuantifier is called when production rangeQuantifier is exited.
func (s *BaseSimpleRegexpListener) ExitRangeQuantifier(ctx *RangeQuantifierContext) {}

// EnterRangeQuantifierExpression is called when production rangeQuantifierExpression is entered.
func (s *BaseSimpleRegexpListener) EnterRangeQuantifierExpression(ctx *RangeQuantifierExpressionContext) {
}

// ExitRangeQuantifierExpression is called when production rangeQuantifierExpression is exited.
func (s *BaseSimpleRegexpListener) ExitRangeQuantifierExpression(ctx *RangeQuantifierExpressionContext) {
}

// EnterRangeQuantifierExact is called when production rangeQuantifierExact is entered.
func (s *BaseSimpleRegexpListener) EnterRangeQuantifierExact(ctx *RangeQuantifierExactContext) {}

// ExitRangeQuantifierExact is called when production rangeQuantifierExact is exited.
func (s *BaseSimpleRegexpListener) ExitRangeQuantifierExact(ctx *RangeQuantifierExactContext) {}

// EnterRangeQuantifierMinMax is called when production rangeQuantifierMinMax is entered.
func (s *BaseSimpleRegexpListener) EnterRangeQuantifierMinMax(ctx *RangeQuantifierMinMaxContext) {}

// ExitRangeQuantifierMinMax is called when production rangeQuantifierMinMax is exited.
func (s *BaseSimpleRegexpListener) ExitRangeQuantifierMinMax(ctx *RangeQuantifierMinMaxContext) {}

// EnterRangeQuantifierMin is called when production rangeQuantifierMin is entered.
func (s *BaseSimpleRegexpListener) EnterRangeQuantifierMin(ctx *RangeQuantifierMinContext) {}

// ExitRangeQuantifierMin is called when production rangeQuantifierMin is exited.
func (s *BaseSimpleRegexpListener) ExitRangeQuantifierMin(ctx *RangeQuantifierMinContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseSimpleRegexpListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseSimpleRegexpListener) ExitNumber(ctx *NumberContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSimpleRegexpListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSimpleRegexpListener) ExitTerm(ctx *TermContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseSimpleRegexpListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseSimpleRegexpListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseSimpleRegexpListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseSimpleRegexpListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleRegexpListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleRegexpListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAtomEscaped is called when production atomEscaped is entered.
func (s *BaseSimpleRegexpListener) EnterAtomEscaped(ctx *AtomEscapedContext) {}

// ExitAtomEscaped is called when production atomEscaped is exited.
func (s *BaseSimpleRegexpListener) ExitAtomEscaped(ctx *AtomEscapedContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSimpleRegexpListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSimpleRegexpListener) ExitAtom(ctx *AtomContext) {}

// EnterCharacterClassSimple is called when production characterClassSimple is entered.
func (s *BaseSimpleRegexpListener) EnterCharacterClassSimple(ctx *CharacterClassSimpleContext) {}

// ExitCharacterClassSimple is called when production characterClassSimple is exited.
func (s *BaseSimpleRegexpListener) ExitCharacterClassSimple(ctx *CharacterClassSimpleContext) {}
