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

var MintAndAucCmd = &cobra.Command{
	Use:   "mintAndAuc -d=2020-06-23",
	Short: "will mintAndAuc history at specified dayslot, secretary only.",
	Long: `will mintAndAuc history at specified dayslot, secretary only.  
e.g.
    GoChainHis history mintAndAuc -d=2020-06-23
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("MintAndAucCmd - just one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		executor.GetExecutor().HistoryMintAndAuction(argDaySlots[0], argFestivals)
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
