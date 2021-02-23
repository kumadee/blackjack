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

var playerCases = []struct {
	description               string
	player                    Player
	expectedStatsOutput       string
	expectedCardsInHandOutput string
}{
	{
		description: "Player zero score",
		player: Player{
			name:         "P0",
			currentScore: 10,
			wins:         0,
			loss:         1,
		},
		expectedStatsOutput: `
Player P0 stats:
CurrentScore: 10, Wins: 0, Loss: 1
`,
		expectedCardsInHandOutput: `
Player P0 cards in hand:
`,
	},
}

func TestShowCardsInHand(t *testing.T) {
	for _, tc := range playerCases {
		output := captureOutput(tc.player.ShowCardsInHand)
		assert.Equalf(t, tc.expectedCardsInHandOutput, output, "%s", tc.description)
	}
}

func TestShowStats(t *testing.T) {
	for _, tc := range playerCases {
		output := captureOutput(tc.player.ShowStats)
		assert.Equalf(t, tc.expectedStatsOutput, output, "%s", tc.description)
	}
}
