// Code generated from Wyst.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Wyst

import "github.com/antlr4-go/antlr/v4"

// BaseWystListener is a complete listener for a parse tree produced by WystParser.
type BaseWystListener struct{}

var _ WystListener = &BaseWystListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWystListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWystListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWystListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWystListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRound_def is called when production round_def is entered.
func (s *BaseWystListener) EnterRound_def(ctx *Round_defContext) {}

// ExitRound_def is called when production round_def is exited.
func (s *BaseWystListener) ExitRound_def(ctx *Round_defContext) {}

// EnterEnum_curly is called when production enum_curly is entered.
func (s *BaseWystListener) EnterEnum_curly(ctx *Enum_curlyContext) {}

// ExitEnum_curly is called when production enum_curly is exited.
func (s *BaseWystListener) ExitEnum_curly(ctx *Enum_curlyContext) {}

// EnterRound_call is called when production round_call is entered.
func (s *BaseWystListener) EnterRound_call(ctx *Round_callContext) {}

// ExitRound_call is called when production round_call is exited.
func (s *BaseWystListener) ExitRound_call(ctx *Round_callContext) {}

// EnterFn_call is called when production fn_call is entered.
func (s *BaseWystListener) EnterFn_call(ctx *Fn_callContext) {}

// ExitFn_call is called when production fn_call is exited.
func (s *BaseWystListener) ExitFn_call(ctx *Fn_callContext) {}

// EnterFunc_def is called when production func_def is entered.
func (s *BaseWystListener) EnterFunc_def(ctx *Func_defContext) {}

// ExitFunc_def is called when production func_def is exited.
func (s *BaseWystListener) ExitFunc_def(ctx *Func_defContext) {}

// EnterVar_def is called when production var_def is entered.
func (s *BaseWystListener) EnterVar_def(ctx *Var_defContext) {}

// ExitVar_def is called when production var_def is exited.
func (s *BaseWystListener) ExitVar_def(ctx *Var_defContext) {}

// EnterVar_def_set is called when production var_def_set is entered.
func (s *BaseWystListener) EnterVar_def_set(ctx *Var_def_setContext) {}

// ExitVar_def_set is called when production var_def_set is exited.
func (s *BaseWystListener) ExitVar_def_set(ctx *Var_def_setContext) {}

// EnterCall_tree is called when production call_tree is entered.
func (s *BaseWystListener) EnterCall_tree(ctx *Call_treeContext) {}

// ExitCall_tree is called when production call_tree is exited.
func (s *BaseWystListener) ExitCall_tree(ctx *Call_treeContext) {}

// EnterStruct_def is called when production struct_def is entered.
func (s *BaseWystListener) EnterStruct_def(ctx *Struct_defContext) {}

// ExitStruct_def is called when production struct_def is exited.
func (s *BaseWystListener) ExitStruct_def(ctx *Struct_defContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseWystListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseWystListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterImport_statement is called when production import_statement is entered.
func (s *BaseWystListener) EnterImport_statement(ctx *Import_statementContext) {}

// ExitImport_statement is called when production import_statement is exited.
func (s *BaseWystListener) ExitImport_statement(ctx *Import_statementContext) {}

// EnterUse_statement is called when production use_statement is entered.
func (s *BaseWystListener) EnterUse_statement(ctx *Use_statementContext) {}

// ExitUse_statement is called when production use_statement is exited.
func (s *BaseWystListener) ExitUse_statement(ctx *Use_statementContext) {}

// EnterGo_import is called when production go_import is entered.
func (s *BaseWystListener) EnterGo_import(ctx *Go_importContext) {}

// ExitGo_import is called when production go_import is exited.
func (s *BaseWystListener) ExitGo_import(ctx *Go_importContext) {}

// EnterMap is called when production map is entered.
func (s *BaseWystListener) EnterMap(ctx *MapContext) {}

// ExitMap is called when production map is exited.
func (s *BaseWystListener) ExitMap(ctx *MapContext) {}

// EnterGo_call is called when production go_call is entered.
func (s *BaseWystListener) EnterGo_call(ctx *Go_callContext) {}

// ExitGo_call is called when production go_call is exited.
func (s *BaseWystListener) ExitGo_call(ctx *Go_callContext) {}

// EnterIf_expr is called when production if_expr is entered.
func (s *BaseWystListener) EnterIf_expr(ctx *If_exprContext) {}

// ExitIf_expr is called when production if_expr is exited.
func (s *BaseWystListener) ExitIf_expr(ctx *If_exprContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseWystListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseWystListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElseif_statement is called when production elseif_statement is entered.
func (s *BaseWystListener) EnterElseif_statement(ctx *Elseif_statementContext) {}

// ExitElseif_statement is called when production elseif_statement is exited.
func (s *BaseWystListener) ExitElseif_statement(ctx *Elseif_statementContext) {}

// EnterElse_statement is called when production else_statement is entered.
func (s *BaseWystListener) EnterElse_statement(ctx *Else_statementContext) {}

// ExitElse_statement is called when production else_statement is exited.
func (s *BaseWystListener) ExitElse_statement(ctx *Else_statementContext) {}

// EnterIf_tree is called when production if_tree is entered.
func (s *BaseWystListener) EnterIf_tree(ctx *If_treeContext) {}

// ExitIf_tree is called when production if_tree is exited.
func (s *BaseWystListener) ExitIf_tree(ctx *If_treeContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseWystListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseWystListener) ExitExpr(ctx *ExprContext) {}

// EnterCode_block is called when production code_block is entered.
func (s *BaseWystListener) EnterCode_block(ctx *Code_blockContext) {}

// ExitCode_block is called when production code_block is exited.
func (s *BaseWystListener) ExitCode_block(ctx *Code_blockContext) {}

// EnterTop is called when production top is entered.
func (s *BaseWystListener) EnterTop(ctx *TopContext) {}

// ExitTop is called when production top is exited.
func (s *BaseWystListener) ExitTop(ctx *TopContext) {}
