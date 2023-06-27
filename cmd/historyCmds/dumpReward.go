/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"GoChainHis/executor"

	"github.com/spf13/cobra"
)

var DumpRewardCmd = &cobra.Command{
	Use:   "dumpReward",
	Short: "will dump how many card reward the earlier one (by -l option) could drain from later one (by -r option). NOTE:: the reward value is 0 Before settleCard called.",
	Long: `will dump how many card reward the earlier one (by -l option) could drain from later one (by -r option). 
NOTE:: the reward value is 0 Before settleCard called.
e.g.
this will show card reward from nft@2023-06-25 to nft@2023-06-23
    GoChainHis history dumpReward -l=2023-06-23 -r=2023-06-25

this will show reward summary only for yourself.
    GoChainHis history dumpReward
`,
	Run: func(cmd *cobra.Command, args []string) {
		executor.GetExecutor().HistoryDumpReward(argEarlierSlotStr, argLaterSlotStr)
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
