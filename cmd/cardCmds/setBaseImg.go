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

// SetBaseImgCmd represents the SetBaseImg command
var SetBaseImgCmd = &cobra.Command{
	Use:   "setBaseImg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var slotStr string
		if len(argDaySlots) > 1 {
			log.Printf("BidCmd - at most one --dayslot/-d param is allowed, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}
		if len(argDaySlots) == 1 {
			slotStr = argDaySlots[0]
		}
		if argBaseImg == "" {
			log.Printf("BidCmd - need --baseImg/-b arg")
			os.Exit(1)
		}
		executor.GetExecutor().CardSetBaseImg(slotStr, argBaseImg)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// SetBaseImgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// SetBaseImgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
