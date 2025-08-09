package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var Phonebook = make(map[int]Contact)

// Contacts format
type Contact struct {
	Name    string
	Surname string
	Address string
	Phone   int
	Email   string
}

// Print Phonebook
func PrintBook(Phonebook map[int]Contact) string {
	var keys []int
	for id := range Phonebook {
		keys = append(keys, id)
	}
	sort.Ints(keys)

	var builder strings.Builder

	for _, id := range keys {
		contact := Phonebook[id]
		builder.WriteString(fmt.Sprintf("Όνομα/Επώνυμο: %s %s\nΔιεύθηνση: %s\nΤηλέφωνο: %d\nEmail: %s\n", contact.Name, contact.Surname, contact.Address, contact.Phone, contact.Email))
	}
	return builder.String()
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
func AddPhone(Phonebook map[int]Contact) {
	id := len(Phonebook) + 1
	var contact Contact

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Δώσε Όνομα:")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Δώσε Επώνυμο:")
	surname, _ := reader.ReadString('\n')
	contact.Surname = strings.TrimSpace(surname)
	fmt.Print("Δώσε Διεύθηνση:")
	address, _ := reader.ReadString('\n')
	contact.Address = strings.TrimSpace(address)
	fmt.Print("Δώσε Τηλέφωνο:")
	var phone int
	fmt.Scanln(&phone)
	contact.Phone = phone
	fmt.Print("Δώσε Email:")
	Email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(Email)

	Phonebook[id] = contact
}

func LoadPhonebook(path string) (map[int]Contact, error) {
	phonebook := make(map[int]Contact)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 5 {
			fmt.Println("Σφάλμα στη γραμμή:", line)
			continue
		}

		phone, err := strconv.Atoi((strings.TrimSpace(parts[3])))
		if err != nil {
			fmt.Println("Μη έγκυρος αριθμός τηλεφώνου", parts[3])
			continue
		}

		contact := Contact{
			Name:    strings.TrimSpace(parts[0]),
			Surname: strings.TrimSpace(parts[1]),
			Address: strings.TrimSpace(parts[2]),
			Phone:   phone,
			Email:   strings.TrimSpace(parts[4]),
		}

		phonebook[id] = contact
		id++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return phonebook, nil
}

func SavePhonebook(book map[int]Contact) {

}

func main() {

	for {
		var input int
		fmt.Print("Επέλεξε: \nΑνάγνωση καταλόγου (1),\nΠροσθήκη στον κατάλογο (2),\nΑναζήτηση στον κατάλογο (3)\nΕισαγωγή καταλόγου (4)\nΑποθήκευση καταλόγου (5)\nΈξοδος (6)\n")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Μη έγκυρη είσοδος")
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		switch input {
		case 1:
			output := PrintBook(Phonebook)
			fmt.Print(output)
		case 2:
			AddPhone(Phonebook)
			//	case 3:
			//		SearchName(phonebook)
		case 4:
			LoadPhonebook("phonebook.txt")
		case 5:
			data := PrintBook(Phonebook)
			err := os.WriteFile("phonebook.txt", []byte(data), 0644)
			if err != nil {
				fmt.Println("Σφάλμα αποθηκευσης")
			}
		case 6:
			fmt.Println("Έξοδος...")
			return
		default:
			fmt.Println("Λάθος είσοδος")
		}
	}
}
