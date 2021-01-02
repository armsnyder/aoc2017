package aocutil

import (
	"errors"
	"fmt"
	"os"
)

func GenerateStub(day int) {
	if _, err := os.Lstat(fmt.Sprintf("day%02d.go", day)); !errors.Is(err, os.ErrNotExist) {
		panic(fmt.Errorf("Did you not build the full package? Lstat error: %v", err))
	}

	code, err := os.Create(fmt.Sprintf("day%02d.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	test, err := os.Create(fmt.Sprintf("day%02d_test.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	if _, err := fmt.Fprintf(code, `package main

import (
	"io"
)

var _ = declareDay(%[1]d, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day%02[1]dPart2(inputReader)
	}
	return day%02[1]dPart1(inputReader)
})

func day%02[1]dPart1(inputReader io.Reader) interface{} {
	panic("no solution")
}

func day%02[1]dPart2(inputReader io.Reader) interface{} {
	panic("no solution")
}
`, day); err != nil {
		panic(err)
	}

	if _, err := fmt.Fprintf(test, `package main

import (
	"testing"
)

func TestDay%02[1]dPart1(t *testing.T) {
	runDayTests(t, %[1]d, []dayTest{
		{
			input: %[2]s%[2]s,
			want:  nil,
		},
	})
}

func TestDay%02[1]dPart2(t *testing.T) {
	runDayTests(t, %[1]d, []dayTest{
		//{
		//	part2: true,
		//	input: %[2]s%[2]s,
		//	want:  nil,
		//},
	})
}
`, day, "`"); err != nil {
		panic(err)
	}
}
