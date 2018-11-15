package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 0, solve(10, []int{1, 2, 3}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 2, solve(10, []int{1, 3, 2}))
}

func Test3(t *testing.T) {
	assert.Equal(t, 14, solve(10, []int{7, 6, 2, 1, 3, 8, 5, 4}))
}
