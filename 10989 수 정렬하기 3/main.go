package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const max = 10000
	var counter [max]int

	r := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	// read and analyze the input numbers
	n := scanInt(s)
	for i := 0; i < n; i++ {
		counter[scanInt(s)-1]++
	}

	// print the sorted numbers
	w := bufio.NewWriter(os.Stdout)
	for i := 0; i < max; i++ {
		for j := 0; j < counter[i]; j++ {
			fmt.Fprintln(w, i+1)
		}
	}
	w.Flush()
}

func scanInt(s *bufio.Scanner) int {
	s.Scan()
	num, _ := strconv.Atoi(s.Text())
	return num
}
