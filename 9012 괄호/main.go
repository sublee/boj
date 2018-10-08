package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	var ps string
	for i := 0; i < t; i++ {
		fmt.Scanln(&ps)
		fmt.Println(solve(ps))
	}
}

func solve(ps string) string {
	var stack []bool

	for _, c := range ps {
		switch c {
		case '(':
			stack = append(stack, true)
		case ')':
			if len(stack) == 0 {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	switch len(stack) {
	case 0:
		return "YES"
	default:
		return "NO"
	}
}
