package main

import "fmt"

func main() {
	phone := map[string]int{
		"Christos": 697712346,
		"Giorgos":  654321987,
		"Kostas":   698754123,
	}

	phone["Antonis"] = 693258741

	for i, j := range phone {
		fmt.Println(i, ":", j)
	}
}
