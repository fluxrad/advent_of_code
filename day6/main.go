package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

var grid [][]bool

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	grid = make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		instruction := scanner.Text()
		if err := playWithLights(instruction); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("%d lights are on\n", countLights(grid))
}

func playWithLights(i string) error {
	var oper string

	switch {
	case regexp.MustCompile("^turn on").MatchString(i):
		oper = "on"
	case regexp.MustCompile("^turn off").MatchString(i):
		oper = "off"
	case regexp.MustCompile("^toggle").MatchString(i):
		oper = "toggle"
	default:
		log.Fatal("Invalid instruction: %s", i)
	}

	pairs := regexp.MustCompile("[0-9]+,[0-9]+").FindAllString(i, 2)
	turnLights(oper, pairs[0], pairs[1])

	return nil
}

func turnLights(oper string, start string, end string) {
	s := strings.Split(start, ",")
	e := strings.Split(end, ",")

	// Don't throw away err in prod
	startx, _ := strconv.Atoi(s[0])
	endx, _ := strconv.Atoi(e[0])

	starty, _ := strconv.Atoi(s[1])
	endy, _ := strconv.Atoi(e[1])

	log.Debugf("%s lights,  %d,%d - %d,%d", oper, startx, starty, endx, endy)

	switch oper {
	case "on":
		for x := startx; x <= endx; x++ {
			for y := starty; y <= endy; y++ {
				grid[x][y] = true
				log.Debugf("%d,%d set to %t", x, y, grid[x][y])
			}

		}

	case "off":
		for x := startx; x <= endx; x++ {
			for y := starty; y <= endy; y++ {
				grid[x][y] = false
				log.Debugf("%d,%d set to %t", x, y, grid[x][y])
			}

		}

	case "toggle":
		for x := startx; x <= endx; x++ {
			for y := starty; y <= endy; y++ {
				grid[x][y] = !grid[x][y]
				log.Debugf("%d,%d set to %t", x, y, grid[x][y])
			}

		}
	}
}

func countLights(l [][]bool) int {
	count := 0
	for i := range l {
		for j := range l[i] {
			if l[i][j] {
				count++
			}
		}
	}
	return count
}
