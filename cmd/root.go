/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"GoChainHis/executor"

	"GoChainHis/cmd/cardCmds"
	"GoChainHis/cmd/historyCmds"

	"github.com/spf13/cobra"
)

var FlagConfPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GoChainHis",
	Short: "CLI Client for ChainHis Dapp",
	Long: `CLI Client to access ChainHis dapp. You can :
A). interact with escrow contract, such as :
    1). check escrow details with : 
	    GoChainHis escrow dump
    2). deposit in escrow contract with (unit in ETH) :
	    GoChainHis escrow deposit --amount=x.xxx
    3). withdraw from escrow contract with (unit in ETH) :
	    GoChainHis escrow withdraw --amount=x.xxx
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		confPath := FlagConfPath
		if strings.HasPrefix(confPath, "~/") {
			home, _ := os.UserHomeDir()
			confPath = filepath.Join(home, confPath[2:])
		}
		exe := executor.GetExecutor()
		err := exe.LoadConfigure(confPath)
		if nil != err {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(historyCmds.HistoryCmd)
	rootCmd.AddCommand(cardCmds.CardCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.GoChainHis.yaml)")
	rootCmd.PersistentFlags().StringVar(&FlagConfPath, "conf", "~/.GoChainHis/conf.json", "config file (default is $HOME/.GoChainHis/conf.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
