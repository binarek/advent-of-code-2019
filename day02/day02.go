package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	part2Result = 19690720
)

func main() {
	program := readProgram()

	result1, err1 := SolvePart1(program)
	printSolveResult("Part 1:", result1, err1)

	result2, err2 := SolvePart2(program)
	printSolveResult("Part 2:", result2, err2)
}

func SolvePart1(program []int) (int, error) {
	programCopy := append(make([]int, 0, len(program)), program...)
	programCopy[1] = 12
	programCopy[2] = 2
	err := executeProgram(programCopy)
	return programCopy[0], err
}

func SolvePart2(program []int) (int, error) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			programCopy := append(make([]int, 0, len(program)), program...)
			programCopy[1] = noun
			programCopy[2] = verb
			err := executeProgram(programCopy)
			if err == nil && programCopy[0] == part2Result {
				return 100*noun + verb, nil
			}
		}
	}
	return 0, errors.New("Cannot find solution")
}

func readProgram() []int {
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day02_input"
	}())
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strNums := strings.Split(scanner.Text(), ",")
		program := make([]int, len(strNums))
		for i, strNum := range strNums {
			program[i], _ = strconv.Atoi(strNum)
		}
		return program
	}
	return nil
}

func printSolveResult(header string, result int, err error) {
	if err != nil {
		fmt.Println(header, err)
	} else {
		fmt.Println(header, result)
	}
}

func executeProgram(program []int) error {
	programLen := len(program)
executeLoop:
	for i := 0; i < programLen; i = i + 4 {
		opcode := program[i]
		switch opcode {
		case 1, 2:
			if !validateParams(programLen, i) {
				return errors.New("Invalid parameter - index out of range")
			}
			if opcode == 1 {
				program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
			} else if opcode == 2 {
				program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
			}
		case 99:
			break executeLoop
		default:
			return errors.New("Invalid opcode " + strconv.Itoa(opcode))
		}
	}
	return nil
}

func validateParams(programLen int, idx int) bool {
	if idx+1 < 0 || idx+1 >= programLen || idx+2 >= programLen || idx+3 >= programLen {
		return false
	}
	return true
}
