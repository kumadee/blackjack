package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
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
	for _, tc := range deckCases {
		actual := NewDeck(tc.input)
		if len(actual) != int(tc.input) {
			t.Fatalf("Fail: %s; invalid length: NewDeck(%d)=%d", tc.description, tc.input, len(actual))
		}
		for index, card := range actual {
			expectedValue := uint((index + 1) % 13)
			if expectedValue == 0 {
				expectedValue = 13
			}
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

func TestResetDiscardPile(t *testing.T) {
	var cases = []struct {
		description  string
		pile         DiscardPile
		expectedPile DiscardPile
	}{
		{
			description: "Blank Pile.",
		},
		{
			description:  "Empty discard pile.",
			pile:         DiscardPile{-1, -1, -1},
			expectedPile: DiscardPile{-1, -1, -1},
		},
		{
			description:  "Discard pile is partially filled.",
			pile:         DiscardPile{4, 1, 2, -1, -1},
			expectedPile: DiscardPile{-1, -1, -1, -1, -1},
		},
		{
			description:  "Discard pile is fully filled.",
			pile:         DiscardPile{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expectedPile: DiscardPile{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
	}
	for _, tc := range cases {
		ResetDiscardPile(tc.pile)
		assert.Equal(t, tc.expectedPile, tc.pile, tc.description)
	}
}

func TestCardExistsInDiscardPile(t *testing.T) {
	var cases = []struct {
		description      string
		inputCardIndex   int
		inputDiscardPile DiscardPile
		expected         bool
	}{
		{
			description:    "Discard pile empty",
			inputCardIndex: 1,
			expected:       false,
		},
		{
			description:      "Card Exists in discard pile",
			inputCardIndex:   10,
			inputDiscardPile: DiscardPile{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expected:         true,
		},
		{
			description:      "Negative card index does not exist in pile",
			inputCardIndex:   -100,
			inputDiscardPile: DiscardPile{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expected:         true,
		},
		{
			description:      "Positive card index does not exist",
			inputCardIndex:   10,
			inputDiscardPile: DiscardPile{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1, -1},
			expected:         false,
		},
		{
			description:      "Negative card index exists in pile.",
			inputCardIndex:   -1,
			inputDiscardPile: DiscardPile{-1, -1, -1},
			expected:         true,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected,
			CardExistsInDiscardPile(tc.inputCardIndex, tc.inputDiscardPile),
			tc.description)
	}
}

func TestCheckDiscardPileFull(t *testing.T) {
	var cases = []struct {
		description string
		pile        DiscardPile
		expected    bool
	}{
		{
			description: "Discard pile empty",
			expected:    false,
		},
		{
			description: "Discard pile is full",
			pile:        DiscardPile{5, 7, 3, 4, 2, 1, 0, 6},
			expected:    true,
		},
		{
			description: "1. Dicard pile is partially filled",
			pile:        DiscardPile{5, 7, 3, 4, 2, 1, 0, -1},
			expected:    false,
		},
		{
			description: "2. Dicard pile is partially filled",
			pile:        DiscardPile{5, 7, 3, -1, 2, 1, 0, -1},
			expected:    false,
		},
		{
			description: "Dicard pile is not filled",
			pile:        DiscardPile{-100, -1, -1, -1},
			expected:    false,
		},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expected,
			CheckDiscardPileFull(tc.pile), tc.description)
	}
}
