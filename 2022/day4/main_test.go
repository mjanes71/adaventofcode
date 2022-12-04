package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareRanges(t *testing.T) {
	t.Run("TestOverlapTrue", func(t *testing.T) {
		a := assert.New(t)
		sut := compareRanges(6, 6, 4, 6)
		a.True(sut)
	})
	t.Run("TestOverlapFalse", func(t *testing.T) {
		a := assert.New(t)
		sut := compareRanges(2,4,6,8)
		a.False(sut)
	})

}
