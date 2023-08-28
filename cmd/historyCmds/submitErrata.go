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

var _ArgStoryIdx uint8
var _ArgOldContent string
var _ArgNewContent string
var _ArgReason string

var SubmitErrataCmd = &cobra.Command{
	Use:   "submitErrata",
	Short: "will submit errata to update specified dayslot's story.",
	Long: `will submit errata to update specified dayslot's story.  
e.g.
    GoChainHis history submitErrata -d=2020-06-23 -idx=3 -old="<CN>xxxxxxxxxxxxxxxxxxxx</CN><EN>xxxxxxxxxxxxxxxxxx</EN>"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(argDaySlots) != 1 {
			log.Printf("SubmitErrataCmd - just one --dayslot/-d param is required, e.g. '-d=2023-08-08'.")
			os.Exit(1)
		}

		if _ArgStoryIdx >= 10 {
			log.Printf("SubmitErrataCmd - 'storyIdx' param must within [0,9]")
			os.Exit(1)
		}

		if _ArgOldContent == "" {
			log.Printf("SubmitErrataCmd - 'oldContent' param required")
			os.Exit(1)
		}

		if _ArgNewContent == "" {
			log.Printf("SubmitErrataCmd - 'newContent' param required")
			os.Exit(1)
		}

		if _ArgReason == "" {
			log.Printf("SubmitErrataCmd - 'reason' param required")
			os.Exit(1)
		}

		if checkNews(_ArgNewContent) {
			log.Printf("SubmitErrataCmd - 'newContent' param bad format, invalid")
			os.Exit(1)
		}

		executor.GetExecutor().HistorySubmitErrata(argDaySlots[0], _ArgStoryIdx, _ArgOldContent, _ArgNewContent, _ArgReason)
	},
}

func init() {
}
