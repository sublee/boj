package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 3, solve(0, 3))
}

func Test2(t *testing.T) {
	assert.Equal(t, 3, solve(1, 5))
}

func Test3(t *testing.T) {
	assert.Equal(t, 4, solve(45, 50))
}

func Test4(t *testing.T) {
	assert.Equal(t, 92681, solve(0, 1<<31))
}
