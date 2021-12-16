// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 23, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 5,
	3, 15, 10, 3, 3, 3, 3, 3, 3, 4, 6, 4, 20, 10, 4, 13, 4, 14, 4, 21, 2, 2,
	5, 3, 3, 5, 4, 7, 5, 3, 2, 3, 4, 2, 11, 12, 15, 15, 2, 24, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 3, 9, 3, 2, 2, 2, 5, 14, 3, 2,
	2, 2, 7, 19, 3, 2, 2, 2, 9, 10, 7, 11, 2, 2, 10, 11, 3, 2, 2, 2, 11, 12,
	8, 2, 2, 2, 12, 4, 3, 2, 2, 2, 13, 15, 7, 15, 2, 2, 14, 13, 3, 2, 2, 2,
	14, 15, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 17, 7, 12, 2, 2, 17, 6, 3,
	2, 2, 2, 18, 20, 10, 2, 2, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21,
	19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 8, 3, 2, 2, 2, 5, 2, 14, 21, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\t'",
}

var lexerSymbolicNames = []string{
	"", "TAB", "NL", "STUFF",
}

var lexerRuleNames = []string{
	"TAB", "NL", "STUFF",
}

type RowsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewRowsLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *RowsLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRowsLexer(input antlr.CharStream) *RowsLexer {
	l := new(RowsLexer)
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
	l.GrammarFileName = "Rows.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RowsLexer tokens.
const (
	RowsLexerTAB   = 1
	RowsLexerNL    = 2
	RowsLexerSTUFF = 3
)
