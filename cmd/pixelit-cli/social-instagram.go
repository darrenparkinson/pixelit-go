package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var instagramCmd = &cobra.Command{
	Use:   "instagram",
	Short: "Send instagram follower count to PixelIt",
	Run:   sendInstagram,
	PreRun: func(cmd *cobra.Command, args []string) {
		// https://github.com/spf13/viper/issues/233
		viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	},
}

func init() {
	socialCmd.AddCommand(instagramCmd)

	instagramCmd.PersistentFlags().StringP("username", "u", "", "instagram username (required)")
	instagramCmd.MarkPersistentFlagRequired("username")
}

func sendInstagram(cmd *cobra.Command, args []string) {
	username := viper.GetString("username")
	err := pxc.SendInstagramFollowers(username)
	if err != nil {
		log.Fatal(err)
	}
}
