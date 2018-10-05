package main

import "fmt"

func main() {
	var (
		ascending  = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
		descending = [8]int{8, 7, 6, 5, 4, 3, 2, 1}

		input [8]int
	)

	for i := 0; i < 8; i++ {
		fmt.Scan(&input[i])
	}

	switch input {
	case ascending:
		fmt.Println("ascending")
	case descending:
		fmt.Println("descending")
	default:
		fmt.Println("mixed")
	}
}
