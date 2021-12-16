// Code generated from Hello.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseHelloVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHelloVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHelloVisitor) VisitParenthesis(ctx *ParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHelloVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHelloVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHelloVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}
