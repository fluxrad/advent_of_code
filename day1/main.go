package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Couldn't read file %s\n", err)
	}

	floor := 0

	for _, v := range data {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		case '\n':
			break
		default:
			fmt.Printf("Invalid character: %q, skipping\n", v)
		}
	}
	fmt.Println("Floor: ", floor)
}
