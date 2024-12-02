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
	var l1, l2 []int
	count := make(map[int]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(strings.Trim(scanner.Text(), " "), "   ")
		num1, _ := strconv.Atoi(data[0])
		num2, _ := strconv.Atoi(data[1])

		l1 = append(l1, num1)
		l2 = append(l2, num2)
		count[num2] = count[num2] + 1
	}
	slices.Sort(l1)
	slices.Sort(l2)

	var dist int
	var similarity int
	for i := range l1 {
		v := l1[i] - l2[i]
		if v < 0 {
			v = v * -1
		}
		dist += v

		similarity += count[l1[i]] * l1[i]
	}

	fmt.Println(dist)
	fmt.Println(similarity)
}
