package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tiktokCmd = &cobra.Command{
	Use:   "tiktok",
	Short: "Send tiktok follower count to PixelIt",
	Run:   sendTiktok,
	PreRun: func(cmd *cobra.Command, args []string) {
		// https://github.com/spf13/viper/issues/233
		viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	},
}

func init() {
	socialCmd.AddCommand(tiktokCmd)

	tiktokCmd.PersistentFlags().StringP("username", "u", "", "titktok username (required)")
	tiktokCmd.MarkPersistentFlagRequired("username")
	// viper.BindPFlag("username", tiktokCmd.PersistentFlags().Lookup("username"))
}

func sendTiktok(cmd *cobra.Command, args []string) {
	username := viper.GetString("username")
	pxc.SendTiktokFollowers(username)
}
