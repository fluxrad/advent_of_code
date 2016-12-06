package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// 3x3
var keypad = [][]string{
	{"", "", "", "", "", "", ""},
	{"", "", "", "1", "", "", ""},
	{"", "", "2", "3", "4", "", ""},
	{"", "5", "6", "7", "8", "9", ""},
	{"", "", "A", "B", "C", "", ""},
	{"", "", "", "D", "", "", ""},
	{"", "", "", "", "", "", ""},
}

// instruction map. note x and y coords are reversed
var imap = map[string][2]int{
	// Y operations
	"U": {-1, 0},
	"D": {1, 0},
	// X operations
	"L": {0, -1},
	"R": {0, 1},
}

func move(old, new []int) []int {
	if keypad[new[0]][new[1]] == "" {
		return old
	}
	return new
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	code := []string{}

	// starting point
	yx := []int{3, 1}

	scanner := bufio.NewScanner(input)

	// Get each line
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), "")
		for _, i := range instructions {
			newYX := []int{yx[0] + imap[i][0], yx[1] + imap[i][1]}

			log.WithFields(log.Fields{
				"number":      keypad[yx[0]][yx[1]],
				"instruction": i,
				"old coords":  yx,
				"new coords":  newYX,
				"new number":  keypad[newYX[0]][newYX[1]],
			}).Debug("Parsing instruction")

			yx = move(yx, newYX)
		}
		code = append(code, keypad[yx[0]][yx[1]])
	}

	fmt.Println("Key Code is:", code)
}
