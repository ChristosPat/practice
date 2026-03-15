package main

import "fmt"

func main() {
	num1 := 15
	num2 := 4
	score := 85
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	if score >= 90 {
		fmt.Println("Excelent")
	} else if score >= 80 {
		fmt.Println("Very good")
	} else if score >= 70 {
		fmt.Println("Good")
	} else {
		fmt.Println("Fail")
	}
}
