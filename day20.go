package main

import (
	"bufio"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

var _ = declareDay(20, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day20Part2(inputReader)
	}
	return day20Part1(inputReader)
})

func day20Part1(inputReader io.Reader) interface{} {
	scanner := bufio.NewScanner(inputReader)
	closestParticleIndex := 0
	closestDistance := math.MaxInt32

	for i := 0; scanner.Scan(); i++ {
		if len(scanner.Bytes()) == 0 {
			i--
			continue
		}

		text := scanner.Text()
		lt := strings.LastIndexByte(text, '<')
		gt := strings.LastIndexByte(text, '>')
		split := strings.SplitN(text[lt+1:gt], ",", 3)
		dist := 0

		for i := 0; i < 3; i++ {
			v, err := strconv.Atoi(split[i])
			if err != nil {
				panic(err)
			}
			if v < 0 {
				v = -v
			}
			dist += v
		}

		if dist < closestDistance {
			closestDistance = dist
			closestParticleIndex = i
		}
	}

	return closestParticleIndex
}

func day20Part2(inputReader io.Reader) interface{} {
	ps := day20Parse(inputReader)
	for t := 0; ; t++ {
		ps.removeCollidingParticles()
		ps.tick()
		if ps.noMoreCollisions() {
			break
		}
	}
	return len(ps)
}

func day20Parse(inputReader io.Reader) day20ParticleSystem {
	var result day20ParticleSystem

	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		text := scanner.Text()

		nextVector := func() day20Vector {
			lt := strings.IndexByte(text, '<')
			gt := strings.IndexByte(text, '>')
			split := strings.SplitN(text[lt+1:gt], ",", 3)
			p, _ := strconv.Atoi(split[0])
			v, _ := strconv.Atoi(split[1])
			a, _ := strconv.Atoi(split[2])

			text = text[gt+1:]

			return day20Vector{int64(p), int64(v), int64(a)}
		}

		result = append(result, day20Particle{nextVector(), nextVector(), nextVector()})
	}

	sort.Sort(result)

	return result
}

type day20Vector struct{ x, y, z int64 }

func (v day20Vector) distToOrigin() int64 {
	x, y, z := v.x, v.y, v.z
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}
	return x + y + z
}

type day20Particle struct{ p, v, a day20Vector }

func (p day20Particle) tick() day20Particle {
	vx, vy, vz := p.v.x+p.a.x, p.v.y+p.a.y, p.v.z+p.a.z
	return day20Particle{
		p: day20Vector{p.p.x + vx, p.p.y + vy, p.p.z + vz},
		v: day20Vector{vx, vy, vz},
		a: p.a,
	}
}

type day20ParticleSystem []day20Particle

func (s day20ParticleSystem) Len() int {
	return len(s)
}

func (s day20ParticleSystem) Less(i, j int) bool {
	di := s[i].p.distToOrigin()
	dj := s[j].p.distToOrigin()
	if di != dj {
		return di < dj
	}
	if s[i].p.x != s[j].p.x {
		return s[i].p.x < s[j].p.x
	}
	if s[i].p.y != s[j].p.y {
		return s[i].p.y < s[j].p.y
	}
	return s[i].p.z < s[j].p.z
}

func (s day20ParticleSystem) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s day20ParticleSystem) tick() {
	for i := range s {
		s[i] = s[i].tick()
	}
	sort.Sort(s)
}

func (s *day20ParticleSystem) removeCollidingParticles() {
	lastPosition := day20Vector{math.MinInt64, math.MinInt64, math.MinInt64}
	for i := 0; i < len(*s); {
		if (*s)[i].p == lastPosition {
			j := i + 1
			for ; j < len(*s) && (*s)[j].p == lastPosition; j++ {
			}
			*s = append((*s)[:i-1], (*s)[j:]...)
			lastPosition = day20Vector{math.MinInt64, math.MinInt64, math.MinInt64}
			i--
		} else {
			lastPosition = (*s)[i].p
			i++
		}
	}
}

func (s day20ParticleSystem) noMoreCollisions() bool {
	lastVelocity, lastAcceleration := int64(math.MinInt64), int64(math.MinInt64)
	for i := 0; i < len(s); i++ {
		a := s[i].a
		v := s[i].v
		p := s[i].p
		if !eqSign(a.x, v.x) || !eqSign(a.y, v.y) || !eqSign(a.z, v.z) {
			return false
		}
		if !eqSign(v.x, p.x) || !eqSign(v.y, p.y) || !eqSign(v.z, p.z) {
			return false
		}
		thisVelocity, thisAcceleration := v.distToOrigin(), a.distToOrigin()
		if thisVelocity < lastVelocity || thisAcceleration < lastAcceleration {
			return false
		}
		lastVelocity, lastAcceleration = thisVelocity, thisAcceleration
	}
	return true
}

func eqSign(a, b int64) bool {
	return a == 0 || b == 0 || (a > 0) == (b > 0)
}
