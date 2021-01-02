package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/armsnyder/aoc2017/aocutil"
)

type dayTest struct {
	name  string
	part2 bool
	input string
	want  interface{}
}

func runDayTests(t *testing.T, day int, tests []dayTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := days[day](tt.part2, strings.NewReader(tt.input))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day%02d(%v, ...) = %v, want %v", day, tt.part2, got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	dayNumbers := make([]int, 0, len(days))
	for dayNum := range days {
		dayNumbers = append(dayNumbers, dayNum)
	}
	sort.Ints(dayNumbers)
	for _, dayNum := range dayNumbers {
		b.Run(fmt.Sprintf("Day %02d", dayNum), func(b *testing.B) {
			input := aocutil.GetInput(dayNum)
			inputBytes, err := ioutil.ReadAll(input)
			if err != nil {
				panic(err)
			}
			input.Close()
			inputReader := bytes.NewReader(inputBytes)

			for part := 1; part <= 2; part++ {
				if dayNum == 25 && part == 2 {
					continue
				}

				b.Run(fmt.Sprintf("Part %d", part), func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						inputReader.Reset(inputBytes)
						days[dayNum](part == 2, inputReader)
					}
				})
			}
		})
	}
}
