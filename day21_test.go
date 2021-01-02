package main

import (
	"strings"
	"testing"
)

func TestDay21(t *testing.T) {
	input := strings.NewReader(`
../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#
`)
	want := 12
	got := day21N(input, 2)
	if got != want {
		t.Errorf("got %d; wanted %d", got, want)
	}
}
