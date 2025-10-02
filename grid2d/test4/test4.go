package main

import "fmt"

func main() {

	var grid [5][5]string
	type Tetromino struct {
		Block [4][2]int
		Label string
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	part := Tetromino{
		Block: [4][2]int{
			{0, 0},
			{1, 0},
			{1, 1},
			{2, 0},
		},
		Label: "A",
	}

	fmt.Println(grid)
	fmt.Println(part)

	rows := len(grid)
	cols := len(grid[0])

	maxRow, maxCol := 0, 0
	offsetRow := 0
	OffsetCol := 2
	for i := 0; i < 4; i++ {
		tRow := part.Block[i][0]
		tCol := part.Block[i][1]
		if tRow > maxRow {
			maxRow = tRow
		}
		if tCol > maxCol {
			maxCol = tCol
		}
	}

	if maxRow+offsetRow >= rows || maxCol+OffsetCol >= cols {
		fmt.Println("Shape doesn't fit")
	} else {
		for _, block := range part.Block {
			row := block[0] + offsetRow
			col := block[1] + OffsetCol
			grid[row][col] = part.Label
		}

	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}

}
