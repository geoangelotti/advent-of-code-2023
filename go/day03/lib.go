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
	p    Point
	char rune
}

type Span struct {
	start int
	end   int
}

type EnginePart struct {
	line  int
	span  Span
	value uint64
}

func (e EnginePart) Id() string {
	return fmt.Sprintf("%d,%d-%d", e.span.start, e.span.end, e.line)
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

func getEmptyGrid(input string) [][]*EnginePart {
	grid := [][]*EnginePart{}
	for _, line := range strings.Split(input, "\n") {
		row := []*EnginePart{}
		for _, _ = range line {
			row = append(row, nil)
		}
		grid = append(grid, row)
	}
	return grid
}

func ProcessPart1(input string) string {
	var sum uint64
	symbols := getSymbols(input)
	engineParts := getEngineParts(input)
	grid := getEmptyGrid(input)
	for _, part := range engineParts {
		p := part
		for i := part.span.start; i <= part.span.end; i++ {
			grid[part.line][i] = &p
		}
	}
	parts := map[string]EnginePart{}
	for _, symbol := range symbols {
		points := []Point{
			{symbol.p.x - 1, symbol.p.y - 1},
			{symbol.p.x, symbol.p.y - 1},
			{symbol.p.x + 1, symbol.p.y - 1},
			{symbol.p.x - 1, symbol.p.y},
			{symbol.p.x + 1, symbol.p.y},
			{symbol.p.x - 1, symbol.p.y + 1},
			{symbol.p.x, symbol.p.y + 1},
			{symbol.p.x + 1, symbol.p.y + 1},
		}
		for _, point := range points {
			isNumber(point, grid, parts)
		}
	}
	for _, part := range parts {
		sum += part.value
	}
	return fmt.Sprint(sum)
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
	symbols := getSymbols(input)
	gears := []Symbol{}
	for _, symbol := range symbols {
		if symbol.char == '*' {
			gears = append(gears, symbol)
		}
	}
	engineParts := getEngineParts(input)
	grid := getEmptyGrid(input)
	for _, part := range engineParts {
		p := part
		for i := part.span.start; i <= part.span.end; i++ {
			grid[part.line][i] = &p
		}
	}
	for _, gear := range gears {
		parts := map[string]EnginePart{}
		points := []Point{
			{gear.p.x - 1, gear.p.y - 1},
			{gear.p.x, gear.p.y - 1},
			{gear.p.x + 1, gear.p.y - 1},
			{gear.p.x - 1, gear.p.y},
			{gear.p.x + 1, gear.p.y},
			{gear.p.x - 1, gear.p.y + 1},
			{gear.p.x, gear.p.y + 1},
			{gear.p.x + 1, gear.p.y + 1},
		}
		for _, point := range points {
			isNumber(point, grid, parts)
		}
		if len(parts) == 2 {
			var product uint64 = 1
			for _, part := range parts {
				product *= part.value
			}
			sum += uint64(product)
		}
	}
	return fmt.Sprint(sum)
}
