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
		"", "'('", "','", "')'", "'='", "'.'", "'%'", "';'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "WS", "NUMBER", "HEX", "IDENTIFIER",
		"MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"round_def", "round_call", "fn_call", "func_def", "var_def", "var_def_set",
		"call_tree", "asm", "expr", "code_block", "top",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 132, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 27, 8, 0, 10, 0, 12, 0, 30, 9, 0, 1,
		0, 3, 0, 33, 8, 0, 3, 0, 35, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 43, 8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 3, 1, 49, 8, 1, 3, 1, 51,
		8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 72, 8, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 77, 8, 6, 5, 6, 79, 8, 6, 10, 6, 12, 6, 82, 9, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 101, 8, 8, 1, 8, 1, 8, 5, 8, 105,
		8, 8, 10, 8, 12, 8, 108, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 114, 8, 9,
		1, 9, 1, 9, 5, 9, 118, 8, 9, 10, 9, 12, 9, 121, 9, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 5, 10, 127, 8, 10, 10, 10, 12, 10, 130, 9, 10, 1, 10, 0, 0, 11,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 0, 140, 0, 22, 1, 0, 0, 0, 2,
		38, 1, 0, 0, 0, 4, 54, 1, 0, 0, 0, 6, 57, 1, 0, 0, 0, 8, 62, 1, 0, 0, 0,
		10, 65, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0, 16, 100, 1,
		0, 0, 0, 18, 109, 1, 0, 0, 0, 20, 128, 1, 0, 0, 0, 22, 34, 5, 1, 0, 0,
		23, 28, 3, 8, 4, 0, 24, 25, 5, 2, 0, 0, 25, 27, 3, 8, 4, 0, 26, 24, 1,
		0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29,
		32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 33, 5, 2, 0, 0, 32, 31, 1, 0, 0,
		0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 23, 1, 0, 0, 0, 34, 35,
		1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 5, 3, 0, 0, 37, 1, 1, 0, 0, 0,
		38, 50, 5, 1, 0, 0, 39, 44, 3, 16, 8, 0, 40, 41, 5, 2, 0, 0, 41, 43, 3,
		16, 8, 0, 42, 40, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44,
		45, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 49, 5, 2, 0,
		0, 48, 47, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 39,
		1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 5, 3, 0, 0,
		53, 3, 1, 0, 0, 0, 54, 55, 5, 13, 0, 0, 55, 56, 3, 2, 1, 0, 56, 5, 1, 0,
		0, 0, 57, 58, 5, 13, 0, 0, 58, 59, 5, 13, 0, 0, 59, 60, 3, 0, 0, 0, 60,
		61, 3, 18, 9, 0, 61, 7, 1, 0, 0, 0, 62, 63, 5, 13, 0, 0, 63, 64, 5, 13,
		0, 0, 64, 9, 1, 0, 0, 0, 65, 66, 3, 8, 4, 0, 66, 67, 5, 4, 0, 0, 67, 68,
		3, 16, 8, 0, 68, 11, 1, 0, 0, 0, 69, 72, 3, 4, 2, 0, 70, 72, 5, 13, 0,
		0, 71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0, 0, 72, 80, 1, 0, 0, 0, 73, 76,
		5, 5, 0, 0, 74, 77, 3, 4, 2, 0, 75, 77, 5, 13, 0, 0, 76, 74, 1, 0, 0, 0,
		76, 75, 1, 0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 73, 1, 0, 0, 0, 79, 82, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 13, 1, 0, 0, 0, 82,
		80, 1, 0, 0, 0, 83, 84, 5, 6, 0, 0, 84, 85, 5, 13, 0, 0, 85, 90, 3, 16,
		8, 0, 86, 87, 5, 2, 0, 0, 87, 89, 3, 16, 8, 0, 88, 86, 1, 0, 0, 0, 89,
		92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0,
		0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 7, 0, 0, 94, 15, 1, 0, 0, 0, 95, 101,
		3, 12, 6, 0, 96, 101, 5, 11, 0, 0, 97, 101, 5, 12, 0, 0, 98, 101, 5, 13,
		0, 0, 99, 101, 5, 15, 0, 0, 100, 95, 1, 0, 0, 0, 100, 96, 1, 0, 0, 0, 100,
		97, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 106, 1,
		0, 0, 0, 102, 103, 5, 14, 0, 0, 103, 105, 3, 16, 8, 0, 104, 102, 1, 0,
		0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 17, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 119, 5, 8, 0, 0, 110, 114,
		3, 16, 8, 0, 111, 114, 3, 8, 4, 0, 112, 114, 3, 10, 5, 0, 113, 110, 1,
		0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0,
		0, 115, 116, 5, 7, 0, 0, 116, 118, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 118,
		121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122,
		1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 9, 0, 0, 123, 19, 1, 0,
		0, 0, 124, 127, 3, 14, 7, 0, 125, 127, 3, 6, 3, 0, 126, 124, 1, 0, 0, 0,
		126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128,
		129, 1, 0, 0, 0, 129, 21, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 16, 28, 32,
		34, 44, 48, 50, 71, 76, 80, 90, 100, 106, 113, 119, 126, 128,
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
	WystParserT__8       = 9
	WystParserWS         = 10
	WystParserNUMBER     = 11
	WystParserHEX        = 12
	WystParserIDENTIFIER = 13
	WystParserMATH       = 14
	WystParserSTRING     = 15
)

// WystParser rules.
const (
	WystParserRULE_round_def   = 0
	WystParserRULE_round_call  = 1
	WystParserRULE_fn_call     = 2
	WystParserRULE_func_def    = 3
	WystParserRULE_var_def     = 4
	WystParserRULE_var_def_set = 5
	WystParserRULE_call_tree   = 6
	WystParserRULE_asm         = 7
	WystParserRULE_expr        = 8
	WystParserRULE_code_block  = 9
	WystParserRULE_top         = 10
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
		p.SetState(22)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WystParserIDENTIFIER {
		{
			p.SetState(23)
			p.Var_def()
		}
		p.SetState(28)
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
					p.SetState(24)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(25)
					p.Var_def()
				}

			}
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(31)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(36)
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

