package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ino2cpp/parser"
	"os"
)

const (
	//cBUILDDATE = "10 Mar 2023"
	cVERSION = "0.1"
)

// rootCmd represents the base command when called without any subcommands
var (
	inoFilename string

	rootCmd = &cobra.Command{
		Use:   "ino2cpp",
		Short: "Convert Arduino INO sketches to C++",
		Long: `Arduino sketches and C++ are very similar.
However, an INO file cannot be compiled as-is by C/C++ compilers (e.g. GCC).
This tool converts INO sketches to C++ code such that off-the-shelf compilers and static analysis tools can be executed on the code.
`,
		Example: "ino2cpp -i example.ino",

		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Ino2Cpp Converter v%s\n", cVERSION)
			startParsing(inoFilename)
		},
		Version: cVERSION,
	}
)

func startParsing(fn string) {
	p := parser.NewParse(fn)
	fmt.Println("Working, please wait...")
	p.Start()
	//fmt.Println("Done.")
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
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&inoFilename, "input", "i", "", "name of .ino file to convert.")
}
