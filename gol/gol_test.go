package gol_test

import (
	"testing"

	"github.com/stanleynguyen/coderetreat/gol"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		g        gol.Grid
		expected bool
	}{
		{gol.Grid{[]uint8{0, 0, 0, 0}, 2}, true},
		{gol.Grid{[]uint8{0, 0, 0, 0}, 3}, false},
	}

	for _, c := range testCases {
		if c.g.IsValid() != c.expected {
			t.Errorf("g.IsValid() = %v, want = %v", c.g.IsValid(), c.expected)
		}
	}
}

func TestString(t *testing.T) {
	testCases := []struct {
		g        gol.Grid
		expected string
	}{
		{gol.Grid{[]uint8{0, 1, 0, 1}, 2}, ".*\n.*\n"},
		{gol.Grid{[]uint8{1, 0, 0, 0}, 2}, "*.\n..\n"},
	}

	for _, c := range testCases {
		if c.g.String() != c.expected {
			t.Errorf("grid.String() = %#v, want = %#v", c.g.String(), c.expected)
		}
	}
}

func TestEmpty(t *testing.T) {
	g := &gol.Grid{
		Arr: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		W:   3,
	}
	nextG := g.Next()
	if g.String() != nextG.String() {
		t.Errorf("g.String() = %#v, want = %#v", nextG.String(), g.String())
	}
}
