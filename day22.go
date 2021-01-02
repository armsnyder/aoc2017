package main

import (
	"bytes"
	"io"
	"io/ioutil"
)

var _ = declareDay(22, func(part2 bool, inputReader io.Reader) interface{} {
	n := 10000
	if part2 {
		n = 10000000
	}

	virus := day22Virus{cluster: day22Parse(inputReader)}

	for i := 0; i < n; i++ {
		virus.burst(part2)
	}

	return virus.infections
})

func day22Parse(inputReader io.Reader) day22Cluster {
	input, _ := ioutil.ReadAll(inputReader)
	rows := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	cluster := make(day22Cluster)

	for i, row := range rows {
		y := i - len(rows)/2
		for j, ch := range row {
			if ch == '#' {
				x := j - len(row)/2
				cluster[day22Coord{x, y}] = day22Infected
			}
		}
	}

	return cluster
}

type day22Cluster map[day22Coord]day22NodeState

type day22Coord struct{ x, y int }

type day22Virus struct {
	cluster    day22Cluster
	infections int
	facing     int
	pos        day22Coord
}

func (v *day22Virus) burst(evolved bool) {
	switch v.cluster[v.pos] {
	case day22Clean:
		v.facing += 3
		v.facing %= 4
		if evolved {
			v.cluster[v.pos] = day22Weakened
		} else {
			v.cluster[v.pos] = day22Infected
			v.infections++
		}

	case day22Weakened:
		v.cluster[v.pos] = day22Infected
		v.infections++

	case day22Infected:
		v.facing += 1
		v.facing %= 4
		if evolved {
			v.cluster[v.pos] = day22Flagged
		} else {
			delete(v.cluster, v.pos)
		}

	case day22Flagged:
		v.facing += 2
		v.facing %= 4
		delete(v.cluster, v.pos)
	}

	switch v.facing {
	case 0:
		v.pos.y--
	case 1:
		v.pos.x++
	case 2:
		v.pos.y++
	case 3:
		v.pos.x--
	}
}

type day22NodeState int

const (
	day22Clean day22NodeState = iota
	day22Weakened
	day22Infected
	day22Flagged
)
