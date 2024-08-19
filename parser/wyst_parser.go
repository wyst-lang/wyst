// Code generated from Wyst.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Wyst

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type WystParser struct {
	*antlr.BaseParser
}

var WystParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func wystParserInit() {
	staticData := &WystParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "','", "')'", "'='", "'%'", "';'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "WS", "NUMBER", "HEX", "IDENTIFIER",
		"MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"round_def", "func_def", "var_def", "var_def_set", "asm", "expr", "code_block",
		"top",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 21,
		8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 0, 3, 0, 27, 8, 0, 3, 0, 29, 8, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 50, 8, 4, 10, 4, 12, 4, 53,
		9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 60, 8, 5, 10, 5, 12, 5, 63, 9,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 69, 8, 6, 1, 6, 1, 6, 5, 6, 73, 8, 6,
		10, 6, 12, 6, 76, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 82, 8, 7, 10, 7,
		12, 7, 85, 9, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 2, 0,
		10, 12, 14, 14, 88, 0, 16, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 37, 1, 0,
		0, 0, 6, 40, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 64,
		1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 28, 5, 1, 0, 0, 17, 22, 3, 4, 2, 0,
		18, 19, 5, 2, 0, 0, 19, 21, 3, 4, 2, 0, 20, 18, 1, 0, 0, 0, 21, 24, 1,
		0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 26, 1, 0, 0, 0, 24,
		22, 1, 0, 0, 0, 25, 27, 5, 2, 0, 0, 26, 25, 1, 0, 0, 0, 26, 27, 1, 0, 0,
		0, 27, 29, 1, 0, 0, 0, 28, 17, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30,
		1, 0, 0, 0, 30, 31, 5, 3, 0, 0, 31, 1, 1, 0, 0, 0, 32, 33, 5, 12, 0, 0,
		33, 34, 5, 12, 0, 0, 34, 35, 3, 0, 0, 0, 35, 36, 3, 12, 6, 0, 36, 3, 1,
		0, 0, 0, 37, 38, 5, 12, 0, 0, 38, 39, 5, 12, 0, 0, 39, 5, 1, 0, 0, 0, 40,
		41, 3, 4, 2, 0, 41, 42, 5, 4, 0, 0, 42, 43, 3, 10, 5, 0, 43, 7, 1, 0, 0,
		0, 44, 45, 5, 5, 0, 0, 45, 46, 5, 12, 0, 0, 46, 51, 3, 10, 5, 0, 47, 48,
		5, 2, 0, 0, 48, 50, 3, 10, 5, 0, 49, 47, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1,
		0, 0, 0, 54, 55, 5, 6, 0, 0, 55, 9, 1, 0, 0, 0, 56, 61, 7, 0, 0, 0, 57,
		58, 5, 13, 0, 0, 58, 60, 3, 10, 5, 0, 59, 57, 1, 0, 0, 0, 60, 63, 1, 0,
		0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 11, 1, 0, 0, 0, 63, 61,
		1, 0, 0, 0, 64, 74, 5, 7, 0, 0, 65, 69, 3, 10, 5, 0, 66, 69, 3, 4, 2, 0,
		67, 69, 3, 6, 3, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67, 1,
		0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 5, 6, 0, 0, 71, 73, 1, 0, 0, 0, 72,
		68, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0,
		0, 75, 77, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 77, 78, 5, 8, 0, 0, 78, 13,
		1, 0, 0, 0, 79, 82, 3, 8, 4, 0, 80, 82, 3, 2, 1, 0, 81, 79, 1, 0, 0, 0,
		81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1,
		0, 0, 0, 84, 15, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 9, 22, 26, 28, 51, 61,
		68, 74, 81, 83,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// WystParserInit initializes any static state used to implement WystParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWystParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WystParserInit() {
	staticData := &WystParserStaticData
	staticData.once.Do(wystParserInit)
}

// NewWystParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWystParser(input antlr.TokenStream) *WystParser {
	WystParserInit()
	this := new(WystParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &WystParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Wyst.g4"

	return this
}

// WystParser tokens.
const (
	WystParserEOF        = antlr.TokenEOF
	WystParserT__0       = 1
	WystParserT__1       = 2
	WystParserT__2       = 3
	WystParserT__3       = 4
	WystParserT__4       = 5
	WystParserT__5       = 6
	WystParserT__6       = 7
	WystParserT__7       = 8
	WystParserWS         = 9
	WystParserNUMBER     = 10
	WystParserHEX        = 11
	WystParserIDENTIFIER = 12
	WystParserMATH       = 13
	WystParserSTRING     = 14
)

// WystParser rules.
const (
	WystParserRULE_round_def   = 0
	WystParserRULE_func_def    = 1
	WystParserRULE_var_def     = 2
	WystParserRULE_var_def_set = 3
	WystParserRULE_asm         = 4
	WystParserRULE_expr        = 5
	WystParserRULE_code_block  = 6
	WystParserRULE_top         = 7
)

