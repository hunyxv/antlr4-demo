// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Rows

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 20, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 6, 2, 10, 10, 2, 13, 2, 14, 2, 11,
	3, 3, 3, 3, 6, 3, 16, 10, 3, 13, 3, 14, 3, 17, 3, 3, 2, 2, 4, 2, 4, 2,
	2, 2, 19, 2, 9, 3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8,
	7, 4, 2, 2, 8, 10, 3, 2, 2, 2, 9, 6, 3, 2, 2, 2, 10, 11, 3, 2, 2, 2, 11,
	9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13, 14, 7, 5, 2,
	2, 14, 16, 8, 3, 1, 2, 15, 13, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 15,
	3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 5, 3, 2, 2, 2, 4, 11, 17,
}
var literalNames = []string{
	"", "'\t'",
}
var symbolicNames = []string{
	"", "TAB", "NL", "STUFF",
}

var ruleNames = []string{
	"file", "row",
}

type RowsParser struct {
	*antlr.BaseParser
}

// NewRowsParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *RowsParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewRowsParser(input antlr.TokenStream) *RowsParser {
	this := new(RowsParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rows.g4"

	return this
}

// add members to generated RowsParser

var col int

func NewMyRowsParser(input antlr.TokenStream, c int) *RowsParser {
	col = c
	return NewRowsParser(input)
}

// RowsParser tokens.
const (
	RowsParserEOF   = antlr.TokenEOF
	RowsParserTAB   = 1
	RowsParserNL    = 2
	RowsParserSTUFF = 3
)

// RowsParser rules.
const (
	RowsParserRULE_file = 0
	RowsParserRULE_row  = 1
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RowsParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RowsParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *FileContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *FileContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(RowsParserNL)
}

func (s *FileContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(RowsParserNL, i)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RowsListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RowsListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RowsVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RowsParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RowsParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RowsParserSTUFF {
		{
			p.SetState(4)
			p.Row()
		}
		{
			p.SetState(5)
			p.Match(RowsParserNL)
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STUFF returns the _STUFF token.
	Get_STUFF() antlr.Token

	// Set_STUFF sets the _STUFF token.
	Set_STUFF(antlr.Token)

	// GetI returns the i attribute.
	GetI() int

	// SetI sets the i attribute.
	SetI(int)

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	i      int // TODO = 0
	_STUFF antlr.Token
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RowsParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RowsParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) Get_STUFF() antlr.Token { return s._STUFF }

func (s *RowContext) Set_STUFF(v antlr.Token) { s._STUFF = v }

func (s *RowContext) GetI() int { return s.i }

func (s *RowContext) SetI(v int) { s.i = v }

func (s *RowContext) AllSTUFF() []antlr.TerminalNode {
	return s.GetTokens(RowsParserSTUFF)
}

func (s *RowContext) STUFF(i int) antlr.TerminalNode {
	return s.GetToken(RowsParserSTUFF, i)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RowsListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RowsListener); ok {
		listenerT.ExitRow(s)
	}
}

func (s *RowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RowsVisitor:
		return t.VisitRow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RowsParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RowsParserRULE_row)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RowsParserSTUFF {
		{
			p.SetState(11)

			var _m = p.Match(RowsParserSTUFF)

			localctx.(*RowContext)._STUFF = _m
		}

		localctx.(*RowContext).i++
		if localctx.(*RowContext).i == col {
			fmt.Println((func() string {
				if localctx.(*RowContext).Get_STUFF() == nil {
					return ""
				} else {
					return localctx.(*RowContext).Get_STUFF().GetText()
				}
			}()))
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
