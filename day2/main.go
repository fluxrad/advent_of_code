package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Day 2

The elves are running low on wrapping paper, and so they need to submit an
order for more. They have a list of the dimensions (length l, width w, and
height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which
makes calculating the required wrapping paper for each gift a little easier:
find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves
also need a little extra paper for each present: the area of the smallest side.
*/

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Couldn't open file %s", err)
	}
	defer f.Close()

	totalWrappingPaper := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		totalWrappingPaper += calculateWrappingPaper(scanner.Text())
	}

	fmt.Println("Total wrapping paper needed is ", totalWrappingPaper, " square feet")
}

// Accept LxWxH as a string
func calculateWrappingPaper(s string) int {
	lwh := strings.Split(s, "x")
	l := measure(lwh[0])
	w := measure(lwh[1])
	h := measure(lwh[2])

	surfaces := []int{l * w, w * h, h * l}

	// Find the smallest surface for the extra
	extra := surfaces[0]
	for _, surface := range surfaces {
		if extra > surface {
			extra = surface
		}
	}

	return (2*(surfaces[0]+surfaces[1]+surfaces[2]) + extra)
}

func measure(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
