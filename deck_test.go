package blackjack

import "testing"

var deckCases = []struct {
	description string
	input       uint
}{
	{
		"Standard deck of size 52.",
		52,
	},
	{
		"Deck of size 13.",
		13,
	},
	{
		"Deck of size 0.",
		0,
	},
	{
		"Deck of size 1.",
		1,
	},
	{
		"Deck of size 12.",
		12,
	},
}

func TestNewDeck(t *testing.T) {
	for _, tc := range deckCases {
		actual := NewDeck(tc.input)
		if len(actual) != int(tc.input) {
			t.Fatalf("Fail: %s; invalid length: NewDeck(%d)=%d", tc.description, tc.input, len(actual))
		}
		for index, card := range actual {
			expectedValue := uint((index + 1) % 13)
			expectedFace := valueToFaceCardMap[expectedValue]
			if card.face != expectedFace {
				t.Fatalf("Fail: %s; invalid face card in deck: Deck[%d].face=%c, expected=%c", tc.description, index, card.face, expectedFace)
			}

			if expectedValue == 0 {
				expectedValue = 13
			}
			if card.value != expectedValue {
				t.Fatalf("Fail: %s; invalid value of card in deck: Deck[%d].value=%d, expected=%d", tc.description, index, card.value, expectedValue)
			}
		}
		t.Log("Pass:", tc.description)
	}
}
