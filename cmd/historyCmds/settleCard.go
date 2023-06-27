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

var SettleCardCmd = &cobra.Command{
	Use:   "settleCard",
	Short: "settle the card rewards for specified days.",
	Long: `settle the card rewards for specified days. 
e.g.
    GoChainHis history settleCard -b=2020-06-01,2020-06-03,2020-06-04
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) == 1 {
			log.Printf("SettleCardCmd - only one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}

		if len(argFromSlots) == 0 {
			log.Printf("SettleCardCmd - at least one --fromSlots/-f param is required, e.g. '-d=2023-08-08,2023-08-02'.")
			os.Exit(1)
		}
		executor.GetExecutor().HistorySettleCardReward(argDaySlots[0], argFromSlots)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
