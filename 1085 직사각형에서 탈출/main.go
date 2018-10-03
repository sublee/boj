package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, w, h int
	fmt.Scan(&x, &y, &w, &h)

	// Find the minimum.
	a := []int{x, y, w - x, h - y}
	sort.Ints(a)
	min := a[0]

	fmt.Println(min)
}
