package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var r1, c1, r2, c2 int
	fmt.Scan(&r1, &c1, &r2, &c2)
	solve(r1, c1, r2, c2)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func spiral(radius int) [][]int {
	distance := 1 + 2*radius

	// Init a spiral table.
	tab := make([][]int, distance)
	for i := 0; i < distance; i++ {
		tab[i] = make([]int, distance)
	}

	dirs := [][2]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

	n := 1
	k := 0
	i, j := radius, radius
	far, stay := 1, 1

	for i >= 0 && j >= 0 && i < distance && j < distance {
		tab[i][j] = n
		n++

		d := dirs[k]
		i += d[0]
		j += d[1]
		stay--

		// Bounded.
		if stay == 0 {
			k = (k + 1) % len(dirs)

			if k == 0 || k == 2 {
				far++
			}

			stay = far
		}
	}

	return tab
}

func solve(r1, c1, r2, c2 int) {
	radius := max(abs(r1), max(abs(c1), max(abs(r2), abs(c2))))

	// Make a spiral table.
	tab := spiral(radius)

	// Fix indices.
	r1 += radius
	c1 += radius
	r2 += radius
	c2 += radius

	// Slice the table to contain numbers only in r1~r2/c1~c2.
	view := make([][]int, r2-r1+1)
	for i := 0; i < len(view); i++ {
		view[i] = tab[r1+i][c1 : c2+1]
	}

	// Determine the printing format.
	m := 0
	for _, row := range view {
		for _, col := range row {
			m = max(m, col)
		}
	}

	digits := int(math.Log10(float64(m))) + 1
	format := "%" + strconv.Itoa(digits+1) + "d"

	// Print the view.
	for _, row := range view {
		for _, col := range row {
			fmt.Printf(format, col)
		}
		fmt.Println()
	}
}
