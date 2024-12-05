package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to read input.txt")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid [][]rune

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	m := len(grid)
	n := len(grid[0])
	count := 0
	count_x := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				count += searchDirections(&grid, i, j, "XMAS")
			} else if grid[i][j] == 'S' {
				count += searchDirections(&grid, i, j, "SAMX")
			} else if grid[i][j] == 'A' && checkX(&grid, i, j) {
				count_x++
			}
		}
	}

	fmt.Println("Total count:", count)
	fmt.Println("Total X:", count_x)
}

func checkX(grid *[][]rune, r, c int) bool {
	if r <= 0 || c <= 0 || r >= len(*grid)-1 || c >= len((*grid)[0])-1 {
		return false
	}
	st := string([]rune{(*grid)[r-1][c-1], (*grid)[r][c], (*grid)[r+1][c+1]})
	if st != "MAS" && st != "SAM" {
		return false
	}
	st = string([]rune{(*grid)[r-1][c+1], (*grid)[r][c], (*grid)[r+1][c-1]})
	if st != "MAS" && st != "SAM" {
		return false
	}
	return true
}

func searchDirections(grid *[][]rune, r, c int, s string) int {
	count := 0
	if down(grid, r, c, s) {
		count++
	}
	if right(grid, r, c, s) {
		count++
	}
	if diagonal(grid, r, c, s, 1) {
		count++
	}
	if diagonal(grid, r, c, s, -1) {
		count++
	}
	return count
}

func down(grid *[][]rune, r, c int, s string) bool {
	if r+len(s)-1 >= len(*grid) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if rune(s[i]) != (*grid)[r+i][c] {
			return false
		}
	}
	return true
}

func right(grid *[][]rune, r, c int, s string) bool {
	if c+len(s)-1 >= len((*grid)[0]) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if rune(s[i]) != (*grid)[r][c+i] {
			return false
		}
	}
	return true
}

func diagonal(grid *[][]rune, r, c int, s string, dir int) bool {
	if dir == 1 {
		if r+len(s)-1 >= len(*grid) || c+len(s)-1 >= len((*grid)[0]) {
			return false
		}
		for i := 0; i < len(s); i++ {
			if rune(s[i]) != (*grid)[r+i][c+i] {
				return false
			}
		}
	} else if dir == -1 {
		if r+len(s)-1 >= len(*grid) || c-len(s)+1 < 0 {
			return false
		}
		for i := 0; i < len(s); i++ {
			if rune(s[i]) != (*grid)[r+i][c-i] {
				return false
			}
		}
	}
	return true
}
