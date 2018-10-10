package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(r, &n)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &s[i])
	}

	sort.Sort(boj1181(s))

	for _, c := range unique(s) {
		fmt.Fprintln(w, c)
	}
	w.Flush()
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

func unique(s []string) []string {
	u := make([]string, len(s))

	j := 0
	for _, c := range s {
		if j > 0 && c == u[j-1] {
			continue
		}
		u[j] = c
		j++
	}

	return u[:j]
}
