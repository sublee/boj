package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(r, &n)

	var x int
	var counter [10000]int

	for i := 0; i < n; i++ {
		x = readInt(r)
		counter[x-1]++
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < counter[i]; j++ {
			fmt.Fprintln(w, i+1)
		}
	}
	w.Flush()
}

func readInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	num, _ := strconv.Atoi(string(line))
	return num
}
