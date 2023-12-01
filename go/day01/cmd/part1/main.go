package main

import (
	"day01"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(day01.ProcessPart1(string(content)))
}
