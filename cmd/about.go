/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Extra info about ino2cpp",
	Long: `Arduino sketches and C++ are very similar.
However, an INO file cannot be compiled as-is by C/C++ compilers (e.g. GCC).
This tool converts INO sketches to C++ code such that off-the-shelf compilers and static analysis tools can be executed on the code.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s - %s\n", AppName, AppVersion, BuildDate)
		fmt.Println("GitHub: " + GitHubRepo)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
