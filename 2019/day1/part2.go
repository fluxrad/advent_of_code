package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Fuel(mass int) int {
	f := (mass / 3) - 2
	if f <= 0 {
		return 0
	}
	return f + Fuel(f)
}

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
		panic(fmt.Sprintf("Couldn't get input.", err))
	}

	fuel := 0
	for _, i := range input {
		fuel += Fuel(i)
	}

	fmt.Printf("The answer is: %d\n", fuel)
	return
}
