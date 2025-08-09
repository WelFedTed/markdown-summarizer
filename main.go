package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: finish struct
// Summary SYSTEMSERVICE contains WINDOWS as child that contains SOFTWARE and HARDWARE as children
// loop through Summary recursively and print to markdown file
type Summary struct {
	Title    string
	Children []Summary
}

func main() {
	input := "temp/manual.md"
	maxHeadingLevel := 6
	summaries := make([]string, 0)

	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		words := strings.Fields(string(line))
		for _, word := range words {
			currentHeadingLevel := "#"
			for range maxHeadingLevel {
				if word == currentHeadingLevel {
					// TODO: if currentlevel = # assign new child slice to summaries[]
					if word == "#" {
						summaries = append(summaries, string(line))
					}
					// TODO: save line to summaries[newSlice]
				}
				currentHeadingLevel = currentHeadingLevel + "#"
			}
		}
	}
	for _, s := range summaries {
		fmt.Println(s)
	}
	fmt.Printf("\nFound %d headings in %s\n\n", len(summaries), input)
	// TODO: build new .md files in ./summaries
	// TODO: save one .md file per element in summaries
	// TODO: filename convention: tags "# System Service > ## Windows" become "summaries/system-service_windows.md"
}
