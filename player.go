package blackjack

import (
	"fmt"
)

// Player represents the user playing the game
type Player struct {
	name         string
	currentScore int
	wins         int
	loss         int
	gamesPlayed  int
	cardsInHand  []Card
}

// ShowCardsInHand prints the cards of the player
func (p *Player) ShowCardsInHand() {
	fmt.Println("Player", p.name, "cards in hand:")
	for _, card := range p.cardsInHand {
		fmt.Printf("%c ", card.face)
	}
}
