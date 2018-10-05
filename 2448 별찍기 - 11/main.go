package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	disp := initDisp(n)
	drawPyramid(disp, 0, 0, n)
	printDisp(disp, n)
}

func initDisp(n int) [][]bool {
	disp := make([][]bool, n)

	for i := 0; i < n; i++ {
		disp[i] = make([]bool, n*2)
	}

	return disp
}

func printDisp(disp [][]bool, n int) {
	w := bufio.NewWriter(os.Stdout)

	for i := 0; i < n; i++ {
		for j := 0; j < n*2; j++ {
			if disp[i][j] {
				fmt.Fprint(w, "*")
			} else {
				fmt.Fprint(w, " ")
			}
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}

// drawTriangle draws the granular triangle.
//
//     *
//    * *
//   *****
//
func drawTriangle(disp [][]bool, i, j int) {
	disp[i+0][j+2] = true
	disp[i+1][j+1] = true
	disp[i+1][j+3] = true
	disp[i+2][j+0] = true
	disp[i+2][j+1] = true
	disp[i+2][j+2] = true
	disp[i+2][j+3] = true
	disp[i+2][j+4] = true
}

// drawPyramid draws the triangles recursively.
func drawPyramid(disp [][]bool, i, j, n int) {
	if n == 3 {
		drawTriangle(disp, i, j)
		return
	}
	drawPyramid(disp, i, j+n/2, n/2)
	drawPyramid(disp, i+n/2, j, n/2)
	drawPyramid(disp, i+n/2, j+n, n/2)
}
