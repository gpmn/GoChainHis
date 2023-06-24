/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package historyCmds

import (
	"github.com/spf13/cobra"
)

var argDaySlots []string
var argFromSlots []string
var argPrefers []int
var dumpCnt int

// historyCmd represents the history command
var HistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	HistoryCmd.AddCommand(DumpCmd)
	HistoryCmd.AddCommand(SubmitCmd)
	HistoryCmd.AddCommand(VoteCmd)
	HistoryCmd.AddCommand(SettleVoterCmd)
	HistoryCmd.AddCommand(SettleCardCmd)
	HistoryCmd.AddCommand(SettleOpsCmd)
	HistoryCmd.AddCommand(ResolveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	HistoryCmd.PersistentFlags().StringSliceVarP(&argDaySlots, "dayslot", "d", nil, "the day to be displayed, can be unix sec of the day, or string of '2023-08-08'")

	VoteCmd.Flags().IntSliceVarP(&argPrefers, "prefers", "p", []int{}, "prefered big news index, must three items, e.a. : 0,1,2")

	DumpCmd.Flags().IntVarP(&dumpCnt, "count", "n", 1, "dump continuos count since specified daySlot.")
	SettleCardCmd.Flags().StringSliceVarP(&argFromSlots, "fromSlots", "f", nil, "settle rewards from the slots of cards.")
}
