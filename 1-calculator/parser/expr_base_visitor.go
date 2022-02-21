// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitIdentity(ctx *IdentityContext) interface{} {
	return v.VisitChildren(ctx)
}
