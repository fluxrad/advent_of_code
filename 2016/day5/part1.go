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
	log.SetLevel(log.DebugLevel)
}

func startsWithZeroes(s string) bool {
	for i := 0; i < 5; i++ {
		if string(s[i]) != "0" {
			return false
		}
	}

	log.WithFields(log.Fields{
		"hash": s,
	}).Debug("found hash")

	return true
}

func findPassword(id string) string {
	var password bytes.Buffer

	// The bug is here
	for i := 0; password.Len() < 8; i++ {
		var test bytes.Buffer

		test.WriteString(id)
		test.WriteString(strconv.Itoa(i))

		sum := fmt.Sprintf("%x", md5.Sum(test.Bytes()))

		if startsWithZeroes(sum) {
			password.WriteByte(sum[5])
			log.Debug(fmt.Sprintf("%s", test.String()))
		}
	}

	return password.String()
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
