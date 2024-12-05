package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	m := make(map[int]map[int]bool)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		nums_strs := strings.Split(scanner.Text(), "|")
		nums := make([]int, len(nums_strs))
		for i, v := range nums_strs {
			n, err := strconv.Atoi(v)
			if err != nil {
				return
			}
			nums[i] = n
		}
		if m[nums[1]] == nil {
			m[nums[1]] = make(map[int]bool)
		}
		m[nums[1]][nums[0]] = true
	}
	var val int
	var val_2 int
	for scanner.Scan() {
		nums_strs := strings.Split(scanner.Text(), ",")
		nums := make([]int, len(nums_strs))
		for i, v := range nums_strs {
			n, err := strconv.Atoi(v)
			if err != nil {
				return
			}
			nums[i] = n
		}
		ok := true
		for i := range nums {
			for j := i + 1; j < len(nums); j++ {
				if m[nums[j]] == nil || !m[nums[j]][nums[i]] {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			val += nums[len(nums)/2]
		} else {
			slices.SortFunc(nums, func(x, y int) int {
				if m[x] != nil && m[x][y] {
					return 1
				} else if m[y] != nil && m[y][x] {
					return -1
				} else {
					return 0
				}
			})
			val_2 += nums[len(nums)/2]
		}
	}
	fmt.Println("val 1 : ", val)
	fmt.Println("val 2 : ", val_2)
}
