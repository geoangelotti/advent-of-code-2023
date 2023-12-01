package day01

import (
	"fmt"
	"strings"
)

func ProcessPart1(input string) string {
	lines := strings.Split(input, "\n")
	var sum int64
	for _, line := range lines {
		values := []int64{}
		for _, char := range line {
			if 57 >= char && 48 <= char {
				values = append(values, int64(char)-48)
			}
		}
		if len(values) > 0 {
			sum += values[0]*10 + values[len(values)-1]*1
		}
	}
	return fmt.Sprint(sum)
}

func ProcessPart2(input string) string {
	letterNumbers := map[string]string{
		"zero":  "z0ero",
		"one":   "o1ne",
		"two":   "t2wo",
		"three": "t3hree",
		"four":  "f4our",
		"five":  "f5ive",
		"six":   "s6ix",
		"seven": "s7even",
		"eight": "e8ight",
		"nine":  "n9ine",
	}
	lines := strings.Split(input, "\n")
	var sb strings.Builder
	for _, line := range lines {
		for k, v := range letterNumbers {
			line = strings.Replace(line, k, v, -1)
		}
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return ProcessPart1(sb.String())
}
