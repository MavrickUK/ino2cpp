package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Info about ino2cpp",
	Long:  `Got fed up doing the conversion manually? So did I. :D`,
	Run: func(cmd *cobra.Command, args []string) {
		// Constant values in rootCmd.go
		fmt.Printf("%s v%s - %s\n", AppName, AppVersion, BuildDate)
		fmt.Println("GitHub: " + GitHubRepo)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
