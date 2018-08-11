package token2

import "strconv"

type LiteralToken struct {
	BaseToken
	Value LiteralValue `json:"literalValue"`
}

type LiteralValue struct {
	ValueType ValueType `json:"type"`
	// TODO: do we need an acting type on a literal ??
	ActingType  ValueType   `json:"acting"`
	TrueValue   interface{} `json:"true"`
	StringValue string      `json:"string"`
}

func (dv *LiteralToken) GetValueType() ValueType {
	return dv.Value.ValueType
}

func (dv *LiteralToken) GetActingType() ValueType {
	return dv.Value.ActingType
}

// func (dv *LiteralToken) GetTrueValue() interface{} {
// 	return dv
// }

func (dv *LiteralToken) GetStringValue() string {
	return dv.Value.StringValue
}

func (l *LiteralToken) SetValue(value LiteralValue) {
	l.Value = value
}

func NewLiteral() *LiteralToken {
	return &LiteralToken{
		BaseToken: BaseToken{
			ID:        0,
			TokenType: Literal,
		},
	}
}

// Should we have the end user be able to instantiate with their own mixed values?
func NewLiteralFromValue(value LiteralValue) *LiteralToken {
	l := NewLiteral()
	l.SetValue(value)

	return l
}

// Put the int default value macro instead of 0 here
func NewInt() *LiteralToken {
	return NewIntFromInt(0)
}

func NewIntFromInt(value int) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		ValueType:   IntValue,
		TrueValue:   value,
		StringValue: strconv.Itoa(value),
	})
}

func NewBool() *LiteralToken {
	return NewBoolFromBool(false)
}

func NewBoolFromBool(value bool) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		ValueType:   BoolValue,
		TrueValue:   value,
		StringValue: strconv.FormatBool(value),
	})
}

func NewFloat() *LiteralToken {
	return NewFloatFromFloat(0.0)
}

func NewFloatFromFloat(value float64) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		ValueType:   FloatValue,
		TrueValue:   value,
		StringValue: strconv.FormatFloat(value, 'f', -1, 64),
	})
}

func NewString() *LiteralToken {
	return NewStringFromString("")
}

func NewStringFromString(value string) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		ValueType:   StringValue,
		TrueValue:   value,
		StringValue: value,
	})
}
