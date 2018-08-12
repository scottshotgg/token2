package token2

func newLiteral(valueType ValueType, trueValue interface{}) *Token {
	return newToken("", Literal, &Value{
		Type:  valueType,
		True:  trueValue,
		Props: map[string]Token{},
	})
}

// Function for instantiating default values
// TODO: need to add void?

func NewInt() *Token {
	// return newLiteral(IntValue, 0)
	return NewIntFromInt(0)
}

func NewBool() *Token {
	return NewBoolFromBool(false)
}

func NewChar() *Token {
	return NewCharFromChar([1]string{})
}

func NewFloat() *Token {
	return NewFloatFromFloat(0)
}

func NewString() *Token {
	return NewStringFromString("")
}

// NewVar creates a new var with a default value
// vars are default initialized to an int type
// although this function should never be used directly
func NewVar() *Token {
	return NewVarFromInt(0)
}

// TODO: need to add arrays, objects, and structs

func NewIntFromInt(newInt int) *Token {
	return newLiteral(IntValue, newInt)
}

func NewBoolFromBool(b bool) *Token {
	return newLiteral(BoolValue, b)
}

func NewCharFromChar(newChar [1]string) *Token {
	return newLiteral(CharValue, newChar)
}

func NewFloatFromFloat(newFloat float64) *Token {
	return newLiteral(FloatValue, newFloat)
}

func NewStringFromString(newString string) *Token {
	return newLiteral(StringValue, newString)
}

// func NewVarFromVar(newVar interface{}) *Token {
// 	return newLiteral(VarValue, newVar)
// }

// NewVarFromInt creates a new var type from an integer
func NewVarFromInt(newInt int) *Token {
	return newLiteral(VarValue, newInt)
}

// TODO: need to make all of these when needed
