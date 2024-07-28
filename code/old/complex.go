package complexity

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

func formatCTuple(data string, prevPoint string) string {
	pairs := strings.Split(strings.TrimSpace(data), ",")
	if len(pairs) != 3 {
		return data
	}
	formattedPairs := make([]string, 4)
	formattedPairs[0] = prevPoint

	for i, pair := range pairs {
		numbers := strings.Split(strings.TrimSpace(pair), " ")
		if len(numbers) != 2 {
			return data
		}
		formattedPairs[i+1] = fmt.Sprintf("(%s)", strings.Join(numbers, ", "))
	}
	return fmt.Sprintf("(1−t)3%s+3(1−t)2t%s​+3(1−t)t2%s​+t3%s​",
		formattedPairs[0], formattedPairs[1], formattedPairs[2], formattedPairs[3])
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
	var prevPoint string
	for i, match := range matches {
		if len(match) == 3 {
			capital := match[1]
			rest := strings.TrimSpace(match[2])
			if capital == "C" {
				rest = formatCTuple(rest, prevPoint)
			} else if capital == "M" {
				numbers := strings.Fields(rest)
				if len(numbers) >= 2 {
					prevPoint = fmt.Sprintf("(%s, %s)", numbers[0], numbers[1])
				}
			}

			tuples = append(tuples, Tuple{Capital: capital, Rest: rest})
		}
		if i > 0 && prevPoint == "" {
			numbers := strings.Fields(tuples[i-1].Rest)
			if len(numbers) >= 2 {
				prevPoint = fmt.Sprintf("(%s, %s)", numbers[len(numbers)-2], numbers[len(numbers)-1])
			}
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