// IRound_callContext is an interface to support dynamic dispatch.
type IRound_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsRound_callContext differentiates from other interfaces.
	IsRound_callContext()
}

type Round_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRound_callContext() *Round_callContext {
	var p = new(Round_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_round_call
	return p
}

func InitEmptyRound_callContext(p *Round_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_round_call
}

func (*Round_callContext) IsRound_callContext() {}

func NewRound_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Round_callContext {
	var p = new(Round_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_round_call

	return p
}

func (s *Round_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Round_callContext) AllExpr() []IExprContext {
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

func (s *Round_callContext) Expr(i int) IExprContext {
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

func (s *Round_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Round_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Round_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterRound_call(s)
	}
}

func (s *Round_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitRound_call(s)
	}
}

func (p *WystParser) Round_call() (localctx IRound_callContext) {
	localctx = NewRound_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WystParserRULE_round_call)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&47104) != 0 {
		{
			p.SetState(39)
			p.Expr()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(40)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(41)
					p.Expr()
				}

			}
			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(47)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(52)
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

// IFn_callContext is an interface to support dynamic dispatch.
type IFn_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Round_call() IRound_callContext

	// IsFn_callContext differentiates from other interfaces.
	IsFn_callContext()
}

type Fn_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_callContext() *Fn_callContext {
	var p = new(Fn_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_fn_call
	return p
}

func InitEmptyFn_callContext(p *Fn_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_fn_call
}

func (*Fn_callContext) IsFn_callContext() {}

func NewFn_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_callContext {
	var p = new(Fn_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_fn_call

	return p
}

func (s *Fn_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_callContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *Fn_callContext) Round_call() IRound_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRound_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRound_callContext)
}

func (s *Fn_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterFn_call(s)
	}
}

func (s *Fn_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitFn_call(s)
	}
}

func (p *WystParser) Fn_call() (localctx IFn_callContext) {
	localctx = NewFn_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WystParserRULE_fn_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Round_call()
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
	p.EnterRule(localctx, 6, WystParserRULE_func_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Round_def()
	}
	{
		p.SetState(60)
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
	p.EnterRule(localctx, 8, WystParserRULE_var_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
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
	p.EnterRule(localctx, 10, WystParserRULE_var_def_set)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Var_def()
	}
	{
		p.SetState(66)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
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

// ICall_treeContext is an interface to support dynamic dispatch.
type ICall_treeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFn_call() []IFn_callContext
	Fn_call(i int) IFn_callContext
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsCall_treeContext differentiates from other interfaces.
	IsCall_treeContext()
}

type Call_treeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_treeContext() *Call_treeContext {
	var p = new(Call_treeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_call_tree
	return p
}

func InitEmptyCall_treeContext(p *Call_treeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_call_tree
}

func (*Call_treeContext) IsCall_treeContext() {}

func NewCall_treeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_treeContext {
	var p = new(Call_treeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_call_tree

	return p
}

func (s *Call_treeContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_treeContext) AllFn_call() []IFn_callContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_callContext); ok {
			len++
		}
	}

	tst := make([]IFn_callContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_callContext); ok {
			tst[i] = t.(IFn_callContext)
			i++
		}
	}

	return tst
}

func (s *Call_treeContext) Fn_call(i int) IFn_callContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_callContext); ok {
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

	return t.(IFn_callContext)
}

func (s *Call_treeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(WystParserIDENTIFIER)
}

func (s *Call_treeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, i)
}

func (s *Call_treeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_treeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_treeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterCall_tree(s)
	}
}

func (s *Call_treeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitCall_tree(s)
	}
}

func (p *WystParser) Call_tree() (localctx ICall_treeContext) {
	localctx = NewCall_treeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, WystParserRULE_call_tree)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(69)
			p.Fn_call()
		}

	case 2:
		{
			p.SetState(70)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__4 {
		{
			p.SetState(73)
			p.Match(WystParserT__4)
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

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(74)
				p.Fn_call()
			}

		case 2:
			{
				p.SetState(75)
				p.Match(WystParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(82)
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
	p.EnterRule(localctx, 14, WystParserRULE_asm)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(WystParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.Expr()
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__1 {
		{
			p.SetState(86)
			p.Match(WystParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Expr()
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(WystParserT__6)
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
	Call_tree() ICall_treeContext
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

func (s *ExprContext) Call_tree() ICall_treeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_treeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_treeContext)
}

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
	p.EnterRule(localctx, 16, WystParserRULE_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(95)
			p.Call_tree()
		}

	case 2:
		{
			p.SetState(96)
			p.Match(WystParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(97)
			p.Match(WystParserHEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(98)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(99)
			p.Match(WystParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(102)
				p.Match(WystParserMATH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(103)
				p.Expr()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, WystParserRULE_code_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(WystParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&47104) != 0 {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(110)
				p.Expr()
			}

		case 2:
			{
				p.SetState(111)
				p.Var_def()
			}

		case 3:
			{
				p.SetState(112)
				p.Var_def_set()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(115)
			p.Match(WystParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
		p.Match(WystParserT__8)
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
	p.EnterRule(localctx, 20, WystParserRULE_top)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__5 || _la == WystParserIDENTIFIER {
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case WystParserT__5:
			{
				p.SetState(124)
				p.Asm()
			}

		case WystParserIDENTIFIER:
			{
				p.SetState(125)
				p.Func_def()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(130)
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
