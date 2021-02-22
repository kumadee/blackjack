package blackjack

// Card holds the value and face of a physical card
// dealt in the game of blackjack.
type Card struct {
	face  rune
	value uint
}

var valueToFaceCardMap = map[uint]rune{
	1:  'A',
	2:  '2',
	3:  '3',
	4:  '4',
	5:  '5',
	6:  '6',
	7:  '7',
	8:  '8',
	9:  '9',
	10: 'X',
	11: 'J',
	12: 'Q',
	0:  'K',
}

// NewDeck returns a collection of cards arranged in
// a manner.
func NewDeck(size uint) []Card {
	deck := make([]Card, size)
	for i := 0; i < len(deck); i++ {
		value := uint((i + 1) % 13)
		face := valueToFaceCardMap[value]
		// since we get a remainder of 13
		// the value should be 13 in card
		if value == 0 {
			value = 13
		}
		deck[i] = Card{face, value}
	}
	return deck
}
