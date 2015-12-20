package main

import (
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.ErrorLevel)

	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Couldn't read file %s", err)
	}

	// we've seen the first point
	point := [2]int{0, 0}
	seen := make(map[[2]int]bool)
	seen[point] = true

	for _, d := range data {
		p := [2]int{}
		d := string(d)

		switch d {
		case ">":
			log.Debugf("Got %s, going east", d)
			p[0], p[1] = point[0]+1, point[1]
		case "<":
			log.Debugf("Got %s, going west", d)
			p[0], p[1] = point[0]-1, point[1]
		case "^":
			log.Debugf("Got %s, going north", d)
			p[0], p[1] = point[0], point[1]+1
		case "v":
			log.Debugf("Got %s, going south", d)
			p[0], p[1] = point[0], point[1]-1
		}

		if seen[p] == true {
			log.Warnf("%v == %v, we've been to this house before.", p, seen[p])
		}
		seen[p] = true

		log.Debugf("We've been to %v houses", len(seen))
		point = p
	}

	fmt.Printf("Visited %d houses\n", len(seen))
}
