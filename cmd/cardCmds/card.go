/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cardCmds

import (
	"github.com/spf13/cobra"
)

var argDaySlots []string
var argRenderOpt uint8
var argGreeting string
var argGreetingImg string

// cardCmd represents the card command
var CardCmd = &cobra.Command{
	Use:   "card",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
	CardCmd.PersistentFlags().StringSliceVarP(&argDaySlots, "dayslot", "d", nil, "the day to be displayed, can be unix sec of the day, or string of '2023-08-08'")
	CardCmd.AddCommand(CustomizeCmd)
	CustomizeCmd.Flags().Uint8VarP(&argRenderOpt, "render", "r", 0, "specify renderOpt for your NFT. 0:show image of the day. 1:show common base image. 2:show combo image from chainhis.net. others: same as 0.")
	CustomizeCmd.Flags().StringVarP(&argGreeting, "greet", "g", "", "greeting message of the NFT.")
	CustomizeCmd.Flags().StringVarP(&argGreetingImg, "image", "i", "", "greeting image of the NFT, e.g. ipfs://QmbWu8kkUZ4i5rgCftFJPVGfhrCmbQ7F5LS4VArNyMTb7x/D2.png, or https://ipfs.io/ipfs/QmYCr6dokRW2QQSp6xVRJrMQz2HJuF2ESE5SEozi9BZmzF")

}
