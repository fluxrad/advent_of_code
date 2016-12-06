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

	scanner := bufio.NewScanner(input)

	trs := [][]int{}
	for scanner.Scan() {

		for i, v := range r.FindAllString(scanner.Text(), -1) {
			s, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("wtf")
			}

			// New batch of triangles
			if len(trs) < 3 {
				trs = append(trs, []int{s})
			} else {
				trs[i] = append(trs[i], s)
			}
		}

		// Are all three triangles fully populated?
		if len(trs) == 3 && len(trs[2]) == 3 {
			for _, triangle := range trs {
				sort.Ints(triangle)
				if triangle[0]+triangle[1] > triangle[2] {
					possibleTriangles++
				}
			}

			// Start over
			trs = trs[:0]
		}
	}

	fmt.Printf("There are %d possible triangles\n", possibleTriangles)
}
