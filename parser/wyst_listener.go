// Code generated from Wyst.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Wyst

import "github.com/antlr4-go/antlr/v4"

// WystListener is a complete listener for a parse tree produced by WystParser.
type WystListener interface {
	antlr.ParseTreeListener

	// EnterRound_def is called when entering the round_def production.
	EnterRound_def(c *Round_defContext)

	// EnterFunc_def is called when entering the func_def production.
	EnterFunc_def(c *Func_defContext)

	// EnterVar_def is called when entering the var_def production.
	EnterVar_def(c *Var_defContext)

	// EnterVar_def_set is called when entering the var_def_set production.
	EnterVar_def_set(c *Var_def_setContext)

	// EnterAsm is called when entering the asm production.
	EnterAsm(c *AsmContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCode_block is called when entering the code_block production.
	EnterCode_block(c *Code_blockContext)

	// EnterTop is called when entering the top production.
	EnterTop(c *TopContext)

	// ExitRound_def is called when exiting the round_def production.
	ExitRound_def(c *Round_defContext)

	// ExitFunc_def is called when exiting the func_def production.
	ExitFunc_def(c *Func_defContext)

	// ExitVar_def is called when exiting the var_def production.
	ExitVar_def(c *Var_defContext)

	// ExitVar_def_set is called when exiting the var_def_set production.
	ExitVar_def_set(c *Var_def_setContext)

	// ExitAsm is called when exiting the asm production.
	ExitAsm(c *AsmContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCode_block is called when exiting the code_block production.
	ExitCode_block(c *Code_blockContext)

	// ExitTop is called when exiting the top production.
	ExitTop(c *TopContext)
}
