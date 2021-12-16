// Code generated from Hello.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by HelloParser.
type HelloVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HelloParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by HelloParser#Parenthesis.
	VisitParenthesis(ctx *ParenthesisContext) interface{}

	// Visit a parse tree produced by HelloParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by HelloParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by HelloParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}
}
