package token2

import "github.com/pkg/errors"

type (
	TokenType          string
	ValueType          string
	AccessModifierType string

	Token struct {
		Name  string
		ID    int
		Type  TokenType
		Value *Value
	}
	// SetID(int) error
	// SetTokenType(TokenType) error
	// SetValue(Value) error

	// String() string
	// Express() string
	// Translate() string

	Value struct {
		Type  ValueType
		True  interface{}
		Props map[string]Token // might need a mutex here later
	}
	// String() string // don't know if we need this
)

func (t *Token) GetName() string         { return t.Name }
func (t *Token) GetID() int              { return t.ID }
func (t *Token) GetTokenType() TokenType { return t.Type }
func (t *Token) GetValue() *Value        { return t.Value }

func (t *Token) SetName(name string) {
	t.Name = name
}
func (t *Token) SetID(id int) {}
func (t *Token) SetTokenType(tokenType TokenType) {
	t.Type = tokenType
}
func (t *Token) SetValue(value *Value) error {
	tValue := t.GetValue().GetValueType()
	switch tValue {
	case VarValue:
		value.SetValueType(VarValue)
		fallthrough
	case SetValue:
	default:
		// TODO: this is where we would add type coersions, type degradations, etc
		if tValue != value.GetValueType() {
			return errors.Errorf("Value type does not match declared token's value type; value: %s token: %s", value.GetValueType(), tValue)
		}
	}

	t.Value = value
	return nil
}

func (v *Value) GetValueType() ValueType         { return v.Type }
func (v *Value) GetTrue() interface{}            { return v.True }
func (v *Value) GetProperty(prop string) Token   { return v.Props[prop] }
func (v *Value) GetProperties() map[string]Token { return v.Props }

func (v *Value) SetValueType(valueType ValueType)             {}
func (v *Value) SetTrue(trueValue interface{})                {}
func (v *Value) SetProperty(propName string, propValue Token) {}
func (v *Value) SetProperties(props map[string]Token)         {}

func newToken(name string, tokenType TokenType, value *Value) *Token {
	return &Token{
		Name:  name,
		Type:  tokenType,
		Value: value,
	}
}

// func newValue() Value {
// 	return Value{}
// }

// FIXME: downcase all of these
const (
	// Base         TokenType = "BASE"
	Custom       TokenType = "CUSTOM"
	Literal      TokenType = "LITERAL"
	Var          TokenType = "VAR"
	Ident        TokenType = "IDENT"
	Type         TokenType = "TYPE"
	Whitespace   TokenType = "WS"
	Attribute    TokenType = "ATTRIBUTE"
	Keyword      TokenType = "KEYWORD"
	SQL          TokenType = "SQL"
	Comma        TokenType = "COMMA"
	EOS          TokenType = "EOS"
	Separator    TokenType = "SEPARATOR"
	Bang         TokenType = "BANG"
	At           TokenType = "AT"
	Hash         TokenType = "HASH"
	Block        TokenType = "BLOCK"
	Function     TokenType = "FUNCTION"
	Group        TokenType = "GROUP"
	Array        TokenType = "ARRAY"
	Set          TokenType = "SET"
	Assign       TokenType = "ASSIGN"
	Init         TokenType = "INIT"
	PriOp        TokenType = "PRI_OP"
	SecOp        TokenType = "SEC_OP"
	Mult         TokenType = "MULT"
	LBrace       TokenType = "L_BRACE"
	LBracket     TokenType = "L_BRACKET"
	LParen       TokenType = "L_PAREN"
	LThan        TokenType = "L_THAN"
	RBrace       TokenType = "R_BRACE"
	RBracket     TokenType = "R_BRACKET"
	RParen       TokenType = "R_PAREN"
	GThan        TokenType = "G_THAN"
	DQuote       TokenType = "D_QUOTE"
	SQuote       TokenType = "S_QUOTE"
	Pipe         TokenType = "PIPE"
	Ampersand    TokenType = "AMPERSAND"
	Dollar       TokenType = "DOLLAR"
	Underscore   TokenType = "UNDERSCORE"
	QuestionMark TokenType = "QM"
	Accessor     TokenType = "ACCESSOR"
	Increment    TokenType = "INCREMENT"
)

const (
	// NotSetValue ValueType = "not_set"
	SetValue      ValueType = "set"
	VoidValue     ValueType = "void"
	IntValue      ValueType = "int"
	BoolValue     ValueType = "bool"
	CharValue     ValueType = "char"
	FloatValue    ValueType = "float"
	StringValue   ValueType = "string"
	VarValue      ValueType = "var"
	FunctionValue ValueType = "function"
	ArrayValue    ValueType = "array"
	ObjectValue   ValueType = "object"
	StructValue   ValueType = "struct"

	// IntArrayValue ValueType = "int[]"
)

const (
	PublicAccessModifier  AccessModifierType = "public"
	PrivateAccessModifier AccessModifierType = "private"
)
