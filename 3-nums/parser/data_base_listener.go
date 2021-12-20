// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Data

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDataListener is a complete listener for a parse tree produced by DataParser.
type BaseDataListener struct{}

var _ DataListener = &BaseDataListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDataListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDataListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDataListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDataListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseDataListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseDataListener) ExitFile(ctx *FileContext) {}

// EnterGroup is called when production group is entered.
func (s *BaseDataListener) EnterGroup(ctx *GroupContext) {}

// ExitGroup is called when production group is exited.
func (s *BaseDataListener) ExitGroup(ctx *GroupContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BaseDataListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BaseDataListener) ExitSequence(ctx *SequenceContext) {}
