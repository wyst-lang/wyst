// Code generated from Wyst.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Wyst

import "github.com/antlr4-go/antlr/v4"

// WystListener is a complete listener for a parse tree produced by WystParser.
type WystListener interface {
	antlr.ParseTreeListener

	// EnterRound_def is called when entering the round_def production.
	EnterRound_def(c *Round_defContext)

	// EnterEnum_curly is called when entering the enum_curly production.
	EnterEnum_curly(c *Enum_curlyContext)

	// EnterRound_call is called when entering the round_call production.
	EnterRound_call(c *Round_callContext)

	// EnterFn_call is called when entering the fn_call production.
	EnterFn_call(c *Fn_callContext)

	// EnterFunc_def is called when entering the func_def production.
	EnterFunc_def(c *Func_defContext)

	// EnterVar_def is called when entering the var_def production.
	EnterVar_def(c *Var_defContext)

	// EnterVar_def_set is called when entering the var_def_set production.
	EnterVar_def_set(c *Var_def_setContext)

	// EnterCall_tree is called when entering the call_tree production.
	EnterCall_tree(c *Call_treeContext)

	// EnterStruct_def is called when entering the struct_def production.
	EnterStruct_def(c *Struct_defContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterImport_statement is called when entering the import_statement production.
	EnterImport_statement(c *Import_statementContext)

	// EnterUse_statement is called when entering the use_statement production.
	EnterUse_statement(c *Use_statementContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCode_block is called when entering the code_block production.
	EnterCode_block(c *Code_blockContext)

	// EnterTop is called when entering the top production.
	EnterTop(c *TopContext)

	// ExitRound_def is called when exiting the round_def production.
	ExitRound_def(c *Round_defContext)

	// ExitEnum_curly is called when exiting the enum_curly production.
	ExitEnum_curly(c *Enum_curlyContext)

	// ExitRound_call is called when exiting the round_call production.
	ExitRound_call(c *Round_callContext)

	// ExitFn_call is called when exiting the fn_call production.
	ExitFn_call(c *Fn_callContext)

	// ExitFunc_def is called when exiting the func_def production.
	ExitFunc_def(c *Func_defContext)

	// ExitVar_def is called when exiting the var_def production.
	ExitVar_def(c *Var_defContext)

	// ExitVar_def_set is called when exiting the var_def_set production.
	ExitVar_def_set(c *Var_def_setContext)

	// ExitCall_tree is called when exiting the call_tree production.
	ExitCall_tree(c *Call_treeContext)

	// ExitStruct_def is called when exiting the struct_def production.
	ExitStruct_def(c *Struct_defContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitImport_statement is called when exiting the import_statement production.
	ExitImport_statement(c *Import_statementContext)

	// ExitUse_statement is called when exiting the use_statement production.
	ExitUse_statement(c *Use_statementContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCode_block is called when exiting the code_block production.
	ExitCode_block(c *Code_blockContext)

	// ExitTop is called when exiting the top production.
	ExitTop(c *TopContext)
}
