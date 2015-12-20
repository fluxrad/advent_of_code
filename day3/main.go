package main

import (
	"fmt"
	"io/ioutil"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type Houses struct {
	sync.Mutex
	Visited map[[2]int]bool
}

func main() {
	log.SetLevel(log.ErrorLevel)

	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Couldn't read file %s", err)
	}

	houses := &Houses{
		Visited: make(map[[2]int]bool),
	}
	santa := make(chan string)
	roboSanta := make(chan string)

	go deliverPresents(santa, houses)
	go deliverPresents(roboSanta, houses)

	for i, d := range data {
		if i%2 == 0 {
			santa <- string(d)
		} else {
			roboSanta <- string(d)
		}
	}
	close(santa)
	close(roboSanta)

	fmt.Printf("Total houses visited: %d\n", len(houses.Visited))
}

func deliverPresents(direction <-chan string, h *Houses) {
	point := [2]int{0, 0}

	for d := range direction {
		p := [2]int{}
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

		if h.Visited[p] == true {
			log.Warnf("%v == %v, we've been to this house before.", p, h.Visited[p])
		}

		h.Lock()
		h.Visited[p] = true
		h.Unlock()
		point = p
	}
}
