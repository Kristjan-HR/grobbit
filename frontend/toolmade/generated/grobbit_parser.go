// Code generated from GrobbitParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // GrobbitParser
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

type GrobbitParser struct {
	*antlr.BaseParser
}

var GrobbitParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grobbitparserParserInit() {
	staticData := &GrobbitParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'{'", "'}'", "','", "'.'", "';'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'='", "'&&'", "'||'", "'!'", "'=='", "'<'", "'>'", "'!='",
		"'<='", "'>='", "':='", "'break'", "'case'", "'chan'", "'const'", "'continue'",
		"'default'", "'defer'", "'else'", "'fallthrough'", "'for'", "'func'",
		"'go'", "'goto'", "'if'", "'import'", "'interface'", "'map'", "'package'",
		"'range'", "'return'", "'select'", "'struct'", "'switch'", "'type'",
		"'var'",
	}
	staticData.SymbolicNames = []string{
		"", "TtLParen", "TtRParen", "TtLBrace", "TtRBrace", "TtComma", "TtPeriod",
		"TtSemicolon", "TtOpAdd", "TtOpSub", "TtOpMul", "TtOpDiv", "TtOpMod",
		"TtOpAssign", "TtOpAnd", "TtOpOr", "TtOpNot", "TtOpEq", "TtOpLt", "TtOpGt",
		"TtOpNe", "TtOpLe", "TtOpGe", "TtOpDefine", "TtKwBreak", "TtKwCase",
		"TtKwChan", "TtKwConst", "TtKwContinue", "TtKwDefault", "TtKwDefer",
		"TtKwElse", "TtKwFallthrough", "TtKwFor", "TtKwFunc", "TtKwGo", "TtKwGoto",
		"TtKwIf", "TtKwImport", "TtKwInterface", "TtKwMap", "TtKwPackage", "TtKwRange",
		"TtKwReturn", "TtKwSelect", "TtKwStruct", "TtKwSwitch", "TtKwType",
		"TtKwVar", "TtBool", "TtIdentifier", "TtFloat", "TtInt", "TtString",
		"TtChar", "WS", "BLOCK_COMMENT", "LINE_COMMENT", "TtUnknown",
	}
	staticData.RuleNames = []string{
		"program", "top_decl", "pkg_name", "import_decl", "import_spec", "const_decl",
		"const_spec", "var_decl", "var_spec", "basic_literals", "basic_literal",
		"func_decl", "param_decl", "param_spec", "identifiers", "identifier",
		"type", "block", "statement", "decl_stmt", "simple_stmt", "return_stmt",
		"continue_stmt", "break_stmt", "if_stmt", "for_stmt", "for_clause",
		"condition", "shortvar_decl", "assign_stmt", "expr_stmt", "expressions",
		"expression", "unary_expr", "primary_expr", "call_expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 349, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 79, 8, 0, 10, 0, 12, 0, 82, 9, 0, 1, 0, 1,
		0, 1, 0, 5, 0, 87, 8, 0, 10, 0, 12, 0, 90, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 3, 1, 97, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5,
		3, 107, 8, 3, 10, 3, 12, 3, 110, 9, 3, 1, 3, 3, 3, 113, 8, 3, 1, 4, 3,
		4, 116, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 126,
		8, 5, 10, 5, 12, 5, 129, 9, 5, 1, 5, 3, 5, 132, 8, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 145, 8, 7, 10, 7,
		12, 7, 148, 9, 7, 1, 7, 3, 7, 151, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		157, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 162, 8, 9, 10, 9, 12, 9, 165, 9, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 173, 8, 11, 1, 11, 1,
		11, 3, 11, 177, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 184, 8,
		12, 10, 12, 12, 12, 187, 9, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		5, 14, 195, 8, 14, 10, 14, 12, 14, 198, 9, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 208, 8, 17, 10, 17, 12, 17, 211,
		9, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 223, 8, 18, 1, 19, 1, 19, 3, 19, 227, 8, 19, 1, 20, 1, 20, 1,
		20, 3, 20, 232, 8, 20, 1, 21, 1, 21, 3, 21, 236, 8, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 246, 8, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 3, 24, 253, 8, 24, 3, 24, 255, 8, 24, 1, 25, 1, 25,
		1, 25, 3, 25, 260, 8, 25, 1, 25, 1, 25, 1, 26, 3, 26, 265, 8, 26, 1, 26,
		1, 26, 3, 26, 269, 8, 26, 1, 26, 1, 26, 3, 26, 273, 8, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 5, 31, 290, 8, 31, 10, 31, 12, 31, 293, 9, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 313, 8, 32, 10, 32,
		12, 32, 316, 9, 32, 1, 33, 1, 33, 1, 33, 3, 33, 321, 8, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 3, 34, 327, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 3, 34, 335, 8, 34, 1, 35, 1, 35, 1, 35, 3, 35, 340, 8, 35, 1, 35, 1,
		35, 1, 35, 3, 35, 345, 8, 35, 1, 35, 1, 35, 1, 35, 0, 1, 64, 36, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 5, 2, 0,
		49, 49, 51, 52, 1, 0, 10, 12, 1, 0, 8, 9, 1, 0, 17, 22, 2, 0, 8, 9, 16,
		16, 362, 0, 72, 1, 0, 0, 0, 2, 96, 1, 0, 0, 0, 4, 98, 1, 0, 0, 0, 6, 100,
		1, 0, 0, 0, 8, 115, 1, 0, 0, 0, 10, 119, 1, 0, 0, 0, 12, 133, 1, 0, 0,
		0, 14, 138, 1, 0, 0, 0, 16, 152, 1, 0, 0, 0, 18, 158, 1, 0, 0, 0, 20, 166,
		1, 0, 0, 0, 22, 168, 1, 0, 0, 0, 24, 180, 1, 0, 0, 0, 26, 188, 1, 0, 0,
		0, 28, 191, 1, 0, 0, 0, 30, 199, 1, 0, 0, 0, 32, 201, 1, 0, 0, 0, 34, 203,
		1, 0, 0, 0, 36, 222, 1, 0, 0, 0, 38, 226, 1, 0, 0, 0, 40, 231, 1, 0, 0,
		0, 42, 233, 1, 0, 0, 0, 44, 237, 1, 0, 0, 0, 46, 239, 1, 0, 0, 0, 48, 241,
		1, 0, 0, 0, 50, 256, 1, 0, 0, 0, 52, 264, 1, 0, 0, 0, 54, 274, 1, 0, 0,
		0, 56, 276, 1, 0, 0, 0, 58, 280, 1, 0, 0, 0, 60, 284, 1, 0, 0, 0, 62, 286,
		1, 0, 0, 0, 64, 294, 1, 0, 0, 0, 66, 320, 1, 0, 0, 0, 68, 334, 1, 0, 0,
		0, 70, 339, 1, 0, 0, 0, 72, 73, 5, 41, 0, 0, 73, 74, 3, 4, 2, 0, 74, 80,
		5, 7, 0, 0, 75, 76, 3, 6, 3, 0, 76, 77, 5, 7, 0, 0, 77, 79, 1, 0, 0, 0,
		78, 75, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 88, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 3, 2, 1, 0, 84,
		85, 5, 7, 0, 0, 85, 87, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 87, 90, 1, 0, 0,
		0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88,
		1, 0, 0, 0, 91, 92, 5, 0, 0, 1, 92, 1, 1, 0, 0, 0, 93, 97, 3, 10, 5, 0,
		94, 97, 3, 14, 7, 0, 95, 97, 3, 22, 11, 0, 96, 93, 1, 0, 0, 0, 96, 94,
		1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 3, 1, 0, 0, 0, 98, 99, 3, 30, 15, 0,
		99, 5, 1, 0, 0, 0, 100, 112, 5, 38, 0, 0, 101, 113, 3, 8, 4, 0, 102, 108,
		5, 1, 0, 0, 103, 104, 3, 8, 4, 0, 104, 105, 5, 7, 0, 0, 105, 107, 1, 0,
		0, 0, 106, 103, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		113, 5, 2, 0, 0, 112, 101, 1, 0, 0, 0, 112, 102, 1, 0, 0, 0, 113, 7, 1,
		0, 0, 0, 114, 116, 3, 4, 2, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 53, 0, 0, 118, 9, 1, 0, 0, 0, 119,
		131, 5, 27, 0, 0, 120, 132, 3, 12, 6, 0, 121, 127, 5, 1, 0, 0, 122, 123,
		3, 12, 6, 0, 123, 124, 5, 7, 0, 0, 124, 126, 1, 0, 0, 0, 125, 122, 1, 0,
		0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0,
		128, 130, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 132, 5, 2, 0, 0, 131,
		120, 1, 0, 0, 0, 131, 121, 1, 0, 0, 0, 132, 11, 1, 0, 0, 0, 133, 134, 3,
		28, 14, 0, 134, 135, 3, 32, 16, 0, 135, 136, 5, 13, 0, 0, 136, 137, 3,
		18, 9, 0, 137, 13, 1, 0, 0, 0, 138, 150, 5, 48, 0, 0, 139, 151, 3, 16,
		8, 0, 140, 146, 5, 1, 0, 0, 141, 142, 3, 16, 8, 0, 142, 143, 5, 7, 0, 0,
		143, 145, 1, 0, 0, 0, 144, 141, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146,
		1, 0, 0, 0, 149, 151, 5, 2, 0, 0, 150, 139, 1, 0, 0, 0, 150, 140, 1, 0,
		0, 0, 151, 15, 1, 0, 0, 0, 152, 153, 3, 28, 14, 0, 153, 156, 3, 32, 16,
		0, 154, 155, 5, 13, 0, 0, 155, 157, 3, 18, 9, 0, 156, 154, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 17, 1, 0, 0, 0, 158, 163, 3, 20, 10, 0, 159,
		160, 5, 5, 0, 0, 160, 162, 3, 20, 10, 0, 161, 159, 1, 0, 0, 0, 162, 165,
		1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 19, 1, 0,
		0, 0, 165, 163, 1, 0, 0, 0, 166, 167, 7, 0, 0, 0, 167, 21, 1, 0, 0, 0,
		168, 169, 5, 34, 0, 0, 169, 170, 3, 30, 15, 0, 170, 172, 5, 1, 0, 0, 171,
		173, 3, 24, 12, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174,
		1, 0, 0, 0, 174, 176, 5, 2, 0, 0, 175, 177, 3, 32, 16, 0, 176, 175, 1,
		0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 3, 34, 17,
		0, 179, 23, 1, 0, 0, 0, 180, 185, 3, 26, 13, 0, 181, 182, 5, 5, 0, 0, 182,
		184, 3, 26, 13, 0, 183, 181, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183,
		1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 25, 1, 0, 0, 0, 187, 185, 1, 0,
		0, 0, 188, 189, 3, 28, 14, 0, 189, 190, 3, 32, 16, 0, 190, 27, 1, 0, 0,
		0, 191, 196, 3, 30, 15, 0, 192, 193, 5, 5, 0, 0, 193, 195, 3, 30, 15, 0,
		194, 192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196,
		197, 1, 0, 0, 0, 197, 29, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 5,
		50, 0, 0, 200, 31, 1, 0, 0, 0, 201, 202, 3, 30, 15, 0, 202, 33, 1, 0, 0,
		0, 203, 209, 5, 3, 0, 0, 204, 205, 3, 36, 18, 0, 205, 206, 5, 7, 0, 0,
		206, 208, 1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 212, 1, 0, 0, 0, 211, 209,
		1, 0, 0, 0, 212, 213, 5, 4, 0, 0, 213, 35, 1, 0, 0, 0, 214, 223, 3, 38,
		19, 0, 215, 223, 3, 40, 20, 0, 216, 223, 3, 42, 21, 0, 217, 223, 3, 46,
		23, 0, 218, 223, 3, 44, 22, 0, 219, 223, 3, 34, 17, 0, 220, 223, 3, 48,
		24, 0, 221, 223, 3, 50, 25, 0, 222, 214, 1, 0, 0, 0, 222, 215, 1, 0, 0,
		0, 222, 216, 1, 0, 0, 0, 222, 217, 1, 0, 0, 0, 222, 218, 1, 0, 0, 0, 222,
		219, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 221, 1, 0, 0, 0, 223, 37, 1,
		0, 0, 0, 224, 227, 3, 10, 5, 0, 225, 227, 3, 14, 7, 0, 226, 224, 1, 0,
		0, 0, 226, 225, 1, 0, 0, 0, 227, 39, 1, 0, 0, 0, 228, 232, 3, 60, 30, 0,
		229, 232, 3, 58, 29, 0, 230, 232, 3, 56, 28, 0, 231, 228, 1, 0, 0, 0, 231,
		229, 1, 0, 0, 0, 231, 230, 1, 0, 0, 0, 232, 41, 1, 0, 0, 0, 233, 235, 5,
		43, 0, 0, 234, 236, 3, 64, 32, 0, 235, 234, 1, 0, 0, 0, 235, 236, 1, 0,
		0, 0, 236, 43, 1, 0, 0, 0, 237, 238, 5, 28, 0, 0, 238, 45, 1, 0, 0, 0,
		239, 240, 5, 24, 0, 0, 240, 47, 1, 0, 0, 0, 241, 245, 5, 37, 0, 0, 242,
		243, 3, 40, 20, 0, 243, 244, 5, 7, 0, 0, 244, 246, 1, 0, 0, 0, 245, 242,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 3, 54,
		27, 0, 248, 254, 3, 34, 17, 0, 249, 252, 5, 31, 0, 0, 250, 253, 3, 48,
		24, 0, 251, 253, 3, 34, 17, 0, 252, 250, 1, 0, 0, 0, 252, 251, 1, 0, 0,
		0, 253, 255, 1, 0, 0, 0, 254, 249, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255,
		49, 1, 0, 0, 0, 256, 259, 5, 33, 0, 0, 257, 260, 3, 54, 27, 0, 258, 260,
		3, 52, 26, 0, 259, 257, 1, 0, 0, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1,
		0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 3, 34, 17, 0, 262, 51, 1, 0, 0,
		0, 263, 265, 3, 40, 20, 0, 264, 263, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 268, 5, 7, 0, 0, 267, 269, 3, 54, 27, 0, 268,
		267, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270, 272,
		5, 7, 0, 0, 271, 273, 3, 40, 20, 0, 272, 271, 1, 0, 0, 0, 272, 273, 1,
		0, 0, 0, 273, 53, 1, 0, 0, 0, 274, 275, 3, 64, 32, 0, 275, 55, 1, 0, 0,
		0, 276, 277, 3, 28, 14, 0, 277, 278, 5, 23, 0, 0, 278, 279, 3, 62, 31,
		0, 279, 57, 1, 0, 0, 0, 280, 281, 3, 28, 14, 0, 281, 282, 5, 13, 0, 0,
		282, 283, 3, 62, 31, 0, 283, 59, 1, 0, 0, 0, 284, 285, 3, 70, 35, 0, 285,
		61, 1, 0, 0, 0, 286, 291, 3, 64, 32, 0, 287, 288, 5, 5, 0, 0, 288, 290,
		3, 64, 32, 0, 289, 287, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1,
		0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 63, 1, 0, 0, 0, 293, 291, 1, 0, 0,
		0, 294, 295, 6, 32, -1, 0, 295, 296, 3, 66, 33, 0, 296, 314, 1, 0, 0, 0,
		297, 298, 10, 5, 0, 0, 298, 299, 7, 1, 0, 0, 299, 313, 3, 64, 32, 6, 300,
		301, 10, 4, 0, 0, 301, 302, 7, 2, 0, 0, 302, 313, 3, 64, 32, 5, 303, 304,
		10, 3, 0, 0, 304, 305, 7, 3, 0, 0, 305, 313, 3, 64, 32, 4, 306, 307, 10,
		2, 0, 0, 307, 308, 5, 14, 0, 0, 308, 313, 3, 64, 32, 3, 309, 310, 10, 1,
		0, 0, 310, 311, 5, 15, 0, 0, 311, 313, 3, 64, 32, 2, 312, 297, 1, 0, 0,
		0, 312, 300, 1, 0, 0, 0, 312, 303, 1, 0, 0, 0, 312, 306, 1, 0, 0, 0, 312,
		309, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315,
		1, 0, 0, 0, 315, 65, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 321, 3, 68,
		34, 0, 318, 319, 7, 4, 0, 0, 319, 321, 3, 66, 33, 0, 320, 317, 1, 0, 0,
		0, 320, 318, 1, 0, 0, 0, 321, 67, 1, 0, 0, 0, 322, 335, 3, 20, 10, 0, 323,
		324, 3, 4, 2, 0, 324, 325, 5, 6, 0, 0, 325, 327, 1, 0, 0, 0, 326, 323,
		1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 335, 3, 30,
		15, 0, 329, 335, 3, 70, 35, 0, 330, 331, 5, 1, 0, 0, 331, 332, 3, 64, 32,
		0, 332, 333, 5, 2, 0, 0, 333, 335, 1, 0, 0, 0, 334, 322, 1, 0, 0, 0, 334,
		326, 1, 0, 0, 0, 334, 329, 1, 0, 0, 0, 334, 330, 1, 0, 0, 0, 335, 69, 1,
		0, 0, 0, 336, 337, 3, 4, 2, 0, 337, 338, 5, 6, 0, 0, 338, 340, 1, 0, 0,
		0, 339, 336, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341,
		342, 3, 30, 15, 0, 342, 344, 5, 1, 0, 0, 343, 345, 3, 62, 31, 0, 344, 343,
		1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 347, 5, 2,
		0, 0, 347, 71, 1, 0, 0, 0, 36, 80, 88, 96, 108, 112, 115, 127, 131, 146,
		150, 156, 163, 172, 176, 185, 196, 209, 222, 226, 231, 235, 245, 252, 254,
		259, 264, 268, 272, 291, 312, 314, 320, 326, 334, 339, 344,
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

// GrobbitParserInit initializes any static state used to implement GrobbitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrobbitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrobbitParserInit() {
	staticData := &GrobbitParserParserStaticData
	staticData.once.Do(grobbitparserParserInit)
}

// NewGrobbitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrobbitParser(input antlr.TokenStream) *GrobbitParser {
	GrobbitParserInit()
	this := new(GrobbitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrobbitParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GrobbitParser.g4"

	return this
}

// GrobbitParser tokens.
const (
	GrobbitParserEOF             = antlr.TokenEOF
	GrobbitParserTtLParen        = 1
	GrobbitParserTtRParen        = 2
	GrobbitParserTtLBrace        = 3
	GrobbitParserTtRBrace        = 4
	GrobbitParserTtComma         = 5
	GrobbitParserTtPeriod        = 6
	GrobbitParserTtSemicolon     = 7
	GrobbitParserTtOpAdd         = 8
	GrobbitParserTtOpSub         = 9
	GrobbitParserTtOpMul         = 10
	GrobbitParserTtOpDiv         = 11
	GrobbitParserTtOpMod         = 12
	GrobbitParserTtOpAssign      = 13
	GrobbitParserTtOpAnd         = 14
	GrobbitParserTtOpOr          = 15
	GrobbitParserTtOpNot         = 16
	GrobbitParserTtOpEq          = 17
	GrobbitParserTtOpLt          = 18
	GrobbitParserTtOpGt          = 19
	GrobbitParserTtOpNe          = 20
	GrobbitParserTtOpLe          = 21
	GrobbitParserTtOpGe          = 22
	GrobbitParserTtOpDefine      = 23
	GrobbitParserTtKwBreak       = 24
	GrobbitParserTtKwCase        = 25
	GrobbitParserTtKwChan        = 26
	GrobbitParserTtKwConst       = 27
	GrobbitParserTtKwContinue    = 28
	GrobbitParserTtKwDefault     = 29
	GrobbitParserTtKwDefer       = 30
	GrobbitParserTtKwElse        = 31
	GrobbitParserTtKwFallthrough = 32
	GrobbitParserTtKwFor         = 33
	GrobbitParserTtKwFunc        = 34
	GrobbitParserTtKwGo          = 35
	GrobbitParserTtKwGoto        = 36
	GrobbitParserTtKwIf          = 37
	GrobbitParserTtKwImport      = 38
	GrobbitParserTtKwInterface   = 39
	GrobbitParserTtKwMap         = 40
	GrobbitParserTtKwPackage     = 41
	GrobbitParserTtKwRange       = 42
	GrobbitParserTtKwReturn      = 43
	GrobbitParserTtKwSelect      = 44
	GrobbitParserTtKwStruct      = 45
	GrobbitParserTtKwSwitch      = 46
	GrobbitParserTtKwType        = 47
	GrobbitParserTtKwVar         = 48
	GrobbitParserTtBool          = 49
	GrobbitParserTtIdentifier    = 50
	GrobbitParserTtFloat         = 51
	GrobbitParserTtInt           = 52
	GrobbitParserTtString        = 53
	GrobbitParserTtChar          = 54
	GrobbitParserWS              = 55
	GrobbitParserBLOCK_COMMENT   = 56
	GrobbitParserLINE_COMMENT    = 57
	GrobbitParserTtUnknown       = 58
)

