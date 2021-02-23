package blackjack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestStartGame(t *testing.T) {
	for _, tc := range gameCases {
		output := captureOutput(StartGame)
		assert.Equalf(t, tc.expectedOuput, output, "%s", tc.description)
	}
}
