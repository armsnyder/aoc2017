package main

import (
	"testing"
)

func TestDay19Part1(t *testing.T) {
	runDayTests(t, 19, []dayTest{
		{
			input: `
     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`,
			want: "ABCDEF",
		},
	})
}

func TestDay19Part2(t *testing.T) {
	runDayTests(t, 19, []dayTest{
		{
			part2: true,
			input: `
     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ 
`,
			want: 38,
		},
	})
}
