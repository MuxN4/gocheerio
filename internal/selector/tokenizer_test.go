package selector

import "testing"

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "div.class#id",
			expected: []Token{
				{Type: TokenTypeTag, Value: "div"},
				{Type: TokenTypeClass, Value: "class"},
				{Type: TokenTypeID, Value: "id"},
			},
		},
		{
			input: "div[data-test='value']",
			expected: []Token{
				{Type: TokenTypeTag, Value: "div"},
				{Type: TokenTypeAttribute, Value: "[data-test='value']"},
			},
		},
		{
			input: "div > p",
			expected: []Token{
				{Type: TokenTypeTag, Value: "div"},
				{Type: TokenTypeCombinator, Value: ">"},
				{Type: TokenTypeTag, Value: "p"},
			},
		},
		// token type Invalid test
		{
			input: "div@",
			expected: []Token{
				{Type: TokenTypeTag, Value: "div"},
				{Type: TokenTypeInvalid, Value: ""},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			tokenizer := NewTokenizer(tc.input)
			for _, expected := range tc.expected {
				token := tokenizer.nextToken()
				if token == nil {
					t.Fatalf("Expected token %v, got nil", expected)
				}
				if token.Type != expected.Type || token.Value != expected.Value {
					t.Errorf("Expected token %v, got %v", expected, token)
				}
			}
		})
	}
}
