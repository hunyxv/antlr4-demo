package main

import (
	"antlr4-demo/1-calculator/parser"
	"errors"
	"log"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CalcListener struct {
	*parser.BaseExprListener

	stack []int
	i     int
}

func (l *CalcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *CalcListener) Pop() (int, error) {
	if len(l.stack) == 0 {
		return 0, errors.New("stack is empty")
	}

	i := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return i, nil
}

func (l *CalcListener) ExitNumber(ctx *parser.NumberContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}

	l.push(i)
	l.i++
	log.Println("listener exit number", l.i, l.stack)
}

func (l *CalcListener) ExitMulDiv(ctx *parser.MulDivContext) {
	var token antlr.Token = ctx.GetOp()
	right, err := l.Pop()
	if err != nil {
		panic(err)
	}
	left, err := l.Pop()
	if err != nil {
		panic(err)
	}
	switch token.GetTokenType() {
	case parser.ExprParserMUL:
		l.push(left * right)
	case parser.ExprParserDIV:
		l.push(left / right)
	default:
		panic("exceptions error")
	}
	l.i++
	log.Println("listener exit muldiv", l.i, l.stack)
}

func (l *CalcListener) ExitAddSub(ctx *parser.AddSubContext) {
	var token antlr.Token = ctx.GetOp()
	right, err := l.Pop()
	if err != nil {
		panic(err)
	}
	left, err := l.Pop()
	if err != nil {
		panic(err)
	}
	switch token.GetTokenType() {
	case parser.ExprParserADD:
		l.push(left + right)
	case parser.ExprParserSUB:
		l.push(left - right)
	default:
		panic("exceptions error")
	}
	l.i++
	log.Println("listener exit addsub", l.i, l.stack)
}

func (l *CalcListener) EnterParenthesis(ctx *parser.ParenthesisContext) {
	l.i++
	log.Println("listener enter parenthesis", l.i, l.stack)
}

func (l *CalcListener) ExitParenthesis(ctx *parser.ParenthesisContext) {
	l.i++
	log.Println("listener exit parenthesis", l.i, l.stack)
}
