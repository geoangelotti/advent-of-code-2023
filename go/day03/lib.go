package day03

import (
	"fmt"
	"strings"
)

type Point struct {
	x int
	y int
}

type Cell struct {
	value   int8
	special bool
}

func isEnginePart(p Point, g [][]Cell, result *bool) {
	defer func() {
		recover()
	}()
	if g[p.y][p.x].value == -1 {
		*result = true
	}
}

func ProcessPart1(input string) string {
	grid := [][]Cell{}
	var sum uint64
	for _, line := range strings.Split(input, "\n") {
		chars := []Cell{}
		for _, char := range line {
			if char >= '0' && char <= '9' {
				chars = append(chars, Cell{int8(char - '0'), false})
			} else if char == '.' {
				chars = append(chars, Cell{-2, false})
			} else {
				chars = append(chars, Cell{-1, false})
			}
		}
		grid = append(grid, chars)
	}
	for y, line := range grid {
		for x, cell := range line {
			special := false
			if cell.value >= 0 {
				points := []Point{
					{x - 1, y - 1},
					{x, y - 1},
					{x + 1, y - 1},
					{x - 1, y},
					{x + 1, y},
					{x - 1, y + 1},
					{x, y + 1},
					{x + 1, y + 1}}
				for _, point := range points {
					isEnginePart(point, grid, &special)
				}
				cell.special = special
				grid[y][x] = cell
			}
		}
	}
	for _, line := range grid {
		groups := [][]Cell{}
		group := []Cell{}
		for _, cell := range line {
			if cell.value < 0 {
				groups = append(groups, group)
				group = []Cell{}
			} else {
				group = append(group, cell)
			}
		}
		groups = append(groups, group)
		for _, group := range groups {
			if len(group) > 0 {
				special := false
				var acc uint64
				for _, cell := range group {
					if cell.special {
						special = true
					}
					acc = acc*10 + uint64(cell.value)
				}
				if special {
					sum += acc
				}
			}
		}
	}
	return fmt.Sprint(sum)
}