// GrobbitParser rules.
const (
	GrobbitParserRULE_program        = 0
	GrobbitParserRULE_top_decl       = 1
	GrobbitParserRULE_pkg_name       = 2
	GrobbitParserRULE_import_decl    = 3
	GrobbitParserRULE_import_spec    = 4
	GrobbitParserRULE_const_decl     = 5
	GrobbitParserRULE_const_spec     = 6
	GrobbitParserRULE_var_decl       = 7
	GrobbitParserRULE_var_spec       = 8
	GrobbitParserRULE_basic_literals = 9
	GrobbitParserRULE_basic_literal  = 10
	GrobbitParserRULE_func_decl      = 11
	GrobbitParserRULE_param_decl     = 12
	GrobbitParserRULE_param_spec     = 13
	GrobbitParserRULE_identifiers    = 14
	GrobbitParserRULE_identifier     = 15
	GrobbitParserRULE_type           = 16
	GrobbitParserRULE_block          = 17
	GrobbitParserRULE_statement      = 18
	GrobbitParserRULE_decl_stmt      = 19
	GrobbitParserRULE_simple_stmt    = 20
	GrobbitParserRULE_return_stmt    = 21
	GrobbitParserRULE_continue_stmt  = 22
	GrobbitParserRULE_break_stmt     = 23
	GrobbitParserRULE_if_stmt        = 24
	GrobbitParserRULE_for_stmt       = 25
	GrobbitParserRULE_for_clause     = 26
	GrobbitParserRULE_condition      = 27
	GrobbitParserRULE_shortvar_decl  = 28
	GrobbitParserRULE_assign_stmt    = 29
	GrobbitParserRULE_expr_stmt      = 30
	GrobbitParserRULE_expressions    = 31
	GrobbitParserRULE_expression     = 32
	GrobbitParserRULE_unary_expr     = 33
	GrobbitParserRULE_primary_expr   = 34
	GrobbitParserRULE_call_expr      = 35
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwPackage() antlr.TerminalNode
	Pkg_name() IPkg_nameContext
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllImport_decl() []IImport_declContext
	Import_decl(i int) IImport_declContext
	AllTop_decl() []ITop_declContext
	Top_decl(i int) ITop_declContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) TtKwPackage() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwPackage, 0)
}

func (s *ProgramContext) Pkg_name() IPkg_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkg_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkg_nameContext)
}

func (s *ProgramContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *ProgramContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrobbitParserEOF, 0)
}

