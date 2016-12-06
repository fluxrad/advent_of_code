package part1

import (
	"bufio"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func Run() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	possible := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		possible += 1
		tr := strings.Split(scanner.Text(), " ")

		log.WithFields(log.Fields{
			"sides":    tr,
			"possible": possible,
		}).Debug()
	}

}
