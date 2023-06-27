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

var BidCmd = &cobra.Command{
	Use:   "bid",
	Short: "will bid nft for specified dayslot.",
	Long: `will bid nft for specified dayslot.  
e.g.
    GoChainHis history bid -d=2020-06-23
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("BidCmd - just one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		executor.GetExecutor().HistoryBid(argDaySlots[0])
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
