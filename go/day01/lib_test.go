package day01_test

import (
	"day01"
	"testing"
)

const INPUT1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const INPUT2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestProcessPart1(t *testing.T) {
	expected := "142"
	answer := day01.ProcessPart1(INPUT1)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}

func TestProcessPart2(t *testing.T) {
	expected := "281"
	answer := day01.ProcessPart2(INPUT2)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}

func TestProcessPart2Naive(t *testing.T) {
	expected := "281"
	answer := day01.NaiveProcessPart2(INPUT2)
	if answer != expected {
		t.Errorf("Expected %s got %s", expected, answer)
	}
}
