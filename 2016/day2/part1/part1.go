package part1

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func move(old, new int) int {
	if (0 <= new) && (new <= 2) {
		return new
	}
	return old
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func Run() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	code := []int{}

	// 3x3
	var keypad = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
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

	// starting point
	x := 1
	y := 1

	scanner := bufio.NewScanner(input)

	// Get each line
	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), "")
		for _, i := range instructions {
			newY := move(y, y+imap[i][0])
			newX := move(x, x+imap[i][1])

			log.WithFields(log.Fields{
				"number":      keypad[y][x],
				"instruction": i,
				"old coords":  fmt.Sprintf("[%d,%d]", x, y),
				"new coords":  fmt.Sprintf("[%d,%d]", newX, newY),
				"new number":  keypad[newY][newX],
			}).Debug("Parsing instruction")

			x = newX
			y = newY

		}
		code = append(code, keypad[y][x])
	}

	fmt.Println("Key Code is:", code)
}
