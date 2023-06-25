/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cardCmds

import (
	"GoChainHis/executor"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var ClaimCmd = &cobra.Command{
	Use:   "GoChainHis card claim -d=2020-06-01",
	Short: "claim the card for the bid winner, then the card will transfer from card contract to the winner.",
	Long: `claim the card for the bid winner, then the card will transfer from card contract to the winner.
e.g.
    GoChainHis card claim -d=2020-06-01
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("ClaimCmd - at least one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		executor.GetExecutor().CardClaim(argDaySlots[0])
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// voteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// voteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
