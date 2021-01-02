package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

var _ = declareDay(23, func(part2 bool, inputReader io.Reader) interface{} {
	if part2 {
		return day23Part2(inputReader)
	}
	return day23Part1(inputReader)
})

func day23Part1(inputReader io.Reader) interface{} {
	cpu := day23CPU{instructions: day23ParseInstructions(inputReader)}
	cpu.run()
	return cpu.mulInvokeCount
}

func day23Part2(inputReader io.Reader) interface{} {
	cpu := day23CPU{instructions: day23ParseInstructions(inputReader)}
	cpu.write('a', 1)
	for cpu.read('f') == 0 {
		cpu.step()
	}

	min := cpu.read('b')
	max := cpu.read('c')
	nonPrimeCount := 0

	for n := min; n <= max; n += 17 {
		if !day23Prime(n) {
			nonPrimeCount++
		}
	}

	return nonPrimeCount
}

func day23Prime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func day23ParseInstructions(inputReader io.Reader) (instructions []day23Instruction) {
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		fields := strings.Fields(scanner.Text())

		argValue := func(arg int) func(*day23CPU) int {
			if asValue, err := strconv.Atoi(fields[arg]); err == nil {
				return func(_ *day23CPU) int {
					return asValue
				}
			}
			return func(c *day23CPU) int {
				return c.read(fields[arg][0])
			}
		}

		noJump := func(instruction day23Instruction) day23Instruction {
			return func(cpu *day23CPU) {
				instruction(cpu)
				cpu.ip++
			}
		}

		instructions = append(instructions, func() day23Instruction {
			switch fields[0] {
			case "set":
				register := fields[1][0]
				arg := argValue(2)

				return noJump(func(cpu *day23CPU) {
					cpu.write(register, arg(cpu))
				})

			case "sub":
				register := fields[1][0]
				arg := argValue(2)

				return noJump(func(cpu *day23CPU) {
					cpu.write(register, cpu.read(register)-arg(cpu))
				})

			case "mul":
				register := fields[1][0]
				arg := argValue(2)

				return noJump(func(cpu *day23CPU) {
					cpu.write(register, cpu.read(register)*arg(cpu))
					cpu.mulInvokeCount++
				})

			case "jnz":
				arg1 := argValue(1)
				arg2 := argValue(2)

				return func(cpu *day23CPU) {
					if arg1(cpu) == 0 {
						cpu.ip++
					} else {
						cpu.ip += arg2(cpu)
					}
				}

			default:
				panic(fields[0])
			}
		}())
	}

	return instructions
}

type day23CPU struct {
	instructions   []day23Instruction
	registers      [8]int
	ip             int
	mulInvokeCount int
}

func (c *day23CPU) run() {
	for c.ip >= 0 && c.ip < len(c.instructions) {
		c.step()
	}
}

func (c *day23CPU) step() {
	c.instructions[c.ip](c)
}

func (c *day23CPU) write(letter byte, value int) {
	c.registers[letter-'a'] = value
}

func (c *day23CPU) read(letter byte) int {
	return c.registers[letter-'a']
}

type day23Instruction func(cpu *day23CPU)
