package main

import (
	"bufio"
	"io"
	"strings"
)

var _ = declareDay(21, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day21N(inputReader, 18)
	}
	return day21N(inputReader, 5)
})

func day21N(inputReader io.Reader, n int) int {
	ruleBook := day21Parse(inputReader)
	image := day21Image{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	for i := 0; i < n; i++ {
		image.process(ruleBook)
	}
	return image.countOnPixels()
}

func day21Parse(inputReader io.Reader) day21RuleBook {
	ruleBook := day21RuleBook{
		rules2x2: make(map[day212x2]day213x3),
		rules3x3: make(map[day213x3]day214x4),
	}
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		inOutSplit := strings.Split(scanner.Text(), " => ")
		in := inOutSplit[0]
		out := inOutSplit[1]
		inSplit := strings.Split(in, "/")
		outSplit := strings.Split(out, "/")

		parseImage := func(dest []bool, src []string) {
			for i, row := range src {
				for j, ch := range row {
					dest[i*len(row)+j] = ch == '#'
				}
			}
		}

		switch len(inSplit) {
		case 2:
			var ruleIn day212x2
			var ruleOut day213x3
			parseImage(ruleIn[:], inSplit)
			parseImage(ruleOut[:], outSplit)
			ruleIn.transformations(func(ruleIn day212x2) {
				ruleBook.rules2x2[ruleIn] = ruleOut
			})

		case 3:
			var ruleIn day213x3
			var ruleOut day214x4
			parseImage(ruleIn[:], inSplit)
			parseImage(ruleOut[:], outSplit)
			ruleIn.transformations(func(ruleIn day213x3) {
				ruleBook.rules3x3[ruleIn] = ruleOut
			})
		}
	}

	return ruleBook
}

type day21RuleBook struct {
	rules2x2 map[day212x2]day213x3
	rules3x3 map[day213x3]day214x4
}

func (b day21RuleBook) process2x2(images day212x2s) day213x3s {
	result := make(day213x3s, len(images))
	for i, row := range images {
		result[i] = make([]day213x3, len(row))
		for j, image := range row {
			result[i][j] = b.rules2x2[image]
		}
	}
	return result
}

func (b day21RuleBook) process3x3(images day213x3s) day214x4s {
	result := make(day214x4s, len(images))
	for i, row := range images {
		result[i] = make([]day214x4, len(row))
		for j, image := range row {
			result[i][j] = b.rules3x3[image]
		}
	}
	return result
}

type day21Image [][]bool

func (m *day21Image) process(ruleBook day21RuleBook) {
	if len(*m)%2 == 0 {
		*m = ruleBook.process2x2(m.split2x2()).join()
	} else {
		*m = ruleBook.process3x3(m.split3x3()).join()
	}
}

func (m day21Image) split2x2() day212x2s {
	dim := 2
	width := len(m) / dim
	result := make(day212x2s, width)

	for i := 0; i < width; i++ {
		result[i] = make([]day212x2, width)
		for j := 0; j < width; j++ {
			for p := 0; p < dim; p++ {
				for q := 0; q < dim; q++ {
					result[i][j][p*dim+q] = m[i*dim+p][j*dim+q]
				}
			}
		}
	}

	return result
}

func (m day21Image) split3x3() day213x3s {
	dim := 3
	width := len(m) / dim
	result := make(day213x3s, width)

	for i := 0; i < width; i++ {
		result[i] = make([]day213x3, width)
		for j := 0; j < width; j++ {
			for p := 0; p < dim; p++ {
				for q := 0; q < dim; q++ {
					result[i][j][p*dim+q] = m[i*dim+p][j*dim+q]
				}
			}
		}
	}

	return result
}

func (m day21Image) countOnPixels() (count int) {
	for _, row := range m {
		for _, v := range row {
			if v {
				count++
			}
		}
	}

	return count
}

type day212x2 [4]bool

type day212x2s [][]day212x2

func (x day212x2) transformations(fn func(day212x2)) {
	flipX := func() {
		x[1], x[0], x[3], x[2] = x[0], x[1], x[2], x[3]
	}

	flipXY := func() {
		x[0], x[3] = x[3], x[0]
	}

	for i := 0; i < 4; i++ {
		fn(x)
		flipX()
		fn(x)
		if i < 3 {
			flipXY()
		}
	}
}

type day213x3 [9]bool

func (x day213x3) transformations(fn func(day213x3)) {
	flipX := func() {
		x[2], x[0], x[5], x[3], x[8], x[6] = x[0], x[2], x[3], x[5], x[6], x[8]
	}

	flipXY := func() {
		x[8], x[5], x[7], x[1], x[3], x[0] = x[0], x[1], x[3], x[5], x[7], x[8]
	}

	for i := 0; i < 4; i++ {
		fn(x)
		flipX()
		fn(x)
		if i < 3 {
			flipXY()
		}
	}
}

type day213x3s [][]day213x3

func (x day213x3s) join() day21Image {
	dim := 3
	width := len(x) * dim
	result := make(day21Image, width)

	for i, row := range x {
		for p := 0; p < dim; p++ {
			result[i*dim+p] = make([]bool, width)
		}

		for j, image := range row {
			for p := 0; p < dim; p++ {
				for q := 0; q < dim; q++ {
					result[i*dim+p][j*dim+q] = image[p*dim+q]
				}
			}
		}
	}

	return result
}

type day214x4 [16]bool

type day214x4s [][]day214x4

func (x day214x4s) join() day21Image {
	dim := 4
	width := len(x) * dim
	result := make(day21Image, width)

	for i, row := range x {
		for p := 0; p < dim; p++ {
			result[i*dim+p] = make([]bool, width)
		}

		for j, image := range row {
			for p := 0; p < dim; p++ {
				for q := 0; q < dim; q++ {
					result[i*dim+p][j*dim+q] = image[p*dim+q]
				}
			}
		}
	}

	return result
}
