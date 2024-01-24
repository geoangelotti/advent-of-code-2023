package day03

import (
	"fmt"
	"regexp"
	"strings"
)

type Point struct {
	x int
	y int
}

type Symbol struct {
	p Point
	c rune
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

func getSymbols(input string) []Symbol {
	lines := strings.Split(input, "\n")
	symbols := []Symbol{}
	indeces := regexp.MustCompile(`[^.0-9]`)
	for y, line := range lines {
		xs := indeces.FindAllStringIndex(line, -1)
		for _, x := range xs {
			if len(x) > 0 {
				symbols = append(symbols, Symbol{Point{x[0], y}, rune(line[x[0]])})
			}
		}
	}
	return symbols
}

func getEngineParts(input string) []EnginePart {
	engineParts := []EnginePart{}
	for y, line := range strings.Split(input, "\n") {
		span := Span{-1, -1}
		var acc uint64
		for x, char := range line {
			if char >= '0' && char <= '9' {
				if span.start == -1 {
					span.start = x
				}
				acc = acc*10 + uint64(char-'0')
				span.end = x
			} else {
				if span.end != -1 {
					engineParts = append(engineParts, EnginePart{y, span, acc})
					span = Span{-1, -1}
					acc = 0
				}
			}
		}
		if span.end != -1 {
			engineParts = append(engineParts, EnginePart{y, span, acc})
		}
	}
	return engineParts
}

func Process(input string) string {
	var sum uint64
	symbols := getSymbols(input)
	fmt.Println(symbols)
	engineParts := getEngineParts(input)
	fmt.Println(engineParts)
	fmt.Println()
	return fmt.Sprint(sum)
}

func ProcessPart1(input string) string {
	Process(input)
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

type Span struct {
	start int
	end   int
}

type EnginePart struct {
	line  int
	span  Span
	Value uint64
}

func (e EnginePart) Id() string {
	return fmt.Sprintf("%d,%d-%d", e.span.start, e.span.end, e.line)
}

func isNumber(point Point, grid [][]*EnginePart, numbers map[string]EnginePart) {
	defer func() {
		recover()
	}()
	if part := grid[point.y][point.x]; part != nil {
		numbers[part.Id()] = *part
	}
}

func ProcessPart2(input string) string {
	var sum uint64
	lines := strings.Split(input, "\n")
	gears := []Point{}
	gearIndeces := regexp.MustCompile(`\*`)
	for y, line := range lines {
		xs := gearIndeces.FindAllStringIndex(line, -1)
		fmt.Println(xs)
		for _, x := range xs {
			if len(x) > 0 {
				gears = append(gears, Point{x[0], y})
			}
		}

	}
	parts := []EnginePart{}
	grid := [][]*EnginePart{}
	for y, line := range lines {
		row := []*EnginePart{}
		startIndex := -1
		var acc uint64
		for x, char := range line {
			row = append(row, nil)
			if '0' <= char && '9' >= char {
				if startIndex == -1 {
					startIndex = x
				}
				acc = acc*10 + uint64(char-'0')

			} else if acc != 0 {
				span := Span{startIndex, x - 1}
				part := EnginePart{y, span, acc}
				parts = append(parts, part)
				acc = 0
				startIndex = -1
			}
		}
		if startIndex != -1 {
			parts = append(parts, EnginePart{y, Span{startIndex, len(line) - 1}, acc})
		}
		grid = append(grid, row)
	}
	for _, part := range parts {
		p := part
		for i := part.span.start; i <= part.span.end; i++ {
			grid[part.line][i] = &p
		}
	}
	fmt.Println(parts)
	fmt.Println(gears)
	//fmt.Println(grid)
	for _, gear := range gears {
		parts := map[string]EnginePart{}
		points := []Point{
			{gear.x - 1, gear.y - 1},
			{gear.x, gear.y - 1},
			{gear.x + 1, gear.y - 1},
			{gear.x - 1, gear.y},
			{gear.x + 1, gear.y},
			{gear.x - 1, gear.y + 1},
			{gear.x, gear.y + 1},
			{gear.x + 1, gear.y + 1},
		}
		for _, point := range points {
			isNumber(point, grid, parts)
		}
		if len(parts) == 2 {
			var product uint64 = 1
			fmt.Println(gear, parts)
			for _, part := range parts {
				product *= part.Value
			}
			sum += uint64(product)
		}
	}
	return fmt.Sprint(sum)
}
