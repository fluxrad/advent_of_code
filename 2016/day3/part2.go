package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	// Input is a little janky so regex it out.
	r := regexp.MustCompile("[0-9]+")

	possibleTriangles := 0
	candidates := [][]int{}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		for i, v := range r.FindAllString(scanner.Text(), -1) {
			s, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("wtf")
			}

			// New batch of candidates
			if len(candidates) < 3 {
				candidates = append(candidates, []int{s})
			} else {
				candidates[i] = append(candidates[i], s)
			}
		}

		// Are all three candidates fully populated?
		if len(candidates) == 3 && len(candidates[2]) == 3 {
			for _, candidate := range candidates {
				sort.Ints(candidate)
				if candidate[0]+candidate[1] > candidate[2] {
					possibleTriangles++
				}
			}

			// Start over
			candidates = candidates[:0]
		}
	}

	fmt.Printf("There are %d possible triangles\n", possibleTriangles)
}
