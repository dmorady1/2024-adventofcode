package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// lines := readLines(os.Args[1])
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	result := calculatePart1(string(data))
	fmt.Printf("part1: %d\n", result)

	result2 := calculatePart2(string(data))
	fmt.Printf("part2: %d\n", result2)
}
func mul(num1 int, num2 int) int {
	return num1 * num2
}

// func calculatePart1(text string) int {
// 	// re := regexp.MustCompile(`mul(\(\d{1,3},\d{1,3}\))`)
// 	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
// 	// matches := re.FindAllString(text, -1)
// 	subMatches := re.FindAllStringSubmatch(text, -1)
//
// 	var result int = 0
// 	for _, subMatch := range subMatches {
// 		num1, err := strconv.Atoi(subMatch[1])
// 		if err != nil {
// 			panic(err)
// 		}
// 		num2, err := strconv.Atoi(subMatch[2])
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		result += num1 * num2
// 	}
// 	return result
//
// }
//

func calculatePart1(text string) int {

	state := StateIdle
	var num1 int
	var num2 int
	var result int = 0
	for _, ch := range text {
		if state == StateIdle {
			num1 = 0
			num2 = 0
		}
		state = transition(state, ch)
		// fmt.Println(state)
		switch state {
		case StateDigit1:
			num1 = int(ch - '0')
			fmt.Println(num1)
		case StateDigit2:
			num1 = num1*10 + int(ch-'0')
			fmt.Println(num1)
		case StateDigit3:
			num1 = num1*10 + int(ch-'0')
			fmt.Println(num1)
		case StateDigit4:
			num2 = int(ch - '0')
		case StateDigit5:
			num2 = num2*10 + int(ch-'0')
		case StateDigit6:
			num2 = num2*10 + int(ch-'0')
		case StateEB:
			fmt.Println(num1, num2)
			result += num1 * num2
			state = StateIdle
		}

	}

	return result
}

type State int

const (
	StateIdle State = iota
	StateM
	StateU
	StateL
	StateOB
	StateDigit1
	StateDigit2
	StateDigit3
	StateCOMMA
	StateDigit4
	StateDigit5
	StateDigit6
	StateEB
	StateD
	StateDO
	StateDON
	StateDONA
	StateDONAT
	StateDONATOB
	StateDONATEB
	StateDOOB
	StateDOEB
)

func transition(state State, ch rune) State {
	// fmt.Println(state)
	switch state {
	case StateIdle:
		if ch == 'm' {
			return StateM
		}
		return StateIdle
	case StateM:
		fmt.Println(ch)
		if ch == 'u' {
			return StateU
		}
		fmt.Println(ch)
		return StateIdle

	case StateU:
		if ch == 'l' {
			return StateL
		}
		return StateIdle

	case StateL:
		if ch == '(' {
			return StateOB
		}
		return StateIdle

	case StateOB:
		if unicode.IsDigit(ch) {
			return StateDigit1
		}
		return StateIdle
	case StateDigit1:
		if unicode.IsDigit(ch) {
			return StateDigit2
		}
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle
	case StateDigit2:
		if unicode.IsDigit(ch) {
			return StateDigit3
		}
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle

	case StateDigit3:
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle
	case StateCOMMA:
		if unicode.IsDigit(ch) {
			return StateDigit4
		}

		return StateIdle
	case StateDigit4:
		if unicode.IsDigit(ch) {
			return StateDigit5
		}

		if ch == ')' {
			return StateEB
		}
		return StateIdle
	case StateDigit5:
		if unicode.IsDigit(ch) {
			return StateDigit6
		}

		if ch == ')' {
			return StateEB
		}
		return StateIdle

	case StateDigit6:
		if ch == ')' {
			return StateEB
		}
		return StateIdle

	case StateEB:
		return StateIdle

	default:
		panic(fmt.Errorf("unknown state: %d", state))
	}
}

