package cmd

import (
	"blackjack"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start New Game.",
	Long: `Starts a new game with a human and CPU player.
Human player can provide the input from keyboard.
CPU player is automatically inputted by the game.`,
	Run: func(cmd *cobra.Command, args []string) {
		blackjack.StartGame()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
