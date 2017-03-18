package org.jsonstructure.jackson.validator.simpleregexp;

// Generated from SimpleRegexp.g4 by ANTLR 4.6
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SimpleRegexpParser}.
 */
public interface SimpleRegexpListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#start}.
	 * @param ctx the parse tree
	 */
	void enterStart(SimpleRegexpParser.StartContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#start}.
	 * @param ctx the parse tree
	 */
	void exitStart(SimpleRegexpParser.StartContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClass}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClass(SimpleRegexpParser.CharacterClassContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClass}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClass(SimpleRegexpParser.CharacterClassContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClassExpression}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClassExpression(SimpleRegexpParser.CharacterClassExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClassExpression}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClassExpression(SimpleRegexpParser.CharacterClassExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClassRange}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClassRange(SimpleRegexpParser.CharacterClassRangeContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClassRange}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClassRange(SimpleRegexpParser.CharacterClassRangeContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClassAtom}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClassAtom(SimpleRegexpParser.CharacterClassAtomContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClassAtom}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClassAtom(SimpleRegexpParser.CharacterClassAtomContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClassEscaped}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClassEscaped(SimpleRegexpParser.CharacterClassEscapedContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClassEscaped}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClassEscaped(SimpleRegexpParser.CharacterClassEscapedContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#quantifier}.
	 * @param ctx the parse tree
	 */
	void enterQuantifier(SimpleRegexpParser.QuantifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#quantifier}.
	 * @param ctx the parse tree
	 */
	void exitQuantifier(SimpleRegexpParser.QuantifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#simpleQuantifier}.
	 * @param ctx the parse tree
	 */
	void enterSimpleQuantifier(SimpleRegexpParser.SimpleQuantifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#simpleQuantifier}.
	 * @param ctx the parse tree
	 */
	void exitSimpleQuantifier(SimpleRegexpParser.SimpleQuantifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#rangeQuantifier}.
	 * @param ctx the parse tree
	 */
	void enterRangeQuantifier(SimpleRegexpParser.RangeQuantifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#rangeQuantifier}.
	 * @param ctx the parse tree
	 */
	void exitRangeQuantifier(SimpleRegexpParser.RangeQuantifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierExpression}.
	 * @param ctx the parse tree
	 */
	void enterRangeQuantifierExpression(SimpleRegexpParser.RangeQuantifierExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierExpression}.
	 * @param ctx the parse tree
	 */
	void exitRangeQuantifierExpression(SimpleRegexpParser.RangeQuantifierExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierExact}.
	 * @param ctx the parse tree
	 */
	void enterRangeQuantifierExact(SimpleRegexpParser.RangeQuantifierExactContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierExact}.
	 * @param ctx the parse tree
	 */
	void exitRangeQuantifierExact(SimpleRegexpParser.RangeQuantifierExactContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierMinMax}.
	 * @param ctx the parse tree
	 */
	void enterRangeQuantifierMinMax(SimpleRegexpParser.RangeQuantifierMinMaxContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierMinMax}.
	 * @param ctx the parse tree
	 */
	void exitRangeQuantifierMinMax(SimpleRegexpParser.RangeQuantifierMinMaxContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierMin}.
	 * @param ctx the parse tree
	 */
	void enterRangeQuantifierMin(SimpleRegexpParser.RangeQuantifierMinContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#rangeQuantifierMin}.
	 * @param ctx the parse tree
	 */
	void exitRangeQuantifierMin(SimpleRegexpParser.RangeQuantifierMinContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(SimpleRegexpParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(SimpleRegexpParser.NumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#term}.
	 * @param ctx the parse tree
	 */
	void enterTerm(SimpleRegexpParser.TermContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#term}.
	 * @param ctx the parse tree
	 */
	void exitTerm(SimpleRegexpParser.TermContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#parenExpression}.
	 * @param ctx the parse tree
	 */
	void enterParenExpression(SimpleRegexpParser.ParenExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#parenExpression}.
	 * @param ctx the parse tree
	 */
	void exitParenExpression(SimpleRegexpParser.ParenExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#orExpression}.
	 * @param ctx the parse tree
	 */
	void enterOrExpression(SimpleRegexpParser.OrExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#orExpression}.
	 * @param ctx the parse tree
	 */
	void exitOrExpression(SimpleRegexpParser.OrExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(SimpleRegexpParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(SimpleRegexpParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#atomEscaped}.
	 * @param ctx the parse tree
	 */
	void enterAtomEscaped(SimpleRegexpParser.AtomEscapedContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#atomEscaped}.
	 * @param ctx the parse tree
	 */
	void exitAtomEscaped(SimpleRegexpParser.AtomEscapedContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#atom}.
	 * @param ctx the parse tree
	 */
	void enterAtom(SimpleRegexpParser.AtomContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#atom}.
	 * @param ctx the parse tree
	 */
	void exitAtom(SimpleRegexpParser.AtomContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleRegexpParser#characterClassSimple}.
	 * @param ctx the parse tree
	 */
	void enterCharacterClassSimple(SimpleRegexpParser.CharacterClassSimpleContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleRegexpParser#characterClassSimple}.
	 * @param ctx the parse tree
	 */
	void exitCharacterClassSimple(SimpleRegexpParser.CharacterClassSimpleContext ctx);
}