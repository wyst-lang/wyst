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
		"", "'('", "','", "')'", "'{'", "';'", "'}'", "'='", "'.'", "'struct'",
		"'namespace'", "'include'", "'use'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "NUMBER",
		"HEX", "IDENTIFIER", "MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"round_def", "enum_curly", "round_call", "fn_call", "func_def", "var_def",
		"var_def_set", "call_tree", "struct_def", "namespace", "import_statement",
		"use_statement", "expr", "code_block", "top",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 171, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 5, 0, 35, 8, 0, 10, 0, 12, 0, 38, 9, 0, 1, 0, 3, 0, 41, 8,
		0, 3, 0, 43, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 51, 8, 1,
		10, 1, 12, 1, 54, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 62, 8,
		2, 10, 2, 12, 2, 65, 9, 2, 1, 2, 3, 2, 68, 8, 2, 3, 2, 70, 8, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 3, 7, 91, 8, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 96, 8, 7, 5, 7, 98, 8, 7, 10, 7, 12, 7, 101, 9, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 112, 8, 9, 1, 9, 1, 9, 1, 9,
		5, 9, 117, 8, 9, 10, 9, 12, 9, 120, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 137, 8, 12, 1, 12, 1, 12, 5, 12, 141, 8, 12, 10, 12, 12, 12, 144,
		9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 150, 8, 13, 1, 13, 1, 13, 5,
		13, 154, 8, 13, 10, 13, 12, 13, 157, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 5, 14, 166, 8, 14, 10, 14, 12, 14, 169, 9, 14, 1,
		14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0,
		0, 181, 0, 30, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 73,
		1, 0, 0, 0, 8, 76, 1, 0, 0, 0, 10, 81, 1, 0, 0, 0, 12, 84, 1, 0, 0, 0,
		14, 90, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 106, 1, 0, 0, 0, 20, 123,
		1, 0, 0, 0, 22, 127, 1, 0, 0, 0, 24, 136, 1, 0, 0, 0, 26, 145, 1, 0, 0,
		0, 28, 167, 1, 0, 0, 0, 30, 42, 5, 1, 0, 0, 31, 36, 3, 10, 5, 0, 32, 33,
		5, 2, 0, 0, 33, 35, 3, 10, 5, 0, 34, 32, 1, 0, 0, 0, 35, 38, 1, 0, 0, 0,
		36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1,
		0, 0, 0, 39, 41, 5, 2, 0, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41,
		43, 1, 0, 0, 0, 42, 31, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 45, 5, 3, 0, 0, 45, 1, 1, 0, 0, 0, 46, 52, 5, 4, 0, 0, 47, 48, 3,
		10, 5, 0, 48, 49, 5, 5, 0, 0, 49, 51, 1, 0, 0, 0, 50, 47, 1, 0, 0, 0, 51,
		54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 55, 56, 5, 6, 0, 0, 56, 3, 1, 0, 0, 0, 57, 69, 5,
		1, 0, 0, 58, 63, 3, 24, 12, 0, 59, 60, 5, 2, 0, 0, 60, 62, 3, 24, 12, 0,
		61, 59, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1,
		0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 66, 68, 5, 2, 0, 0, 67,
		66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 58, 1, 0, 0,
		0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 3, 0, 0, 72, 5, 1,
		0, 0, 0, 73, 74, 5, 16, 0, 0, 74, 75, 3, 4, 2, 0, 75, 7, 1, 0, 0, 0, 76,
		77, 5, 16, 0, 0, 77, 78, 5, 16, 0, 0, 78, 79, 3, 0, 0, 0, 79, 80, 3, 26,
		13, 0, 80, 9, 1, 0, 0, 0, 81, 82, 5, 16, 0, 0, 82, 83, 5, 16, 0, 0, 83,
		11, 1, 0, 0, 0, 84, 85, 3, 10, 5, 0, 85, 86, 5, 7, 0, 0, 86, 87, 3, 24,
		12, 0, 87, 13, 1, 0, 0, 0, 88, 91, 3, 6, 3, 0, 89, 91, 5, 16, 0, 0, 90,
		88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 99, 1, 0, 0, 0, 92, 95, 5, 8, 0,
		0, 93, 96, 3, 6, 3, 0, 94, 96, 5, 16, 0, 0, 95, 93, 1, 0, 0, 0, 95, 94,
		1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0,
		99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 15, 1, 0, 0, 0, 101, 99,
		1, 0, 0, 0, 102, 103, 5, 9, 0, 0, 103, 104, 5, 16, 0, 0, 104, 105, 3, 2,
		1, 0, 105, 17, 1, 0, 0, 0, 106, 107, 5, 10, 0, 0, 107, 108, 5, 16, 0, 0,
		108, 118, 5, 4, 0, 0, 109, 112, 3, 10, 5, 0, 110, 112, 3, 12, 6, 0, 111,
		109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 114,
		5, 5, 0, 0, 114, 117, 1, 0, 0, 0, 115, 117, 3, 8, 4, 0, 116, 111, 1, 0,
		0, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121,
		122, 5, 6, 0, 0, 122, 19, 1, 0, 0, 0, 123, 124, 5, 11, 0, 0, 124, 125,
		5, 16, 0, 0, 125, 126, 5, 5, 0, 0, 126, 21, 1, 0, 0, 0, 127, 128, 5, 12,
		0, 0, 128, 129, 5, 16, 0, 0, 129, 130, 5, 5, 0, 0, 130, 23, 1, 0, 0, 0,
		131, 137, 3, 14, 7, 0, 132, 137, 5, 14, 0, 0, 133, 137, 5, 15, 0, 0, 134,
		137, 5, 16, 0, 0, 135, 137, 5, 18, 0, 0, 136, 131, 1, 0, 0, 0, 136, 132,
		1, 0, 0, 0, 136, 133, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0,
		0, 0, 137, 142, 1, 0, 0, 0, 138, 139, 5, 17, 0, 0, 139, 141, 3, 24, 12,
		0, 140, 138, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 25, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0, 145, 155, 5,
		4, 0, 0, 146, 150, 3, 24, 12, 0, 147, 150, 3, 10, 5, 0, 148, 150, 3, 12,
		6, 0, 149, 146, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 148, 1, 0, 0, 0,
		150, 151, 1, 0, 0, 0, 151, 152, 5, 5, 0, 0, 152, 154, 1, 0, 0, 0, 153,
		149, 1, 0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 6,
		0, 0, 159, 27, 1, 0, 0, 0, 160, 166, 3, 8, 4, 0, 161, 166, 3, 16, 8, 0,
		162, 166, 3, 18, 9, 0, 163, 166, 3, 20, 10, 0, 164, 166, 3, 22, 11, 0,
		165, 160, 1, 0, 0, 0, 165, 161, 1, 0, 0, 0, 165, 162, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 29, 1, 0, 0, 0, 169, 167, 1, 0,
		0, 0, 19, 36, 40, 42, 52, 63, 67, 69, 90, 95, 99, 111, 116, 118, 136, 142,
		149, 155, 165, 167,
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
	WystParserT__9       = 10
	WystParserT__10      = 11
	WystParserT__11      = 12
	WystParserWS         = 13
	WystParserNUMBER     = 14
	WystParserHEX        = 15
	WystParserIDENTIFIER = 16
	WystParserMATH       = 17
	WystParserSTRING     = 18
)

