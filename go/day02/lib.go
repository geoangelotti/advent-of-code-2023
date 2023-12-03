package day02

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Red   int
	Green int
	Blue  int
}

func getGames(input string) []Game {
	games := []Game{}

	fmt.Println(games)
	return games
}

func (g Game) IsPossible() bool {
	return g.Red <= target.Red && g.Green <= target.Green && g.Blue <= target.Blue
}

var target Game = Game{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func ProcessPart1(input string) string {
	var sum int64
	lines := strings.Split(input, "\n")
	getGames("")
	for _, line := range lines {
		gameLine := strings.Split(line, ":")
		game := strings.Split(gameLine[0], " ")
		gameNumber, _ := strconv.ParseInt(game[1], 10, 64)
		games := strings.Split(gameLine[1], ";")
		fmt.Println(gameNumber, games)

	}
	return fmt.Sprint(sum)
}

func ProcessPart2(input string) string {
	return ""
}
