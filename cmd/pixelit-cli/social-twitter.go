package main

// var twitterCmd = &cobra.Command{
// 	Use:   "twitter",
// 	Short: "Send twitter follower count to PixelIt",
// 	Run:   sendTwitter,
// 	PreRun: func(cmd *cobra.Command, args []string) {
// 		// https://github.com/spf13/viper/issues/233
// 		viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
// 	},
// }

// func init() {
// 	socialCmd.AddCommand(twitterCmd)

// 	twitterCmd.PersistentFlags().StringP("username", "u", "", "twitter username (required)")
// 	twitterCmd.MarkPersistentFlagRequired("username")
// }

// func sendTwitter(cmd *cobra.Command, args []string) {
// 	username := viper.GetString("username")
// 	pxc.SendTwitterFollowers(username)
// }