// WystParser rules.
const (
	WystParserRULE_round_def        = 0
	WystParserRULE_enum_curly       = 1
	WystParserRULE_round_call       = 2
	WystParserRULE_fn_call          = 3
	WystParserRULE_func_def         = 4
	WystParserRULE_var_def          = 5
	WystParserRULE_var_def_set      = 6
	WystParserRULE_call_tree        = 7
	WystParserRULE_struct_def       = 8
	WystParserRULE_namespace        = 9
	WystParserRULE_import_statement = 10
	WystParserRULE_use_statement    = 11
	WystParserRULE_expr             = 12
	WystParserRULE_code_block       = 13
	WystParserRULE_top              = 14
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
		p.SetState(30)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WystParserIDENTIFIER {
		{
			p.SetState(31)
			p.Var_def()
		}
		p.SetState(36)
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
					p.SetState(32)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(33)
					p.Var_def()
				}

			}
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(39)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(44)
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

// IEnum_curlyContext is an interface to support dynamic dispatch.
type IEnum_curlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_def() []IVar_defContext
	Var_def(i int) IVar_defContext

	// IsEnum_curlyContext differentiates from other interfaces.
	IsEnum_curlyContext()
}

type Enum_curlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_curlyContext() *Enum_curlyContext {
	var p = new(Enum_curlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_enum_curly
	return p
}

func InitEmptyEnum_curlyContext(p *Enum_curlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_enum_curly
}

func (*Enum_curlyContext) IsEnum_curlyContext() {}

func NewEnum_curlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_curlyContext {
	var p = new(Enum_curlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_enum_curly

	return p
}

func (s *Enum_curlyContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_curlyContext) AllVar_def() []IVar_defContext {
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

func (s *Enum_curlyContext) Var_def(i int) IVar_defContext {
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

func (s *Enum_curlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_curlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_curlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterEnum_curly(s)
	}
}

func (s *Enum_curlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitEnum_curly(s)
	}
}

func (p *WystParser) Enum_curly() (localctx IEnum_curlyContext) {
	localctx = NewEnum_curlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WystParserRULE_enum_curly)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserIDENTIFIER {
		{
			p.SetState(47)
			p.Var_def()
		}
		{
			p.SetState(48)
			p.Match(WystParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
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
	p.EnterRule(localctx, 4, WystParserRULE_round_call)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&376832) != 0 {
		{
			p.SetState(58)
			p.Expr()
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
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(59)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(60)
					p.Expr()
				}

			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(66)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(71)
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
	p.EnterRule(localctx, 6, WystParserRULE_fn_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
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
	p.EnterRule(localctx, 8, WystParserRULE_func_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Round_def()
	}
	{
		p.SetState(79)
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
	p.EnterRule(localctx, 10, WystParserRULE_var_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
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
	p.EnterRule(localctx, 12, WystParserRULE_var_def_set)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Var_def()
	}
	{
		p.SetState(85)
		p.Match(WystParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
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
	p.EnterRule(localctx, 14, WystParserRULE_call_tree)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(88)
			p.Fn_call()
		}

	case 2:
		{
			p.SetState(89)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__7 {
		{
			p.SetState(92)
			p.Match(WystParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(93)
				p.Fn_call()
			}

		case 2:
			{
				p.SetState(94)
				p.Match(WystParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(101)
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

// IStruct_defContext is an interface to support dynamic dispatch.
type IStruct_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Enum_curly() IEnum_curlyContext

	// IsStruct_defContext differentiates from other interfaces.
	IsStruct_defContext()
}

type Struct_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_defContext() *Struct_defContext {
	var p = new(Struct_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_struct_def
	return p
}

func InitEmptyStruct_defContext(p *Struct_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_struct_def
}

func (*Struct_defContext) IsStruct_defContext() {}

func NewStruct_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_defContext {
	var p = new(Struct_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_struct_def

	return p
}

func (s *Struct_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_defContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *Struct_defContext) Enum_curly() IEnum_curlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_curlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_curlyContext)
}

func (s *Struct_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterStruct_def(s)
	}
}

func (s *Struct_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitStruct_def(s)
	}
}

func (p *WystParser) Struct_def() (localctx IStruct_defContext) {
	localctx = NewStruct_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, WystParserRULE_struct_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(WystParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Enum_curly()
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

// INamespaceContext is an interface to support dynamic dispatch.
type INamespaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllFunc_def() []IFunc_defContext
	Func_def(i int) IFunc_defContext
	AllVar_def() []IVar_defContext
	Var_def(i int) IVar_defContext
	AllVar_def_set() []IVar_def_setContext
	Var_def_set(i int) IVar_def_setContext

	// IsNamespaceContext differentiates from other interfaces.
	IsNamespaceContext()
}

type NamespaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceContext() *NamespaceContext {
	var p = new(NamespaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_namespace
	return p
}

func InitEmptyNamespaceContext(p *NamespaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_namespace
}

func (*NamespaceContext) IsNamespaceContext() {}

func NewNamespaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceContext {
	var p = new(NamespaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_namespace

	return p
}

func (s *NamespaceContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *NamespaceContext) AllFunc_def() []IFunc_defContext {
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

func (s *NamespaceContext) Func_def(i int) IFunc_defContext {
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

func (s *NamespaceContext) AllVar_def() []IVar_defContext {
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

func (s *NamespaceContext) Var_def(i int) IVar_defContext {
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

func (s *NamespaceContext) AllVar_def_set() []IVar_def_setContext {
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

func (s *NamespaceContext) Var_def_set(i int) IVar_def_setContext {
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

func (s *NamespaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterNamespace(s)
	}
}

func (s *NamespaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitNamespace(s)
	}
}

func (p *WystParser) Namespace() (localctx INamespaceContext) {
	localctx = NewNamespaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, WystParserRULE_namespace)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(WystParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserIDENTIFIER {
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(109)
					p.Var_def()
				}

			case 2:
				{
					p.SetState(110)
					p.Var_def_set()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			{
				p.SetState(113)
				p.Match(WystParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(115)
				p.Func_def()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
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

// IImport_statementContext is an interface to support dynamic dispatch.
type IImport_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsImport_statementContext differentiates from other interfaces.
	IsImport_statementContext()
}

type Import_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_statementContext() *Import_statementContext {
	var p = new(Import_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_import_statement
	return p
}

func InitEmptyImport_statementContext(p *Import_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_import_statement
}

func (*Import_statementContext) IsImport_statementContext() {}

func NewImport_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_statementContext {
	var p = new(Import_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_import_statement

	return p
}

func (s *Import_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_statementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *Import_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterImport_statement(s)
	}
}

func (s *Import_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitImport_statement(s)
	}
}

func (p *WystParser) Import_statement() (localctx IImport_statementContext) {
	localctx = NewImport_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, WystParserRULE_import_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(WystParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(WystParserT__4)
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

// IUse_statementContext is an interface to support dynamic dispatch.
type IUse_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsUse_statementContext differentiates from other interfaces.
	IsUse_statementContext()
}

type Use_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUse_statementContext() *Use_statementContext {
	var p = new(Use_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_use_statement
	return p
}

func InitEmptyUse_statementContext(p *Use_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_use_statement
}

func (*Use_statementContext) IsUse_statementContext() {}

func NewUse_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Use_statementContext {
	var p = new(Use_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_use_statement

	return p
}

func (s *Use_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Use_statementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, 0)
}

func (s *Use_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Use_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Use_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterUse_statement(s)
	}
}

func (s *Use_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitUse_statement(s)
	}
}

func (p *WystParser) Use_statement() (localctx IUse_statementContext) {
	localctx = NewUse_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, WystParserRULE_use_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(WystParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(WystParserT__4)
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
	p.EnterRule(localctx, 24, WystParserRULE_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(131)
			p.Call_tree()
		}

	case 2:
		{
			p.SetState(132)
			p.Match(WystParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(133)
			p.Match(WystParserHEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(134)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(135)
			p.Match(WystParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(138)
				p.Match(WystParserMATH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(139)
				p.Expr()
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 26, WystParserRULE_code_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&376832) != 0 {
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(146)
				p.Expr()
			}

		case 2:
			{
				p.SetState(147)
				p.Var_def()
			}

		case 3:
			{
				p.SetState(148)
				p.Var_def_set()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(151)
			p.Match(WystParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
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

// ITopContext is an interface to support dynamic dispatch.
type ITopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFunc_def() []IFunc_defContext
	Func_def(i int) IFunc_defContext
	AllStruct_def() []IStruct_defContext
	Struct_def(i int) IStruct_defContext
	AllNamespace() []INamespaceContext
	Namespace(i int) INamespaceContext
	AllImport_statement() []IImport_statementContext
	Import_statement(i int) IImport_statementContext
	AllUse_statement() []IUse_statementContext
	Use_statement(i int) IUse_statementContext

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

func (s *TopContext) AllStruct_def() []IStruct_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_defContext); ok {
			len++
		}
	}

	tst := make([]IStruct_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_defContext); ok {
			tst[i] = t.(IStruct_defContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Struct_def(i int) IStruct_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_defContext); ok {
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

	return t.(IStruct_defContext)
}

func (s *TopContext) AllNamespace() []INamespaceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INamespaceContext); ok {
			len++
		}
	}

	tst := make([]INamespaceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INamespaceContext); ok {
			tst[i] = t.(INamespaceContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Namespace(i int) INamespaceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceContext); ok {
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

	return t.(INamespaceContext)
}

func (s *TopContext) AllImport_statement() []IImport_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_statementContext); ok {
			len++
		}
	}

	tst := make([]IImport_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_statementContext); ok {
			tst[i] = t.(IImport_statementContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Import_statement(i int) IImport_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_statementContext); ok {
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

	return t.(IImport_statementContext)
}

func (s *TopContext) AllUse_statement() []IUse_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUse_statementContext); ok {
			len++
		}
	}

	tst := make([]IUse_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUse_statementContext); ok {
			tst[i] = t.(IUse_statementContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Use_statement(i int) IUse_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUse_statementContext); ok {
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

	return t.(IUse_statementContext)
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
	p.EnterRule(localctx, 28, WystParserRULE_top)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&73216) != 0 {
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case WystParserIDENTIFIER:
			{
				p.SetState(160)
				p.Func_def()
			}

		case WystParserT__8:
			{
				p.SetState(161)
				p.Struct_def()
			}

		case WystParserT__9:
			{
				p.SetState(162)
				p.Namespace()
			}

		case WystParserT__10:
			{
				p.SetState(163)
				p.Import_statement()
			}

		case WystParserT__11:
			{
				p.SetState(164)
				p.Use_statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(169)
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
