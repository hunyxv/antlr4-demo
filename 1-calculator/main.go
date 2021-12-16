package main

import (
	"antlr4-demo/1-calculator/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func runVisitor() {
	stream := antlr.NewInputStream("(1 + 2) * (3+2) *  (2  - 3)")

	lexer := parser.NewHelloLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewHelloParser(tokenStream)

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
	lexer := parser.NewHelloLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewHelloParser(tokenStream)
	listener := &CalcListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Start())

	result, err := listener.Pop()
	if err != nil {
		panic(err)
	}
	fmt.Printf("listener: %d\n", result)
}

func main() {
	runVisitor()
	fmt.Print("\n --------------------\n\n")
	runListener()
}
