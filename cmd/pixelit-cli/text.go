package main

import "github.com/spf13/cobra"

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Send text to PixelIt",
	Run:   sendText,
}

func init() {
	rootCmd.AddCommand(textCmd)
}

func sendText(cmd *cobra.Command, args []string) {
	pxc.SendText("Hello hello")
}
