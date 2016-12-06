package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/fluxrad/advent_of_code/2016/day3/part1"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	fmt.Println("Part 1")

	part1.Run()
}
