package main

import "fmt"

type momentum struct {
	i, j int
}

func main() {
	var x int
	fmt.Scan(&x)

	var i, j int
	m := momentum{-1, +1}
	for ; x != 1; x-- {
		if m.i == 0 {
			m.i = +1
			m.j = -1
		}
		if m.j == 0 {
			m.i = -1
			m.j = +1
		}

		if m.i == -1 && i == 0 {
			m.i = 0
			m.j = +1
		}
		if m.j == -1 && j == 0 {
			m.i = +1
			m.j = 0
		}

		i += m.i
		j += m.j
	}

	fmt.Printf("%d/%d\n", i+1, j+1)
}
