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

	for i := 0; i < 4; i++ {
		grid[image[i][0]][image[i][1]] = "#"
	}
	for row := range grid {
		for col := range grid[row] {
			fmt.Print(grid[row][col])
		}
		fmt.Println()
	}

}
