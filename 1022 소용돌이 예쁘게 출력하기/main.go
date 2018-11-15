package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func spiral(r1, c1, r2, c2 int) [][]int {
	radius := max(abs(r1), max(abs(c1), max(abs(r2), abs(c2))))

	// Init a spiral table.
	tab := make([][]int, r2-r1+1)
	for i := 0; i < len(tab); i++ {
		tab[i] = make([]int, c2-c1+1)
	}

	// Anticlockwise
	var (
		right = [2]int{0, 1}
		up    = [2]int{-1, 0}
		left  = [2]int{0, -1}
		down  = [2]int{1, 0}
	)
	dirs := [][2]int{right, up, left, down}

	n := 1
	k := 0
	i, j := 0, 0
	far := 1

	distance := 1 + 2*radius
	for n <= distance*distance {
		d := dirs[k]

		for l := 0; l < far; l++ {
			if i >= r1 && i <= r2 && j >= c1 && j <= c2 {
				tab[i-r1][j-c1] = n
			}
			n++

			i += d[0]
			j += d[1]
		}

		// Bounded.
		k = (k + 1) % len(dirs)

		if k == 0 || k == 2 {
			far++
		}
	}

	return tab
}

func solve(r1, c1, r2, c2 int) {
	// Make a spiral table.
	tab := spiral(r1, c1, r2, c2)

	// Determine the printing format.
	m := 0
	for _, row := range tab {
		for _, col := range row {
			m = max(m, col)
		}
	}

	digits := int(math.Log10(float64(m))) + 1
	format := "%" + strconv.Itoa(digits) + "d"

	// Print the view.
	w := bufio.NewWriter(os.Stdout)
	for _, row := range tab {
		for j, col := range row {
			if j > 0 {
				fmt.Fprint(w, " ")
			}
			fmt.Fprintf(w, format, col)
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}
