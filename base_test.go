package token2

import (
	"encoding/json"
	"fmt"
	"testing"
)

type MapValue struct {
	Type   int
	String string
	True   interface{}
}

type MapToken struct {
	Type       int
	Properties map[string]MapValue
}

func TestNewBaseToken(t *testing.T) {
	// mt := &MapToken{
	// 	Type: 2,
	// 	Properties: map[string]MapValue{
	// 		"value": MapValue{
	// 			Type:   1,
	// 			String: "7",
	// 			True:   7,
	// 		},
	// 	},
	// }
	boolVar := NewBoolFromBool(false)
	tokens := []Token{
		NewBaseToken(),
		boolVar,
	}

	tokenJSON, _ := json.Marshal(tokens)
	fmt.Println("tokenJSON", string(tokenJSON))
}
