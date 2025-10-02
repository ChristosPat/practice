package main

import "fmt"

func main() {
	//declare grid and tetro
	var grid [5][5]string
	type Tetro struct {
		Block [4][2]int
		Label string
	}

	//fill the grid with dots
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	//create the "T"
	part := Tetro{
		Block: [4][2]int{
			{0, 0},
			{0, 1},
			{0, 2},
			{1, 1},
		},
		Label: "A",
	}
	//	fmt.Println(grid, "\n", part)
	offsetRow := 1
	offsetCol := 1
	maxRow := 0
	maxCol := 0
	//find max values
	for _, block := range part.Block {
		if block[0] > maxRow {
			maxRow = block[0]
		}
		if block[1] > maxCol {
			maxCol = block[1]
		}
	}

	//find if it fits
	if maxRow+offsetRow >= len(grid) || maxCol+offsetCol >= len(grid[0]) {
		fmt.Println("The tetromino doesn't fit")
	} else {
		for _, block := range part.Block {
			row := block[0] + offsetRow
			col := block[1] + offsetCol
			grid[row][col] = part.Label
		}
	}

	//Print the board
	for row := range grid {
		for col := range grid[row] {
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}

}
