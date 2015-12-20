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
	beenInBasement := false

	for i, v := range data {
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

		if floor == -1 && beenInBasement == false {
			fmt.Printf("Entered basement at position: %d\n", i+1)
			beenInBasement = true
		}
	}
	fmt.Println("Floor: ", floor)
}
