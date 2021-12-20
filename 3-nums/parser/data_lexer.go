// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 4, 19, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 3, 2, 6, 2, 9, 10, 2, 13, 2, 14, 2, 10, 3, 3,
	6, 3, 14, 10, 3, 13, 3, 14, 3, 15, 3, 3, 3, 3, 2, 2, 4, 3, 3, 5, 4, 3,
	2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 20, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 3, 8, 3, 2, 2, 2, 5, 13, 3, 2, 2, 2, 7, 9, 9, 2, 2,
	2, 8, 7, 3, 2, 2, 2, 9, 10, 3, 2, 2, 2, 10, 8, 3, 2, 2, 2, 10, 11, 3, 2,
	2, 2, 11, 4, 3, 2, 2, 2, 12, 14, 9, 3, 2, 2, 13, 12, 3, 2, 2, 2, 14, 15,
	3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2,
	17, 18, 8, 3, 2, 2, 18, 6, 3, 2, 2, 2, 5, 2, 10, 15, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "INT", "WS",
}

var lexerRuleNames = []string{
	"INT", "WS",
}

type DataLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDataLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DataLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDataLexer(input antlr.CharStream) *DataLexer {
	l := new(DataLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Data.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DataLexer tokens.
const (
	DataLexerINT = 1
	DataLexerWS  = 2
)
