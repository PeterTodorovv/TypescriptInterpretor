package token

const (
	//types
	EOF     		= "EOF"
	IDENT   		= "ident"
	NUMBER  		= "number"
	ILLEGAL 		= "illegal"
	STRING  		= "string" 

	ASSIGN          = "="

	//math operators
	COMMA           = ","
	FULLSTOP        = "."
	WHITESPACE      = " "
	PLUS            = "+"
	MINUS           = "-"
	LARGER          = ">"
	SMALLER         = "<"
	EQUALS  		= "=="
	LARGEROREQUAL   = ">="
	SMALLEROREQUAL  = "<="
	NOT_EQUAL		= "!="

	OPENING_CURLY   = "{"
	CLOSING_CURLY   = "}"
	OPENING_BRACKET = "("
	CLOSING_BRAKET  = ")"
	OPENING_SQUARE  = "["
	CLOSING_SQUARE  = "]"
	EXCLAMATIONMARK = "!"
	BACKTICK        = "`"
	SINGLEQUOTE     = "'"
	DOUBLEQUOTE     = "\""
	SEMICOLUN       = ";"
	COLON           = ":"
	NEW_LINE        = "\n"

	LET             = "let"
	CONST           = "const"
	CLASS           = "class"
	IMPORT          = "import"
	RETURN          = "return"
)

var declarations = map[string]string{
	"let": LET,
	"const": CONST,
	"class": CLASS,
	"import": IMPORT,
	"return": RETURN,
}

// var operations = map[string]string{

// }

type Token struct {
	Type  string
	Value string
}

func GetIdentifier(identifier string) string {
	if token, ok := declarations[identifier]; ok {
		return token
	}

	return IDENT
}
