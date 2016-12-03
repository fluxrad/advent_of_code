package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Location struct {
	Orientation int
	X           int
	Y           int
}

func (l *Location) Go(turn string, blocks int) {
	log.Debugf("Turning %s and traveling %d blocks", turn, blocks)

	// Mod 4
	switch turn {
	case "R":
		l.Orientation++
	case "L":
		l.Orientation--
	}

	switch math.Mod(float64(l.Orientation), 4) {
	case 0: // North
		l.Y += blocks
	case 1:
		l.X += blocks
	case 2:
		l.Y -= blocks
	case 3:
		l.X -= blocks
	}

	log.Debugf("You are at coordinates: %d,%d", l.X, l.Y)
	log.Debugf("You are presently %d blocks from start", l.DistanceFromStart())
}

func (l *Location) DistanceFromStart() int {
	return int((math.Abs(float64(l.X)) + math.Abs(float64(l.Y))))
}

func main() {
	log.SetLevel(log.InfoLevel)

	l := &Location{
		Orientation: 0,
		X:           0,
		Y:           0,
	}

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(input), "\n")

	directions := strings.Split(s, ", ")
	for _, d := range directions {
		direction := string(d[0])
		distance, err := strconv.Atoi(d[1:])
		if err != nil {
			log.Fatal(err)
		}

		l.Go(direction, distance)
	}

	fmt.Printf("You are: %d blocks away\n", l.DistanceFromStart())
}
