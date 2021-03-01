package blackjack

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func(io.Reader), rd io.Reader) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, err := io.Copy(&buf, reader)
		if err == nil {
			out <- buf.String()
		}
	}()
	wg.Wait()
	f(rd)
	writer.Close()
	return <-out
}

func TestStartGame(t *testing.T) {
	var gameCases = []struct {
		description   string
		expectedOuput string
		rd            io.Reader
	}{
		{
			description: "New Game 0",
			expectedOuput: `
Player Human cards in hand:

Player Human stats:
CurrentScore: 0, Wins: 0, Loss: 0

Player CPU cards in hand:

Player CPU stats:
CurrentScore: 0, Wins: 0, Loss: 0
`,
			rd: strings.NewReader("quit\n"),
		},
	}
	for _, tc := range gameCases {
		output := captureOutput(StartGame, tc.rd)
		assert.Contains(t, output, tc.expectedOuput, tc.description)
	}
}

func TestDealCardFromDeck(t *testing.T) {
	var cases = []struct {
		description         string
		inputGame           Game
		expectedDiscardPile DiscardPile
	}{
		{
			description: "Deal cards from a deck with 13 cards.",
			inputGame: Game{
				deck:    NewDeck(13),
				discard: make(DiscardPile, 13),
			},
			expectedDiscardPile: DiscardPile{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}
	for _, tc := range cases {
		actual := make(Deck, len(tc.inputGame.deck))
		for i := 0; i < len(tc.inputGame.deck); i++ {
			actual[i] = tc.inputGame.DealCardFromDeck()
		}
		assert.ElementsMatchf(t, tc.inputGame.deck, actual,
			"%s - deck test", tc.description)
		assert.ElementsMatchf(t, tc.expectedDiscardPile, tc.inputGame.discard,
			"%s - discard pile test", tc.description)
	}
}

func TestRoundWinner(t *testing.T) {
	var cases = []struct {
		description    string
		game           Game
		expectedWinner Player
		expectedError  error
	}{
		{
			description:    "CPU score more than 21 but greater than human's score.",
			expectedWinner: Player{name: "Human", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 10},
					{name: "CPU", currentScore: 25},
				},
			},
		},
		{
			description:    "Human score more than 21 but greater than CPU's score.",
			expectedWinner: Player{name: "CPU", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 25},
					{name: "CPU", currentScore: 10},
				},
			},
		},
		{
			description:    "Human score greater than CPU and both less than 21.",
			expectedWinner: Player{name: "Human", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 11},
					{name: "CPU", currentScore: 10},
				},
			},
		},
		{
			description:    "CPU score greater than Human and both less than 21.",
			expectedWinner: Player{name: "CPU", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 17},
					{name: "CPU", currentScore: 20},
				},
			},
		},
		{
			description:    "CPU score greater than Human and equal to 21.",
			expectedWinner: Player{name: "CPU", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 17},
					{name: "CPU", currentScore: 21},
				},
			},
		},
		{
			description:    "Human score greater than CPU and equal to 21.",
			expectedWinner: Player{name: "Human", loss: 0, wins: 1},
			expectedError:  nil,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 21},
					{name: "CPU", currentScore: 14},
				},
			},
		},
		{
			description:    "No winner.",
			expectedWinner: Player{loss: 0, wins: 1},
			expectedError:  ErrNoWinner,
			game: Game{
				players: []Player{
					{name: "Human", currentScore: 25},
					{name: "CPU", currentScore: 22},
				},
			},
		},
	}
	for _, tc := range cases {
		name, err := tc.game.RoundWinner()
		if err == nil {
			assert.Equalf(t, tc.expectedWinner.name, name, "%s - winner name test", tc.description)
		}
		for _, player := range tc.game.players {
			if player.name == tc.expectedWinner.name {
				assert.Equalf(t, tc.expectedWinner.loss, player.loss, "%s - Winner loss count test", tc.description)
				assert.Equalf(t, tc.expectedWinner.wins, player.wins, "%s - Winner wins count test", tc.description)
			} else {
				assert.Equalf(t, tc.expectedWinner.loss+1, player.loss, "%s - Loser loss count test", tc.description)
				assert.Equalf(t, tc.expectedWinner.wins-1, player.wins, "%s - Loser wins count test", tc.description)
			}
		}
		assert.Equalf(t, tc.expectedError, err, "%s - error test", tc.description)
	}
}

func BenchmarkStartGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// benchmark for one round
		StartGame(strings.NewReader("stay\nquit\n"))
	}
}
