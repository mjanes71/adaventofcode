package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleString(t *testing.T) {
	t.Run("TestSingleDigits", func(t *testing.T) {
		a := assert.New(t)
		move, from, to := handleString("move 7 from 3 to 9")
		expectedMove := 7
		expectedFrom := 3
		expectedTo := 9
		a.Equal(expectedMove, move)
		a.Equal(expectedFrom, from)
		a.Equal(expectedTo, to)
	})
	t.Run("TestDoubleDigits", func(t *testing.T) {
		a := assert.New(t)
		move, from, to := handleString("move 21 from 29 to 91")
		expectedMove := 21
		expectedFrom := 29
		expectedTo := 91
		a.Equal(expectedMove, move)
		a.Equal(expectedFrom, from)
		a.Equal(expectedTo, to)
	})

}

func TestMoveCrate9000(t *testing.T) {
	t.Run("TestPass", func(t *testing.T) {
		a := assert.New(t)
		move := 2
		var stack6 = []string{"C", "B", "P", "T"}
		var stack7 = []string{"B", "J", "R", "P", "L"}
		expected6 := []string{"C", "B"}
		expected7 := []string{"B", "J", "R", "P", "L", "T", "P"}
		//move 2 from 6 to 7
		stack6, stack7 = moveCrates9000(stack6, stack7, move)
		a.ElementsMatch(expected6, stack6)
		a.ElementsMatch(expected7, stack7)

	})
}
func TestMoveCrate9001(t *testing.T) {
	t.Run("TestPass", func(t *testing.T) {
		a := assert.New(t)
		move := 2
		var stack6 = []string{"C", "B", "P", "T"}
		var stack7 = []string{"B", "J", "R", "P", "L"}
		expected6 := []string{"C", "B"}
		expected7 := []string{"B", "J", "R", "P", "L", "P", "T"}
		//move 2 from 6 to 7
		stack6, stack7 = moveCrates9001(stack6, stack7, move)
		a.ElementsMatch(expected6, stack6)
		a.ElementsMatch(expected7, stack7)

	})
}