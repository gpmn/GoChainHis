/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package escrowSubs

import (
	"GoChainHis/executor"

	"github.com/spf13/cobra"
)

var FlagDumpAddr string

// dumpCmd represents the dump command
var DumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		executor.GetExecutor().EscrowDump(FlagDumpAddr)
	},
}

func init() {
	DumpCmd.Flags().StringVarP(&FlagDumpAddr, "addr", "a", "", "target address")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
