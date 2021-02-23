package blackjack

import (
	"fmt"
	"strings"
)

// Player represents the user playing the game
type Player struct {
	name         string
	currentScore int
	wins         int
	loss         int
	cardsInHand  []Card
}

var cardSeparator rune = '|'

// ShowCardsInHand prints the cards of the player
func (p *Player) ShowCardsInHand() {
	var builder strings.Builder
	var format string = `
Player %s cards in hand:`
	fmt.Fprintf(&builder, format, p.name)
	for _, card := range p.cardsInHand {
		builder.WriteRune(card.face)
		builder.WriteRune(cardSeparator)
	}
	fmt.Println(builder.String())
}

// ShowStats displays the stats of the player
func (p *Player) ShowStats() {
	var builder strings.Builder
	var format string = `
Player %s stats:
CurrentScore: %d, Wins: %d, Loss: %d`
	fmt.Fprintf(&builder, format, p.name, p.currentScore, p.wins, p.loss)
	fmt.Println(builder.String())
}
