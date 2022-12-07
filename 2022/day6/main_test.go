package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	t.Run("TestPassNoDuplicates", func(t *testing.T) {
		a := assert.New(t)
		testSlice := []string{"b", "v", "w", "x"}
		got := containsDuplicates(testSlice)
		a.False(got)

	})
	t.Run("TestPassDuplicates", func(t *testing.T) {
		a := assert.New(t)
		testSlice := []string{"b", "v", "w", "w"}
		got := containsDuplicates(testSlice)
		a.True(got)
	})
}