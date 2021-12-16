package main

import (
	"antlr4-demo/1-calculator/parser"
	"errors"
	"log"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// 访问者模式
type CalVisitor struct {
	parser.BaseHelloVisitor
	stack []int
	i     int
}

func (v *CalVisitor) push(i int) {
	v.stack = append(v.stack, i)
}

func (v *CalVisitor) Pop() (int, error) {
	if len(v.stack) == 0 {
		return 0, errors.New("stack is empty")
	}

	i := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]
	return i, nil
}

// func (v *CalVisitor) visitRule(node antlr.RuleNode) interface{} {
// 	node.Accept(v)
// 	return nil
// }

func (v *CalVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	v.i++
	log.Println("visit start", v.i)
	return ctx.Expr().Accept(v)
}

func (v *CalVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err)
	}
	v.push(i)
	v.i++
	log.Println("visit number", v.i, v.stack)
	return nil
}

func (v *CalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	ctx.Expr(0).Accept(v)
	ctx.Expr(1).Accept(v)

	var token antlr.Token = ctx.GetOp()
	right, err := v.Pop()
	if err != nil {
		panic(err)
	}
	left, err := v.Pop()
	if err != nil {
		panic(err)
	}
	switch token.GetTokenType() {
	case parser.HelloParserMUL:
		v.push(left * right)
	case parser.HelloParserDIV:
		v.push(left / right)
	default:
		panic("exceptions error")
	}
	v.i++
	log.Println("visit mul div", v.i, v.stack)

	return nil
}

func (v *CalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	ctx.Expr(0).Accept(v)
	ctx.Expr(1).Accept(v)

	var token antlr.Token = ctx.GetOp()
	right, err := v.Pop()
	if err != nil {
		panic(err)
	}
	left, err := v.Pop()
	if err != nil {
		panic(err)
	}
	switch token.GetTokenType() {
	case parser.HelloParserADD:
		v.push(left + right)
	case parser.HelloParserSUB:
		v.push(left - right)
	default:
		panic("exceptions error")
	}
	v.i++
	log.Println("visit add sub", v.i, v.stack)
	return nil
}

func (v *CalVisitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	log.Println("visit parenthesis start", v.i, v.stack)
	ctx.Expr().Accept(v)
	v.i++
	log.Println("visit parenthesis end", v.i, v.stack)
	return nil
}
