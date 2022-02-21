// Code generated from Expr.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterIdentity is called when entering the Identity production.
	EnterIdentity(c *IdentityContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitIdentity is called when exiting the Identity production.
	ExitIdentity(c *IdentityContext)
}
