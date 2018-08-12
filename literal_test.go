package token2_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/scottshotgg/token2"
)

func TestNewIntFromInt(t *testing.T) {
	intToken := token2.NewIntFromInt(77)
	intTokenJSON, err := json.Marshal(intToken)
	if err != nil {

	}

	fmt.Println(string(intTokenJSON))
}

func TestNewInt(t *testing.T) {
	intToken := token2.NewInt()
	intTokenJSON, err := json.Marshal(intToken)
	if err != nil {

	}

	fmt.Println(string(intTokenJSON))
}
