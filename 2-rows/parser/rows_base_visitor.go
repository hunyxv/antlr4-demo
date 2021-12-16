// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Rows

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRowsVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRowsVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRowsVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}
