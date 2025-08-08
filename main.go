package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "temp/manual.md"
	maxHeadingLevel := 3
	// summaries := make([]string, 0)

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
					fmt.Println(string(line))
					// if currentlevel = # assign new child slice to summaries[]
					// save line to summaries[newSlice]
				}
				currentHeadingLevel = currentHeadingLevel + "#"
			}
		}
		// build new .md files in .\temp\summaries\
		// save one .md file per element in summaries
	}
}
