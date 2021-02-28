package blackjack

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

// Game represents the deck and dicard pile of cards.
type Game struct {
	deck    Deck
	discard DiscardPile
	players []Player
}

// StartGame runs the game loop
func StartGame() {
	cardsDeck := NewDeck(13)
	game := Game{
		deck:    cardsDeck,
		discard: make(DiscardPile, len(cardsDeck)),
		players: []Player{{name: "Human"}, {name: "CPU"}},
	}
	human := &game.players[0]
	cpu := &game.players[1]
	// Remove timestamp from logger
	log.SetFlags(0)
	log.Printf("%s", ShowPlayersStatsAndCards(game.players))
	for {
		// Deal card to CPU player
		cpu.UpdateCardsInHand(game.DealCardFromDeck())
		// Show 1 hidden and 1 face-up card of CPU player
		// Deal card to human player
		human.UpdateCardsInHand(game.DealCardFromDeck())
		//log.Printf("%s", ShowPlayersStatsAndCards(game.players))
		//log.Printf("\n%s", "Enter h ('hit') or s ('stay'):")
		break
		// Show all face-up card to human player
		// 'hit' or 'stay' loop start until both CPU & human
		// says 'stay' or deck is empty
		//	for {
		// CPU choose 'hit' or 'stay'
		// Player choose 'hit' or 'stay'
		//	}
		// Find and reveal winner
		// Update rounds win or lose for both players
	}
}

// DealCardFromDeck returns a card from deck and also moves it to discard pile
// if all cards in discard pile, then move from discard pile to deck
func (game *Game) DealCardFromDeck() Card {
	rand.Seed(time.Now().UnixNano())
	cardIndex := -1
	if CheckDiscardPileFull(game.discard) {
		ResetDiscardPile(game.discard)
	}
	for CardExistsInDiscardPile(cardIndex, game.discard) {
		cardIndex = rand.Intn(len(game.deck))
	}
	MoveCardToDiscardPile(cardIndex, game.discard)
	return game.deck[cardIndex]
}

// ShowPlayersStatsAndCards returns string with all players'
// stats and cards in hand.
func ShowPlayersStatsAndCards(players []Player) string {
	var output strings.Builder
	for _, player := range players {
		output.WriteString(player.ShowCardsInHand())
		output.WriteString(player.ShowStats())
	}
	return output.String()
}
