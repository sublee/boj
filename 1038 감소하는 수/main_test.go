package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 42, solve(18))
}

func Test2(t *testing.T) {
	assert.Equal(t, -1, solve(10000))
}

func Test3(t *testing.T) {
	assert.Equal(t, 0, solve(0))
	assert.Equal(t, 1, solve(1))
	assert.Equal(t, 10, solve(10))
}

func Test4(t *testing.T) {
	assert.Equal(t, 9876543210, solve(1022))
	assert.Equal(t, -1, solve(1023))
}

func Test5(t *testing.T) {
	assert.Equal(t, 3210, solve(175))
}
