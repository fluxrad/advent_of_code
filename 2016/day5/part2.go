package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
}

func startsWithZeroes(s string) bool {
	for i := 0; i < 5; i++ {
		if string(s[i]) != "0" {
			return false
		}
	}

	/*
		log.WithFields(log.Fields{
			"hash": s,
		}).Debug("found hash")
	*/

	return true
}

func findPassword(id string) string {
	password := make([]byte, 8, 8)
	found := 0

	// The bug is here
	for i := 0; found < 8; i++ {
		var test bytes.Buffer

		test.WriteString(id)
		test.WriteString(strconv.Itoa(i))

		sum := fmt.Sprintf("%x", md5.Sum(test.Bytes()))

		// This block is gross. I feel shame.
		if startsWithZeroes(sum) {
			// Only fill the password slot if:
			//   - we have found a valid slot
			//   - that slot hasn't already been filled
			if v, err := strconv.Atoi(string(sum[5])); err == nil && v < 8 {
				log.Debugf("Using %s", sum)

				if password[int(v)] == 0 {
					password[int(v)] = sum[6]
					found++
				}
			}
		}
	}

	return string(password)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var doorID string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		doorID = scanner.Text()
	}
	log.WithFields(log.Fields{"doorID": doorID}).Debug()

	password := findPassword(doorID)
	fmt.Printf("Password is: %s\n", password)
}
