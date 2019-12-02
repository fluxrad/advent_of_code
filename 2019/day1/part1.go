package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput() ([]int, error) {
	input := []int{}

	infile, err := os.Open("input")
	if err != nil {
		return input, err
	}
	defer infile.Close()

	inputScanner := bufio.NewScanner(infile)
	for inputScanner.Scan() {
		x, err := strconv.Atoi(inputScanner.Text())
		if err != nil {
			return input, err
		}
		input = append(input, x)
	}

	return input, nil
}

func main() {
	input, err := getInput()
	if err != nil {
		log.Fatal("Couldn't get input.", err)
	}

	fuel := 0
	for i, _ := range input {
		fuel += (i / 3) - 2
	}

	fmt.Printf("The answer is: %d\n", fuel)
	return
}
