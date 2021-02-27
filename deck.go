package blackjack

// Card holds the value and face of a physical card
// dealt in the game of blackjack.
type Card struct {
	face  rune
	value uint
}

// Deck represents a list of cards.
type Deck []Card

// DiscardPile represents the cards removed from Deck.
type DiscardPile []int

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
	13: 'K',
}

// NewDeck returns a collection of cards arranged in
// a manner.
func NewDeck(size uint) Deck {
	deck := make(Deck, size)
	for i := 0; i < len(deck); i++ {
		value := uint((i + 1) % 13)
		// since we get a remainder of 13
		// the value should be 13 in card
		if value == 0 {
			value = 13
		}
		face := valueToFaceCardMap[value]
		deck[i] = Card{face, value}
	}
	return deck
}

// CardExistsInDiscardPile return true if card exists in
// discarded pile of cards otherwise false.
func CardExistsInDiscardPile(cardIndex int, pile DiscardPile) bool {
	for _, discardValue := range pile {
		// negative card index is returned true, because
		// we want to always draw a new card in this case.
		if cardIndex == discardValue || cardIndex < 0 {
			return true
		} else if discardValue < 0 {
			// If negative value is reached, then values after it
			// in the slice will also be negative. No need to check
			// further.
			return false
		}
	}
	return false
}

// ResetDiscardPile removes cards from discard pile
func ResetDiscardPile(pile DiscardPile) {
	for i := 0; i < len(pile); i++ {
		pile[i] = -1
	}
}

// MoveCardToDiscardPile moves card to discard pile
func MoveCardToDiscardPile(cardIndex int, pile DiscardPile) {
	for i := 0; i < len(pile); i++ {
		if pile[i] < 0 {
			pile[i] = cardIndex
			return
		}
	}
}

// CheckDiscardPileFull checks if the dicard pile is full
// i.e., no further cards can be dealt from deck
func CheckDiscardPileFull(pile DiscardPile) bool {
	for _, discardValue := range pile {
		if discardValue < 0 {
			// negative value represents that discard pile is not full
			// and some cards can be dealt from the deck
			return false
		}
	}
	return true
}
