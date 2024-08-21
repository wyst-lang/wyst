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
		"", "'('", "','", "')'", "'{'", "';'", "'}'", "'='", "'.'", "'struct'",
		"'namespace'", "'include'", "'use'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "NUMBER",
		"HEX", "IDENTIFIER", "MATH", "STRING",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "WS", "NUMBER", "HEX", "IDENTIFIER", "MATH",
		"ESC", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 147, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 86, 8, 12, 11,
		12, 12, 12, 87, 1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 94, 8, 13, 10, 13, 12,
		13, 97, 9, 13, 1, 13, 3, 13, 100, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5,
		14, 106, 8, 14, 10, 14, 12, 14, 109, 9, 14, 1, 15, 1, 15, 5, 15, 113, 8,
		15, 10, 15, 12, 15, 116, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15,
		123, 8, 15, 10, 15, 12, 15, 126, 9, 15, 5, 15, 128, 8, 15, 10, 15, 12,
		15, 131, 9, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		5, 18, 141, 8, 18, 10, 18, 12, 18, 144, 9, 18, 1, 18, 1, 18, 0, 0, 19,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 0, 37, 18, 1, 0, 9,
		3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 48, 57,
		65, 70, 97, 102, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 3, 0, 42, 43, 45, 45, 47, 47, 3, 0, 34, 34, 39, 39, 92, 92,
		2, 0, 34, 34, 92, 92, 154, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 1, 39, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 45, 1, 0,
		0, 0, 9, 47, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13, 51, 1, 0, 0, 0, 15, 53,
		1, 0, 0, 0, 17, 55, 1, 0, 0, 0, 19, 62, 1, 0, 0, 0, 21, 72, 1, 0, 0, 0,
		23, 80, 1, 0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 99, 1, 0, 0, 0, 29, 101, 1,
		0, 0, 0, 31, 110, 1, 0, 0, 0, 33, 132, 1, 0, 0, 0, 35, 134, 1, 0, 0, 0,
		37, 137, 1, 0, 0, 0, 39, 40, 5, 40, 0, 0, 40, 2, 1, 0, 0, 0, 41, 42, 5,
		44, 0, 0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 41, 0, 0, 44, 6, 1, 0, 0, 0, 45,
		46, 5, 123, 0, 0, 46, 8, 1, 0, 0, 0, 47, 48, 5, 59, 0, 0, 48, 10, 1, 0,
		0, 0, 49, 50, 5, 125, 0, 0, 50, 12, 1, 0, 0, 0, 51, 52, 5, 61, 0, 0, 52,
		14, 1, 0, 0, 0, 53, 54, 5, 46, 0, 0, 54, 16, 1, 0, 0, 0, 55, 56, 5, 115,
		0, 0, 56, 57, 5, 116, 0, 0, 57, 58, 5, 114, 0, 0, 58, 59, 5, 117, 0, 0,
		59, 60, 5, 99, 0, 0, 60, 61, 5, 116, 0, 0, 61, 18, 1, 0, 0, 0, 62, 63,
		5, 110, 0, 0, 63, 64, 5, 97, 0, 0, 64, 65, 5, 109, 0, 0, 65, 66, 5, 101,
		0, 0, 66, 67, 5, 115, 0, 0, 67, 68, 5, 112, 0, 0, 68, 69, 5, 97, 0, 0,
		69, 70, 5, 99, 0, 0, 70, 71, 5, 101, 0, 0, 71, 20, 1, 0, 0, 0, 72, 73,
		5, 105, 0, 0, 73, 74, 5, 110, 0, 0, 74, 75, 5, 99, 0, 0, 75, 76, 5, 108,
		0, 0, 76, 77, 5, 117, 0, 0, 77, 78, 5, 100, 0, 0, 78, 79, 5, 101, 0, 0,
		79, 22, 1, 0, 0, 0, 80, 81, 5, 117, 0, 0, 81, 82, 5, 115, 0, 0, 82, 83,
		5, 101, 0, 0, 83, 24, 1, 0, 0, 0, 84, 86, 7, 0, 0, 0, 85, 84, 1, 0, 0,
		0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89,
		1, 0, 0, 0, 89, 90, 6, 12, 0, 0, 90, 26, 1, 0, 0, 0, 91, 95, 7, 1, 0, 0,
		92, 94, 7, 2, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1,
		0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 100, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98,
		100, 5, 48, 0, 0, 99, 91, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 28, 1, 0,
		0, 0, 101, 102, 5, 48, 0, 0, 102, 103, 5, 120, 0, 0, 103, 107, 1, 0, 0,
		0, 104, 106, 7, 3, 0, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107,
		105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 30, 1, 0, 0, 0, 109, 107, 1,
		0, 0, 0, 110, 114, 7, 4, 0, 0, 111, 113, 7, 5, 0, 0, 112, 111, 1, 0, 0,
		0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		129, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 58, 0, 0, 118, 119,
		5, 58, 0, 0, 119, 120, 1, 0, 0, 0, 120, 124, 7, 4, 0, 0, 121, 123, 7, 5,
		0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0,
		124, 125, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127,
		117, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 32, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 7, 6,
		0, 0, 133, 34, 1, 0, 0, 0, 134, 135, 5, 92, 0, 0, 135, 136, 7, 7, 0, 0,
		136, 36, 1, 0, 0, 0, 137, 142, 5, 34, 0, 0, 138, 141, 3, 35, 17, 0, 139,
		141, 8, 8, 0, 0, 140, 138, 1, 0, 0, 0, 140, 139, 1, 0, 0, 0, 141, 144,
		1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0,
		0, 0, 144, 142, 1, 0, 0, 0, 145, 146, 5, 34, 0, 0, 146, 38, 1, 0, 0, 0,
		10, 0, 87, 95, 99, 107, 114, 124, 129, 140, 142, 1, 6, 0, 0,
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
	WystLexerT__9       = 10
	WystLexerT__10      = 11
	WystLexerT__11      = 12
	WystLexerWS         = 13
	WystLexerNUMBER     = 14
	WystLexerHEX        = 15
	WystLexerIDENTIFIER = 16
	WystLexerMATH       = 17
	WystLexerSTRING     = 18
)
