package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result1, _ := SolvePart1(356261, 846303)
	fmt.Println(result1)
}

func SolvePart1(min int32, max int32) (int, error) {
	if min < 100_000 || max > 999_999 {
		return 0, errors.New("Arguments out of range")
	}
	count := 0
main_loop:
	for i := min; i <= max; i++ {
		code := mapToDigitArray(i)
		adjacent := false
		for i := 1; i < 6; i++ {
			if code[i] < code[i-1] {
				continue main_loop
			}
			if code[i] == code[i-1] {
				adjacent = true
			}
		}
		if adjacent {
			count++
		}
	}
	return count, nil
}

func mapToDigitArray(number int32) [6]int8 {
	var code [6]int8
	putDigits(float64(number), 5, &code)
	return code
}

func putDigits(number float64, idx int8, code *[6]int8) {
	if idx < 0 {
		return
	}
	code[5-idx] = int8(int(number/math.Pow(10, float64(idx))) % 10)
	putDigits(number, idx-1, code)
}
