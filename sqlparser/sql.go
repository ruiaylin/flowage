//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const NAMES = 57413
const REPLACE = 57414
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const PROXY = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const TABLE = 57424
const INDEX = 57425
const VIEW = 57426
const TO = 57427
const IGNORE = 57428
const IF = 57429
const UNIQUE = 57430
const USING = 57431

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
	"PROXY",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 215
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 584

var yyAct = []int{

	109, 315, 146, 382, 70, 185, 350, 117, 106, 269,
	136, 225, 260, 307, 186, 3, 107, 95, 200, 195,
	72, 160, 161, 96, 262, 391, 391, 208, 391, 88,
	155, 313, 155, 59, 155, 84, 223, 223, 100, 283,
	284, 285, 286, 287, 74, 288, 289, 79, 61, 44,
	81, 46, 73, 77, 85, 47, 49, 112, 50, 93,
	276, 252, 116, 332, 334, 122, 101, 34, 35, 36,
	37, 139, 99, 113, 114, 115, 91, 393, 392, 135,
	390, 104, 340, 312, 302, 120, 300, 143, 253, 222,
	336, 361, 138, 147, 360, 150, 359, 153, 78, 157,
	152, 80, 333, 214, 103, 51, 75, 187, 118, 119,
	97, 188, 52, 53, 54, 123, 56, 57, 58, 261,
	191, 149, 212, 254, 67, 215, 194, 74, 145, 341,
	74, 198, 294, 204, 203, 73, 159, 132, 73, 121,
	160, 161, 183, 184, 205, 261, 127, 305, 160, 161,
	202, 153, 60, 242, 218, 343, 134, 71, 101, 231,
	204, 356, 229, 308, 219, 235, 272, 142, 240, 241,
	151, 244, 245, 246, 247, 248, 249, 250, 251, 230,
	236, 221, 83, 358, 211, 213, 210, 173, 174, 175,
	232, 125, 101, 101, 128, 243, 357, 74, 74, 233,
	234, 265, 129, 201, 330, 73, 267, 256, 258, 273,
	264, 326, 144, 308, 324, 329, 327, 268, 328, 325,
	274, 74, 129, 223, 201, 278, 279, 367, 345, 73,
	34, 35, 36, 37, 264, 229, 228, 86, 18, 293,
	271, 296, 297, 280, 277, 227, 281, 154, 18, 19,
	20, 21, 131, 18, 295, 377, 376, 375, 116, 148,
	101, 122, 171, 172, 173, 174, 175, 129, 75, 113,
	114, 115, 304, 220, 196, 301, 22, 148, 314, 311,
	310, 120, 292, 228, 197, 197, 192, 158, 190, 229,
	229, 155, 227, 322, 323, 189, 94, 124, 291, 60,
	306, 339, 75, 60, 118, 119, 369, 370, 342, 337,
	335, 123, 319, 318, 74, 388, 347, 217, 216, 348,
	351, 199, 346, 68, 140, 137, 27, 28, 29, 352,
	30, 32, 31, 389, 126, 121, 23, 24, 26, 25,
	133, 362, 130, 82, 38, 364, 363, 344, 168, 169,
	170, 171, 172, 173, 174, 175, 87, 66, 153, 18,
	299, 373, 365, 371, 40, 41, 42, 43, 395, 379,
	351, 206, 141, 381, 380, 55, 383, 383, 383, 74,
	384, 385, 263, 386, 257, 18, 112, 73, 64, 62,
	316, 116, 396, 372, 122, 374, 397, 355, 398, 317,
	112, 99, 113, 114, 115, 116, 270, 354, 122, 237,
	104, 238, 239, 321, 120, 75, 113, 114, 115, 283,
	284, 285, 286, 287, 104, 288, 289, 89, 120, 201,
	92, 69, 39, 103, 394, 378, 18, 118, 119, 97,
	90, 17, 16, 15, 123, 112, 14, 103, 13, 12,
	116, 118, 119, 122, 207, 45, 275, 209, 123, 48,
	75, 113, 114, 115, 116, 76, 266, 122, 121, 104,
	387, 255, 368, 120, 75, 113, 114, 115, 349, 353,
	320, 303, 121, 148, 193, 259, 111, 120, 108, 110,
	309, 105, 103, 162, 102, 331, 118, 119, 226, 163,
	167, 165, 166, 123, 282, 224, 98, 290, 156, 63,
	118, 119, 33, 366, 65, 11, 10, 123, 179, 180,
	181, 182, 9, 176, 177, 178, 8, 121, 168, 169,
	170, 171, 172, 173, 174, 175, 7, 6, 5, 4,
	2, 121, 1, 0, 0, 164, 168, 169, 170, 171,
	172, 173, 174, 175, 338, 0, 0, 168, 169, 170,
	171, 172, 173, 174, 175, 298, 0, 0, 168, 169,
	170, 171, 172, 173, 174, 175, 168, 169, 170, 171,
	172, 173, 174, 175,
}
var yyPact = []int{

	243, -1000, -1000, 181, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -48, -43, 8, 15, -1000, -1000, -1000,
	-1000, 26, 264, 431, 372, -1000, -1000, -1000, 370, -1000,
	328, 288, 422, 71, -49, 0, 264, -1000, 4, 264,
	-1000, 308, -67, 264, -67, 327, 417, 421, 264, 252,
	-1000, -1000, -1000, 37, -1000, 258, 288, 301, 70, 288,
	169, 307, -1000, 207, -1000, 61, 305, 89, 264, -1000,
	290, -1000, -29, 289, 352, 103, 264, 288, -1000, 425,
	439, 417, 439, 421, 439, 238, -1000, -1000, 268, 60,
	83, 478, -1000, 425, 380, -1000, -1000, -1000, 439, 251,
	244, -1000, 242, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 439, -1000, 241, 267, 286, 419, 267,
	-1000, 439, 264, -1000, 351, -77, -1000, 90, -1000, 283,
	-1000, -1000, 282, -1000, 240, 83, 478, 508, 233, -1000,
	508, 417, -16, 508, 201, 37, -1000, -1000, 264, 117,
	425, 425, 439, 215, 388, 439, 439, 128, 439, 439,
	439, 439, 439, 439, 439, 439, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -44, -17, 18, 478, -1000, 366,
	37, -1000, 431, 40, 508, 354, 267, 267, 214, -1000,
	393, 425, -1000, 508, -1000, -1000, -1000, 102, 264, -1000,
	-40, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 354,
	267, -1000, -1000, 439, 193, 365, 263, 248, 56, -1000,
	-1000, -1000, -1000, -1000, -1000, 508, -1000, 215, 439, 439,
	508, 500, -1000, 335, 191, 191, 191, 114, 114, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -19, 37, -21, 66,
	-1000, 425, 99, 215, 181, 149, -22, -1000, 393, 375,
	385, 83, 278, -1000, -1000, 277, -1000, -1000, 169, 508,
	402, 201, 201, -1000, -1000, 160, 157, 164, 161, 150,
	1, -1000, 275, -15, 274, -1000, 508, 489, 439, -1000,
	-1000, -23, -1000, 47, -1000, 439, 75, -1000, 317, 175,
	-1000, -1000, -1000, 267, 375, -1000, 439, 439, -1000, -1000,
	395, 383, 365, 97, -1000, 142, -1000, 129, -1000, -1000,
	-1000, -1000, -2, -4, -7, -1000, -1000, -1000, 439, 508,
	-1000, -1000, 508, 439, 314, 215, -1000, -1000, 460, 174,
	-1000, 280, -1000, 393, 425, 439, 425, -1000, -1000, 213,
	212, 211, 508, 508, 428, -1000, 439, 439, -1000, -1000,
	-1000, 375, 83, 170, 83, 264, 264, 264, 267, 508,
	-1000, 299, -25, -1000, -27, -28, 169, -1000, 427, 347,
	-1000, 264, -1000, -1000, -1000, 264, -1000, 264, -1000,
}
var yyPgo = []int{

	0, 542, 540, 14, 539, 538, 537, 536, 526, 522,
	516, 515, 344, 514, 512, 509, 17, 23, 508, 507,
	506, 505, 11, 504, 498, 124, 495, 3, 18, 38,
	494, 493, 24, 491, 2, 16, 5, 490, 489, 7,
	488, 8, 486, 485, 12, 484, 481, 480, 479, 9,
	478, 6, 472, 1, 470, 19, 466, 13, 4, 20,
	182, 465, 459, 457, 456, 455, 454, 0, 10, 449,
	448, 446, 443, 442, 441, 76, 29, 432,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 73, 8, 8, 8, 9, 9,
	9, 10, 11, 11, 11, 77, 12, 13, 13, 14,
	14, 14, 14, 14, 15, 15, 16, 16, 17, 17,
	17, 20, 20, 18, 18, 18, 21, 21, 22, 22,
	22, 22, 19, 19, 19, 23, 23, 23, 23, 23,
	23, 23, 23, 23, 24, 24, 24, 25, 25, 26,
	26, 26, 26, 27, 27, 28, 28, 76, 76, 76,
	75, 75, 29, 29, 29, 29, 29, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 30, 31, 31, 31,
	31, 31, 31, 31, 32, 32, 37, 37, 35, 35,
	39, 36, 36, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	38, 38, 40, 40, 40, 42, 45, 45, 43, 43,
	44, 46, 46, 41, 41, 33, 33, 33, 33, 47,
	47, 48, 48, 49, 49, 50, 50, 51, 52, 52,
	52, 53, 53, 53, 54, 54, 54, 55, 55, 56,
	56, 57, 57, 58, 58, 59, 60, 60, 61, 61,
	62, 62, 63, 63, 63, 63, 63, 64, 64, 65,
	65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 3, 4, 5, 5, 8, 4, 6, 7,
	4, 5, 4, 5, 5, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 0, 2, 2,
	0, 2, 1, 3, 3, 2, 3, 3, 3, 4,
	3, 4, 5, 6, 3, 4, 2, 1, 1, 1,
	1, 1, 1, 1, 2, 1, 1, 3, 3, 1,
	3, 1, 3, 1, 1, 1, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 4, 5, 4, 1,
	1, 1, 1, 1, 1, 5, 0, 1, 1, 2,
	4, 0, 2, 1, 3, 1, 1, 1, 1, 0,
	3, 0, 2, 0, 3, 1, 3, 2, 0, 1,
	1, 0, 2, 4, 0, 2, 4, 0, 3, 1,
	3, 0, 5, 1, 3, 3, 0, 2, 0, 3,
	0, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 93, 94, 96, 95, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -77,
	-12, -12, -12, -12, 97, -65, 99, 103, -62, 99,
	101, 97, 97, 98, 99, -12, 90, 91, 92, -67,
	35, -3, 17, -15, 18, -13, 29, -25, 35, 9,
	-58, 86, -59, -41, -67, 35, -61, 102, 98, -67,
	97, -67, 35, -60, 102, -67, -60, 29, -76, 10,
	23, -75, 9, -67, 44, -16, -17, 73, -20, 35,
	-29, -34, -30, 67, 44, -33, -41, -35, -40, -67,
	-38, -42, 20, 36, 37, 38, 25, -39, 71, 72,
	48, 102, 28, 78, 39, -25, 33, 76, -25, 53,
	35, 45, 76, 35, 67, -67, -68, 35, -68, 100,
	35, 20, 64, -67, -25, -29, -34, -34, 44, -76,
	-34, -75, -36, -34, 9, 53, -18, -67, 19, 76,
	65, 66, -31, 21, 67, 23, 24, 22, 68, 69,
	70, 71, 72, 73, 74, 75, 45, 46, 47, 40,
	41, 42, 43, -29, -29, -36, -3, -34, -34, 44,
	44, -39, 44, -45, -34, -55, 33, 44, -58, 35,
	-28, 10, -59, -34, -67, -68, 20, -66, 104, -63,
	96, 94, 32, 95, 13, 35, 35, 35, -68, -55,
	33, -76, 105, 53, -21, -22, -24, 44, 35, -39,
	-17, -67, 73, -29, -29, -34, -35, 21, 23, 24,
	-34, -34, 25, 67, -34, -34, -34, -34, -34, -34,
	-34, -34, 105, 105, 105, 105, -16, 18, -16, -43,
	-44, 79, -32, 28, -3, -58, -56, -41, -28, -49,
	13, -29, 64, -67, -68, -64, 100, -32, -58, -34,
	-28, 53, -23, 54, 55, 56, 57, 58, 60, 61,
	-19, 35, 19, -22, 76, -35, -34, -34, 65, 25,
	105, -16, 105, -46, -44, 81, -29, -57, 64, -37,
	-35, -57, 105, 53, -49, -53, 15, 14, 35, 35,
	-47, 11, -22, -22, 54, 59, 54, 59, 54, 54,
	54, -26, 62, 101, 63, 35, 105, 35, 65, -34,
	105, 82, -34, 80, 30, 53, -41, -53, -34, -50,
	-51, -34, -68, -48, 12, 14, 64, 54, 54, 98,
	98, 98, -34, -34, 31, -35, 53, 53, -52, 26,
	27, -49, -29, -36, -29, 44, 44, 44, 7, -34,
	-51, -53, -27, -67, -27, -27, -58, -54, 16, 34,
	105, 53, 105, 105, 7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 209, 200, 0, 0, 28, 29, 30,
	45, 0, 0, 0, 49, 51, 52, 53, 54, 47,
	0, 0, 0, 0, 198, 0, 0, 210, 0, 0,
	201, 0, 196, 0, 196, 0, 97, 100, 0, 0,
	213, 19, 50, 0, 55, 46, 0, 0, 87, 0,
	26, 0, 193, 0, 163, 213, 0, 0, 0, 214,
	0, 214, 0, 0, 0, 0, 0, 0, 32, 0,
	0, 97, 0, 100, 0, 17, 56, 58, 63, 213,
	61, 62, 102, 0, 0, 133, 134, 135, 0, 163,
	0, 149, 0, 165, 166, 167, 168, 129, 152, 153,
	154, 150, 151, 156, 48, 187, 0, 0, 95, 0,
	27, 0, 0, 214, 0, 211, 37, 0, 40, 0,
	42, 197, 0, 214, 187, 98, 0, 99, 0, 33,
	101, 97, 0, 131, 0, 0, 59, 64, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 117, 118, 119, 120,
	121, 122, 123, 105, 0, 0, 0, 131, 144, 0,
	0, 116, 0, 0, 157, 0, 0, 0, 95, 88,
	173, 0, 194, 195, 164, 35, 199, 0, 0, 214,
	207, 202, 203, 204, 205, 206, 41, 43, 44, 0,
	0, 34, 31, 0, 95, 66, 72, 0, 84, 86,
	57, 65, 60, 103, 104, 107, 108, 0, 0, 0,
	110, 0, 114, 0, 136, 137, 138, 139, 140, 141,
	142, 143, 106, 128, 130, 145, 0, 0, 0, 161,
	158, 0, 191, 0, 125, 191, 0, 189, 173, 181,
	0, 96, 0, 212, 38, 0, 208, 22, 23, 132,
	169, 0, 0, 75, 76, 0, 0, 0, 0, 0,
	89, 73, 0, 0, 0, 109, 111, 0, 0, 115,
	146, 0, 148, 0, 159, 0, 0, 20, 0, 124,
	126, 21, 188, 0, 181, 25, 0, 0, 214, 39,
	171, 0, 67, 70, 77, 0, 79, 0, 81, 82,
	83, 68, 0, 0, 0, 74, 69, 85, 0, 112,
	147, 155, 162, 0, 0, 0, 190, 24, 182, 174,
	175, 178, 36, 173, 0, 0, 0, 78, 80, 0,
	0, 0, 113, 160, 0, 127, 0, 0, 177, 179,
	180, 181, 172, 170, 71, 0, 0, 0, 0, 183,
	176, 184, 0, 93, 0, 0, 192, 18, 0, 0,
	90, 0, 91, 92, 185, 0, 94, 0, 186,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 105, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:170
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:176
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:196
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:200
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:204
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:211
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:215
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:227
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:231
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:244
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:250
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:256
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:260
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:266
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:272
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:278
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:284
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 32:
		//line sql.y:290
		{
			yyVAL.statement = &Show{Section: "databases", LikeOrWhere: yyS[yypt-0].expr}
		}
	case 33:
		//line sql.y:294
		{
			yyVAL.statement = &Show{Section: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 34:
		//line sql.y:298
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyS[yypt-2].bytes), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 35:
		//line sql.y:304
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 36:
		//line sql.y:308
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:313
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 38:
		//line sql.y:319
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 39:
		//line sql.y:323
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:328
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 41:
		//line sql.y:334
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:340
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 43:
		//line sql.y:344
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 44:
		//line sql.y:349
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 45:
		//line sql.y:354
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		//line sql.y:358
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		//line sql.y:364
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		//line sql.y:368
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 49:
		//line sql.y:374
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		//line sql.y:378
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		//line sql.y:382
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		//line sql.y:386
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		//line sql.y:390
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		//line sql.y:395
		{
			yyVAL.str = ""
		}
	case 55:
		//line sql.y:399
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		//line sql.y:405
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 57:
		//line sql.y:409
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 58:
		//line sql.y:415
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		//line sql.y:419
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 60:
		//line sql.y:423
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 61:
		//line sql.y:429
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 62:
		//line sql.y:433
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 63:
		//line sql.y:438
		{
			yyVAL.bytes = nil
		}
	case 64:
		//line sql.y:442
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		//line sql.y:446
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:452
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 67:
		//line sql.y:456
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 68:
		//line sql.y:462
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 69:
		//line sql.y:466
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 70:
		//line sql.y:470
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 71:
		//line sql.y:474
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 72:
		//line sql.y:479
		{
			yyVAL.bytes = nil
		}
	case 73:
		//line sql.y:483
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		//line sql.y:487
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:493
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		//line sql.y:497
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		//line sql.y:501
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		//line sql.y:505
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		//line sql.y:509
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		//line sql.y:513
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		//line sql.y:517
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		//line sql.y:521
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		//line sql.y:525
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		//line sql.y:531
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:535
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:539
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 87:
		//line sql.y:545
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:549
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:554
		{
			yyVAL.indexHints = nil
		}
	case 90:
		//line sql.y:558
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:562
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:566
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 93:
		//line sql.y:572
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 94:
		//line sql.y:576
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 95:
		//line sql.y:581
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		//line sql.y:585
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 97:
		//line sql.y:590
		{
			yyVAL.expr = nil
		}
	case 98:
		//line sql.y:594
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 99:
		//line sql.y:598
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 100:
		//line sql.y:603
		{
			yyVAL.valExpr = nil
		}
	case 101:
		//line sql.y:607
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 102:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 103:
		//line sql.y:614
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 104:
		//line sql.y:618
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 105:
		//line sql.y:622
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 106:
		//line sql.y:626
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 107:
		//line sql.y:632
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 108:
		//line sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 109:
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 110:
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 111:
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 112:
		//line sql.y:652
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 113:
		//line sql.y:656
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 114:
		//line sql.y:660
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 115:
		//line sql.y:664
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 116:
		//line sql.y:668
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 117:
		//line sql.y:674
		{
			yyVAL.str = AST_EQ
		}
	case 118:
		//line sql.y:678
		{
			yyVAL.str = AST_LT
		}
	case 119:
		//line sql.y:682
		{
			yyVAL.str = AST_GT
		}
	case 120:
		//line sql.y:686
		{
			yyVAL.str = AST_LE
		}
	case 121:
		//line sql.y:690
		{
			yyVAL.str = AST_GE
		}
	case 122:
		//line sql.y:694
		{
			yyVAL.str = AST_NE
		}
	case 123:
		//line sql.y:698
		{
			yyVAL.str = AST_NSE
		}
	case 124:
		//line sql.y:704
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 125:
		//line sql.y:708
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 126:
		//line sql.y:714
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 127:
		//line sql.y:718
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 128:
		//line sql.y:724
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 129:
		//line sql.y:728
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 130:
		//line sql.y:734
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 131:
		//line sql.y:740
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:744
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 133:
		//line sql.y:750
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 134:
		//line sql.y:754
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 135:
		//line sql.y:758
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 136:
		//line sql.y:762
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 138:
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 139:
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 140:
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 141:
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 142:
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 143:
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 144:
		//line sql.y:794
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 145:
		//line sql.y:809
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 146:
		//line sql.y:813
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 147:
		//line sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 148:
		//line sql.y:821
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 149:
		//line sql.y:825
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 150:
		//line sql.y:831
		{
			yyVAL.bytes = IF_BYTES
		}
	case 151:
		//line sql.y:835
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 152:
		//line sql.y:841
		{
			yyVAL.byt = AST_UPLUS
		}
	case 153:
		//line sql.y:845
		{
			yyVAL.byt = AST_UMINUS
		}
	case 154:
		//line sql.y:849
		{
			yyVAL.byt = AST_TILDA
		}
	case 155:
		//line sql.y:855
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 156:
		//line sql.y:860
		{
			yyVAL.valExpr = nil
		}
	case 157:
		//line sql.y:864
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 158:
		//line sql.y:870
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 159:
		//line sql.y:874
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 160:
		//line sql.y:880
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 161:
		//line sql.y:885
		{
			yyVAL.valExpr = nil
		}
	case 162:
		//line sql.y:889
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 163:
		//line sql.y:895
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 164:
		//line sql.y:899
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 165:
		//line sql.y:905
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 166:
		//line sql.y:909
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 167:
		//line sql.y:913
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 168:
		//line sql.y:917
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 169:
		//line sql.y:922
		{
			yyVAL.valExprs = nil
		}
	case 170:
		//line sql.y:926
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 171:
		//line sql.y:931
		{
			yyVAL.boolExpr = nil
		}
	case 172:
		//line sql.y:935
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 173:
		//line sql.y:940
		{
			yyVAL.orderBy = nil
		}
	case 174:
		//line sql.y:944
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 175:
		//line sql.y:950
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 176:
		//line sql.y:954
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 177:
		//line sql.y:960
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 178:
		//line sql.y:965
		{
			yyVAL.str = AST_ASC
		}
	case 179:
		//line sql.y:969
		{
			yyVAL.str = AST_ASC
		}
	case 180:
		//line sql.y:973
		{
			yyVAL.str = AST_DESC
		}
	case 181:
		//line sql.y:978
		{
			yyVAL.limit = nil
		}
	case 182:
		//line sql.y:982
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 183:
		//line sql.y:986
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 184:
		//line sql.y:991
		{
			yyVAL.str = ""
		}
	case 185:
		//line sql.y:995
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 186:
		//line sql.y:999
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 187:
		//line sql.y:1012
		{
			yyVAL.columns = nil
		}
	case 188:
		//line sql.y:1016
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 189:
		//line sql.y:1022
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 190:
		//line sql.y:1026
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 191:
		//line sql.y:1031
		{
			yyVAL.updateExprs = nil
		}
	case 192:
		//line sql.y:1035
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 193:
		//line sql.y:1041
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 194:
		//line sql.y:1045
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 195:
		//line sql.y:1051
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 196:
		//line sql.y:1056
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1058
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1061
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1063
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1066
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1068
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1078
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1083
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1085
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1088
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1090
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1093
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1095
		{
			yyVAL.empty = struct{}{}
		}
	case 213:
		//line sql.y:1099
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 214:
		//line sql.y:1104
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