func (s *ProgramContext) AllImport_decl() []IImport_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_declContext); ok {
			len++
		}
	}

	tst := make([]IImport_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_declContext); ok {
			tst[i] = t.(IImport_declContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Import_decl(i int) IImport_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_declContext); ok {
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

	return t.(IImport_declContext)
}

func (s *ProgramContext) AllTop_decl() []ITop_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITop_declContext); ok {
			len++
		}
	}

	tst := make([]ITop_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITop_declContext); ok {
			tst[i] = t.(ITop_declContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Top_decl(i int) ITop_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITop_declContext); ok {
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

	return t.(ITop_declContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrobbitParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(GrobbitParserTtKwPackage)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Pkg_name()
	}
	{
		p.SetState(74)
		p.Match(GrobbitParserTtSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrobbitParserTtKwImport {
		{
			p.SetState(75)
			p.Import_decl()
		}
		{
			p.SetState(76)
			p.Match(GrobbitParserTtSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281492290797568) != 0 {
		{
			p.SetState(83)
			p.Top_decl()
		}
		{
			p.SetState(84)
			p.Match(GrobbitParserTtSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(GrobbitParserEOF)
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

// ITop_declContext is an interface to support dynamic dispatch.
type ITop_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_decl() IConst_declContext
	Var_decl() IVar_declContext
	Func_decl() IFunc_declContext

	// IsTop_declContext differentiates from other interfaces.
	IsTop_declContext()
}

type Top_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_declContext() *Top_declContext {
	var p = new(Top_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_top_decl
	return p
}

func InitEmptyTop_declContext(p *Top_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_top_decl
}

func (*Top_declContext) IsTop_declContext() {}

func NewTop_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_declContext {
	var p = new(Top_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_top_decl

	return p
}

func (s *Top_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Top_declContext) Const_decl() IConst_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_declContext)
}

func (s *Top_declContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Top_declContext) Func_decl() IFunc_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declContext)
}

func (s *Top_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Top_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitTop_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Top_decl() (localctx ITop_declContext) {
	localctx = NewTop_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrobbitParserRULE_top_decl)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtKwConst:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Const_decl()
		}

	case GrobbitParserTtKwVar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Var_decl()
		}

	case GrobbitParserTtKwFunc:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.Func_decl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IPkg_nameContext is an interface to support dynamic dispatch.
type IPkg_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsPkg_nameContext differentiates from other interfaces.
	IsPkg_nameContext()
}

type Pkg_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkg_nameContext() *Pkg_nameContext {
	var p = new(Pkg_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_pkg_name
	return p
}

func InitEmptyPkg_nameContext(p *Pkg_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_pkg_name
}

func (*Pkg_nameContext) IsPkg_nameContext() {}

func NewPkg_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pkg_nameContext {
	var p = new(Pkg_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_pkg_name

	return p
}

func (s *Pkg_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Pkg_nameContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Pkg_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pkg_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pkg_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitPkg_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Pkg_name() (localctx IPkg_nameContext) {
	localctx = NewPkg_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrobbitParserRULE_pkg_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Identifier()
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

// IImport_declContext is an interface to support dynamic dispatch.
type IImport_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwImport() antlr.TerminalNode
	AllImport_spec() []IImport_specContext
	Import_spec(i int) IImport_specContext
	TtLParen() antlr.TerminalNode
	TtRParen() antlr.TerminalNode
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode

	// IsImport_declContext differentiates from other interfaces.
	IsImport_declContext()
}

type Import_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_declContext() *Import_declContext {
	var p = new(Import_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_import_decl
	return p
}

func InitEmptyImport_declContext(p *Import_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_import_decl
}

func (*Import_declContext) IsImport_declContext() {}

func NewImport_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_declContext {
	var p = new(Import_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_import_decl

	return p
}

func (s *Import_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_declContext) TtKwImport() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwImport, 0)
}

func (s *Import_declContext) AllImport_spec() []IImport_specContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImport_specContext); ok {
			len++
		}
	}

	tst := make([]IImport_specContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImport_specContext); ok {
			tst[i] = t.(IImport_specContext)
			i++
		}
	}

	return tst
}

func (s *Import_declContext) Import_spec(i int) IImport_specContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImport_specContext); ok {
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

	return t.(IImport_specContext)
}

func (s *Import_declContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Import_declContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Import_declContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *Import_declContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *Import_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitImport_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Import_decl() (localctx IImport_declContext) {
	localctx = NewImport_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrobbitParserRULE_import_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GrobbitParserTtKwImport)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtIdentifier, GrobbitParserTtString:
		{
			p.SetState(101)
			p.Import_spec()
		}

	case GrobbitParserTtLParen:
		{
			p.SetState(102)
			p.Match(GrobbitParserTtLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrobbitParserTtIdentifier || _la == GrobbitParserTtString {
			{
				p.SetState(103)
				p.Import_spec()
			}
			{
				p.SetState(104)
				p.Match(GrobbitParserTtSemicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(111)
			p.Match(GrobbitParserTtRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IImport_specContext is an interface to support dynamic dispatch.
type IImport_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtString() antlr.TerminalNode
	Pkg_name() IPkg_nameContext

	// IsImport_specContext differentiates from other interfaces.
	IsImport_specContext()
}

type Import_specContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_specContext() *Import_specContext {
	var p = new(Import_specContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_import_spec
	return p
}

func InitEmptyImport_specContext(p *Import_specContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_import_spec
}

func (*Import_specContext) IsImport_specContext() {}

func NewImport_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_specContext {
	var p = new(Import_specContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_import_spec

	return p
}

func (s *Import_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_specContext) TtString() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtString, 0)
}

func (s *Import_specContext) Pkg_name() IPkg_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkg_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkg_nameContext)
}

func (s *Import_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitImport_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Import_spec() (localctx IImport_specContext) {
	localctx = NewImport_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrobbitParserRULE_import_spec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtIdentifier {
		{
			p.SetState(114)
			p.Pkg_name()
		}

	}
	{
		p.SetState(117)
		p.Match(GrobbitParserTtString)
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

// IConst_declContext is an interface to support dynamic dispatch.
type IConst_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwConst() antlr.TerminalNode
	AllConst_spec() []IConst_specContext
	Const_spec(i int) IConst_specContext
	TtLParen() antlr.TerminalNode
	TtRParen() antlr.TerminalNode
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode

	// IsConst_declContext differentiates from other interfaces.
	IsConst_declContext()
}

type Const_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_declContext() *Const_declContext {
	var p = new(Const_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_const_decl
	return p
}

func InitEmptyConst_declContext(p *Const_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_const_decl
}

func (*Const_declContext) IsConst_declContext() {}

func NewConst_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_declContext {
	var p = new(Const_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_const_decl

	return p
}

func (s *Const_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_declContext) TtKwConst() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwConst, 0)
}

func (s *Const_declContext) AllConst_spec() []IConst_specContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_specContext); ok {
			len++
		}
	}

	tst := make([]IConst_specContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_specContext); ok {
			tst[i] = t.(IConst_specContext)
			i++
		}
	}

	return tst
}

func (s *Const_declContext) Const_spec(i int) IConst_specContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_specContext); ok {
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

	return t.(IConst_specContext)
}

func (s *Const_declContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Const_declContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Const_declContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *Const_declContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *Const_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitConst_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Const_decl() (localctx IConst_declContext) {
	localctx = NewConst_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrobbitParserRULE_const_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GrobbitParserTtKwConst)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtIdentifier:
		{
			p.SetState(120)
			p.Const_spec()
		}

	case GrobbitParserTtLParen:
		{
			p.SetState(121)
			p.Match(GrobbitParserTtLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrobbitParserTtIdentifier {
			{
				p.SetState(122)
				p.Const_spec()
			}
			{
				p.SetState(123)
				p.Match(GrobbitParserTtSemicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(130)
			p.Match(GrobbitParserTtRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IConst_specContext is an interface to support dynamic dispatch.
type IConst_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext
	Type_() ITypeContext
	TtOpAssign() antlr.TerminalNode
	Basic_literals() IBasic_literalsContext

	// IsConst_specContext differentiates from other interfaces.
	IsConst_specContext()
}

type Const_specContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_specContext() *Const_specContext {
	var p = new(Const_specContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_const_spec
	return p
}

func InitEmptyConst_specContext(p *Const_specContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_const_spec
}

func (*Const_specContext) IsConst_specContext() {}

func NewConst_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_specContext {
	var p = new(Const_specContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_const_spec

	return p
}

func (s *Const_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_specContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *Const_specContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Const_specContext) TtOpAssign() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAssign, 0)
}

func (s *Const_specContext) Basic_literals() IBasic_literalsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_literalsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasic_literalsContext)
}

func (s *Const_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitConst_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Const_spec() (localctx IConst_specContext) {
	localctx = NewConst_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrobbitParserRULE_const_spec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Identifiers()
	}
	{
		p.SetState(134)
		p.Type_()
	}
	{
		p.SetState(135)
		p.Match(GrobbitParserTtOpAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Basic_literals()
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

// IVar_declContext is an interface to support dynamic dispatch.
type IVar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwVar() antlr.TerminalNode
	AllVar_spec() []IVar_specContext
	Var_spec(i int) IVar_specContext
	TtLParen() antlr.TerminalNode
	TtRParen() antlr.TerminalNode
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode

	// IsVar_declContext differentiates from other interfaces.
	IsVar_declContext()
}

type Var_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declContext() *Var_declContext {
	var p = new(Var_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_var_decl
	return p
}

func InitEmptyVar_declContext(p *Var_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_var_decl
}

func (*Var_declContext) IsVar_declContext() {}

func NewVar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declContext {
	var p = new(Var_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_var_decl

	return p
}

func (s *Var_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declContext) TtKwVar() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwVar, 0)
}

func (s *Var_declContext) AllVar_spec() []IVar_specContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_specContext); ok {
			len++
		}
	}

	tst := make([]IVar_specContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_specContext); ok {
			tst[i] = t.(IVar_specContext)
			i++
		}
	}

	return tst
}

func (s *Var_declContext) Var_spec(i int) IVar_specContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_specContext); ok {
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

	return t.(IVar_specContext)
}

func (s *Var_declContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Var_declContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Var_declContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *Var_declContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *Var_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitVar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Var_decl() (localctx IVar_declContext) {
	localctx = NewVar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrobbitParserRULE_var_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(GrobbitParserTtKwVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtIdentifier:
		{
			p.SetState(139)
			p.Var_spec()
		}

	case GrobbitParserTtLParen:
		{
			p.SetState(140)
			p.Match(GrobbitParserTtLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GrobbitParserTtIdentifier {
			{
				p.SetState(141)
				p.Var_spec()
			}
			{
				p.SetState(142)
				p.Match(GrobbitParserTtSemicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(GrobbitParserTtRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IVar_specContext is an interface to support dynamic dispatch.
type IVar_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext
	Type_() ITypeContext
	TtOpAssign() antlr.TerminalNode
	Basic_literals() IBasic_literalsContext

	// IsVar_specContext differentiates from other interfaces.
	IsVar_specContext()
}

type Var_specContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_specContext() *Var_specContext {
	var p = new(Var_specContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_var_spec
	return p
}

func InitEmptyVar_specContext(p *Var_specContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_var_spec
}

func (*Var_specContext) IsVar_specContext() {}

func NewVar_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_specContext {
	var p = new(Var_specContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_var_spec

	return p
}

func (s *Var_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_specContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *Var_specContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Var_specContext) TtOpAssign() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAssign, 0)
}

func (s *Var_specContext) Basic_literals() IBasic_literalsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_literalsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasic_literalsContext)
}

func (s *Var_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitVar_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Var_spec() (localctx IVar_specContext) {
	localctx = NewVar_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrobbitParserRULE_var_spec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Identifiers()
	}
	{
		p.SetState(153)
		p.Type_()
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtOpAssign {
		{
			p.SetState(154)
			p.Match(GrobbitParserTtOpAssign)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Basic_literals()
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

// IBasic_literalsContext is an interface to support dynamic dispatch.
type IBasic_literalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBasic_literal() []IBasic_literalContext
	Basic_literal(i int) IBasic_literalContext
	AllTtComma() []antlr.TerminalNode
	TtComma(i int) antlr.TerminalNode

	// IsBasic_literalsContext differentiates from other interfaces.
	IsBasic_literalsContext()
}

type Basic_literalsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasic_literalsContext() *Basic_literalsContext {
	var p = new(Basic_literalsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_basic_literals
	return p
}

func InitEmptyBasic_literalsContext(p *Basic_literalsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_basic_literals
}

func (*Basic_literalsContext) IsBasic_literalsContext() {}

func NewBasic_literalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Basic_literalsContext {
	var p = new(Basic_literalsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_basic_literals

	return p
}

func (s *Basic_literalsContext) GetParser() antlr.Parser { return s.parser }

func (s *Basic_literalsContext) AllBasic_literal() []IBasic_literalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBasic_literalContext); ok {
			len++
		}
	}

	tst := make([]IBasic_literalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBasic_literalContext); ok {
			tst[i] = t.(IBasic_literalContext)
			i++
		}
	}

	return tst
}

func (s *Basic_literalsContext) Basic_literal(i int) IBasic_literalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_literalContext); ok {
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

	return t.(IBasic_literalContext)
}

func (s *Basic_literalsContext) AllTtComma() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtComma)
}

func (s *Basic_literalsContext) TtComma(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtComma, i)
}

func (s *Basic_literalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Basic_literalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Basic_literalsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitBasic_literals(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Basic_literals() (localctx IBasic_literalsContext) {
	localctx = NewBasic_literalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrobbitParserRULE_basic_literals)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Basic_literal()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrobbitParserTtComma {
		{
			p.SetState(159)
			p.Match(GrobbitParserTtComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Basic_literal()
		}

		p.SetState(165)
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

// IBasic_literalContext is an interface to support dynamic dispatch.
type IBasic_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtInt() antlr.TerminalNode
	TtFloat() antlr.TerminalNode
	TtBool() antlr.TerminalNode

	// IsBasic_literalContext differentiates from other interfaces.
	IsBasic_literalContext()
}

type Basic_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasic_literalContext() *Basic_literalContext {
	var p = new(Basic_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_basic_literal
	return p
}

func InitEmptyBasic_literalContext(p *Basic_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_basic_literal
}

func (*Basic_literalContext) IsBasic_literalContext() {}

func NewBasic_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Basic_literalContext {
	var p = new(Basic_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_basic_literal

	return p
}

func (s *Basic_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Basic_literalContext) TtInt() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtInt, 0)
}

func (s *Basic_literalContext) TtFloat() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtFloat, 0)
}

func (s *Basic_literalContext) TtBool() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtBool, 0)
}

func (s *Basic_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Basic_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Basic_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitBasic_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Basic_literal() (localctx IBasic_literalContext) {
	localctx = NewBasic_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrobbitParserRULE_basic_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7318349394477056) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IFunc_declContext is an interface to support dynamic dispatch.
type IFunc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwFunc() antlr.TerminalNode
	Identifier() IIdentifierContext
	TtLParen() antlr.TerminalNode
	TtRParen() antlr.TerminalNode
	Block() IBlockContext
	Param_decl() IParam_declContext
	Type_() ITypeContext

	// IsFunc_declContext differentiates from other interfaces.
	IsFunc_declContext()
}

type Func_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declContext() *Func_declContext {
	var p = new(Func_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_func_decl
	return p
}

func InitEmptyFunc_declContext(p *Func_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_func_decl
}

func (*Func_declContext) IsFunc_declContext() {}

func NewFunc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declContext {
	var p = new(Func_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_func_decl

	return p
}

func (s *Func_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declContext) TtKwFunc() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwFunc, 0)
}

func (s *Func_declContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Func_declContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Func_declContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Func_declContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_declContext) Param_decl() IParam_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_declContext)
}

func (s *Func_declContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Func_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitFunc_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Func_decl() (localctx IFunc_declContext) {
	localctx = NewFunc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrobbitParserRULE_func_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(GrobbitParserTtKwFunc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Identifier()
	}
	{
		p.SetState(170)
		p.Match(GrobbitParserTtLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtIdentifier {
		{
			p.SetState(171)
			p.Param_decl()
		}

	}
	{
		p.SetState(174)
		p.Match(GrobbitParserTtRParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtIdentifier {
		{
			p.SetState(175)
			p.Type_()
		}

	}
	{
		p.SetState(178)
		p.Block()
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

// IParam_declContext is an interface to support dynamic dispatch.
type IParam_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam_spec() []IParam_specContext
	Param_spec(i int) IParam_specContext
	AllTtComma() []antlr.TerminalNode
	TtComma(i int) antlr.TerminalNode

	// IsParam_declContext differentiates from other interfaces.
	IsParam_declContext()
}

type Param_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_declContext() *Param_declContext {
	var p = new(Param_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_param_decl
	return p
}

func InitEmptyParam_declContext(p *Param_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_param_decl
}

func (*Param_declContext) IsParam_declContext() {}

func NewParam_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_declContext {
	var p = new(Param_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_param_decl

	return p
}

func (s *Param_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_declContext) AllParam_spec() []IParam_specContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParam_specContext); ok {
			len++
		}
	}

	tst := make([]IParam_specContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParam_specContext); ok {
			tst[i] = t.(IParam_specContext)
			i++
		}
	}

	return tst
}

func (s *Param_declContext) Param_spec(i int) IParam_specContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_specContext); ok {
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

	return t.(IParam_specContext)
}

func (s *Param_declContext) AllTtComma() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtComma)
}

func (s *Param_declContext) TtComma(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtComma, i)
}

func (s *Param_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitParam_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Param_decl() (localctx IParam_declContext) {
	localctx = NewParam_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrobbitParserRULE_param_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Param_spec()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrobbitParserTtComma {
		{
			p.SetState(181)
			p.Match(GrobbitParserTtComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Param_spec()
		}

		p.SetState(187)
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

// IParam_specContext is an interface to support dynamic dispatch.
type IParam_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext
	Type_() ITypeContext

	// IsParam_specContext differentiates from other interfaces.
	IsParam_specContext()
}

type Param_specContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_specContext() *Param_specContext {
	var p = new(Param_specContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_param_spec
	return p
}

func InitEmptyParam_specContext(p *Param_specContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_param_spec
}

func (*Param_specContext) IsParam_specContext() {}

func NewParam_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_specContext {
	var p = new(Param_specContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_param_spec

	return p
}

func (s *Param_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_specContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *Param_specContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *Param_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitParam_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Param_spec() (localctx IParam_specContext) {
	localctx = NewParam_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrobbitParserRULE_param_spec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Identifiers()
	}
	{
		p.SetState(189)
		p.Type_()
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

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	AllTtComma() []antlr.TerminalNode
	TtComma(i int) antlr.TerminalNode

	// IsIdentifiersContext differentiates from other interfaces.
	IsIdentifiersContext()
}

type IdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiersContext() *IdentifiersContext {
	var p = new(IdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_identifiers
	return p
}

func InitEmptyIdentifiersContext(p *IdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_identifiers
}

func (*IdentifiersContext) IsIdentifiersContext() {}

func NewIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiersContext {
	var p = new(IdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_identifiers

	return p
}

func (s *IdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiersContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *IdentifiersContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *IdentifiersContext) AllTtComma() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtComma)
}

func (s *IdentifiersContext) TtComma(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtComma, i)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitIdentifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrobbitParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Identifier()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrobbitParserTtComma {
		{
			p.SetState(192)
			p.Match(GrobbitParserTtComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Identifier()
		}

		p.SetState(198)
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtIdentifier() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) TtIdentifier() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrobbitParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		p.Match(GrobbitParserTtIdentifier)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrobbitParserRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Identifier()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtLBrace() antlr.TerminalNode
	TtRBrace() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) TtLBrace() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLBrace, 0)
}

func (s *BlockContext) TtRBrace() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRBrace, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *BlockContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrobbitParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(GrobbitParserTtLBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1416317424893960) != 0 {
		{
			p.SetState(204)
			p.Statement()
		}
		{
			p.SetState(205)
			p.Match(GrobbitParserTtSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
		p.Match(GrobbitParserTtRBrace)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Simple_stmt() ISimple_stmtContext
	Return_stmt() IReturn_stmtContext
	Break_stmt() IBreak_stmtContext
	Continue_stmt() IContinue_stmtContext
	Block() IBlockContext
	If_stmt() IIf_stmtContext
	For_stmt() IFor_stmtContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StatementContext) Simple_stmt() ISimple_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *StatementContext) Return_stmt() IReturn_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_stmtContext)
}

func (s *StatementContext) Break_stmt() IBreak_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *StatementContext) Continue_stmt() IContinue_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StatementContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrobbitParserRULE_statement)
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtKwConst, GrobbitParserTtKwVar:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Decl_stmt()
		}

	case GrobbitParserTtIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Simple_stmt()
		}

	case GrobbitParserTtKwReturn:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Return_stmt()
		}

	case GrobbitParserTtKwBreak:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Break_stmt()
		}

	case GrobbitParserTtKwContinue:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.Continue_stmt()
		}

	case GrobbitParserTtLBrace:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(219)
			p.Block()
		}

	case GrobbitParserTtKwIf:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(220)
			p.If_stmt()
		}

	case GrobbitParserTtKwFor:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(221)
			p.For_stmt()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_decl() IConst_declContext
	Var_decl() IVar_declContext

	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) Const_decl() IConst_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_declContext)
}

