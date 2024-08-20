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

// EnterPower_keyword is called when production power_keyword is entered.
func (s *BaseWystListener) EnterPower_keyword(ctx *Power_keywordContext) {}

// ExitPower_keyword is called when production power_keyword is exited.
func (s *BaseWystListener) ExitPower_keyword(ctx *Power_keywordContext) {}

// EnterPower_keyword_call is called when production power_keyword_call is entered.
func (s *BaseWystListener) EnterPower_keyword_call(ctx *Power_keyword_callContext) {}

// ExitPower_keyword_call is called when production power_keyword_call is exited.
func (s *BaseWystListener) ExitPower_keyword_call(ctx *Power_keyword_callContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseWystListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseWystListener) ExitNamespace(ctx *NamespaceContext) {}

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
