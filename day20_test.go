package main

import (
	"testing"
)

func TestDay20Part1(t *testing.T) {
	runDayTests(t, 20, []dayTest{
		{
			input: `
p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
`,
			want: 0,
		},
		{
			input: `
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
`,
			want: 1,
		},
	})
}

func TestDay20Part2(t *testing.T) {
	runDayTests(t, 20, []dayTest{
		{
			part2: true,
			input: `
p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>
`,
			want: 1,
		},
	})
}
