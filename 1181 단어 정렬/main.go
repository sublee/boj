package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&s[i])
	}

	sort.Sort(boj1181(s))

	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}
}

type boj1181 []string

func (s boj1181) Len() int {
	return len(s)
}

func (s boj1181) Less(i, j int) bool {
	a, b := s[i], s[j]

	if len(a) < len(b) {
		return true
	} else if len(a) > len(b) {
		return false
	}

	return (a < b)
}

func (s boj1181) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
