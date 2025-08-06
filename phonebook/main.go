package main

import (
	"fmt"
	"sort"
)

// Print Phonebook
func PrintBook(Phonebook map[string]string) {
	var keys []string
	for name := range Phonebook {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	for _, name := range keys {
		fmt.Printf("Όνομα: %s\tΤηλέφωνο: %s\n", name, Phonebook[name])
	}
}

// Shearch Name on Phonebook
func SearchName(Phonebook map[string]string) {
	var input string
	fmt.Print("Δώσε όνομα: ")
	fmt.Scanln(&input)
	if number, exists := Phonebook[input]; exists {
		fmt.Println("O αριθμος ειναι :", number)
	} else {
		fmt.Println("Δεν βρεθηκε αυτό το όνομα στον κατάλογο")
	}
}

// Add a telephone number
func AddPhone(Phonebook map[string]string) {
	var name string
	var number string
	fmt.Print("Δώσε Όνομα:")
	fmt.Scanln(&name)
	fmt.Print("Δώσε Τηλέφωνο:")
	fmt.Scanln(&number)
	Phonebook[name] = number
}

func main() {

	Phonebook := map[string]string{
		"Γιωργος": "123456",
		"Μαρια":   "789456",
		"Αντώνης": "456789",
		"Κώστας":  "741258",
		"Grecory": "951624",
	}
	for {
		var input int
		fmt.Print("Επέλεξε: \nΑνάγνωση καταλόγου (1),\nΠροσθήκη στον κατάλογο (2),\nΑναζήτηση στον κατάλογο (3)\nΈξοδος (4)\n")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Μη έγκυρη είσοδος")
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		switch input {
		case 1:
			PrintBook(Phonebook)
		case 2:
			AddPhone(Phonebook)
		case 3:
			SearchName(Phonebook)
		case 4:
			fmt.Println("Έξοδος...")
			return
		default:
			fmt.Println("Λάθος είσοδος")
		}
	}
}
