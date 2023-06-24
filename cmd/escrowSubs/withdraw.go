/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package escrowSubs

import (
	"GoChainHis/executor"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var withdrawAmount float64

// withdrawCmd represents the withdraw command
var WithdrawCmd = &cobra.Command{
	Use:   "withdraw",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if withdrawAmount <= 0 {
			log.Printf("escrow.Withdraw - parameter 'amount' is '%f', invalid", withdrawAmount)
			os.Exit(1)
		}

		if err := executor.GetExecutor().EscrowWithdraw(withdrawAmount); nil != err {
			log.Printf("escrow.Withdraw - failed.")
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	WithdrawCmd.Flags().Float64VarP(&withdrawAmount, "amount", "n", 0, "amount to withdraw, unit in ETH")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// withdrawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// withdrawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
