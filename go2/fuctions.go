package main

import "fmt"

func main() {
	colours := []string{"red", "blue", "green"}

	colours = append(colours, "yellow", "orange")

	fmt.Print("Colours :")
	for i := range colours {
		fmt.Print(colours[i], " ")
	}
	fmt.Println()
	fmt.Println(len(colours))
}
