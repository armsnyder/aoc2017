package main

import (
	"testing"
)

func TestDay18Part2(t *testing.T) {
	runDayTests(t, 18, []dayTest{
		{
			part2: true,
			input: `
snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`,
			want: 3,
		},
		{
			part2: true,
			input: `
jgz p 6
snd 1
snd 1
snd 1
rcv a
jgz 1 100
snd 1
snd 1
snd 1
`,
			want: 3,
		},
	})
}
