package main

import (
	"testing"
)

func TestDay24Part1(t *testing.T) {
	runDayTests(t, 24, []dayTest{
		{
			input: `
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`,
			want: 31,
		},
	})
}

func TestDay24Part2(t *testing.T) {
	runDayTests(t, 24, []dayTest{
		{
			part2: true,
			input: `
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`,
			want: 19,
		},
	})
}
