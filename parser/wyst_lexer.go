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
		"", "'('", "','", "')'", "'='", "'%'", "';'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "WS", "NUMBER", "HEX", "IDENTIFIER",
		"MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "WS",
		"NUMBER", "HEX", "IDENTIFIER", "MATH", "ESC", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 97, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 4, 8, 49, 8, 8, 11, 8, 12, 8, 50, 1, 8, 1, 8, 1,
		9, 1, 9, 5, 9, 57, 8, 9, 10, 9, 12, 9, 60, 9, 9, 1, 9, 3, 9, 63, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 69, 8, 10, 10, 10, 12, 10, 72, 9, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 78, 8, 11, 10, 11, 12, 11, 81, 9, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 91, 8, 14,
		10, 14, 12, 14, 94, 9, 14, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		0, 29, 14, 1, 0, 9, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 49, 57, 1, 0, 48,
		57, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0,
		48, 57, 65, 90, 97, 122, 3, 0, 42, 43, 45, 45, 47, 47, 3, 0, 34, 34, 39,
		39, 92, 92, 2, 0, 34, 34, 92, 92, 103, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 1, 31, 1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 37,
		1, 0, 0, 0, 9, 39, 1, 0, 0, 0, 11, 41, 1, 0, 0, 0, 13, 43, 1, 0, 0, 0,
		15, 45, 1, 0, 0, 0, 17, 48, 1, 0, 0, 0, 19, 62, 1, 0, 0, 0, 21, 64, 1,
		0, 0, 0, 23, 73, 1, 0, 0, 0, 25, 82, 1, 0, 0, 0, 27, 84, 1, 0, 0, 0, 29,
		87, 1, 0, 0, 0, 31, 32, 5, 40, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 5, 44,
		0, 0, 34, 4, 1, 0, 0, 0, 35, 36, 5, 41, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38,
		5, 61, 0, 0, 38, 8, 1, 0, 0, 0, 39, 40, 5, 37, 0, 0, 40, 10, 1, 0, 0, 0,
		41, 42, 5, 59, 0, 0, 42, 12, 1, 0, 0, 0, 43, 44, 5, 123, 0, 0, 44, 14,
		1, 0, 0, 0, 45, 46, 5, 125, 0, 0, 46, 16, 1, 0, 0, 0, 47, 49, 7, 0, 0,
		0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 8, 0, 0, 53, 18, 1, 0, 0, 0,
		54, 58, 7, 1, 0, 0, 55, 57, 7, 2, 0, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1,
		0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 63, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 61, 63, 5, 48, 0, 0, 62, 54, 1, 0, 0, 0, 62, 61, 1, 0,
		0, 0, 63, 20, 1, 0, 0, 0, 64, 65, 5, 48, 0, 0, 65, 66, 5, 120, 0, 0, 66,
		70, 1, 0, 0, 0, 67, 69, 7, 3, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1, 0, 0,
		0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 22, 1, 0, 0, 0, 72, 70,
		1, 0, 0, 0, 73, 79, 7, 4, 0, 0, 74, 78, 7, 5, 0, 0, 75, 76, 5, 58, 0, 0,
		76, 78, 5, 58, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1,
		0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 24, 1, 0, 0, 0, 81,
		79, 1, 0, 0, 0, 82, 83, 7, 6, 0, 0, 83, 26, 1, 0, 0, 0, 84, 85, 5, 92,
		0, 0, 85, 86, 7, 7, 0, 0, 86, 28, 1, 0, 0, 0, 87, 92, 5, 34, 0, 0, 88,
		91, 3, 27, 13, 0, 89, 91, 8, 8, 0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0,
		0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 95,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 34, 0, 0, 96, 30, 1, 0, 0, 0,
		9, 0, 50, 58, 62, 70, 77, 79, 90, 92, 1, 6, 0, 0,
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
	WystLexerWS         = 9
	WystLexerNUMBER     = 10
	WystLexerHEX        = 11
	WystLexerIDENTIFIER = 12
	WystLexerMATH       = 13
	WystLexerSTRING     = 14
)
