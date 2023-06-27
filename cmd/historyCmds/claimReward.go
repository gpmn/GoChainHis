/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"GoChainHis/executor"

	"github.com/spf13/cobra"
)

var ClaimRewardCmd = &cobra.Command{
	Use:   "claimReward",
	Short: "will claimReward nft for specified dayslot.",
	Long: `will claimReward nft for specified dayslot.  
e.g.
    GoChainHis history claimReward -d=2020-06-23
`,
	Run: func(cmd *cobra.Command, args []string) {
		executor.GetExecutor().HistoryClaimReward()
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
