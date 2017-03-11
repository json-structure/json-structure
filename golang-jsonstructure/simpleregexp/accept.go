package simpleregexp

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CaptureErrorListener struct {
	antlr.DefaultErrorListener
	Err error
}

func (c *CaptureErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Err = errors.New(msg)
}

func Accept(input string) error {
	iStream := antlr.NewInputStream(input)
	lexer := NewSimpleRegexpLexer(iStream)
	tStream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewSimpleRegexpParser(tStream)
	p.RemoveErrorListeners()
	listen := CaptureErrorListener{}
	p.AddErrorListener(&listen)
	p.Start()
	return listen.Err
}
