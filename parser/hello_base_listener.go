// Code generated from Hello.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHelloListener is a complete listener for a parse tree produced by HelloParser.
type BaseHelloListener struct{}

var _ HelloListener = &BaseHelloListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHelloListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHelloListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHelloListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHelloListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseHelloListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseHelloListener) ExitStart(ctx *StartContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseHelloListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseHelloListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseHelloListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseHelloListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseHelloListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseHelloListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseHelloListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseHelloListener) ExitAddSub(ctx *AddSubContext) {}
