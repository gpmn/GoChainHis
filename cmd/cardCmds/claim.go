/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cardCmds

import (
	"log"
	"os"

	"GoChainHis/executor"

	"github.com/spf13/cobra"
)

// claimCmd represents the claim command
var ClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("BidCmd - just one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		executor.GetExecutor().CardClaim(argDaySlots[0])
	},
}

func init() {
	CardCmd.AddCommand(ClaimCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// claimCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// claimCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
