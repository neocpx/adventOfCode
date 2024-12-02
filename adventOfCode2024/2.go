package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safe int32
	var safeDampener int32
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		strLevels := strings.Split(line, " ")
		levels := make([]int32, len(strLevels))
		for i := range strLevels {
			v, err := strconv.Atoi(strLevels[i])
			if err != nil {
				log.Fatalf("Failed to parse value '%s': %v", strLevels[i], err)
			}
			levels[i] = int32(v)
		}

		// Check safety conditions
		if isIncreasing(levels) || isDecreasing(levels) {
			safe++
		}
		if isDampenedIncreasing(levels) || isDampenedDecreasing(levels) {
			safeDampener++
		}
	}

	fmt.Println("Safe values:", safe)
	fmt.Println("Safe values with dampener:", safeDampener)
}

func isIncreasing(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] || a[i]-a[i-1] > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] >= a[i-1] || a[i-1]-a[i] > 3 {
			return false
		}
	}
	return true
}

func isDampenedIncreasing(a []int32) bool {
	for i := range a {
		if isIncreasing(removeIndex(a, i)) {
			return true
		}
	}
	return false
}

func isDampenedDecreasing(a []int32) bool {
	for i := range a {
		if isDecreasing(removeIndex(a, i)) {
			return true
		}
	}
	return false
}

func removeIndex(a []int32, index int) []int32 {
	result := make([]int32, 0, len(a)-1)
	result = append(result, a[:index]...)
	return append(result, a[index+1:]...)
}
