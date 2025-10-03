package main

import "fmt"

func main() {

	type Tetros struct {
		Block [4][2]int
		Label string
	}
	Module := []Tetros{
		{
			Block: [4][2]int{{0, 0}, {0, 1}, {0, 2}, {1, 1}},
			Label: "A",
		},
		{
			Block: [4][2]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}},
			Label: "B",
		},
	}

	//grid
	var grid [5][5]string
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			grid[i][j] = "."
		}
	}

	//offsets
	offsets := [][2]int{
		{0, 0},
		{0, 2},
	}

	//fit-fill or discard
	for i := 0; i < len(offsets); i++ {
		maxRow, maxCol := 0, 0
		for _, j := range Module[i].Block {

			if j[0] > maxRow {
				maxRow = j[0]
			}
			if j[1] > maxCol {
				maxCol = j[1]
			}
		}
		if maxRow+offsets[i][0] >= len(grid) || maxCol+offsets[i][1] >= len(grid[0]) {
			fmt.Println("Tetromino ", Module[i].Label, " doesn't fit!!!")
			return
		} else {
			for _, block := range Module[i].Block {
				row := block[0]
				col := block[1]
				if grid[row+offsets[i][0]][col+offsets[i][1]] != "." {
					fmt.Println("ATTENTION!!! Tetromino ", Module[i].Label, " overlaps tetromino ", grid[row+offsets[i][0]][col+offsets[i][1]])
					grid[row+offsets[i][0]][col+offsets[i][1]] = Module[i].Label
				} else {
					grid[row+offsets[i][0]][col+offsets[i][1]] = Module[i].Label
				}
			}
		}

	}
	//print
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}
}
