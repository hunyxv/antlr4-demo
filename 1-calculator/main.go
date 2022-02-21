package main

import (
	"antlr4-demo/1-calculator/parser"
	"fmt"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func runVisitor() {
	stream := antlr.NewInputStream("(1 + 2) * (3+2) *  (2  - 3)")

	lexer := parser.NewExprLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewExprParser(tokenStream)

	visitor := &CalVisitor{}
	p.Start().Accept(visitor)

	result, err := visitor.Pop()
	if err != nil {
		panic(err)
	}

	fmt.Printf("visitor: %d\n", result)
}

func runListener() {
	stream := antlr.NewInputStream("(1 + 2) * (3+2) *  (2  - 3)")
	lexer := parser.NewExprLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(tokenStream)
	listener := &CalcListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	result, err := listener.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("listener: %d\n", result)
}

func main() {
	now := time.Now()
	runVisitor()
	fmt.Print("\n ----------", time.Since(now) ,"----------\n\n")
	now = time.Now()
	runListener()
	fmt.Print("\n ----------", time.Since(now) ,"----------\n\n")
}
