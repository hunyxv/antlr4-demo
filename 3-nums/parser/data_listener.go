// Code generated from Data.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Data

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DataListener is a complete listener for a parse tree produced by DataParser.
type DataListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterGroup is called when entering the group production.
	EnterGroup(c *GroupContext)

	// EnterSequence is called when entering the sequence production.
	EnterSequence(c *SequenceContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitGroup is called when exiting the group production.
	ExitGroup(c *GroupContext)

	// ExitSequence is called when exiting the sequence production.
	ExitSequence(c *SequenceContext)
}
