package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 33, solve([]int{6, 10, 13, 9, 8, 1}))
}
