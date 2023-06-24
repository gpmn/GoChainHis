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

// dumpCmd represents the dump command
var DumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "the day to be displayed",
	Long: `the day to be displayed, can be unix sec of the day, or string of '2023-08-08,2023-09-09'.
e.g.
    GoChainHis history dump -d=2023-05-01,2023-06-01 -n=5

this command will dump 5 day's details since 2023-05-01 to 2023-05-05, and 2023-06-01 to 2023-06-05
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) == 0 {
			log.Printf("history dump --dayslot/-d params is required, e.g. '-d 2023-08-08'.")
			os.Exit(1)
		}
		for idx := range argDaySlots {
			for offset := 0; offset < argDumpCnt; offset++ {
				executor.GetExecutor().HistoryDump(argDaySlots[idx], offset)
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}
