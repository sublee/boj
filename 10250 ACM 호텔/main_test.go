package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	assert.Equal(t, 402, solve(6, 12, 10))
	assert.Equal(t, 1203, solve(30, 50, 72))
}

func Test101(t *testing.T) {
	assert.Equal(t, 101, solve(1, 1, 1))
	assert.Equal(t, 101, solve(10, 10, 1))
}

func TestMore(t *testing.T) {
	assert.Equal(t, 102, solve(99, 99, 100))
	assert.Equal(t, 9999, solve(99, 99, 9801))
	assert.Equal(t, 304, solve(3, 4, 12))
}
