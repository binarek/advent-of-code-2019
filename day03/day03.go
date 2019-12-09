package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	path1, path2 := readPaths()

	result1 := SolvePart1(path1, path2)
	fmt.Println("Part 1:", result1)
}

type pathElem struct {
	dir   rune
	value int
}

func SolvePart1(path1 []pathElem, path2 []pathElem) int {
	points1 := calcPoints(path1)
	points2 := calcPoints(path2)
	commonPoints := findCrosses(points1, points2)
	return calcMinDistance(commonPoints)
}

func calcPoints(path []pathElem) map[int][]int {
	x, y := 0, 0
	points := make(map[int][]int)
	for _, elem := range path {
		switch elem.dir {
		case 'D':
			limit := y - elem.value
			for y--; y >= limit; y-- {
				points[y] = append(points[y], x)
			}
			y++
		case 'L':
			limit := x - elem.value
			for x--; x >= limit; x-- {
				points[y] = append(points[y], x)
			}
			x++
		case 'U':
			limit := y + elem.value
			for y++; y <= limit; y++ {
				points[y] = append(points[y], x)
			}
			y--
		case 'R':
			limit := x + elem.value
			for x++; x <= limit; x++ {
				points[y] = append(points[y], x)
			}
			x--
		}
	}
	return points
}

func findCrosses(points1 map[int][]int, points2 map[int][]int) [][2]int {
	var crosses [][2]int = nil
	for y, x1List := range points1 {
		x2List := points2[y]
		for _, x1 := range x1List {
			for _, x2 := range x2List {
				if x1 == x2 {
					crosses = append(crosses, [2]int{y, x1})
				}
			}
		}
	}
	return crosses
}

func calcMinDistance(points [][2]int) int {
	minDist := math.MaxInt32
	for _, point := range points {
		dist := abs(point[0]) + abs(point[1])
		if dist < minDist {
			minDist = dist
		}
	}
	return minDist
}

func readPaths() ([]pathElem, []pathElem) {
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day03_input"
	}())
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path1 := parseLineToPath(scanner.Text())
	scanner.Scan()
	path2 := parseLineToPath(scanner.Text())
	return path1, path2
}

func parseLineToPath(line string) []pathElem {
	commands := strings.Split(line, ",")
	var path = make([]pathElem, len(commands))
	for i, cmd := range commands {
		dir := rune(cmd[0])
		value, _ := strconv.Atoi(string(cmd[1:]))
		path[i] = pathElem{dir, value}
	}
	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
