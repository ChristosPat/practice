package main

import "fmt"

func main() {

	var grid [5][5]string
	var image [4][2]int

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			grid[row][col] = "."
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}
	fmt.Println("\n----\n")
	image[0] = [2]int{0, 0}
	image[1] = [2]int{1, 0}
	image[2] = [2]int{2, 0}
	image[3] = [2]int{2, 1}
	rows := len(grid)
	cols := len(grid[0])
	offsetRow := 2
	offsetCol := 4
	maxRow, maxCol := 0, 0
	for i := 0; i < 4; i++ {
		if image[i][0] > maxRow {
			maxRow = image[i][0]
		}
		if image[i][1] > maxCol {
			maxCol = image[i][1]
		}
	}

	if maxRow+offsetRow >= rows || maxCol+offsetCol >= cols {
		fmt.Println("Doesn't fit here")
	} else {
		for i := 0; i < 4; i++ {
			row := image[i][0] + offsetRow
			col := image[i][1] + offsetCol
			grid[row][col] = "#"
		}
		for row := range grid {
			for col := range grid[row] {
				fmt.Print(grid[row][col])
			}
			fmt.Println()
		}
	}

}
