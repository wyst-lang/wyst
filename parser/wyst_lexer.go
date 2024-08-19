// Code generated from Wyst.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type WystLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var WystLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func wystlexerLexerInit() {
	staticData := &WystLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "','", "')'", "'='", "'.'", "'%'", "';'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "WS", "NUMBER", "HEX", "IDENTIFIER",
		"MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"WS", "NUMBER", "HEX", "IDENTIFIER", "MATH", "ESC", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 101, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 4, 9, 53, 8, 9, 11, 9,
		12, 9, 54, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 61, 8, 10, 10, 10, 12, 10,
		64, 9, 10, 1, 10, 3, 10, 67, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11,
		73, 8, 11, 10, 11, 12, 11, 76, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12,
		82, 8, 12, 10, 12, 12, 12, 85, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 5, 15, 95, 8, 15, 10, 15, 12, 15, 98, 9, 15, 1, 15,
		1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 0, 31, 15, 1, 0, 9, 3, 0,
		9, 10, 13, 13, 32, 32, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 48, 57, 65, 70,
		97, 102, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 3, 0, 42, 43, 45, 45, 47, 47, 3, 0, 34, 34, 39, 39, 92, 92, 2, 0,
		34, 34, 92, 92, 107, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 31,
		1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 35, 1, 0, 0, 0, 5, 37, 1, 0, 0, 0, 7,
		39, 1, 0, 0, 0, 9, 41, 1, 0, 0, 0, 11, 43, 1, 0, 0, 0, 13, 45, 1, 0, 0,
		0, 15, 47, 1, 0, 0, 0, 17, 49, 1, 0, 0, 0, 19, 52, 1, 0, 0, 0, 21, 66,
		1, 0, 0, 0, 23, 68, 1, 0, 0, 0, 25, 77, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0,
		29, 88, 1, 0, 0, 0, 31, 91, 1, 0, 0, 0, 33, 34, 5, 40, 0, 0, 34, 2, 1,
		0, 0, 0, 35, 36, 5, 44, 0, 0, 36, 4, 1, 0, 0, 0, 37, 38, 5, 41, 0, 0, 38,
		6, 1, 0, 0, 0, 39, 40, 5, 61, 0, 0, 40, 8, 1, 0, 0, 0, 41, 42, 5, 46, 0,
		0, 42, 10, 1, 0, 0, 0, 43, 44, 5, 37, 0, 0, 44, 12, 1, 0, 0, 0, 45, 46,
		5, 59, 0, 0, 46, 14, 1, 0, 0, 0, 47, 48, 5, 123, 0, 0, 48, 16, 1, 0, 0,
		0, 49, 50, 5, 125, 0, 0, 50, 18, 1, 0, 0, 0, 51, 53, 7, 0, 0, 0, 52, 51,
		1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 6, 9, 0, 0, 57, 20, 1, 0, 0, 0, 58, 62, 7,
		1, 0, 0, 59, 61, 7, 2, 0, 0, 60, 59, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62,
		60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 67, 1, 0, 0, 0, 64, 62, 1, 0, 0,
		0, 65, 67, 5, 48, 0, 0, 66, 58, 1, 0, 0, 0, 66, 65, 1, 0, 0, 0, 67, 22,
		1, 0, 0, 0, 68, 69, 5, 48, 0, 0, 69, 70, 5, 120, 0, 0, 70, 74, 1, 0, 0,
		0, 71, 73, 7, 3, 0, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 24, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0,
		77, 83, 7, 4, 0, 0, 78, 82, 7, 5, 0, 0, 79, 80, 5, 58, 0, 0, 80, 82, 5,
		58, 0, 0, 81, 78, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83,
		81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 26, 1, 0, 0, 0, 85, 83, 1, 0, 0,
		0, 86, 87, 7, 6, 0, 0, 87, 28, 1, 0, 0, 0, 88, 89, 5, 92, 0, 0, 89, 90,
		7, 7, 0, 0, 90, 30, 1, 0, 0, 0, 91, 96, 5, 34, 0, 0, 92, 95, 3, 29, 14,
		0, 93, 95, 8, 8, 0, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 98,
		1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 1, 0, 0, 0,
		98, 96, 1, 0, 0, 0, 99, 100, 5, 34, 0, 0, 100, 32, 1, 0, 0, 0, 9, 0, 54,
		62, 66, 74, 81, 83, 94, 96, 1, 6, 0, 0,
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

// WystLexerInit initializes any static state used to implement WystLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewWystLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func WystLexerInit() {
	staticData := &WystLexerLexerStaticData
	staticData.once.Do(wystlexerLexerInit)
}

// NewWystLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewWystLexer(input antlr.CharStream) *WystLexer {
	WystLexerInit()
	l := new(WystLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &WystLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Wyst.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// WystLexer tokens.
const (
	WystLexerT__0       = 1
	WystLexerT__1       = 2
	WystLexerT__2       = 3
	WystLexerT__3       = 4
	WystLexerT__4       = 5
	WystLexerT__5       = 6
	WystLexerT__6       = 7
	WystLexerT__7       = 8
	WystLexerT__8       = 9
	WystLexerWS         = 10
	WystLexerNUMBER     = 11
	WystLexerHEX        = 12
	WystLexerIDENTIFIER = 13
	WystLexerMATH       = 14
	WystLexerSTRING     = 15
)