func (s *Decl_stmtContext) Var_decl() IVar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_declContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Decl_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitDecl_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrobbitParserRULE_decl_stmt)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtKwConst:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Const_decl()
		}

	case GrobbitParserTtKwVar:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Var_decl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// ISimple_stmtContext is an interface to support dynamic dispatch.
type ISimple_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr_stmt() IExpr_stmtContext
	Assign_stmt() IAssign_stmtContext
	Shortvar_decl() IShortvar_declContext

	// IsSimple_stmtContext differentiates from other interfaces.
	IsSimple_stmtContext()
}

type Simple_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_stmtContext() *Simple_stmtContext {
	var p = new(Simple_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_simple_stmt
	return p
}

func InitEmptySimple_stmtContext(p *Simple_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_simple_stmt
}

func (*Simple_stmtContext) IsSimple_stmtContext() {}

func NewSimple_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_stmtContext {
	var p = new(Simple_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_simple_stmt

	return p
}

func (s *Simple_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_stmtContext) Expr_stmt() IExpr_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpr_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpr_stmtContext)
}

func (s *Simple_stmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *Simple_stmtContext) Shortvar_decl() IShortvar_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShortvar_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShortvar_declContext)
}

func (s *Simple_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitSimple_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Simple_stmt() (localctx ISimple_stmtContext) {
	localctx = NewSimple_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrobbitParserRULE_simple_stmt)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Expr_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Assign_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Shortvar_decl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// IReturn_stmtContext is an interface to support dynamic dispatch.
type IReturn_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwReturn() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_stmtContext differentiates from other interfaces.
	IsReturn_stmtContext()
}

