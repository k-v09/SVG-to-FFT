package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Tuple struct {
	Capital string
	Rest    string
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	text := string(content)
	re := regexp.MustCompile(`([A-Z])([^A-Z]*)`)
	matches := re.FindAllStringSubmatch(text, -1)
	var tuples []Tuple
	for _, match := range matches {
		if len(match) == 3 {
			rest := strings.TrimSpace(match[2])
			tuples = append(tuples, Tuple{Capital: match[1], Rest: rest})
		}
	}
	var output strings.Builder
	output.WriteString(fmt.Sprintf("Number of tuples: %d\n", len(tuples)))
	for _, tuple := range tuples {
		output.WriteString(fmt.Sprintf("('%s', '%s')\n", tuple.Capital, tuple.Rest))
	}
	err = os.WriteFile("output.txt", []byte(output.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	fmt.Println("Tuples have been written to output.txt")
}
