package blackjack

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
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
	f()
	writer.Close()
	return <-out
}

func TestStartGame(t *testing.T) {
	var gameCases = []struct {
		description   string
		expectedOuput string
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
		},
	}
	for _, tc := range gameCases {
		output := captureOutput(StartGame)
		assert.Equalf(t, tc.expectedOuput, output, "%s", tc.description)
	}
}

func TestDealCardFromDeck(t *testing.T) {
	var cases = []struct {
		description string
		inputGame   Game
	}{}
	for _, tc := range cases {
		actual := make(Deck, len(tc.inputGame.deck))
		for i := 0; i < len(tc.inputGame.deck); i++ {
			actual[i] = tc.inputGame.DealCardFromDeck()
		}
		assert.ElementsMatch(t, tc.inputGame.deck, actual, tc.description)
	}
}

func BenchmarkStartGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StartGame()
	}
}
