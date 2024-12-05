package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}
	scan := bufio.NewScanner(file)

	reg, err := regexp.Compile("mul\\(([0-9]+),([0-9]+)\\)")
	if err != nil {
		fmt.Println("failed to compile pattern")
		return
	}
	reg_do, err := regexp.Compile("do\\(\\)")
	if err != nil {
		fmt.Println("failed to compile pattern")
		return
	}
	reg_dont, err := regexp.Compile("don't\\(\\)")
	if err != nil {
		fmt.Println("failed to compile pattern")
		return
	}

	uncorrupted_sum := 0
	uncorrupted_sum_filtered := 0
	state := true
	for scan.Scan() {
		patt_do := reg_do.FindAllIndex([]byte(scan.Text()), -1)
		patt_dont := reg_dont.FindAllIndex([]byte(scan.Text()), -1)
		patt_mul := reg.FindAllIndex([]byte(scan.Text()), -1)

		index := make([][]int, 0, len(patt_do)+len(patt_dont)+len(patt_mul))
		index = append(index, patt_do...)
		index = append(index, patt_dont...)
		index = append(index, patt_mul...)
		sort.Slice(index, func(i, j int) bool {
			if index[i][0] == index[j][0] {
				return index[i][1] < index[j][1]
			}
			return index[i][0] < index[j][0]
		})

		patt := reg.FindAll([]byte(scan.Text()), -1)

		for _, v := range patt {
			cal_v, err := calculate_mul(string(v))
			if err != nil {
				fmt.Println("failed to calculate value")
				return
			}
			uncorrupted_sum += cal_v
		}
		for _, v := range index {
			switch v[1] - v[0] {
			case 4:
				state = true
			case 7:
				state = false
			default:
				if state {
					st := scan.Text()[v[0]:v[1]]
					val, err := calculate_mul(st)
					if err != nil {
						fmt.Println("failed to calculate value")
						return
					}
					uncorrupted_sum_filtered += val
				}
			}
		}
	}
	fmt.Println("uncorrupted sum : ", uncorrupted_sum)
	fmt.Println("uncorrupted sum filtered: ", uncorrupted_sum_filtered)
}

func calculate_mul(s string) (int, error) {
	index := strings.IndexAny(s, ",")
	n1, err := strconv.Atoi(s[4:index])
	if err != nil {
		fmt.Println("failed to parse int")
		return 0, err
	}
	n2, err := strconv.Atoi(s[index+1 : len(s)-1])
	if err != nil {
		fmt.Println("failed to parse int")
		return 0, err
	}
	return n1 * n2, nil
}
