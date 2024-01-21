package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Number struct {
	Value string
}

func (n Number) GetValue() int64 {
	letterNumbers := map[string]int64{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	value, found := letterNumbers[n.Value]
	if found {
		return value
	}
	value, _ = strconv.ParseInt(n.Value, 10, 64)
	return value

}

func NaiveProcessPart2(input string) string {
	lines := strings.Split(input, "\n")
	var sum int64
	numbers := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	for i := 0; i < 10; i++ {
		numbers = append(numbers, fmt.Sprint(i))
	}

	for _, line := range lines {
		var firstValue, lastValue int64
		var firstIndex, lastIndex int
		firstIndex = len(line) + 1
		lastIndex = -1
		for _, number := range numbers {
			for _, i := range regexp.MustCompile(number).FindAllStringIndex(line, -1) {
				i := i[0]
				if i != -1 {
					if i <= firstIndex {
						firstIndex = i
						firstValue = Number{number}.GetValue()
					}
					if i >= lastIndex {
						lastIndex = i
						lastValue = Number{number}.GetValue()
					}
				}
			}
		}
		sum += firstValue*10 + lastValue
	}
	return fmt.Sprint(sum)
}
