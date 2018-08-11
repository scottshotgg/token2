package token2

type BaseToken struct {
	ID        int
	TokenType TokenType
	Expected  TokenType
}

func (b *BaseToken) GetID() int {
	return b.ID
}

func (b *BaseToken) GetTokenType() TokenType {
	return b.TokenType
}

// func (b *BaseToken) GetExpected() TokenType {
// 	return b.Expected
// }

// func (b *BaseToken) MarshalJSON() ([]byte, error) {
// 	fmt.Println(json.Marshal(*b))
// 	fmt.Println("b boy", *b)
// 	return json.Marshal(*b)
// }

func (b *BaseToken) SetID(ID int) {
	b.ID = ID
}

func (b *BaseToken) SetTokenType(TokenType TokenType) {
	b.TokenType = TokenType
}

// func (b *BaseToken) SetExpected(Expected TokenType) {
// 	b.Expected = Expected
// }

func NewBaseToken() *BaseToken {
	return &BaseToken{
		ID:        -1,
		TokenType: Base,
	}
}
