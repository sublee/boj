package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.Equal(t, 8, solve([]int{4, 2}))
}

func Test2(t *testing.T) {
	assert.Equal(t, 16, solve([]int{8, 4, 2}))
}

func Test4(t *testing.T) {
	assert.Equal(t, 42, solve([]int{6, 7, 2, 21, 3}))
}

func Test5(t *testing.T) {
	assert.Equal(t, 999997000002, solve([]int{999999, 999998}))
}

func Test6(t *testing.T) {
	assert.Equal(t, 4, solve([]int{2}))
}

func Test7(t *testing.T) {
	assert.Equal(t, 16, solve([]int{2, 8, 4}))
}

func Test8(t *testing.T) {
	assert.Equal(t, 12, solve([]int{6, 2, 3, 4}))
}

func Test9(t *testing.T) {
	assert.Equal(t, 140, solve([]int{2, 7, 14, 70, 35, 10, 5, 4}))
}

func Test10(t *testing.T) {
	assert.Equal(t, 36, solve([]int{2, 3, 4, 6, 9, 12, 18}))
}

func Test11(t *testing.T) {
	assert.Equal(t, 246, solve([]int{3, 2, 123, 82, 41}))
}

func Test12(t *testing.T) {
	assert.Equal(t, 712984, solve([]int{89123, 2, 4, 178246, 356492, 8}))
}

func Test13(t *testing.T) {
	assert.Equal(t, 156200, solve([]int{2, 4, 5, 8, 10, 11, 20, 22, 25, 40, 44, 50, 55, 71, 88, 100, 110, 142, 200, 220, 275, 284, 355, 440, 550, 568, 710, 781, 1100, 1420, 1562, 1775, 2200, 2840, 3124, 3550, 3905, 6248, 7100, 7810, 14200, 15620, 19525, 31240, 39050, 78100}))
}

func Test14(t *testing.T) {
	assert.Equal(t, 36481, solve([]int{191}))
}
