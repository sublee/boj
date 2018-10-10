package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n, m, v int
	fmt.Fscan(r, &n, &m, &v)

	g := newGraph(n)

	var a, b int
	for i := 0; i < m; i++ {
		fmt.Fscan(r, &a, &b)
		g.link(a-1, b-1)
	}

	print(w, g.dfs(v-1))
	print(w, g.bfs(v-1))
	w.Flush()
}

type graph struct {
	edges [][]int
}

func newGraph(n int) *graph {
	edges := make([][]int, n)
	return &graph{edges}
}

func (g *graph) link(a, b int) {
	g.edges[a] = append(g.edges[a], b)
	g.edges[b] = append(g.edges[b], a)

	sort.Ints(g.edges[a])
	sort.Ints(g.edges[b])
}

type pair struct {
	a, b int
}

func (g *graph) dfs(v int) []int {
	result := []int{}
	marked := map[int]bool{}
	g._dfs(v, &result, marked)
	return result
}

func (g *graph) _dfs(v int, result *[]int, marked map[int]bool) {
	marked[v] = true
	*result = append(*result, v)

	for _, b := range g.edges[v] {
		if marked[b] {
			continue
		}
		g._dfs(b, result, marked)
	}
}

func (g *graph) bfs(v int) []int {
	marked := make(map[int]bool)

	queue := []int{v}
	marked[v] = true

	result := []int{}

	for len(queue) != 0 {
		a := pop(&queue)
		result = append(result, a)

		for _, b := range g.edges[a] {
			if marked[b] {
				continue
			}
			marked[b] = true
			queue = append(queue, b)
		}
	}

	return result
}

func pop(queue *[]int) int {
	a := (*queue)[0]
	*queue = (*queue)[1:]
	return a
}

func print(w io.Writer, a []int) {
	for i, n := range a {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, n+1)
	}
	fmt.Fprintln(w)
}
