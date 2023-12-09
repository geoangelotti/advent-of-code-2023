package day03_test

import (
	"day03"
	"testing"
)

const INPUT = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

const INPUT2 = `.2.
.*.
585`

func TestProcessPart1(t *testing.T) {
	expected := "4361"
	answer := day03.ProcessPart1(INPUT)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}

func TestProcessPart1_again(t *testing.T) {
	expected := "587"
	answer := day03.ProcessPart1(INPUT2)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}

func TestProcessPart2(t *testing.T) {
	expected := "467835"
	answer := day03.ProcessPart2(INPUT)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}

func TestProcessPart2_again(t *testing.T) {
	expected := "1170"
	answer := day03.ProcessPart2(INPUT2)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}
