package main

import (
	"testing"
)

func TestDay22Part1(t *testing.T) {
	runDayTests(t, 22, []dayTest{
		{
			input: `
..#
#..
...
`,
			want: 5587,
		},
	})
}

func TestDay22Part2(t *testing.T) {
	runDayTests(t, 22, []dayTest{
		{
			part2: true,
			input: `
..#
#..
...
`,
			want: 2511944,
		},
	})
}
