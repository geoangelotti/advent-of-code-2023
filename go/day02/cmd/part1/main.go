package main

import (
	"day02"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(day02.ProcessPart1(string(content)))
}
