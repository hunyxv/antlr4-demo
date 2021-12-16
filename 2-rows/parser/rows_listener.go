// Code generated from Rows.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Rows

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RowsListener is a complete listener for a parse tree produced by RowsParser.
type RowsListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)
}
