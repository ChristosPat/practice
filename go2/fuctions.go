package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func printMessage(msg string) {
	fmt.Print(msg)
}

func main() {
	printMessage("Μαθαίνω Go!\n")
	printMessage("Το γινόμενο του 6*7=")
	fmt.Println(multiply(6, 7))
}