func transition2(state State, ch rune) State {
	switch state {
	case StateIdle:
		if ch == 'm' {
			return StateM
		}
		if ch == 'd' {
			return StateD
		}
		return StateIdle

	case StateM:
		if ch == 'u' {
			return StateU
		}
		return StateIdle

	case StateU:
		if ch == 'l' {
			return StateL
		}
		return StateIdle

	case StateL:
		if ch == '(' {
			return StateOB
		}
		return StateIdle

	case StateOB:
		if unicode.IsDigit(ch) {
			return StateDigit1
		}
		return StateIdle

	case StateDigit1:
		if unicode.IsDigit(ch) {
			return StateDigit2
		}
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle

	case StateDigit2:
		if unicode.IsDigit(ch) {
			return StateDigit3
		}
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle

	case StateDigit3:
		if ch == ',' {
			return StateCOMMA
		}
		return StateIdle

	case StateCOMMA:
		if unicode.IsDigit(ch) {
			return StateDigit4
		}
		return StateIdle

	case StateDigit4:
		if unicode.IsDigit(ch) {
			return StateDigit5
		}
		if ch == ')' {
			return StateEB
		}
		return StateIdle

	case StateDigit5:
		if unicode.IsDigit(ch) {
			return StateDigit6
		}
		if ch == ')' {
			return StateEB
		}
		return StateIdle

	case StateDigit6:
		if ch == ')' {
			return StateEB
		}
		return StateIdle

	case StateEB:
		return StateIdle

	case StateD:
		if ch == 'o' {
			return StateDO
		}
		return StateIdle

	case StateDO:
		if ch == 'n' {
			return StateDON
		}
		if ch == '(' {
			return StateDOOB
		}
		return StateIdle

	case StateDON:
		if ch == '\'' {
			return StateDONA
		}
		return StateIdle

	case StateDONA:
		if ch == 't' {
			return StateDONAT
		}
		return StateIdle

	case StateDONAT:
		if ch == '(' {
			return StateDONATOB
		}
		return StateIdle

	case StateDONATOB:
		if ch == ')' {
			return StateDONATEB
		}
		return StateIdle

	case StateDONATEB:
		return StateIdle

	case StateDOOB:
		if ch == ')' {
			return StateDOEB
		}
		return StateIdle

	case StateDOEB:
		return StateIdle

	default:
		return StateIdle
	}
}

func calculatePart2(text string) int {
	state := StateIdle
	var num1 int
	var num2 int
	var result int = 0
	mulEnabled := true

	for _, ch := range text {
		if state == StateIdle {
			num1 = 0
			num2 = 0
		}

		prevState := state
		state = transition2(state, ch)

		// Debug output
		if state != prevState {
			fmt.Printf("Char: %c, New State: %d, MulEnabled: %v\n", ch, state, mulEnabled)
		}

		switch state {
		case StateDigit1:
			num1 = int(ch - '0')
		case StateDigit2:
			num1 = num1*10 + int(ch-'0')
		case StateDigit3:
			num1 = num1*10 + int(ch-'0')
		case StateDigit4:
			num2 = int(ch - '0')
		case StateDigit5:
			num2 = num2*10 + int(ch-'0')
		case StateDigit6:
			num2 = num2*10 + int(ch-'0')
		case StateEB:
			if mulEnabled {
				result += num1 * num2
				fmt.Printf("Multiplication: %d * %d = %d (Running total: %d)\n",
					num1, num2, num1*num2, result)
			} else {
				fmt.Printf("Skipped multiplication: %d * %d (multiplication disabled)\n",
					num1, num2)
			}
			state = StateIdle
		case StateDONATEB: // don't() completed
			mulEnabled = false
			fmt.Println("Multiplication disabled")
			state = StateIdle
		case StateDOEB: // do() completed
			mulEnabled = true
			fmt.Println("Multiplication enabled")
			state = StateIdle
		}
	}

	return result
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
