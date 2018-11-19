package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 75, solve([]int{10, 20, 15, 25, 10, 20}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 14, solve([]int{1, 10, 4}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 3, solve([]int{1, 2}))
}

func Test4(t *testing.T) {
	assert.Equal(t, 3, solve([]int{2, 2, 1}))
}
