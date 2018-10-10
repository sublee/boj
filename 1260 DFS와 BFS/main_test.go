package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	g := newGraph(4)
	g.link(0, 1)
	g.link(0, 2)
	g.link(0, 3)
	g.link(1, 3)
	g.link(2, 3)
	assert.Equal(t, []int{0, 1, 3, 2}, g.dfs(0))
	assert.Equal(t, []int{0, 1, 2, 3}, g.bfs(0))
}

func Test2(t *testing.T) {
	g := newGraph(4)
	g.link(1, 2)
	g.link(2, 0)
	g.link(0, 1)
	assert.Equal(t, []int{1, 0, 2}, g.dfs(1))
	assert.Equal(t, []int{1, 0, 2}, g.bfs(1))
}
