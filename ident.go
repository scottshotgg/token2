package token2

import "errors"

// TODO: alternatively, we could enforce these types to the functions
// through enumerating new types like IntToken, StringToken, etc, and
// use the Token and Value types as an interface

// ToIdent creates a new ident token
func (t *Token) ToIdent(name string) error { // (*Token, error) {
	if name == "" {
		return errors.New("Ident name cannot be empty")
	}

	// Upgrade ident from a literal by setting the type and name
	t.SetTokenType(Ident)
	t.SetName(name)

	return nil
}

// func (t *Token) NewIdentFromInt(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }

// // TODO: these should take actual values
// func (t *Token) NewIdentFromBool(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }

// func (t *Token) NewIdentFromChar(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }

// func (t *Token) NewIdentFromFloat(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }

// func (t *Token) NewIdentFromString(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }

// func (t *Token) NewIdentFromVar(name string) error {
// 	// TODO: implement this later
// 	// if t.GetTokenType() != Literal {

// 	// }
// 	// if t.GetValue().GetValueType() != IntValue {

// 	// }

// 	return toIdent(name, t)
// }
