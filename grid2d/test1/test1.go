package main

import "fmt"

func main() {
	var grid [5][5]string

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			grid[row][col] = "."
		}
	}
	fmt.Println(grid)

	a := "#"
	grid[0][2] = "#"
	grid[1][1], grid[1][3] = a, a
	grid[1][3] = "#"
	grid[2][0] = "#"
	grid[2][1] = "#"
	grid[2][2] = a
	grid[2][3] = a
	grid[2][3] = a
	grid[2][4] = a
	grid[3][0] = a
	grid[3][4] = a
	grid[4][0] = a
	grid[4][4] = a

	for row := range grid {
		for col := range grid[row] {
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}
}
