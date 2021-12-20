// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Data

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DataParser.
type DataVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DataParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by DataParser#group.
	VisitGroup(ctx *GroupContext) interface{}

	// Visit a parse tree produced by DataParser#sequence.
	VisitSequence(ctx *SequenceContext) interface{}
}
