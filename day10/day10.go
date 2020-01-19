package main

import (
	"bufio"
	"fmt"
	"os"
)

type spaceVal rune

const (
	astroid spaceVal = '#'
	empty   spaceVal = '.'
)

func main() {
	regionMap, _ := readRegionMap()
	fmt.Println(SolvePart1(regionMap))
}

func readRegionMap() ([][]spaceVal, error) {
	file, _ := os.Open(func() string {
		if len(os.Args) >= 2 {
			return os.Args[1]
		}
		return "./day10_input"
	}())
	defer file.Close()

	regionMap := make([][]spaceVal, 0, 30)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLine := scanner.Text()
		regionLine := make([]spaceVal, len(fileLine))
		for idx, char := range fileLine {
			regionLine[idx] = spaceVal(char)
		}
		regionMap = append(regionMap, regionLine)
	}
	if err := scanner.Err(); err != nil {
		return regionMap, err
	}
	return regionMap, nil
}

func SolvePart1(regionMap [][]spaceVal) int {
	maxAstroids := 0
	for y, xList := range regionMap {
		for x, space := range xList {
			if space == astroid {
				astroids := calcVisibleAstroids(regionMap, x, y)
				if astroids > maxAstroids {
					maxAstroids = astroids
				}
			}
		}
	}
	return maxAstroids
}

func calcVisibleAstroids(regionMap [][]spaceVal, candidateX, candidateY int) int {
	dirs := make(map[int][]int) // key: x, value: y
	for y, xList := range regionMap {
		for x, space := range xList {

			if space == astroid && (x != candidateX || y != candidateY) {
				dirX, dirY := calcDirection(x, candidateX, y, candidateY)
				if !contains(dirs[dirX], dirY) {
					dirs[dirX] = append(dirs[dirX], dirY)
				}
			}
		}
	}
	dirsCount := 0
	for _, dirsY := range dirs {
		dirsCount += len(dirsY)
	}
	return dirsCount
}

func calcDirection(x1, x2, y1, y2 int) (int, int) {
	dirX, dirY := x1-x2, y1-y2
	gcd := abs(gcd(dirX, dirY))
	return dirX / gcd, dirY / gcd
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func contains(slice []int, val int) bool {
	for _, sliceVal := range slice {
		if sliceVal == val {
			return true
		}
	}
	return false
}
