package main

import (
	"antlr4-demo/2-rows/parser"
	"io/ioutil"
	"os"
	"strconv"

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
	lexer := parser.NewRowsLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	col, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	rowsParser := parser.NewMyRowsParser(tokenStream, col)
	
	rowsParser.BuildParseTrees = false
	rowsParser.File()
}