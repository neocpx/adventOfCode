package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("failed to open file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var safe int32
	var safe_dampener int32
	for scanner.Scan() {
		str_levels := strings.Split(scanner.Text(), " ")
		levels := make([]int32, len(str_levels))
		for i := range str_levels {
			v, err := strconv.Atoi(str_levels[i])
			if err != nil {
				fmt.Errorf("failed to parse value")
				return
			}
			levels[i] = int32(v)
		}
		if inc(levels) || dec(levels) {
			safe++
		}
		if inc_dampener(levels) || dec_dampener(levels) {
			fmt.Println(levels)
			safe_dampener++
		}
	}
	fmt.Println("safe values :", safe)
	fmt.Println("safe values with dampener :", safe_dampener)
}

func inc(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] || a[i]-a[i-1] > 3 {
			return false
		}
	}
	return true
}
func dec(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] >= a[i-1] || a[i-1]-a[i] > 3 {
			return false
		}
	}
	return true
}
func inc_dampener(a []int32) bool {
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] || a[i]-a[i-1] > 3 {
			return inc(append(a[:i-1], a[i:]...)) || inc(append(a[:i], a[i+1:]...))
		}
	}
	return true
}
func dec_dampener(a []int32) bool {
	fmt.Println(a)
	for i := 1; i < len(a); i++ {
		if a[i] >= a[i-1] || a[i-1]-a[i] > 3 {
			return dec(append(a[:i-1], a[i:]...)) || dec(append(a[:i], a[i+1:]...))
		}
	}
	return true
}
