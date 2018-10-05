package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var r int
	var word string
	for i := 0; i < n; i++ {
		fmt.Scan(&word)
		if isGroupWord(word) {
			r++
		}
	}

	fmt.Println(r)
}

func isGroupWord(word string) bool {
	var poss [26]int

	for i, c := range word {
		a := c - 'a'

		if poss[a] != 0 && poss[a] != i {
			return false
		}

		poss[a] = i + 1
	}

	return true
}
