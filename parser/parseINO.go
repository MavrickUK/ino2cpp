package parser

import (
	"bufio"
	"fmt"
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
	sourceName string
}

func NewParse(fname string) *Parse {
	p := &Parse{
		sourceName: strings.TrimSuffix(filepath.Base(fname), filepath.Ext(fname)),
	}
	return p
}

func printHeader(appVersion string) {
	fmt.Printf("Ino2Cpp Converter v%s\n", appVersion)
	fmt.Println("Working, please wait...")
}

func (p *Parse) Start(appVersion string) {
	printHeader(appVersion)
	content, err := os.ReadFile(p.sourceName + ".ino")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input := string(content)
	pattern := cREGEX

	r := regexp.MustCompile(pattern)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) > 1 {
			if !((strings.Contains(match[1], "setup")) ||
				(strings.Contains(match[1], "loop"))) {
				//fmt.Println(match[1] + ";")
				lines = append(lines, match[1]+";")
			}
		}
	}

	p.createHeader(p.sourceName)
	p.modifySourceFile(p.sourceName)
	fmt.Printf("Done!\n%s and %s created.", p.sourceName+".cpp", p.sourceName+".h")
}

// Creates the .h file containg all our exported functions
func (p *Parse) createHeader(fn string) {
	// Create a file for writing
	f, _ := os.Create(fn + ".h")

	// Create a writer
	w := bufio.NewWriter(f)
	//w.WriteString("// " + fn + " - HEADER FILE\n\n")

	for i := range lines {
		_, err := w.WriteString(lines[i] + "\n")
		if err != nil {
			return
		}
	}
	// Very important to invoke after writing a large number of lines
	err := w.Flush()
	if err != nil {
		return
	}
}

// Modify the original file and add two lines to the start
func (p *Parse) modifySourceFile(fn string) {
	// Open the file in read-write mode
	file, err := os.OpenFile(fn+".ino", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	outputFile, err := os.Create(p.sourceName + ".cpp")
	//outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	// Read the current contents of the file
	scanner := bufio.NewScanner(file)
	var contents string
	for scanner.Scan() {
		contents += scanner.Text() + "\n"
	}
	// Prepend the two lines of text
	newContents := cARDUINOH + "\n" + `#include "` + p.sourceName + ".h" + `"` + "\n\n" + contents

	// Write the updated contents of the file to the beginning
	_, err = outputFile.WriteAt([]byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
