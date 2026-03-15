package main

import "fmt"

// backtrack φτιάχνει όλες τις διατάξεις
func backtrack(nums []int, path []int, used []bool) {
	// Αν path έχει το ίδιο μήκος με nums → βρήκαμε λύση
	if len(path) == len(nums) {
		fmt.Println(path)
		return
	}

	// Δοκιμάζουμε κάθε αριθμό
	for i := 0; i < len(nums); i++ {
		if used[i] { // αν τον έχω ήδη βάλει, τον παραλείπω
			continue
		}

		// κάνω επιλογή
		path = append(path, nums[i])
		used[i] = true

		// αναδρομή
		backtrack(nums, path, used)

		// κάνω undo (backtrack)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, used)
}
