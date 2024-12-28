package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseLines(splittedByNewLines []string) ([]uint, []uint, []uint, string) {
	registersString := splittedByNewLines[0]
	programString := splittedByNewLines[1]

	lines := strings.Split(registersString, "\n")

	var registers []uint
	re := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		numbers := re.FindAllString(line, -1)
		num, err := strconv.ParseUint(numbers[0], 10, 32)
		if err != nil {
			fmt.Println("Error:", err)
		}
		registers = append(registers, uint(num))
	}

	programNumbers := re.FindAllString(programString, -1)

	var opCodes []uint
	var operands []uint

	for i := 0; i < len(programNumbers)-1; i = i + 2 {
		opCode, err := strconv.ParseUint(programNumbers[i], 10, 32)
		if err != nil {
			fmt.Println("Error:", err)
		}
		opCodes = append(opCodes, uint(opCode))

		operand, err := strconv.ParseUint(programNumbers[i+1], 10, 32)
		if err != nil {
			fmt.Println("Error:", err)
		}
		operands = append(operands, uint(operand))
	}

	return registers, opCodes, operands, strings.Split(programString, " ")[1]
}

func literalOperand(operand uint) uint {
	return operand
}

func comboOperand(operand uint, registers []uint) uint {
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

func calculatePart1(registers []uint, opcodes []uint, operands []uint) string {
	output := ""

	if len(opcodes) != len(operands) {
		panic("opcodes len is not equal to operands len")
	}

	a := uint(0)
	b := uint(1)
	c := uint(2)

	instruction_pointer := uint(0)
	for instruction_pointer < uint(len(opcodes)) {
		if instruction_pointer >= uint(len(opcodes)) {
			break
		}
		opcode := opcodes[instruction_pointer]
		operand := operands[instruction_pointer]

		switch opcode {
		case 0:
			numerator := registers[a]
			denominator := uint(1) << comboOperand(operand, registers)
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
			output = output + strconv.FormatUint(uint64(number), 10) + ","
		case 6:
			numerator := registers[a]
			denominator := uint(1) << comboOperand(operand, registers)
			registers[b] = numerator / denominator
		case 7:
			numerator := registers[a]
			denominator := uint(1) << comboOperand(operand, registers)
			registers[c] = numerator / denominator
		}

		instruction_pointer++
	}

	return strings.TrimSuffix(output, ",")
}

func testfunction(A uint) uint {
	fmt.Printf("A (decimal): %d\n", A)
	fmt.Printf("A (binary):  %032b\n", A)

	B := A % 8
	fmt.Printf("B after A %% 8:      %03b\n", B)

	B = B ^ 5
	fmt.Printf("B after XOR 5:       %03b\n", B)

	C := A / (1 << B)
	fmt.Printf("C after A / (1<<B):  %032b\n", C)

	B = B ^ 6
	fmt.Printf("B after XOR 6:       %03b\n", B)

	A = A / 8
	fmt.Printf("A after / 8:         %032b\n", A)

	B = B ^ C
	fmt.Printf("B after XOR C:   %032b\n", B)

	output := B % 8
	fmt.Printf("Output:              %032b\n", output)
	fmt.Println("--------------------")

	return output
}

func programFunc(A uint) []uint {
	var result []uint
	for A != 0 {

		B := A % 8

		B = B ^ 5

		C := A / (1 << B)

		B = B ^ 6

		A = A / 8

		B = B ^ C

		output := B % 8
		fmt.Println(output)
		result = append(result, output)

	}
	return result
}

func binaryToUint(binary string) uint {
	i, _ := strconv.ParseUint(binary, 2, 64)
	return uint(i)
}

func uintToBinary(n uint) string {
	binary := strconv.FormatUint(uint64(n), 2)
	// Pad with zeros to ensure at least 3 characters
	for len(binary) < 3 {
		binary = "0" + binary
	}
	// Keep only the last 3 characters
	return binary[len(binary)-3:]
}

func recursiveTest(registers []uint, opcodes []uint, operands []uint, expected string, program []uint, A uint, depth uint) uint {
	// Test current value
	result := calculatePart1([]uint{A, 0, 0}, opcodes, operands)

	// If we found a match, return the current value
	if result == expected {
		return A
	}

	// Limit recursion depth
	if depth >= 16 {
		return 0
	}

	// Try all possible 3-bit values (0-7)
	baseA := A << 3
	for i := uint(0); i < 8; i++ {
		newA := baseA | i
		testResult := calculatePart1([]uint{newA, 0, 0}, opcodes, operands)

		// Check if this combination could lead to the expected result
		if strings.HasSuffix(expected, testResult) {
			result := recursiveTest(registers, opcodes, operands, expected, program, newA, depth+1)
			if result != 0 {
				return result
			}
		}
	}

	return 0
}

func calculatePart2(registers []uint, opcodes []uint, operands []uint, expected string, program []uint) uint {
	return recursiveTest(registers, opcodes, operands, expected, program, 0, 0)
}

func main() {
	startTime := time.Now()

	lines := readLines(os.Args[1])
	registers, opcodes, operands, expected := parseLines(lines)
	expected = strings.ReplaceAll(expected, "\n", "")
	A := registers[0]
	result := calculatePart1(registers, opcodes, operands)

	fmt.Println(expected)
	fmt.Println(programFunc(A))
	fmt.Printf("part1: %v\n", result)

	registers[0] = A
	registers[1] = 0
	registers[2] = 0

	var program []uint

	for index := 0; index < len(opcodes); index++ {
		program = append(program, opcodes[index])
		program = append(program, operands[index])
	}

	result2 := calculatePart2(registers, opcodes, operands, expected, program)
	fmt.Printf("part2: %v\n", result2)
	//
	// intToBinaryMap := make(map[uint][]string)
	// var output uint
	//
	// indexToCandidates := make(map[int][]uint)
	// fmt.Println("hier")
	// for outputIndex := len(program) - 1; outputIndex >= 0; outputIndex-- {
	// 	for try := uint(0); try < 8; try++ {
	// 		registers[0] = try
	// 		registers[1] = 0
	// 		registers[2] = 0
	//
	// 		test := calculatePart1(registers, opcodes, operands)
	// 		fmt.Println(test)
	// 		// if program[outputIndex] == test {
	// 		// 	indexToCandidates[outputIndex] = append(indexToCandidates[outputIndex], try)
	// 		// }
	// 		// test := calculatePart1(registers, opcodes, operands)
	// 	}
	// }
	// fmt.Println(indexToCandidates)
	// return
	//
	// for A := uint(0); A < 10; A++ {
	// 	output = testfunction(A)
	// 	fmt.Println(output)
	// }
	//
	// fmt.Println(testfunction(binaryToUint("10100000000000000000000000000000")))
	// registers[0] = binaryToUint("10000000000000000000000000000000")
	//
	// fmt.Println("--------------asdjfkasjdfsafjaskdfjasdjfasjdfasjdfjaskfdjaskdjfaskdjfkasdj")
	// for i := 0; i < 10; i++ {
	//
	// 	fmt.Println(i)
	// 	registers[0] = 2 << 45
	// 	registers[1] = 0
	// 	registers[2] = 0
	// 	fmt.Println(calculatePart1(registers, opcodes, operands))
	// }
	// fmt.Println(intToBinaryMap)

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
