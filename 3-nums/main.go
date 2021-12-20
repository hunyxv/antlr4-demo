package main

import (
	"antlr4-demo/3-nums/parser"
	"io/ioutil"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	stream := antlr.NewInputStream(string(data))
	lexer := parser.NewDataLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewDataParser(tokenStream)
	//p.BuildParseTrees = false
	listerner := NewDataListerner()
	antlr.ParseTreeWalkerDefault.Walk(listerner, p.File())

	//p.File()

}
