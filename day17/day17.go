package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseLines(splittedByNewLines []string) ([]int, []int, []int) {
	registersString := splittedByNewLines[0]
	programString := splittedByNewLines[1]

	lines := strings.Split(registersString, "\n")

	var registers []int
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		numbers := re.FindAllString(line, -1)
		num, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println("Error:", err)
		}
		registers = append(registers, num)
	}

	programNumbers := re.FindAllString(programString, -1)

	var opCodes []int
	var operands []int

	for i := 0; i < len(programNumbers)-1; i = i + 2 {
		opCode, err := strconv.Atoi(programNumbers[i])
		if err != nil {
			fmt.Println("Error:", err)
		}
		opCodes = append(opCodes, opCode)

		operand, err := strconv.Atoi(programNumbers[i+1])
		if err != nil {
			fmt.Println("Error:", err)
		}
		operands = append(operands, operand)
	}

	return registers, opCodes, operands
}

func literalOperand(operand int) int {
	return operand
}

func comboOperand(operand int, registers []int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	case 7:
		panic("Invalid program: case 7 in combo operand")
	}
	panic("Invalid combo operand")
}

func calculatePart1(registers []int, opcodes []int, operands []int) string {
	output := ""

	if len(opcodes) != len(operands) {
		panic("opcodes len is not equal to operands len")
	}

	a := 0
	b := 1
	c := 2

	instruction_pointer := 0
	for instruction_pointer < len(opcodes) {
		if instruction_pointer >= len(opcodes) {
			break
		}
		opcode := opcodes[instruction_pointer]
		operand := operands[instruction_pointer]
		fmt.Println(instruction_pointer, opcode, operand)

		switch opcode {
		case 0:
			numerator := registers[a]
			denominator := int(1 << comboOperand(operand, registers))
			registers[a] = numerator / denominator
		case 1:
			registers[b] = registers[b] ^ literalOperand(operand)

		case 2:
			registers[b] = comboOperand(operand, registers) % 8
		case 3:
			if registers[a] != 0 {
				instruction_pointer = literalOperand(operand) / 2
				continue
			}

		case 4:
			registers[b] = registers[b] ^ registers[c]

		case 5:
			number := comboOperand(operand, registers) % 8
			output = output + strconv.Itoa(number) + ","
		case 6:
			numerator := registers[a]
			denominator := int(1 << comboOperand(operand, registers))
			registers[b] = numerator / denominator

		case 7:
			numerator := registers[a]
			denominator := int(1 << comboOperand(operand, registers))
			registers[c] = numerator / denominator
		}

		instruction_pointer++
	}

	return strings.TrimSuffix(output, ",")
}

func calculatePart2() string {
	result := ""

	return result
}
func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	registers, opcodes, operands := parseLines(lines)
	result := calculatePart1(registers, opcodes, operands)
	fmt.Printf("part1: %v\n", result)
	//
	// result2 := calculatePart2(grid, walls, boxesNew, robot_start_Position, moves, dots)
	// fmt.Printf("part2: %d\n", result2)

	executionTime := time.Since(startTime)
	fmt.Printf("Execution time: %v\n", executionTime)
}

func readLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return strings.Split(string(file), "\n\n")

}
