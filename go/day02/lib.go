package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	var sum uint64
	lines := strings.Split(input, "\n")
	var red uint64 = 12
	var green uint64 = 13
	var blue uint64 = 14
	for _, line := range lines {
		gameLine := strings.Split(line, ":")
		gameIdentifier := strings.Split(gameLine[0], " ")
		game, _ := strconv.ParseUint(gameIdentifier[1], 10, 64)
		plays := strings.Split(gameLine[1], ";")
		possible := true
		for _, play := range plays {
			cubes := strings.Split(play, ",")
			appearances := map[string]uint64{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, cube := range cubes {
				atoms := strings.Split(cube, " ")
				number, _ := strconv.ParseUint(atoms[1], 10, 64)
				colour := atoms[2]
				appearances[colour] = number
			}
			if appearances["red"] > red || appearances["green"] > green || appearances["blue"] > blue {
				possible = false
			}
		}
		if possible {
			sum += game
		}
	}
	return fmt.Sprint(sum)
}

func ProcessPart2(input string) string {
	var sum uint64
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var red uint64 = 0
		var green uint64 = 0
		var blue uint64 = 0
		gameLine := strings.Split(line, ":")
		plays := strings.Split(gameLine[1], ";")
		possible := true
		for _, play := range plays {
			cubes := strings.Split(play, ",")
			appearances := map[string]uint64{
				"red":   0,
				"green": 0,
				"blue":  0,
			}
			for _, cube := range cubes {
				atoms := strings.Split(cube, " ")
				number, _ := strconv.ParseUint(atoms[1], 10, 64)
				colour := atoms[2]
				appearances[colour] = number
			}
			if appearances["red"] > red {
				red = appearances["red"]
			}
			if appearances["green"] > green {
				green = appearances["green"]
			}
			if appearances["blue"] > blue {
				blue = appearances["blue"]
			}
		}
		if possible {
			sum += red * green * blue
		}
	}
	return fmt.Sprint(sum)
}
