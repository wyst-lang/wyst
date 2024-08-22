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
		"'namespace'", "'include'", "'use'", "'%import'", "'%map'", "'&&'",
		"'||'", "'if'", "'else'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "COMMENT", "MULTILINE_COMMENT", "WS", "NUMBER", "HEX", "IDENTIFIER",
		"GOIDENTIFIER", "MATH", "STRING", "MODULE_NAME",
	}
	staticData.RuleNames = []string{
		"round_def", "enum_curly", "round_call", "fn_call", "func_def", "var_def",
		"var_def_set", "call_tree", "struct_def", "namespace", "import_statement",
		"use_statement", "go_import", "map", "go_call", "if_expr", "if_statement",
		"elseif_statement", "else_statement", "if_tree", "expr", "code_block",
		"top",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 244, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 51, 8, 0, 10, 0,
		12, 0, 54, 9, 0, 1, 0, 3, 0, 57, 8, 0, 3, 0, 59, 8, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 67, 8, 1, 10, 1, 12, 1, 70, 9, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 78, 8, 2, 10, 2, 12, 2, 81, 9, 2, 1, 2, 3, 2,
		84, 8, 2, 3, 2, 86, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 3,
		7, 107, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 112, 8, 7, 5, 7, 114, 8, 7, 10, 7,
		12, 7, 117, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 128, 8, 9, 1, 9, 1, 9, 1, 9, 5, 9, 133, 8, 9, 10, 9, 12, 9, 136,
		9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 165, 8, 15, 10, 15,
		12, 15, 168, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 175, 8, 16,
		10, 16, 12, 16, 178, 9, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 185,
		8, 18, 10, 18, 12, 18, 188, 9, 18, 1, 19, 1, 19, 5, 19, 192, 8, 19, 10,
		19, 12, 19, 195, 9, 19, 1, 19, 3, 19, 198, 8, 19, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 207, 8, 20, 1, 20, 1, 20, 5, 20, 211,
		8, 20, 10, 20, 12, 20, 214, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 220,
		8, 21, 1, 21, 1, 21, 1, 21, 5, 21, 225, 8, 21, 10, 21, 12, 21, 228, 9,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22,
		239, 8, 22, 10, 22, 12, 22, 242, 9, 22, 1, 22, 0, 0, 23, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		0, 1, 1, 0, 15, 16, 256, 0, 46, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 4, 73, 1,
		0, 0, 0, 6, 89, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10, 97, 1, 0, 0, 0, 12,
		100, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 118, 1, 0, 0, 0, 18, 122, 1,
		0, 0, 0, 20, 139, 1, 0, 0, 0, 22, 143, 1, 0, 0, 0, 24, 147, 1, 0, 0, 0,
		26, 151, 1, 0, 0, 0, 28, 157, 1, 0, 0, 0, 30, 160, 1, 0, 0, 0, 32, 171,
		1, 0, 0, 0, 34, 179, 1, 0, 0, 0, 36, 182, 1, 0, 0, 0, 38, 189, 1, 0, 0,
		0, 40, 206, 1, 0, 0, 0, 42, 215, 1, 0, 0, 0, 44, 240, 1, 0, 0, 0, 46, 58,
		5, 1, 0, 0, 47, 52, 3, 10, 5, 0, 48, 49, 5, 2, 0, 0, 49, 51, 3, 10, 5,
		0, 50, 48, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53,
		1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 57, 5, 2, 0, 0,
		56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 47, 1,
		0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 5, 3, 0, 0, 61,
		1, 1, 0, 0, 0, 62, 68, 5, 4, 0, 0, 63, 64, 3, 10, 5, 0, 64, 65, 5, 5, 0,
		0, 65, 67, 1, 0, 0, 0, 66, 63, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66,
		1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0,
		71, 72, 5, 6, 0, 0, 72, 3, 1, 0, 0, 0, 73, 85, 5, 1, 0, 0, 74, 79, 3, 40,
		20, 0, 75, 76, 5, 2, 0, 0, 76, 78, 3, 40, 20, 0, 77, 75, 1, 0, 0, 0, 78,
		81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 83, 1, 0, 0,
		0, 81, 79, 1, 0, 0, 0, 82, 84, 5, 2, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 74, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 87, 1, 0, 0, 0, 87, 88, 5, 3, 0, 0, 88, 5, 1, 0, 0, 0, 89, 90, 5, 24,
		0, 0, 90, 91, 3, 4, 2, 0, 91, 7, 1, 0, 0, 0, 92, 93, 5, 24, 0, 0, 93, 94,
		5, 24, 0, 0, 94, 95, 3, 0, 0, 0, 95, 96, 3, 42, 21, 0, 96, 9, 1, 0, 0,
		0, 97, 98, 5, 24, 0, 0, 98, 99, 5, 24, 0, 0, 99, 11, 1, 0, 0, 0, 100, 101,
		3, 10, 5, 0, 101, 102, 5, 7, 0, 0, 102, 103, 3, 40, 20, 0, 103, 13, 1,
		0, 0, 0, 104, 107, 3, 6, 3, 0, 105, 107, 5, 24, 0, 0, 106, 104, 1, 0, 0,
		0, 106, 105, 1, 0, 0, 0, 107, 115, 1, 0, 0, 0, 108, 111, 5, 8, 0, 0, 109,
		112, 3, 6, 3, 0, 110, 112, 5, 24, 0, 0, 111, 109, 1, 0, 0, 0, 111, 110,
		1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 108, 1, 0, 0, 0, 114, 117, 1, 0,
		0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 15, 1, 0, 0, 0,
		117, 115, 1, 0, 0, 0, 118, 119, 5, 9, 0, 0, 119, 120, 5, 24, 0, 0, 120,
		121, 3, 2, 1, 0, 121, 17, 1, 0, 0, 0, 122, 123, 5, 10, 0, 0, 123, 124,
		5, 24, 0, 0, 124, 134, 5, 4, 0, 0, 125, 128, 3, 10, 5, 0, 126, 128, 3,
		12, 6, 0, 127, 125, 1, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0,
		0, 129, 130, 5, 5, 0, 0, 130, 133, 1, 0, 0, 0, 131, 133, 3, 8, 4, 0, 132,
		127, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0, 136, 134, 1, 0,
		0, 0, 137, 138, 5, 6, 0, 0, 138, 19, 1, 0, 0, 0, 139, 140, 5, 11, 0, 0,
		140, 141, 5, 28, 0, 0, 141, 142, 5, 5, 0, 0, 142, 21, 1, 0, 0, 0, 143,
		144, 5, 12, 0, 0, 144, 145, 5, 24, 0, 0, 145, 146, 5, 5, 0, 0, 146, 23,
		1, 0, 0, 0, 147, 148, 5, 13, 0, 0, 148, 149, 5, 27, 0, 0, 149, 150, 5,
		5, 0, 0, 150, 25, 1, 0, 0, 0, 151, 152, 5, 14, 0, 0, 152, 153, 5, 24, 0,
		0, 153, 154, 5, 2, 0, 0, 154, 155, 5, 24, 0, 0, 155, 156, 5, 5, 0, 0, 156,
		27, 1, 0, 0, 0, 157, 158, 5, 25, 0, 0, 158, 159, 3, 4, 2, 0, 159, 29, 1,
		0, 0, 0, 160, 161, 5, 1, 0, 0, 161, 166, 3, 40, 20, 0, 162, 163, 7, 0,
		0, 0, 163, 165, 3, 40, 20, 0, 164, 162, 1, 0, 0, 0, 165, 168, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 169, 1, 0, 0, 0, 168,
		166, 1, 0, 0, 0, 169, 170, 5, 3, 0, 0, 170, 31, 1, 0, 0, 0, 171, 172, 5,
		17, 0, 0, 172, 176, 3, 30, 15, 0, 173, 175, 3, 42, 21, 0, 174, 173, 1,
		0, 0, 0, 175, 178, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0,
		0, 177, 33, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 179, 180, 5, 18, 0, 0, 180,
		181, 3, 32, 16, 0, 181, 35, 1, 0, 0, 0, 182, 186, 5, 18, 0, 0, 183, 185,
		3, 42, 21, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1,
		0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 37, 1, 0, 0, 0, 188, 186, 1, 0, 0,
		0, 189, 193, 3, 32, 16, 0, 190, 192, 3, 34, 17, 0, 191, 190, 1, 0, 0, 0,
		192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194,
		197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 198, 3, 36, 18, 0, 197, 196,
		1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 39, 1, 0, 0, 0, 199, 207, 3, 14,
		7, 0, 200, 207, 5, 22, 0, 0, 201, 207, 5, 23, 0, 0, 202, 207, 5, 24, 0,
		0, 203, 207, 5, 27, 0, 0, 204, 207, 3, 28, 14, 0, 205, 207, 5, 25, 0, 0,
		206, 199, 1, 0, 0, 0, 206, 200, 1, 0, 0, 0, 206, 201, 1, 0, 0, 0, 206,
		202, 1, 0, 0, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 205,
		1, 0, 0, 0, 207, 212, 1, 0, 0, 0, 208, 209, 5, 26, 0, 0, 209, 211, 3, 40,
		20, 0, 210, 208, 1, 0, 0, 0, 211, 214, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0,
		212, 213, 1, 0, 0, 0, 213, 41, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 226,
		5, 4, 0, 0, 216, 220, 3, 40, 20, 0, 217, 220, 3, 10, 5, 0, 218, 220, 3,
		12, 6, 0, 219, 216, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 218, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 222, 5, 5, 0, 0, 222, 225, 1, 0, 0, 0, 223,
		225, 3, 38, 19, 0, 224, 219, 1, 0, 0, 0, 224, 223, 1, 0, 0, 0, 225, 228,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1, 0,
		0, 0, 228, 226, 1, 0, 0, 0, 229, 230, 5, 6, 0, 0, 230, 43, 1, 0, 0, 0,
		231, 239, 3, 8, 4, 0, 232, 239, 3, 16, 8, 0, 233, 239, 3, 18, 9, 0, 234,
		239, 3, 20, 10, 0, 235, 239, 3, 22, 11, 0, 236, 239, 3, 24, 12, 0, 237,
		239, 3, 26, 13, 0, 238, 231, 1, 0, 0, 0, 238, 232, 1, 0, 0, 0, 238, 233,
		1, 0, 0, 0, 238, 234, 1, 0, 0, 0, 238, 235, 1, 0, 0, 0, 238, 236, 1, 0,
		0, 0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0,
		240, 241, 1, 0, 0, 0, 241, 45, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 25, 52,
		56, 58, 68, 79, 83, 85, 106, 111, 115, 127, 132, 134, 166, 176, 186, 193,
		197, 206, 212, 219, 224, 226, 238, 240,
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
	WystParserEOF               = antlr.TokenEOF
	WystParserT__0              = 1
	WystParserT__1              = 2
	WystParserT__2              = 3
	WystParserT__3              = 4
	WystParserT__4              = 5
	WystParserT__5              = 6
	WystParserT__6              = 7
	WystParserT__7              = 8
	WystParserT__8              = 9
	WystParserT__9              = 10
	WystParserT__10             = 11
	WystParserT__11             = 12
	WystParserT__12             = 13
	WystParserT__13             = 14
	WystParserT__14             = 15
	WystParserT__15             = 16
	WystParserT__16             = 17
	WystParserT__17             = 18
	WystParserCOMMENT           = 19
	WystParserMULTILINE_COMMENT = 20
	WystParserWS                = 21
	WystParserNUMBER            = 22
	WystParserHEX               = 23
	WystParserIDENTIFIER        = 24
	WystParserGOIDENTIFIER      = 25
	WystParserMATH              = 26
	WystParserSTRING            = 27
	WystParserMODULE_NAME       = 28
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
	WystParserRULE_go_import        = 12
	WystParserRULE_map              = 13
	WystParserRULE_go_call          = 14
	WystParserRULE_if_expr          = 15
	WystParserRULE_if_statement     = 16
	WystParserRULE_elseif_statement = 17
	WystParserRULE_else_statement   = 18
	WystParserRULE_if_tree          = 19
	WystParserRULE_expr             = 20
	WystParserRULE_code_block       = 21
	WystParserRULE_top              = 22
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
		p.SetState(46)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WystParserIDENTIFIER {
		{
			p.SetState(47)
			p.Var_def()
		}
		p.SetState(52)
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
					p.SetState(48)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(49)
					p.Var_def()
				}

			}
			p.SetState(54)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(55)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(60)
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
		p.SetState(62)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserIDENTIFIER {
		{
			p.SetState(63)
			p.Var_def()
		}
		{
			p.SetState(64)
			p.Match(WystParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
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
		p.SetState(73)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197132288) != 0 {
		{
			p.SetState(74)
			p.Expr()
		}
		p.SetState(79)
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
					p.SetState(75)
					p.Match(WystParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)
					p.Expr()
				}

			}
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == WystParserT__1 {
			{
				p.SetState(82)
				p.Match(WystParserT__1)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(87)
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
		p.SetState(89)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(90)
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
		p.SetState(92)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Round_def()
	}
	{
		p.SetState(95)
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
		p.SetState(97)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
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
		p.SetState(100)
		p.Var_def()
	}
	{
		p.SetState(101)
		p.Match(WystParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(104)
			p.Fn_call()
		}

	case 2:
		{
			p.SetState(105)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__7 {
		{
			p.SetState(108)
			p.Match(WystParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(109)
				p.Fn_call()
			}

		case 2:
			{
				p.SetState(110)
				p.Match(WystParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(117)
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
		p.SetState(118)
		p.Match(WystParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
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
		p.SetState(122)
		p.Match(WystParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserIDENTIFIER {
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
		case 1:
			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(125)
					p.Var_def()
				}

			case 2:
				{
					p.SetState(126)
					p.Var_def_set()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			{
				p.SetState(129)
				p.Match(WystParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			{
				p.SetState(131)
				p.Func_def()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(137)
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
	MODULE_NAME() antlr.TerminalNode

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

func (s *Import_statementContext) MODULE_NAME() antlr.TerminalNode {
	return s.GetToken(WystParserMODULE_NAME, 0)
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
		p.SetState(139)
		p.Match(WystParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(WystParserMODULE_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
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
		p.SetState(143)
		p.Match(WystParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
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

// IGo_importContext is an interface to support dynamic dispatch.
type IGo_importContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsGo_importContext differentiates from other interfaces.
	IsGo_importContext()
}

type Go_importContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGo_importContext() *Go_importContext {
	var p = new(Go_importContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_go_import
	return p
}

func InitEmptyGo_importContext(p *Go_importContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_go_import
}

func (*Go_importContext) IsGo_importContext() {}

func NewGo_importContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Go_importContext {
	var p = new(Go_importContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_go_import

	return p
}

func (s *Go_importContext) GetParser() antlr.Parser { return s.parser }

func (s *Go_importContext) STRING() antlr.TerminalNode {
	return s.GetToken(WystParserSTRING, 0)
}

func (s *Go_importContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Go_importContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Go_importContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterGo_import(s)
	}
}

func (s *Go_importContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitGo_import(s)
	}
}

func (p *WystParser) Go_import() (localctx IGo_importContext) {
	localctx = NewGo_importContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, WystParserRULE_go_import)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(WystParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(WystParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
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

// IMapContext is an interface to support dynamic dispatch.
type IMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsMapContext differentiates from other interfaces.
	IsMapContext()
}

type MapContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapContext() *MapContext {
	var p = new(MapContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_map
	return p
}

func InitEmptyMapContext(p *MapContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_map
}

func (*MapContext) IsMapContext() {}

func NewMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapContext {
	var p = new(MapContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_map

	return p
}

func (s *MapContext) GetParser() antlr.Parser { return s.parser }

func (s *MapContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(WystParserIDENTIFIER)
}

func (s *MapContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(WystParserIDENTIFIER, i)
}

func (s *MapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterMap(s)
	}
}

func (s *MapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitMap(s)
	}
}

func (p *WystParser) Map_() (localctx IMapContext) {
	localctx = NewMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, WystParserRULE_map)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(WystParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(WystParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(WystParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
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

// IGo_callContext is an interface to support dynamic dispatch.
type IGo_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GOIDENTIFIER() antlr.TerminalNode
	Round_call() IRound_callContext

	// IsGo_callContext differentiates from other interfaces.
	IsGo_callContext()
}

type Go_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGo_callContext() *Go_callContext {
	var p = new(Go_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_go_call
	return p
}

func InitEmptyGo_callContext(p *Go_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_go_call
}

func (*Go_callContext) IsGo_callContext() {}

func NewGo_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Go_callContext {
	var p = new(Go_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_go_call

	return p
}

func (s *Go_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Go_callContext) GOIDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserGOIDENTIFIER, 0)
}

func (s *Go_callContext) Round_call() IRound_callContext {
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

func (s *Go_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Go_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Go_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterGo_call(s)
	}
}

func (s *Go_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitGo_call(s)
	}
}

func (p *WystParser) Go_call() (localctx IGo_callContext) {
	localctx = NewGo_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, WystParserRULE_go_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(WystParserGOIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
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

// IIf_exprContext is an interface to support dynamic dispatch.
type IIf_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsIf_exprContext differentiates from other interfaces.
	IsIf_exprContext()
}

type If_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_exprContext() *If_exprContext {
	var p = new(If_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_expr
	return p
}

func InitEmptyIf_exprContext(p *If_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_expr
}

func (*If_exprContext) IsIf_exprContext() {}

func NewIf_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_exprContext {
	var p = new(If_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_if_expr

	return p
}

func (s *If_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *If_exprContext) AllExpr() []IExprContext {
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

func (s *If_exprContext) Expr(i int) IExprContext {
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

func (s *If_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterIf_expr(s)
	}
}

func (s *If_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitIf_expr(s)
	}
}

func (p *WystParser) If_expr() (localctx IIf_exprContext) {
	localctx = NewIf_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, WystParserRULE_if_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(WystParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Expr()
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__14 || _la == WystParserT__15 {
		{
			p.SetState(162)
			_la = p.GetTokenStream().LA(1)

			if !(_la == WystParserT__14 || _la == WystParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(163)
			p.Expr()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(169)
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_expr() IIf_exprContext
	AllCode_block() []ICode_blockContext
	Code_block(i int) ICode_blockContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) If_expr() IIf_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_exprContext)
}

func (s *If_statementContext) AllCode_block() []ICode_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICode_blockContext); ok {
			len++
		}
	}

	tst := make([]ICode_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICode_blockContext); ok {
			tst[i] = t.(ICode_blockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Code_block(i int) ICode_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
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

	return t.(ICode_blockContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *WystParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, WystParserRULE_if_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(WystParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.If_expr()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__3 {
		{
			p.SetState(173)
			p.Code_block()
		}

		p.SetState(178)
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

// IElseif_statementContext is an interface to support dynamic dispatch.
type IElseif_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_statement() IIf_statementContext

	// IsElseif_statementContext differentiates from other interfaces.
	IsElseif_statementContext()
}

type Elseif_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseif_statementContext() *Elseif_statementContext {
	var p = new(Elseif_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_elseif_statement
	return p
}

func InitEmptyElseif_statementContext(p *Elseif_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_elseif_statement
}

func (*Elseif_statementContext) IsElseif_statementContext() {}

func NewElseif_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elseif_statementContext {
	var p = new(Elseif_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_elseif_statement

	return p
}

func (s *Elseif_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Elseif_statementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Elseif_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elseif_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elseif_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterElseif_statement(s)
	}
}

func (s *Elseif_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitElseif_statement(s)
	}
}

func (p *WystParser) Elseif_statement() (localctx IElseif_statementContext) {
	localctx = NewElseif_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, WystParserRULE_elseif_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(WystParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.If_statement()
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

// IElse_statementContext is an interface to support dynamic dispatch.
type IElse_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCode_block() []ICode_blockContext
	Code_block(i int) ICode_blockContext

	// IsElse_statementContext differentiates from other interfaces.
	IsElse_statementContext()
}

type Else_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_statementContext() *Else_statementContext {
	var p = new(Else_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_else_statement
	return p
}

func InitEmptyElse_statementContext(p *Else_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_else_statement
}

func (*Else_statementContext) IsElse_statementContext() {}

func NewElse_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_statementContext {
	var p = new(Else_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_else_statement

	return p
}

func (s *Else_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_statementContext) AllCode_block() []ICode_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICode_blockContext); ok {
			len++
		}
	}

	tst := make([]ICode_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICode_blockContext); ok {
			tst[i] = t.(ICode_blockContext)
			i++
		}
	}

	return tst
}

func (s *Else_statementContext) Code_block(i int) ICode_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICode_blockContext); ok {
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

	return t.(ICode_blockContext)
}

func (s *Else_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterElse_statement(s)
	}
}

func (s *Else_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitElse_statement(s)
	}
}

func (p *WystParser) Else_statement() (localctx IElse_statementContext) {
	localctx = NewElse_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, WystParserRULE_else_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(WystParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == WystParserT__3 {
		{
			p.SetState(183)
			p.Code_block()
		}

		p.SetState(188)
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

// IIf_treeContext is an interface to support dynamic dispatch.
type IIf_treeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_statement() IIf_statementContext
	AllElseif_statement() []IElseif_statementContext
	Elseif_statement(i int) IElseif_statementContext
	Else_statement() IElse_statementContext

	// IsIf_treeContext differentiates from other interfaces.
	IsIf_treeContext()
}

type If_treeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_treeContext() *If_treeContext {
	var p = new(If_treeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_tree
	return p
}

func InitEmptyIf_treeContext(p *If_treeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = WystParserRULE_if_tree
}

func (*If_treeContext) IsIf_treeContext() {}

func NewIf_treeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_treeContext {
	var p = new(If_treeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = WystParserRULE_if_tree

	return p
}

func (s *If_treeContext) GetParser() antlr.Parser { return s.parser }

func (s *If_treeContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *If_treeContext) AllElseif_statement() []IElseif_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseif_statementContext); ok {
			len++
		}
	}

	tst := make([]IElseif_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseif_statementContext); ok {
			tst[i] = t.(IElseif_statementContext)
			i++
		}
	}

	return tst
}

func (s *If_treeContext) Elseif_statement(i int) IElseif_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseif_statementContext); ok {
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

	return t.(IElseif_statementContext)
}

func (s *If_treeContext) Else_statement() IElse_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_statementContext)
}

func (s *If_treeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_treeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_treeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.EnterIf_tree(s)
	}
}

func (s *If_treeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WystListener); ok {
		listenerT.ExitIf_tree(s)
	}
}

func (p *WystParser) If_tree() (localctx IIf_treeContext) {
	localctx = NewIf_treeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, WystParserRULE_if_tree)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.If_statement()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(190)
				p.Elseif_statement()
			}

		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == WystParserT__17 {
		{
			p.SetState(196)
			p.Else_statement()
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
	Go_call() IGo_callContext
	GOIDENTIFIER() antlr.TerminalNode
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

func (s *ExprContext) Go_call() IGo_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGo_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGo_callContext)
}

func (s *ExprContext) GOIDENTIFIER() antlr.TerminalNode {
	return s.GetToken(WystParserGOIDENTIFIER, 0)
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
	p.EnterRule(localctx, 40, WystParserRULE_expr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(199)
			p.Call_tree()
		}

	case 2:
		{
			p.SetState(200)
			p.Match(WystParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(201)
			p.Match(WystParserHEX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(202)
			p.Match(WystParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(203)
			p.Match(WystParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(204)
			p.Go_call()
		}

	case 7:
		{
			p.SetState(205)
			p.Match(WystParserGOIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(208)
				p.Match(WystParserMATH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(209)
				p.Expr()
			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	AllIf_tree() []IIf_treeContext
	If_tree(i int) IIf_treeContext
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

func (s *Code_blockContext) AllIf_tree() []IIf_treeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_treeContext); ok {
			len++
		}
	}

	tst := make([]IIf_treeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_treeContext); ok {
			tst[i] = t.(IIf_treeContext)
			i++
		}
	}

	return tst
}

func (s *Code_blockContext) If_tree(i int) IIf_treeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_treeContext); ok {
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

	return t.(IIf_treeContext)
}

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
	p.EnterRule(localctx, 42, WystParserRULE_code_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(WystParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&197263360) != 0 {
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case WystParserNUMBER, WystParserHEX, WystParserIDENTIFIER, WystParserGOIDENTIFIER, WystParserSTRING:
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(216)
					p.Expr()
				}

			case 2:
				{
					p.SetState(217)
					p.Var_def()
				}

			case 3:
				{
					p.SetState(218)
					p.Var_def_set()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}
			{
				p.SetState(221)
				p.Match(WystParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case WystParserT__16:
			{
				p.SetState(223)
				p.If_tree()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(229)
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
	AllGo_import() []IGo_importContext
	Go_import(i int) IGo_importContext
	AllMap_() []IMapContext
	Map_(i int) IMapContext

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

func (s *TopContext) AllGo_import() []IGo_importContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGo_importContext); ok {
			len++
		}
	}

	tst := make([]IGo_importContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGo_importContext); ok {
			tst[i] = t.(IGo_importContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Go_import(i int) IGo_importContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGo_importContext); ok {
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

	return t.(IGo_importContext)
}

func (s *TopContext) AllMap_() []IMapContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMapContext); ok {
			len++
		}
	}

	tst := make([]IMapContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMapContext); ok {
			tst[i] = t.(IMapContext)
			i++
		}
	}

	return tst
}

func (s *TopContext) Map_(i int) IMapContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapContext); ok {
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

	return t.(IMapContext)
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
	p.EnterRule(localctx, 44, WystParserRULE_top)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16809472) != 0 {
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case WystParserIDENTIFIER:
			{
				p.SetState(231)
				p.Func_def()
			}

		case WystParserT__8:
			{
				p.SetState(232)
				p.Struct_def()
			}

		case WystParserT__9:
			{
				p.SetState(233)
				p.Namespace()
			}

		case WystParserT__10:
			{
				p.SetState(234)
				p.Import_statement()
			}

		case WystParserT__11:
			{
				p.SetState(235)
				p.Use_statement()
			}

		case WystParserT__12:
			{
				p.SetState(236)
				p.Go_import()
			}

		case WystParserT__13:
			{
				p.SetState(237)
				p.Map_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(242)
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