type Return_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_stmtContext() *Return_stmtContext {
	var p = new(Return_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_return_stmt
	return p
}

func InitEmptyReturn_stmtContext(p *Return_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_return_stmt
}

func (*Return_stmtContext) IsReturn_stmtContext() {}

func NewReturn_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_stmtContext {
	var p = new(Return_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_return_stmt

	return p
}

func (s *Return_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_stmtContext) TtKwReturn() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwReturn, 0)
}

func (s *Return_stmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitReturn_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Return_stmt() (localctx IReturn_stmtContext) {
	localctx = NewReturn_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrobbitParserRULE_return_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(GrobbitParserTtKwReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8444249301385986) != 0 {
		{
			p.SetState(234)
			p.expression(0)
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

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwContinue() antlr.TerminalNode

	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_continue_stmt
	return p
}

func InitEmptyContinue_stmtContext(p *Continue_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_continue_stmt
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_stmtContext) TtKwContinue() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwContinue, 0)
}

func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitContinue_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Continue_stmt() (localctx IContinue_stmtContext) {
	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrobbitParserRULE_continue_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(GrobbitParserTtKwContinue)
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

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwBreak() antlr.TerminalNode

	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_break_stmt
	return p
}

func InitEmptyBreak_stmtContext(p *Break_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_break_stmt
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_stmtContext) TtKwBreak() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwBreak, 0)
}

