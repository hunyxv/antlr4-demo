// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ExprParser#Parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by ExprParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by ExprParser#Identity.
	VisitIdentity(ctx *IdentityContext) interface{}
}
