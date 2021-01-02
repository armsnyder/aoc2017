package main

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

var _ = declareDay(19, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day19Part2(inputReader)
	}
	return day19Part1(inputReader)
})

func day19Part1(inputReader io.Reader) interface{} {
	answer := strings.Builder{}
	day19ParseDiagram(inputReader).walk(func(b byte) {
		if unicode.IsLetter(rune(b)) {
			answer.WriteByte(b)
		}
	})
	return answer.String()
}

func day19Part2(inputReader io.Reader) interface{} {
	answer := 0
	day19ParseDiagram(inputReader).walk(func(b byte) {
		answer++
	})
	return answer
}

func day19ParseDiagram(inputReader io.Reader) (diagram day19Diagram) {
	scanner := bufio.NewScanner(inputReader)
	for scanner.Scan() {
		if len(scanner.Bytes()) > 0 {
			diagram = append(diagram, scanner.Text())
		}
	}
	return diagram
}

type day19Vector struct{ x, y int }

func (v day19Vector) plus(v2 day19Vector) day19Vector {
	return day19Vector{v.x + v2.x, v.y + v2.y}
}

func (v day19Vector) rot90() day19Vector {
	return day19Vector{-v.y, v.x}
}

func (v day19Vector) rot270() day19Vector {
	return day19Vector{v.y, -v.x}
}

type day19Diagram []string

func (d day19Diagram) walk(fn func(byte)) {
	curPos, curDir := func() (curPos, curDir day19Vector) {
		for x, ch := range d[0] {
			if ch == '|' {
				return day19Vector{x, 0}, day19Vector{0, 1}
			}
		}
		panic("no start")
	}()

	charAt := func(pos day19Vector) byte {
		if pos.x < 0 || pos.x >= len(d[0]) || pos.y < 0 || pos.y >= len(d) {
			return ' '
		}
		return d[pos.y][pos.x]
	}

	isOnPath := func(pos day19Vector) bool {
		return charAt(pos) != ' '
	}

	for {
		fn(charAt(curPos))

		switch {
		case isOnPath(curPos.plus(curDir)):
			curPos = curPos.plus(curDir)
		case isOnPath(curPos.plus(curDir.rot90())):
			curDir = curDir.rot90()
			curPos = curPos.plus(curDir)
		case isOnPath(curPos.plus(curDir.rot270())):
			curDir = curDir.rot270()
			curPos = curPos.plus(curDir)
		default:
			return
		}
	}
}
