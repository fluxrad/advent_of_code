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
	for scanner.Scan() {
		sides := []int{}

		for _, v := range r.FindAllString(scanner.Text(), -1) {
			s, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("wtf")
			}
			sides = append(sides, s)
		}

		// Sort so we get O and A, we'll calculate against H
		sort.Ints(sides)

		if sides[0]+sides[1] > sides[2] {
			possibleTriangles++
		}
	}

	fmt.Printf("There are %d possible triangles\n", possibleTriangles)
}
