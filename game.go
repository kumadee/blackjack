package blackjack

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	// ErrNoWinner is returned when no player wins the game.
	ErrNoWinner = errors.New("no game winner")
)

// Game represents the deck and dicard pile of cards.
type Game struct {
	deck    Deck
	discard DiscardPile
	players []Player
}

// StartGame runs the game loop
func StartGame(read io.Reader) {
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

	reader := bufio.NewReader(read)
	for {
		// Deal card to CPU player
		cpu.UpdateCardsInHand(game.DealCardFromDeck())
		// Show 1 hidden and 1 face-up card of CPU player
		// Deal card to human player
		human.UpdateCardsInHand(game.DealCardFromDeck())

		// Show all face-up card to human player
		// 'hit' or 'stay' loop start until both CPU & human
		// says 'stay' or deck is empty
		humanInput := true
		cpuInput := true
	inner:
		for {
			log.Printf("%s", ShowPlayersStatsAndCards(game.players))
			if humanInput {
				log.Printf("\n%s", "Enter h ('hit') or s ('stay') or q ('quit'):")
				input, err := reader.ReadString('\n')
				if err != nil {
					panic(err)
				}
				switch strings.ToLower(strings.TrimRight(input, "\n")) {
				case "h", "hit":
					human.UpdateCardsInHand(game.DealCardFromDeck())
				case "s", "stay":
					humanInput = false
				case "q", "quit":
					return
				}
			}

			if cpuInput {
				switch CPUHit(cpu) {
				case true:
					cpu.UpdateCardsInHand(game.DealCardFromDeck())
				default:
					cpuInput = false
				}
			}

			if !(cpuInput || humanInput) {
				break inner
			}
		}
		// Find and reveal winner
		// Update rounds win or lose for both players
		name, err := game.RoundWinner()
		switch err {
		case ErrNoWinner:
			log.Println(err.Error())
		default:
			log.Printf("\nPlayer %s wins the round.\n", name)
		}
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

// RoundWinner is used to find the game round winner
func (game *Game) RoundWinner() (string, error) {
	var highScore uint = 0
	var winner *Player
	for _, player := range game.players {
		if player.currentScore > highScore && player.currentScore <= 21 {
			highScore = player.currentScore
			winner = &player
		}
		// Since there is only one winner, we intentionally update
		// the loss for all players in this loop. If we successfully
		// find the winner we deduct its loss and increments its wins.
		player.loss++
	}
	if winner == nil {
		return "", ErrNoWinner
	}
	winner.loss--
	winner.wins++
	return winner.name, nil
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

// CPUHit chooses where hit or stay for a
// CPU player. If true then hit else stay.
func CPUHit(p *Player) bool {
	if p.currentScore >= 21 {
		return false
	}
	if len(p.cardsInHand) > 2 && p.currentScore >= 17 {
		return false
	}
	return true
}
