package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompleteOverlap(t *testing.T) {
	t.Run("TestOverlapTrue", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForCompleteOverlap(6, 6, 4, 6)
		a.True(sut)
	})
	t.Run("TestOverlapFalse", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForCompleteOverlap(2,4,6,8)
		a.False(sut)
	})

}

func TestPartialOverlap(t *testing.T) {
	t.Run("TestOverlapSingleDigit", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForPartialOverlap(5,7,7,9)
		a.True(sut)
	})
	t.Run("TestOverlapFirstPairLower", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForPartialOverlap(5,8,7,9)
		a.True(sut)
	})
	t.Run("TestOverlapSecondPairLower", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForPartialOverlap(7,9,5,8)
		a.True(sut)
	})
	t.Run("TestOverlapFalseFirstPairLower", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForPartialOverlap(2,4,6,8)
		a.False(sut)
	})
	t.Run("TestOverlapFalseSecondPairLower", func(t *testing.T) {
		a := assert.New(t)
		sut := compareForPartialOverlap(6,8,2,4)
		a.False(sut)
	})

}
