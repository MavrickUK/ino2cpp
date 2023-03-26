package parser

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//goland:noinspection ALL
const (
	cARDUINOH = "#include <Arduino.h>"
	cREGEX    = `(void \w+\(\s*(?:\w+\s+\w+\s*(?:,\s*)?)*\s*\))\s*\{`
)

var lines []string

type Parse struct {
	sourceFilename string
	outputFilename string
	verboseOutput  bool
}

func NewParse(fname string, oname string, verbose bool) *Parse {
	p := &Parse{
		sourceFilename: strings.TrimSuffix(filepath.Base(fname), filepath.Ext(fname)),
		outputFilename: oname,
		verboseOutput:  verbose,
	}
	return p
}

func (p *Parse) Start(appVersion string) {
	fmt.Printf("Ino2Cpp Converter v%s\n", appVersion)
	fmt.Println("Working, please wait...")

	contentChan := make(chan []byte)
	errChan := make(chan error)

	// Read the file asynchronously
	go func() {
		content, err := os.ReadFile(p.sourceFilename + ".ino")
		if err != nil {
			errChan <- fmt.Errorf("error reading file: %w", err)
			return
		}
		contentChan <- content
	}()

	// Wait for the file to be read
	select {
	case content := <-contentChan:
		matchFunctions(content, p.verboseOutput)

		p.createHeader(p.outputFilename)
		p.createCppFile(p.sourceFilename)

		fmt.Printf("%s.cpp and %s.h created. Done!\n", p.outputFilename, p.outputFilename)

	case err := <-errChan:
		fmt.Println(err)
	}
}

func matchFunctions(content []byte, verbose bool) {
	input := string(content)
	pattern := cREGEX
	var funcsFound int
	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) > 1 {
			if !((strings.Contains(match[1], "setup")) ||
				(strings.Contains(match[1], "loop"))) {
				if verbose {
					fmt.Println("  " + match[1] + ";")
				}
				lines = append(lines, match[1]+";")
				funcsFound++
			}
		}
	}
	fmt.Printf("Funcs exported: %d\n", funcsFound)
}

// Creates the .h file containing all our exported functions
func (p *Parse) createHeader(fn string) {
	// Create a file for writing
	f, err := os.Create(fn + ".h")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer f.Close()

	// Use a strings.Builder instead of a bufio.Writer
	var b strings.Builder

	for i := range lines {
		_, err := b.WriteString(lines[i] + "\n")
		if err != nil {
			log.Fatalf("error writing to file: %v", err)
		}
	}

	// Write the contents of the builder to the file
	_, err = io.WriteString(f, b.String())
	if err != nil {
		log.Fatalf("error writing to file: %v", err)
	}
}

// Modify the original file and add two lines to the start
func (p *Parse) createCppFile(fn string) {
	inputFile, err := os.OpenFile(fn+".ino", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(p.outputFilename + ".cpp")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Write the additional content at the beginning of the output file
	header := []byte(cARDUINOH + "\n" + `#include "` + p.outputFilename + ".h" + `"` + "\n\n")
	_, err = outputFile.Write(header)
	if err != nil {
		panic(err)
	}

	// Copy the contents of the input file to the output file
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		panic(err)
	}

	if p.verboseOutput {
		fmt.Printf("Added: %s and #include \"%s.h\" to %s.cpp\n", cARDUINOH, p.outputFilename, p.outputFilename)
	}
}
