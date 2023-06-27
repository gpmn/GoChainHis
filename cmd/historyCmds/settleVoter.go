/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"GoChainHis/executor"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var SettleVoterCmd = &cobra.Command{
	Use:   "settleVote",
	Short: "settle the vote rewards for specified days.",
	Long: `settle the vote rewards for specified days. 
e.g.
    GoChainHis history settleVote -b=2020-06-01,2020-06-03,2020-06-04
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) == 0 {
			log.Printf("SettleVoteCmd - at least one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		executor.GetExecutor().HistorySettleVoterReward(argDaySlots)
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
