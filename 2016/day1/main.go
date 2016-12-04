package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Location struct {
	Orientation int
	*Coords
}

type Coords struct {
	X, Y int
}

var visited map[Coords]bool

func (c *Coords) String() string {
	return fmt.Sprintf("[%d,%d]", c.X, c.Y)
}

func (l *Location) Go(turn string, blocks int) bool {
	log.Debugf("Turning %s and traveling %d blocks", turn, blocks)

	// Mod 4
	switch turn {
	case "R":
		l.Orientation++
	case "L":
		l.Orientation--
	}

	var axis *int
	var modifier int

	// Figure out which coordinate to change.
	switch math.Mod(float64(l.Orientation), 4) {
	case 0: // North
		axis = &l.Y
		modifier = 1
	case 1: // East
		axis = &l.X
		modifier = 1
	case 2: // South
		axis = &l.Y
		modifier = -1
	case 3: // West
		axis = &l.X
		modifier = -1
	}

	// Move one step in each dirction
	for i := 1; i <= blocks; i++ {
		*axis += 1 * modifier

		log.Debugf("You are at coordinates: %s", l.Coords)
		log.Debugf("You are presently %d blocks from start", l.DistanceFromStart())

		if visited[*l.Coords] {
			log.Debugf("Already visited %s, returning.", l.Coords)
			return true
		}
		visited[*l.Coords] = true
	}

	return false
}

func (l *Location) DistanceFromStart() int {
	return int((math.Abs(float64(l.X)) + math.Abs(float64(l.Y))))
}

func init() {
	visited = make(map[Coords]bool)
}

func main() {
	log.SetLevel(log.InfoLevel)

	l := &Location{
		Coords: &Coords{},
	}

	input, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Trim(string(input), "\n")
	for _, d := range strings.Split(s, ", ") {
		direction := string(d[0])
		distance, err := strconv.Atoi(d[1:])
		if err != nil {
			log.Fatal(err)
		}

		if v := l.Go(direction, distance); v == true {
			fmt.Printf("Visited %s twice\n", l.Coords)
			fmt.Printf("It's %d blocks away\n", l.DistanceFromStart())
			os.Exit(0)
		}
	}

	fmt.Printf("You are: %d blocks away\n", l.DistanceFromStart())
}
