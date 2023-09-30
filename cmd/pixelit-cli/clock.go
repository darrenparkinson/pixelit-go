package main

import (
	"github.com/darrenparkinson/pixelit-go"
	"github.com/spf13/cobra"
)

var clockCmd = &cobra.Command{
	Use:   "clock",
	Short: "Send a clock to PixelIt",
	Run:   sendClock,
}

func init() {
	rootCmd.AddCommand(clockCmd)
}

func sendClock(cmd *cobra.Command, args []string) {
	s := pixelit.Screen{
		Clock: &pixelit.Clock{
			Show:         pixelit.Bool(true),
			SwitchAktiv:  pixelit.Bool(true),
			WithSeconds:  pixelit.Bool(false),
			SwitchSec:    pixelit.Int(7),
			DrawWeekDays: pixelit.Bool(true),
			Color: &pixelit.Color{
				R: 255,
				G: 255,
				B: 255,
			},
		},
	}
	pxc.SendScreen(&s)
}
