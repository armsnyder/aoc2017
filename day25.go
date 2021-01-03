package main

import (
	"bufio"
	"io"
	"strconv"
)

var _ = declareDay(25, func(_ bool, inputReader io.Reader) interface{} {
	program, state, totalSteps := day25Parse(inputReader)
	tape := day25NewTape()

	for i := 0; i < totalSteps; i++ {
		state = program[state](tape)
	}

	return tape.checksum()
})

func day25Parse(inputReader io.Reader) (program map[byte]day25Operation, state byte, totalSteps int) {
	program = make(map[byte]day25Operation)
	scanner := bufio.NewScanner(inputReader)

	for len(scanner.Bytes()) == 0 {
		scanner.Scan()
	}
	state = scanner.Bytes()[len(scanner.Bytes())-2]

	scanner.Scan()
	totalSteps, _ = strconv.Atoi(scanner.Text()[36 : len(scanner.Bytes())-7])

	scanner.Scan()

	parseClause := func() (write bool, move func(tape *day25Tape), continueState byte) {
		scanner.Scan()
		scanner.Scan()
		write = scanner.Bytes()[len(scanner.Bytes())-2] == '1'

		scanner.Scan()
		if scanner.Bytes()[len(scanner.Bytes())-6] == 'r' {
			move = func(tape *day25Tape) { tape.right() }
		} else {
			move = func(tape *day25Tape) { tape.left() }
		}

		scanner.Scan()
		continueState = scanner.Bytes()[len(scanner.Bytes())-2]

		return write, move, continueState
	}

	for scanner.Scan() {
		inState := scanner.Bytes()[len(scanner.Bytes())-2]

		ifFalseWrite, ifFalseMove, ifFalseContinueState := parseClause()
		ifTrueWrite, ifTrueMove, ifTrueContinueState := parseClause()

		program[inState] = func(tape *day25Tape) byte {
			if tape.read() {
				tape.write(ifTrueWrite)
				ifTrueMove(tape)
				return ifTrueContinueState
			}

			tape.write(ifFalseWrite)
			ifFalseMove(tape)
			return ifFalseContinueState
		}

		scanner.Scan()
	}

	return program, state, totalSteps
}

type day25Tape struct {
	cursor *day25Value
}

func day25NewTape() *day25Tape {
	return &day25Tape{cursor: &day25Value{}}
}

func (t *day25Tape) write(value bool) {
	t.cursor.value = value
}

func (t *day25Tape) read() bool {
	return t.cursor.value
}

func (t *day25Tape) left() {
	if t.cursor.left == nil {
		t.cursor.left = &day25Value{}
		t.cursor.left.right = t.cursor
	}

	t.cursor = t.cursor.left
}

func (t *day25Tape) right() {
	if t.cursor.right == nil {
		t.cursor.right = &day25Value{}
		t.cursor.right.left = t.cursor
	}

	t.cursor = t.cursor.right
}

func (t *day25Tape) checksum() (checksum int) {
	for cur := t.cursor; cur != nil; cur = cur.right {
		if cur.value {
			checksum++
		}
	}

	for cur := t.cursor.left; cur != nil; cur = cur.left {
		if cur.value {
			checksum++
		}
	}

	return checksum
}

type day25Value struct {
	value bool
	left  *day25Value
	right *day25Value
}

type day25Operation func(tape *day25Tape) (nextState byte)
