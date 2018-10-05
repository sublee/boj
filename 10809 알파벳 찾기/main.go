package main

import "fmt"

func main() {
	var word string
	var poss [26]int

	fmt.Scan(&word)

	for i, c := range word {
		if poss[c-'a'] == 0 {
			poss[c-'a'] = i + 1
		}
	}

	for _, p := range poss {
		fmt.Printf("%d ", p-1)
	}
}
