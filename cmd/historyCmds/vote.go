/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"log"
	"os"

	"GoChainHis/executor"

	"github.com/spf13/cobra"
)

// voteCmd represents the vote command
var VoteCmd = &cobra.Command{
	Use:   "vote",
	Short: "vote the big news in the candidates for specified dayslot.",
	Long: `vote the big news in the candidates for specified dayslot. 
The dayslot can be unix sec of the day, or string of '2023-08-08'.
The prefers is a int slice seprated with ','.

e.g.
    GoChainHis history vote -d 2023-06-06 -p 0,2,3

This example means to vote the first, the third, the 4th candidates as big news for '2023-06-06'.

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("VoteCmd - only one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}

		if len(argPrefers) != 3 {
			log.Printf("VoteCmd - --prefers/-p param is required, e.g. : -p=1,3,4")
			os.Exit(1)
		}
		executor.GetExecutor().HistoryVote(argDaySlots[0],
			[3]uint8{uint8(argPrefers[0]), uint8(argPrefers[1]), uint8(argPrefers[2])})
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
