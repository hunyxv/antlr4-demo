// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseExprListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseExprListener) ExitStart(ctx *StartContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseExprListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseExprListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseExprListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseExprListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterIdentity is called when production Identity is entered.
func (s *BaseExprListener) EnterIdentity(ctx *IdentityContext) {}

// ExitIdentity is called when production Identity is exited.
func (s *BaseExprListener) ExitIdentity(ctx *IdentityContext) {}
