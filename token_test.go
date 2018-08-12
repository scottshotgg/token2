package token2_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/scottshotgg/token2"
)

func TestNewToken(r *testing.T) {
	intToken := token2.NewIntFromInt(77)
	intTokenJSON, err := json.Marshal(intToken)
	if err != nil {

	}
	fmt.Println(string(intTokenJSON))

	err = intToken.ToIdent("name")
	if err != nil {

	}

	intTokenJSON, err = json.Marshal(intToken)
	if err != nil {

	}
	fmt.Println(string(intTokenJSON))
}