func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitBreak_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Break_stmt() (localctx IBreak_stmtContext) {
	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrobbitParserRULE_break_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(GrobbitParserTtKwBreak)
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

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwIf() antlr.TerminalNode
	Condition() IConditionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Simple_stmt() ISimple_stmtContext
	TtSemicolon() antlr.TerminalNode
	TtKwElse() antlr.TerminalNode
	If_stmt() IIf_stmtContext

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) TtKwIf() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwIf, 0)
}

func (s *If_stmtContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *If_stmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *If_stmtContext) Simple_stmt() ISimple_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_stmtContext)
}

func (s *If_stmtContext) TtSemicolon() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, 0)
}

func (s *If_stmtContext) TtKwElse() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwElse, 0)
}

func (s *If_stmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GrobbitParserRULE_if_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(GrobbitParserTtKwIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(242)
			p.Simple_stmt()
		}
		{
			p.SetState(243)
			p.Match(GrobbitParserTtSemicolon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(247)
		p.Condition()
	}
	{
		p.SetState(248)
		p.Block()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtKwElse {
		{
			p.SetState(249)
			p.Match(GrobbitParserTtKwElse)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GrobbitParserTtKwIf:
			{
				p.SetState(250)
				p.If_stmt()
			}

		case GrobbitParserTtLBrace:
			{
				p.SetState(251)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TtKwFor() antlr.TerminalNode
	Block() IBlockContext
	Condition() IConditionContext
	For_clause() IFor_clauseContext

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) TtKwFor() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtKwFor, 0)
}

func (s *For_stmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *For_stmtContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *For_stmtContext) For_clause() IFor_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_clauseContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitFor_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GrobbitParserRULE_for_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Match(GrobbitParserTtKwFor)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(257)
			p.Condition()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(258)
			p.For_clause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(261)
		p.Block()
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

// IFor_clauseContext is an interface to support dynamic dispatch.
type IFor_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() ISimple_stmtContext

	// GetPost returns the post rule contexts.
	GetPost() ISimple_stmtContext

	// SetInit sets the init rule contexts.
	SetInit(ISimple_stmtContext)

	// SetPost sets the post rule contexts.
	SetPost(ISimple_stmtContext)

	// Getter signatures
	AllTtSemicolon() []antlr.TerminalNode
	TtSemicolon(i int) antlr.TerminalNode
	Condition() IConditionContext
	AllSimple_stmt() []ISimple_stmtContext
	Simple_stmt(i int) ISimple_stmtContext

	// IsFor_clauseContext differentiates from other interfaces.
	IsFor_clauseContext()
}

type For_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	init   ISimple_stmtContext
	post   ISimple_stmtContext
}

func NewEmptyFor_clauseContext() *For_clauseContext {
	var p = new(For_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_for_clause
	return p
}

func InitEmptyFor_clauseContext(p *For_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_for_clause
}

func (*For_clauseContext) IsFor_clauseContext() {}

func NewFor_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_clauseContext {
	var p = new(For_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_for_clause

	return p
}

func (s *For_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *For_clauseContext) GetInit() ISimple_stmtContext { return s.init }

func (s *For_clauseContext) GetPost() ISimple_stmtContext { return s.post }

func (s *For_clauseContext) SetInit(v ISimple_stmtContext) { s.init = v }

func (s *For_clauseContext) SetPost(v ISimple_stmtContext) { s.post = v }

func (s *For_clauseContext) AllTtSemicolon() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtSemicolon)
}

func (s *For_clauseContext) TtSemicolon(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtSemicolon, i)
}

func (s *For_clauseContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *For_clauseContext) AllSimple_stmt() []ISimple_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_stmtContext); ok {
			len++
		}
	}

	tst := make([]ISimple_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_stmtContext); ok {
			tst[i] = t.(ISimple_stmtContext)
			i++
		}
	}

	return tst
}

func (s *For_clauseContext) Simple_stmt(i int) ISimple_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_stmtContext); ok {
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

	return t.(ISimple_stmtContext)
}

func (s *For_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitFor_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) For_clause() (localctx IFor_clauseContext) {
	localctx = NewFor_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GrobbitParserRULE_for_clause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtIdentifier {
		{
			p.SetState(263)

			var _x = p.Simple_stmt()

			localctx.(*For_clauseContext).init = _x
		}

	}
	{
		p.SetState(266)
		p.Match(GrobbitParserTtSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8444249301385986) != 0 {
		{
			p.SetState(267)
			p.Condition()
		}

	}
	{
		p.SetState(270)
		p.Match(GrobbitParserTtSemicolon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GrobbitParserTtIdentifier {
		{
			p.SetState(271)

			var _x = p.Simple_stmt()

			localctx.(*For_clauseContext).post = _x
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GrobbitParserRULE_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.expression(0)
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

// IShortvar_declContext is an interface to support dynamic dispatch.
type IShortvar_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext
	TtOpDefine() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsShortvar_declContext differentiates from other interfaces.
	IsShortvar_declContext()
}

type Shortvar_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortvar_declContext() *Shortvar_declContext {
	var p = new(Shortvar_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_shortvar_decl
	return p
}

func InitEmptyShortvar_declContext(p *Shortvar_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_shortvar_decl
}

func (*Shortvar_declContext) IsShortvar_declContext() {}

func NewShortvar_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shortvar_declContext {
	var p = new(Shortvar_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_shortvar_decl

	return p
}

func (s *Shortvar_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Shortvar_declContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *Shortvar_declContext) TtOpDefine() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpDefine, 0)
}

func (s *Shortvar_declContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Shortvar_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shortvar_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shortvar_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitShortvar_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Shortvar_decl() (localctx IShortvar_declContext) {
	localctx = NewShortvar_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GrobbitParserRULE_shortvar_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Identifiers()
	}
	{
		p.SetState(277)
		p.Match(GrobbitParserTtOpDefine)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.Expressions()
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

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifiers() IIdentifiersContext
	TtOpAssign() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) Identifiers() IIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifiersContext)
}

func (s *Assign_stmtContext) TtOpAssign() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAssign, 0)
}

func (s *Assign_stmtContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitAssign_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GrobbitParserRULE_assign_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Identifiers()
	}
	{
		p.SetState(281)
		p.Match(GrobbitParserTtOpAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Expressions()
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

// IExpr_stmtContext is an interface to support dynamic dispatch.
type IExpr_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call_expr() ICall_exprContext

	// IsExpr_stmtContext differentiates from other interfaces.
	IsExpr_stmtContext()
}

type Expr_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpr_stmtContext() *Expr_stmtContext {
	var p = new(Expr_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expr_stmt
	return p
}

func InitEmptyExpr_stmtContext(p *Expr_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expr_stmt
}

func (*Expr_stmtContext) IsExpr_stmtContext() {}

func NewExpr_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_stmtContext {
	var p = new(Expr_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_expr_stmt

	return p
}

func (s *Expr_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_stmtContext) Call_expr() ICall_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *Expr_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitExpr_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Expr_stmt() (localctx IExpr_stmtContext) {
	localctx = NewExpr_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GrobbitParserRULE_expr_stmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Call_expr()
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

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllTtComma() []antlr.TerminalNode
	TtComma(i int) antlr.TerminalNode

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expressions
	return p
}

func InitEmptyExpressionsContext(p *ExpressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expressions
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionsContext) AllTtComma() []antlr.TerminalNode {
	return s.GetTokens(GrobbitParserTtComma)
}

func (s *ExpressionsContext) TtComma(i int) antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtComma, i)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GrobbitParserRULE_expressions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.expression(0)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrobbitParserTtComma {
		{
			p.SetState(287)
			p.Match(GrobbitParserTtComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.expression(0)
		}

		p.SetState(293)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Unary_expr() IUnary_exprContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	TtOpMul() antlr.TerminalNode
	TtOpDiv() antlr.TerminalNode
	TtOpMod() antlr.TerminalNode
	TtOpAdd() antlr.TerminalNode
	TtOpSub() antlr.TerminalNode
	TtOpEq() antlr.TerminalNode
	TtOpNe() antlr.TerminalNode
	TtOpLt() antlr.TerminalNode
	TtOpLe() antlr.TerminalNode
	TtOpGt() antlr.TerminalNode
	TtOpGe() antlr.TerminalNode
	TtOpAnd() antlr.TerminalNode
	TtOpOr() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionContext) Unary_expr() IUnary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) TtOpMul() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpMul, 0)
}

func (s *ExpressionContext) TtOpDiv() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpDiv, 0)
}

