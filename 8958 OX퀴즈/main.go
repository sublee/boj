package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	var ox string
	for i := 0; i < t; i++ {
		fmt.Scanln(&ox)
		fmt.Println(solve(ox))
	}
}

func solve(ox string) int {
	score := 0
	combo := 0

	for _, a := range ox {
		if a == 'O' {
			combo++
			score += combo
		} else {
			combo = 0
		}
	}

	return score
}
