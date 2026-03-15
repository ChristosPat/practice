package main

import "fmt"

func main() {
	colours := []string{"red", "blue", "green"}

	colours = append(colours, "yellow", "orange")

	fmt.Println(colours)
}