// IRound_defContext is an interface to support dynamic dispatch.
type IRound_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_def() []IVar_defContext
	Var_def(i int) IVar_defContext

	// IsRound_defContext differentiates from other interfaces.
	IsRound_defContext()
}

type Round_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRound_defContext() *Round_defContext {
	var p = new(Round_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_round_def
	return p
}

func InitEmptyRound_defContext(p *Round_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_round_def
}

func (*Round_defContext) IsRound_defContext() {}

func NewRound_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Round_defContext {
	var p = new(Round_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_round_def

	return p
}

func (s *Round_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Round_defContext) AllVar_def() []IVar_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_defContext); ok {
			len++
		}
	}

	tst := make([]IVar_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_defContext); ok {
			tst[i] = t.(IVar_defContext)
			i++
		}
	}

	return tst
}

func (s *Round_defContext) Var_def(i int) IVar_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_defContext)
}

func (s *Round_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Round_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Round_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterRound_def(s)
	}
}

func (s *Round_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitRound_def(s)
	}
}

func (p *WystParser) Round_def() (localctx IRound_defContext) {
	localctx = NewRound_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WystParserRULE_round_def)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WystParserIDENTIFIER {
		{
			p.SetState(17)
			p.Var_def()
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(18)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(19)
					p.Var_def()
				}

			}
			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(25)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(30)
		p.Match(WystParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_defContext is an interface to support dynamic dispatch.
type IFunc_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	Round_def() IRound_defContext
	Code_block() ICode_blockContext

	// IsFunc_defContext differentiates from other interfaces.
	IsFunc_defContext()
}

type Func_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_defContext() *Func_defContext {
	var p = new(Func_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_func_def
	return p
}

func InitEmptyFunc_defContext(p *Func_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_func_def
}

func (*Func_defContext) IsFunc_defContext() {}

func NewFunc_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_defContext {
	var p = new(Func_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_func_def

	return p
}

func (s *Func_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_defContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(WystParserIDENTIFIER)
}

func (s *Func_defContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, i)
}

func (s *Func_defContext) Round_def() IRound_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRound_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRound_defContext)
}

func (s *Func_defContext) Code_block() ICode_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICode_blockContext)
}

func (s *Func_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterFunc_def(s)
	}
}

func (s *Func_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitFunc_def(s)
	}
}

func (p *WystParser) Func_def() (localctx IFunc_defContext) {
	localctx = NewFunc_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WystParserRULE_func_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Round_def()
	}
	{
		p.SetState(35)
		p.Code_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_defContext is an interface to support dynamic dispatch.
type IVar_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsVar_defContext differentiates from other interfaces.
	IsVar_defContext()
}

type Var_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_defContext() *Var_defContext {
	var p = new(Var_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_var_def
	return p
}

func InitEmptyVar_defContext(p *Var_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_var_def
}

func (*Var_defContext) IsVar_defContext() {}

func NewVar_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_defContext {
	var p = new(Var_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_var_def

	return p
}

func (s *Var_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_defContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(WystParserIDENTIFIER)
}

func (s *Var_defContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, i)
}

func (s *Var_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterVar_def(s)
	}
}

func (s *Var_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitVar_def(s)
	}
}

func (p *WystParser) Var_def() (localctx IVar_defContext) {
	localctx = NewVar_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WystParserRULE_var_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_def_setContext is an interface to support dynamic dispatch.
type IVar_def_setContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_def() IVar_defContext
	Expr() IExprContext

	// IsVar_def_setContext differentiates from other interfaces.
	IsVar_def_setContext()
}

type Var_def_setContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_def_setContext() *Var_def_setContext {
	var p = new(Var_def_setContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_var_def_set
	return p
}

func InitEmptyVar_def_setContext(p *Var_def_setContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_var_def_set
}

func (*Var_def_setContext) IsVar_def_setContext() {}

func NewVar_def_setContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_def_setContext {
	var p = new(Var_def_setContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_var_def_set

	return p
}

func (s *Var_def_setContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_def_setContext) Var_def() IVar_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_defContext)
}

func (s *Var_def_setContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_def_setContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_def_setContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_def_setContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterVar_def_set(s)
	}
}

func (s *Var_def_setContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitVar_def_set(s)
	}
}

func (p *WystParser) Var_def_set() (localctx IVar_def_setContext) {
	localctx = NewVar_def_setContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, WystParserRULE_var_def_set)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Var_def()
	}
	{
		p.SetState(41)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsmContext is an interface to support dynamic dispatch.
type IAsmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsAsmContext differentiates from other interfaces.
	IsAsmContext()
}

type AsmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsmContext() *AsmContext {
	var p = new(AsmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_asm
	return p
}

func InitEmptyAsmContext(p *AsmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_asm
}

func (*AsmContext) IsAsmContext() {}

func NewAsmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsmContext {
	var p = new(AsmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_asm

	return p
}

func (s *AsmContext) GetParser() antlr.Parser { return s.parser }

func (s *AsmContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *AsmContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AsmContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterAsm(s)
	}
}

func (s *AsmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitAsm(s)
	}
}

func (p *WystParser) Asm() (localctx IAsmContext) {
	localctx = NewAsmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, WystParserRULE_asm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(WystParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Expr()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__1 {
		{
			p.SetState(47)
			p.Match(WystParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Expr()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(WystParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	HEX() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllMATH() []antlr.TerminalNode
	MATH(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(WystParserNUMBER, 0)
}

func (s *ExprContext) HEX() antlr.TerminalNode {
	return s.GetToken(WystParserHEX, 0)
}

func (s *ExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(WystParserSTRING, 0)
}

func (s *ExprContext) AllMATH() []antlr.TerminalNode {
	return s.GetTokens(WystParserMATH)
}

func (s *ExprContext) MATH(i int) antlr.TerminalNode {
	return s.GetToken(WystParserMATH, i)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *WystParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, WystParserRULE_expr)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&23552) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(57)
				p.Match(WystParserMATH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(58)
				p.Expr()
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICode_blockContext is an interface to support dynamic dispatch.
type ICode_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllVar_def() []IVar_defContext
	Var_def(i int) IVar_defContext
	AllVar_def_set() []IVar_def_setContext
	Var_def_set(i int) IVar_def_setContext

	// IsCode_blockContext differentiates from other interfaces.
	IsCode_blockContext()
}

type Code_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCode_blockContext() *Code_blockContext {
	var p = new(Code_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_code_block
	return p
}

func InitEmptyCode_blockContext(p *Code_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_code_block
}

func (*Code_blockContext) IsCode_blockContext() {}

func NewCode_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Code_blockContext {
	var p = new(Code_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_code_block

	return p
}

func (s *Code_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Code_blockContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Code_blockContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Code_blockContext) AllVar_def() []IVar_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_defContext); ok {
			len++
		}
	}

	tst := make([]IVar_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_defContext); ok {
			tst[i] = t.(IVar_defContext)
			i++
		}
	}

	return tst
}

func (s *Code_blockContext) Var_def(i int) IVar_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_defContext)
}

func (s *Code_blockContext) AllVar_def_set() []IVar_def_setContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_def_setContext); ok {
			len++
		}
	}

	tst := make([]IVar_def_setContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_def_setContext); ok {
			tst[i] = t.(IVar_def_setContext)
			i++
		}
	}

	return tst
}

func (s *Code_blockContext) Var_def_set(i int) IVar_def_setContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_def_setContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_def_setContext)
}

func (s *Code_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Code_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Code_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterCode_block(s)
	}
}

func (s *Code_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitCode_block(s)
	}
}

func (p *WystParser) Code_block() (localctx ICode_blockContext) {
	localctx = NewCode_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WystParserRULE_code_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(WystParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&23552) != 0 {
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(65)
				p.Expr()
			}

		case 2:
			{
				p.SetState(66)
				p.Var_def()
			}

		case 3:
			{
				p.SetState(67)
				p.Var_def_set()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(70)
			p.Match(WystParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(WystParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopContext is an interface to support dynamic dispatch.
type ITopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAsm() []IAsmContext
	Asm(i int) IAsmContext
	AllFunc_def() []IFunc_defContext
	Func_def(i int) IFunc_defContext

	// IsTopContext differentiates from other interfaces.
	IsTopContext()
}

type TopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopContext() *TopContext {
	var p = new(TopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_top
	return p
}

func InitEmptyTopContext(p *TopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_top
}

func (*TopContext) IsTopContext() {}

func NewTopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopContext {
	var p = new(TopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_top

	return p
}

func (s *TopContext) GetParser() antlr.Parser { return s.parser }

func (s *TopContext) AllAsm() []IAsmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAsmContext); ok {
			len++
		}
	}

	tst := make([]IAsmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAsmContext); ok {
			tst[i] = t.(IAsmContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Asm(i int) IAsmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsmContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsmContext)
}

func (s *TopContext) AllFunc_def() []IFunc_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_defContext); ok {
			len++
		}
	}

	tst := make([]IFunc_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_defContext); ok {
			tst[i] = t.(IFunc_defContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Func_def(i int) IFunc_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_defContext)
}

func (s *TopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterTop(s)
	}
}

func (s *TopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitTop(s)
	}
}

func (p *WystParser) Top() (localctx ITopContext) {
	localctx = NewTopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, WystParserRULE_top)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__4 || _la == WystParserIDENTIFIER {
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case WystParserT__4:
			{
				p.SetState(79)
				p.Asm()
			}

		case WystParserIDENTIFIER:
			{
				p.SetState(80)
				p.Func_def()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
