package main

import "fmt"

func main() {
	var grid [5][5]rune
	var tetro [4][2]int

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			grid[row][col] = '.'
		}
	}

	tetro[0] = [2]int{0, 0}
	tetro[1] = [2]int{0, 1}
	tetro[2] = [2]int{1, 0}
	tetro[3] = [2]int{1, 1}

	for _, block := range tetro {
		grid[block[0]][block[1]] = '#'
	}

	for row := 0; row < len(grid[0]); row++ {
		for col := 0; col < len(grid[row]); col++ {
			fmt.Print(string(grid[row][col]))
		}
		fmt.Println()
	}
}
