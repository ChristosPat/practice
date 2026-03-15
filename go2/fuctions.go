package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	sum := 0

	fmt.Print("Στοιχεία:   ")
	for i := range 5 {
		fmt.Print(numbers[i], " ")
		sum += numbers[i]
	}
	fmt.Println()
	fmt.Println("Σύνολο:  ", sum)
}
