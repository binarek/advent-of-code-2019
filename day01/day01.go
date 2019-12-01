package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	masses := readMasses()
	fmt.Println("Part 1:", SolvePart1(masses))
	fmt.Println("Part 2:", SolvePart2(masses))
}

func readMasses() []int {
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day01_input"
	}())
	scanner := bufio.NewScanner(file)
	var masses []int
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		masses = append(masses, mass)
	}
	return masses
}

func SolvePart1(masses []int) int {
	fuel := 0
	for _, mass := range masses {
		fuel += calculateFuel(mass)
	}
	return fuel
}

func SolvePart2(masses []int) int {
	fuel := 0
	for _, mass := range masses {
		fuel += calculateFullFuel(mass)
	}
	return fuel
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFullFuel(mass int) int {
	fullFuel := 0
	fuel := calculateFuel(mass)
	for fuel > 0 {
		fullFuel += fuel
		fuel = calculateFuel(fuel)
	}
	return fullFuel
}
