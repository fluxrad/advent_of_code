package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

const columns = 8

type MessageColumn struct {
	Letters map[rune]int
}

type MessageColumnList []MessageColumn

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var decryptedMessage string
	mcl := make(MessageColumnList, 8)

	scanner := bufio.NewScanner(input)

	// Count up letter counts by column
	for scanner.Scan() {
		msg := scanner.Text()

		log.WithFields(log.Fields{
			"msg": msg,
		}).Debug()

		for i, v := range msg {
			if mcl[i].Letters == nil {
				mcl[i].Letters = make(map[rune]int)
			}

			mcl[i].Letters[v] += 1

			log.WithFields(log.Fields{
				"index":  i,
				"letter": string(v),
			}).Debug()
		}
	}

	// Find the most frequent letter for each column.
	for _, messageColumn := range mcl {
		bestLetterCount := 0
		var letter rune

		for ltr, count := range messageColumn.Letters {
			if count > bestLetterCount {
				letter = ltr
				bestLetterCount = count
			}
		}

		decryptedMessage = decryptedMessage + string(letter)
	}

	fmt.Println("The secret message is:", decryptedMessage)
}
