package main

import "github.com/spf13/cobra"

var socialCmd = &cobra.Command{
	Use:   "social",
	Short: "Send social network follower counts to PixelIt",
}

func init() {
	rootCmd.AddCommand(socialCmd)
}
