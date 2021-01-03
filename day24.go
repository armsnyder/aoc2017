package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"
)

var _ = declareDay(24, func(part2 bool, inputReader io.Reader) interface{} {
	mu := &sync.Mutex{}
	maxStrength := 0
	maxLength := 0

	day24FindPossibleBridges(day24ParseComponents(inputReader), func(length, strength int) {
		mu.Lock()
		defer mu.Unlock()

		stronger := strength > maxStrength
		longer := length > maxLength
		shorter := length < maxLength

		if !part2 && stronger ||
			part2 && longer ||
			part2 && !shorter && stronger {

			maxLength = length
			maxStrength = strength
		}
	})

	return maxStrength
})

func day24ParseComponents(inputReader io.Reader) (components []day24Component) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		split := strings.SplitN(scanner.Text(), "/", 2)
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])

		components = append(components, day24Component{a, b})
	}

	return components
}

func day24FindPossibleBridges(components []day24Component, fn func(length, strength int)) {
	wg := &sync.WaitGroup{}

	for i, c := range components {
		if c.hasPort(0) {
			wg.Add(1)
			go func(zeroIndex int) {
				bridgeBuilder := day24NewBridgeBuilder(components)
				bridgeBuilder.addComponentIndex(zeroIndex)

				for bridgeBuilder.bridge.len() > 0 {
					for bridgeBuilder.findAndAddNextComponent() {
					}

					fn(bridgeBuilder.bridge.len(), bridgeBuilder.strength)

					bridgeBuilder.removeLastComponent()
				}

				wg.Done()
			}(i)
		}
	}

	wg.Wait()
}

type day24BridgeBuilder struct {
	components  []day24Component
	bridge      *day24IndexStack
	searchStart int
	exposedPort int
	strength    int
}

func day24NewBridgeBuilder(components []day24Component) *day24BridgeBuilder {
	return &day24BridgeBuilder{
		components: components,
		bridge:     day24NewIndexStack(len(components)),
	}
}

func (b *day24BridgeBuilder) findAndAddNextComponent() bool {
	for i := b.searchStart; i < len(b.components); i++ {
		if !b.bridge.contains(i) && b.components[i].hasPort(b.exposedPort) {
			b.addComponentIndex(i)
			return true
		}
	}
	return false
}

func (b *day24BridgeBuilder) addComponentIndex(i int) {
	b.bridge.push(i)
	b.strength += b.components[i].strength()
	b.exposedPort = b.components[i].portOppositeTo(b.exposedPort)
	b.searchStart = 0
}

func (b *day24BridgeBuilder) removeLastComponent() {
	i := b.bridge.pop()
	b.strength -= b.components[i].strength()
	b.exposedPort = b.components[i].portOppositeTo(b.exposedPort)
	b.searchStart = i + 1
}

type day24IndexStack struct {
	indexStack       []int
	containedIndices []bool
}

func day24NewIndexStack(capacity int) *day24IndexStack {
	return &day24IndexStack{
		indexStack:       make([]int, 0, capacity),
		containedIndices: make([]bool, capacity),
	}
}

func (s *day24IndexStack) push(index int) {
	s.indexStack = append(s.indexStack, index)
	s.containedIndices[index] = true
}

func (s *day24IndexStack) pop() int {
	end := len(s.indexStack) - 1
	index := s.indexStack[end]
	s.indexStack = s.indexStack[:end]
	s.containedIndices[index] = false
	return index
}

func (s *day24IndexStack) contains(index int) bool {
	return s.containedIndices[index]
}

func (s *day24IndexStack) len() int {
	return len(s.indexStack)
}

type day24Component [2]int

func (c day24Component) strength() int {
	return c[0] + c[1]
}

func (c day24Component) portOppositeTo(port int) int {
	if c[0] == port {
		return c[1]
	}
	return c[0]
}

func (c day24Component) hasPort(port int) bool {
	return c[0] == port || c[1] == port
}
