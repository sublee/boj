package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var count int
	var ds []int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &count)

	ds = make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Fscan(reader, &ds[i])
	}

	fmt.Println(solve(ds))
}

func solve(ds []int) int {
	sort.Ints(ds)
	min := ds[0]
	max := ds[len(ds)-1]
	return min * max
}
