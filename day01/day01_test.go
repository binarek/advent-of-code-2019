package main

import "testing"

var testCases = []struct {
	masses   []int
	fuel     int
	fullFuel int
}{
	{nil, 0, 0},
	{[]int{14}, 2, 2},
	{[]int{12}, 2, 2},
	{[]int{1969}, 654, 966},
	{[]int{100756}, 33583, 50346},
	{[]int{1969, 100756}, 34237, 51312},
}

func TestSolvePart1(t *testing.T) {
	for _, c := range testCases {
		resultFuel := SolvePart1(c.masses)
		if resultFuel != c.fuel {
			t.Errorf("SolvePart1(%d), actual %d, expected %d", c.masses, resultFuel, c.fuel)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	for _, c := range testCases {
		resultFullFuel := SolvePart2(c.masses)
		if resultFullFuel != c.fullFuel {
			t.Errorf("SolvePart2(%d): actual %d, expected %d", c.masses, resultFullFuel, c.fullFuel)
		}
	}
}
