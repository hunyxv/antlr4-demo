// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Rows

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRowsListener is a complete listener for a parse tree produced by RowsParser.
type BaseRowsListener struct{}

var _ RowsListener = &BaseRowsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRowsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRowsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRowsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRowsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseRowsListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseRowsListener) ExitFile(ctx *FileContext) {}

// EnterRow is called when production row is entered.
func (s *BaseRowsListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseRowsListener) ExitRow(ctx *RowContext) {}
