package blackjack

import (
	"fmt"
	"strings"
)

// Player represents the user playing the game
type Player struct {
	name         string
	currentScore uint
	wins         uint
	loss         uint
	cardsInHand  Deck
}

var cardSeparator rune = '|'

// ShowCardsInHand prints the cards of the player
func (p *Player) ShowCardsInHand() string {
	var builder strings.Builder
	var format string = `
Player %s cards in hand:`
	fmt.Fprintf(&builder, format, p.name)
	for _, card := range p.cardsInHand {
		builder.WriteRune(card.face)
		builder.WriteRune(cardSeparator)
	}
	builder.WriteRune('\n')
	return builder.String()
}

// ShowStats displays the stats of the player
func (p *Player) ShowStats() string {
	var builder strings.Builder
	var format string = `
Player %s stats:
CurrentScore: %d, Wins: %d, Loss: %d
`
	fmt.Fprintf(&builder, format, p.name, p.currentScore, p.wins, p.loss)
	return builder.String()
}

// UpdateCardsInHand updates the player's cards
// in hand and updates the current score as well
func (p *Player) UpdateCardsInHand(cards ...Card) {
	for _, card := range cards {
		p.cardsInHand = append(p.cardsInHand, card)
		p.currentScore += card.value
	}
}
