package token2

import (
	"reflect"
	"testing"
)

func TestNewLiteral(t *testing.T) {
	tests := []struct {
		name string
		want *LiteralToken
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLiteral(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiteralToken_SetValue(t *testing.T) {
	type fields struct {
		BaseToken BaseToken
		Value     LiteralValue
	}
	type args struct {
		value LiteralValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LiteralToken{
				BaseToken: tt.fields.BaseToken,
				Value:     tt.fields.Value,
			}
			l.SetValue(tt.args.value)
		})
	}
}
