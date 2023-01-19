package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecideLineType(t *testing.T) {
	t.Run("TestFile", func(t *testing.T) {
		a := assert.New(t)
		sut := decideLineType("14848514 b.txt")
		a.Equal("file", sut)
	})
	t.Run("TestCommand", func(t *testing.T) {
		a := assert.New(t)
		sut := decideLineType("$ cd /")
		a.Equal("command", sut)
	})
	t.Run("TestDirectory", func(t *testing.T) {
		a := assert.New(t)
		sut := decideLineType("dir a")
		a.Equal("directory", sut)
	})

}