func (s *ExpressionContext) TtOpMod() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpMod, 0)
}

func (s *ExpressionContext) TtOpAdd() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAdd, 0)
}

func (s *ExpressionContext) TtOpSub() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpSub, 0)
}

func (s *ExpressionContext) TtOpEq() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpEq, 0)
}

func (s *ExpressionContext) TtOpNe() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpNe, 0)
}

func (s *ExpressionContext) TtOpLt() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpLt, 0)
}

func (s *ExpressionContext) TtOpLe() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpLe, 0)
}

func (s *ExpressionContext) TtOpGt() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpGt, 0)
}

func (s *ExpressionContext) TtOpGe() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpGe, 0)
}

func (s *ExpressionContext) TtOpAnd() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAnd, 0)
}

func (s *ExpressionContext) TtOpOr() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpOr, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GrobbitParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 64
	p.EnterRecursionRule(localctx, 64, GrobbitParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Unary_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GrobbitParserRULE_expression)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(298)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(299)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GrobbitParserRULE_expression)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(301)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrobbitParserTtOpAdd || _la == GrobbitParserTtOpSub) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(302)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GrobbitParserRULE_expression)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(304)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8257536) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(305)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GrobbitParserRULE_expression)
				p.SetState(306)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(307)

					var _m = p.Match(GrobbitParserTtOpAnd)

					localctx.(*ExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(308)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GrobbitParserRULE_expression)
				p.SetState(309)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(310)

					var _m = p.Match(GrobbitParserTtOpOr)

					localctx.(*ExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(311)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnary_exprContext is an interface to support dynamic dispatch.
type IUnary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	Primary_expr() IPrimary_exprContext
	Unary_expr() IUnary_exprContext
	TtOpAdd() antlr.TerminalNode
	TtOpSub() antlr.TerminalNode
	TtOpNot() antlr.TerminalNode

	// IsUnary_exprContext differentiates from other interfaces.
	IsUnary_exprContext()
}

type Unary_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyUnary_exprContext() *Unary_exprContext {
	var p = new(Unary_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_unary_expr
	return p
}

func InitEmptyUnary_exprContext(p *Unary_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_unary_expr
}

func (*Unary_exprContext) IsUnary_exprContext() {}

func NewUnary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_exprContext {
	var p = new(Unary_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_unary_expr

	return p
}

func (s *Unary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_exprContext) GetOp() antlr.Token { return s.op }

func (s *Unary_exprContext) SetOp(v antlr.Token) { s.op = v }

func (s *Unary_exprContext) Primary_expr() IPrimary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *Unary_exprContext) Unary_expr() IUnary_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_exprContext)
}

func (s *Unary_exprContext) TtOpAdd() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpAdd, 0)
}

func (s *Unary_exprContext) TtOpSub() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpSub, 0)
}

func (s *Unary_exprContext) TtOpNot() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtOpNot, 0)
}

func (s *Unary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitUnary_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Unary_expr() (localctx IUnary_exprContext) {
	localctx = NewUnary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GrobbitParserRULE_unary_expr)
	var _la int

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrobbitParserTtLParen, GrobbitParserTtBool, GrobbitParserTtIdentifier, GrobbitParserTtFloat, GrobbitParserTtInt:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Primary_expr()
		}

	case GrobbitParserTtOpAdd, GrobbitParserTtOpSub, GrobbitParserTtOpNot:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Unary_exprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66304) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Unary_exprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(319)
			p.Unary_expr()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IPrimary_exprContext is an interface to support dynamic dispatch.
type IPrimary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Basic_literal() IBasic_literalContext
	Identifier() IIdentifierContext
	Pkg_name() IPkg_nameContext
	TtPeriod() antlr.TerminalNode
	Call_expr() ICall_exprContext
	TtLParen() antlr.TerminalNode
	Expression() IExpressionContext
	TtRParen() antlr.TerminalNode

	// IsPrimary_exprContext differentiates from other interfaces.
	IsPrimary_exprContext()
}

type Primary_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_exprContext() *Primary_exprContext {
	var p = new(Primary_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_primary_expr
	return p
}

func InitEmptyPrimary_exprContext(p *Primary_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_primary_expr
}

func (*Primary_exprContext) IsPrimary_exprContext() {}

func NewPrimary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_exprContext {
	var p = new(Primary_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_primary_expr

	return p
}

func (s *Primary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_exprContext) Basic_literal() IBasic_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasic_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasic_literalContext)
}

func (s *Primary_exprContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Primary_exprContext) Pkg_name() IPkg_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkg_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkg_nameContext)
}

func (s *Primary_exprContext) TtPeriod() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtPeriod, 0)
}

func (s *Primary_exprContext) Call_expr() ICall_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_exprContext)
}

func (s *Primary_exprContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Primary_exprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Primary_exprContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Primary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitPrimary_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Primary_expr() (localctx IPrimary_exprContext) {
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GrobbitParserRULE_primary_expr)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Basic_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(326)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(323)
				p.Pkg_name()
			}
			{
				p.SetState(324)
				p.Match(GrobbitParserTtPeriod)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(328)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(329)
			p.Call_expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(330)
			p.Match(GrobbitParserTtLParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.expression(0)
		}
		{
			p.SetState(332)
			p.Match(GrobbitParserTtRParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
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

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	TtLParen() antlr.TerminalNode
	TtRParen() antlr.TerminalNode
	Pkg_name() IPkg_nameContext
	TtPeriod() antlr.TerminalNode
	Expressions() IExpressionsContext

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_call_expr
	return p
}

func InitEmptyCall_exprContext(p *Call_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrobbitParserRULE_call_expr
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrobbitParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Call_exprContext) TtLParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtLParen, 0)
}

func (s *Call_exprContext) TtRParen() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtRParen, 0)
}

func (s *Call_exprContext) Pkg_name() IPkg_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPkg_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPkg_nameContext)
}

func (s *Call_exprContext) TtPeriod() antlr.TerminalNode {
	return s.GetToken(GrobbitParserTtPeriod, 0)
}

func (s *Call_exprContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrobbitParserVisitor:
		return t.VisitCall_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrobbitParser) Call_expr() (localctx ICall_exprContext) {
	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GrobbitParserRULE_call_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(336)
			p.Pkg_name()
		}
		{
			p.SetState(337)
			p.Match(GrobbitParserTtPeriod)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(341)
		p.Identifier()
	}
	{
		p.SetState(342)
		p.Match(GrobbitParserTtLParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8444249301385986) != 0 {
		{
			p.SetState(343)
			p.Expressions()
		}

	}
	{
		p.SetState(346)
		p.Match(GrobbitParserTtRParen)
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

func (p *GrobbitParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 32:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrobbitParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
