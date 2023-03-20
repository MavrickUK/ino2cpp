package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ino2cpp/parser"
	"ino2cpp/utils"
	"os"
	"strings"
)

const (
	AppName         = "ino2cpp"
	AppVersion      = "0.1" //TODO: Update BEFORE release/push
	BuildDate       = "20 Mar 2023"
	cFilenameSuffix = ".ino"
	GitHubRepo      = "https://github.com/MavrickUK/ino2cpp"
)

var (
	//inoFilename string
	outFilename string
	verbose     bool
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "ino2cpp <filename> [-o <filename>]",
		Short: "Convert Arduino INO sketches to C++",
		Long: `Arduino sketches and C++ are very similar.
However, an INO file cannot be compiled as-is by C/C++ compilers (e.g. GCC).
This tool converts INO sketches to C++ code such that off-the-shelf compilers and static analysis tools can be executed on the code.
`,
		Example: `
  ino2cpp example.ino
  ino2cpp example.ino -o new_file
  ino2cpp example.ino -v`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inoFile := args[0]
			output, _ := cmd.Flags().GetString("output")
			if output != "" {
				processOutputFile(outFilename) //TODO: Add function to do the processing
			}
			test := utils.RemoveInvalidFilenameChars(inoFile)
			processOutputFile(test) //TODO: Remove the processOutFile function.
			//startParsing(inoFilename)
		},
		Version: AppVersion,
	}
)

func processOutputFile(outputFilename string) {
	fmt.Println("Running output from flag. ", outputFilename)
}

func startParsing(fn string) {
	if !strings.HasSuffix(fn, cFilenameSuffix) {
		fn += cFilenameSuffix
	}
	p := parser.NewParse(fn, outFilename, verbose)
	p.Start(AppVersion)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outFilename, "output", "o", "", "output filename.")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose mode")
}
