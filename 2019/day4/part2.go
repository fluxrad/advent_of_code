package main

import (
	"fmt"
	"strconv"
	"strings"
)

func has_double(s string) bool {
	prev := -1
	count := 0
	okay := false

	for _, r := range s {
		n, _ := strconv.Atoi(string(r))
		if n == prev {
			count++
			if count == 2 {
				okay = true
			} else {
				okay = false
			}
		} else {
			if count == 2 {
				return true
			}
			okay = false
			count = 1
		}

		prev = n
	}

	return okay
}

func decreases(s string) bool {
	prev := -1
	for _, r := range s {
		n, _ := strconv.Atoi(string(r))
		if n < prev {
			return true
		}
		prev = n
	}
	return false
}

func main() {
	input := "265275-781584"
	meh := strings.Split(input, "-")

	lower, _ := strconv.Atoi(meh[0])
	upper, _ := strconv.Atoi(meh[1])

	count := 0
	for i := lower + 1; i < upper; i++ {
		num := strconv.Itoa(i)
		if has_double(num) && !decreases(num) {
			count++
		}
	}

	fmt.Printf("The answer is %d\n", count)
}
