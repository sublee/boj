package main

import (
	"fmt"
	"sort"
)

func main() {
	var h int
	dwarfs := make(map[int]bool)

	sum := 0
	for i := 0; i < 9; i++ {
		fmt.Scan(&h)
		sum += h
		dwarfs[h] = true
	}
	diff := sum - 100

	heights := make([]int, 7)
	i := 0
	for h := range dwarfs {
		if len(dwarfs) == 9 {
			// Didn't find the fake dwarfs yet. A fake dwarf has a peer. The
			// sum with the peer is same with the diff.
			pair := diff - h

			if h == pair {
				// Every dwarf has different height.
			} else if dwarfs[pair] {
				// Remove the fake dwarfs.
				delete(dwarfs, h)
				delete(dwarfs, pair)
				continue
			}
		}

		heights[i] = h
		i++
	}

	sort.Ints(heights)
	for _, h := range heights {
		fmt.Println(h)
	}
}
