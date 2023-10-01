package main

import (
	"fmt"
	"strconv"
)

func main() {
	var row, col int

	fmt.Scan(&row, &col)

	fmt.Print("    |")

	for i := 1; i <= col; i++ {
		intend(i)
		fmt.Print(i)
	}

	fmt.Print("\n   --")

	for i := 1; i <= col; i++ {
		fmt.Print("----")
	}

	fmt.Print("\n")

	for i := 1; i <= row; i++ {
		intend(i)
		fmt.Print(i, "|")
		for j := 1; j <= col; j++ {
			intend(i * j)
			fmt.Print(i * j)
		}
		fmt.Println()
	}
}

func intend(n int) {
	switch len(strconv.Itoa(n)) {
	case 1:
		fmt.Print("   ")
	case 2:
		fmt.Print("  ")
	case 3:
		fmt.Print(" ")
	}
}
