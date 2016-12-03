package main

import (
	"bufio"
	"log"
	"os"
)

type Wire struct {
	Name   string
	Signal uint16
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
	}
}
