// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Data

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 4, 25, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 6, 2, 10, 10, 2, 13, 2, 14, 2, 11,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 20, 10, 4, 12, 4, 14, 4, 23,
	11, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 23, 2, 9, 3, 2, 2, 2, 4, 13, 3,
	2, 2, 2, 6, 21, 3, 2, 2, 2, 8, 10, 5, 4, 3, 2, 9, 8, 3, 2, 2, 2, 10, 11,
	3, 2, 2, 2, 11, 9, 3, 2, 2, 2, 11, 12, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13,
	14, 7, 3, 2, 2, 14, 15, 5, 6, 4, 2, 15, 5, 3, 2, 2, 2, 16, 17, 6, 4, 2,
	3, 17, 18, 7, 3, 2, 2, 18, 20, 8, 4, 1, 2, 19, 16, 3, 2, 2, 2, 20, 23,
	3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 7, 3, 2, 2, 2,
	23, 21, 3, 2, 2, 2, 4, 11, 21,
}
var literalNames []string

var symbolicNames = []string{
	"", "INT", "WS",
}

var ruleNames = []string{
	"file", "group", "sequence",
}

type DataParser struct {
	*antlr.BaseParser
}

// NewDataParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DataParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDataParser(input antlr.TokenStream) *DataParser {
	this := new(DataParser)
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
	this.GrammarFileName = "Data.g4"

	return this
}

// DataParser tokens.
const (
	DataParserEOF = antlr.TokenEOF
	DataParserINT = 1
	DataParserWS  = 2
)

// DataParser rules.
const (
	DataParserRULE_file     = 0
	DataParserRULE_group    = 1
	DataParserRULE_sequence = 2
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
	p.RuleIndex = DataParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllGroup() []IGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupContext)(nil)).Elem())
	var tst = make([]IGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupContext)
		}
	}

	return tst
}

func (s *FileContext) Group(i int) IGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DataParserRULE_file)
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

	for ok := true; ok; ok = _la == DataParserINT {
		{
			p.SetState(6)
			p.Group()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_INT   antlr.Token
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DataParserRULE_group
	return p
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) Get_INT() antlr.Token { return s._INT }

func (s *GroupContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *GroupContext) INT() antlr.TerminalNode {
	return s.GetToken(DataParserINT, 0)
}

func (s *GroupContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.ExitGroup(s)
	}
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataVisitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataParser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DataParserRULE_group)

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
	{
		p.SetState(11)

		var _m = p.Match(DataParserINT)

		localctx.(*GroupContext)._INT = _m
	}
	{
		p.SetState(12)
		p.Sequence((func() int {
			if localctx.(*GroupContext).Get_INT() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*GroupContext).Get_INT().GetText())
				return i
			}
		}()))
	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n attribute.
	GetN() int

	// GetI returns the i attribute.
	GetI() int

	// SetN sets the n attribute.
	SetN(int)

	// SetI sets the i attribute.
	SetI(int)

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	n      int
	i      int // TODO = 0
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DataParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, n int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DataParserRULE_sequence

	p.n = n

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) GetN() int { return s.n }

func (s *SequenceContext) GetI() int { return s.i }

func (s *SequenceContext) SetN(v int) { s.n = v }

func (s *SequenceContext) SetI(v int) { s.i = v }

func (s *SequenceContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(DataParserINT)
}

func (s *SequenceContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(DataParserINT, i)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DataListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DataVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DataParser) Sequence(n int) (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState(), n)
	p.EnterRule(localctx, 4, DataParserRULE_sequence)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(14)

			if !(localctx.(*SequenceContext).i < localctx.(*SequenceContext).n) {
				panic(antlr.NewFailedPredicateException(p, "$i<$n", ""))
			}
			{
				p.SetState(15)
				p.Match(DataParserINT)
			}
			localctx.(*SequenceContext).i++

		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

func (p *DataParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *SequenceContext = nil
		if localctx != nil {
			t = localctx.(*SequenceContext)
		}
		return p.Sequence_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DataParser) Sequence_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return localctx.(*SequenceContext).i < localctx.(*SequenceContext).n

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
