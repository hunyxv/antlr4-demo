// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Data

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDataVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDataVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDataVisitor) VisitGroup(ctx *GroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDataVisitor) VisitSequence(ctx *SequenceContext) interface{} {
	return v.VisitChildren(ctx)
}
