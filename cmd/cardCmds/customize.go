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

// customizeCmd represents the customize command
var CustomizeCmd = &cobra.Command{
	Use:   "customize",
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
		executor.GetExecutor().CardCustomize(argDaySlots[0], argRenderOpt, argGreeting, argGreetingImg)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// customizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// customizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
