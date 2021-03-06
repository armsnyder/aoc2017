package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

var _ = declareDay(18, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day18Part2(inputReader)
	}
	return day18Part1(inputReader)
})

func day18Part1(_ io.Reader) interface{} {
	return nil
}

func day18Part2(inputReader io.Reader) interface{} {
	instructions := day18Parse(inputReader)
	p0 := day18Computer{instructions: instructions}
	p1 := day18Computer{instructions: instructions}
	p0.peerQueue, p1.peerQueue = &p1.queue, &p0.queue
	p1.registers[int('p'-'a')] = 1
	for p0.tick() || p1.tick() {
	}
	return p1.sends
}

type day18Computer struct {
	registers    [26]int
	queue        []int
	peerQueue    *[]int
	ip           int
	instructions []day18Instruction
	sends        int
}

func (c *day18Computer) tick() bool {
	return c.ip >= 0 && c.ip < len(c.instructions) && c.instructions[c.ip](c)
}

type day18Instruction func(*day18Computer) bool

func day18Parse(inputReader io.Reader) (instructions []day18Instruction) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		fields := strings.Fields(scanner.Text())
		register := int(fields[1][0] - 'a')
		argValue := func(arg int) func(*day18Computer) int {
			if asValue, err := strconv.Atoi(fields[arg]); err == nil {
				return func(_ *day18Computer) int {
					return asValue
				}
			}
			argRegister := int(fields[arg][0] - 'a')
			return func(c *day18Computer) int {
				return c.registers[argRegister]
			}
		}

		instructions = append(instructions, func() day18Instruction {
			switch fields[0] {
			case "snd":
				arg1 := argValue(1)
				return func(c *day18Computer) bool {
					*c.peerQueue = append(*c.peerQueue, arg1(c))
					c.sends++
					c.ip++
					return true
				}
			case "set":
				arg2 := argValue(2)
				return func(c *day18Computer) bool {
					c.registers[register] = arg2(c)
					c.ip++
					return true
				}
			case "add":
				arg2 := argValue(2)
				return func(c *day18Computer) bool {
					c.registers[register] += arg2(c)
					c.ip++
					return true
				}
			case "mul":
				arg2 := argValue(2)
				return func(c *day18Computer) bool {
					c.registers[register] *= arg2(c)
					c.ip++
					return true
				}
			case "mod":
				arg2 := argValue(2)
				return func(c *day18Computer) bool {
					c.registers[register] %= arg2(c)
					c.ip++
					return true
				}
			case "rcv":
				return func(c *day18Computer) bool {
					if len(c.queue) == 0 {
						return false
					}
					c.registers[register] = c.queue[0]
					c.queue = c.queue[1:]
					c.ip++
					return true
				}
			case "jgz":
				arg1 := argValue(1)
				arg2 := argValue(2)
				return func(c *day18Computer) bool {
					if arg1(c) > 0 {
						c.ip += arg2(c)
					} else {
						c.ip++
					}
					return true
				}
			default:
				panic(fields[0])
			}
		}())
	}

	return instructions
}
