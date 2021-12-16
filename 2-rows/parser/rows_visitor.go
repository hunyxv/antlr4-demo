// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Rows

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RowsParser.
type RowsVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RowsParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by RowsParser#row.
	VisitRow(ctx *RowContext) interface{}
}
