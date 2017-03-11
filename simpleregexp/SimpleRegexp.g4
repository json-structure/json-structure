grammar SimpleRegexp;

/*
 * JSON Schema Validation: A Vocabulary for Structural Validation of JSON
 *
 * Section 3.3
 * 
 * ...
 * ...
 * 
 * Furthermore, given the high disparity in regular expression constructs support, 
 * schema authors SHOULD limit themselves to the following regular expression tokens:
 *
 * individual Unicode characters, as defined by the JSON specification [RFC7159];
 * simple character classes ([abc]), range character classes ([a-z]);
 * complemented character classes ([^abc], [^a-z]);
 * simple quantifiers: "+" (one or more), "*" (zero or more), "?" (zero or one), and their lazy versions ("+?", "*?", "??");
 * range quantifiers: "{x}" (exactly x occurrences), "{x,y}" (at least x, at most y, occurrences), {x,} (x occurrences or more), and their lazy versions;
 * the beginning-of-input ("^") and end-of-input ("$") anchors;
 * simple grouping ("(...)") and alternation ("|").
 *
 */

start : Caret? expression* Dollar? EOF;

characterClass           : LeftBracket Caret? characterClassExpression+ RightBracket;
characterClassExpression : characterClassAtom | characterClassRange;
characterClassRange      : characterClassAtom Dash characterClassAtom;
characterClassAtom       : characterClassSimple | characterClassEscaped;
characterClassEscaped    : Slash (Slash | LeftBracket | RightBracket | Caret | Dash);

quantifier                  : (simpleQuantifier | rangeQuantifier) Question?;
simpleQuantifier            : Plus | Star | Question;
rangeQuantifier             : LeftBrace rangeQuantifierExpression RightBrace;
rangeQuantifierExpression   : rangeQuantifierExact | rangeQuantifierMinMax | rangeQuantifierMin;
rangeQuantifierExact        : number;
rangeQuantifierMinMax       : number Comma number;
rangeQuantifierMin          : number Comma;
number                      : Digit+;

term : (atom | atomEscaped | characterClass)+;

parenExpression  : LeftParen expression RightParen;
orExpression     : Pipe expression;
expression       : (parenExpression | orExpression | term) quantifier?;

atomEscaped: Slash (Slash
    | Dot
    | Plus
    | Star
    | Question
    | LeftParen
    | RightParen
    | Pipe
    | LeftBracket
    | RightBracket
    | LeftBrace
    | RightBrace
    | Caret
    | Dollar);

Dot           : '.';
Plus          : '+';
Star          : '*';
Question      : '?';
LeftParen     : '(';
RightParen    : ')';
Pipe          : '|';
LeftBrace     : '{';
RightBrace    : '}';
Dollar        : '$';
Slash         : '\\';
LeftBracket   : '[';
RightBracket  : ']';
Caret         : '^';
Dash          : '-';
Comma         : ',';

Digit      : [0-9];
Letter     : [A-Za-z];
Space      : ' ';

atom: Letter | Digit | Space | Dot | Dash | Comma;

// Inside a character class only
// Slash, LeftBracket, RightBracket, Caret, and Dash
// need to be escaped.
characterClassSimple : Letter
    | Digit
    | Space
    | Dot
    | Plus
    | Star
    | Question
    | LeftParen
    | RightParen
    | Pipe
    | LeftBrace
    | RightBrace
    | Dollar
    | Slash
    | Comma;
