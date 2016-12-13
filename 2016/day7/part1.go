package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	log "github.com/Sirupsen/logrus"
	_ "github.com/fluxrad/advent_of_code/common"
)

func supportsTLS(addr string) bool {
	re := regexp.MustCompile("\\[?[a-z]+\\]?")
	subs := re.FindAllString(addr, -1)

	abba := false

	for _, sub := range subs {
		found, err := hasABBA(sub)
		if err != nil {
			return false
		}

		if found {
			abba = true
		}
	}

	return abba
}

func hasABBA(s string) (bool, error) {
	for i, _ := range s {
		end := i + 3
		if end > len(s)-1 {
			return false, nil
		}

		if s[i] == s[end] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			if s[0] == '[' {
				return true, errors.New("ABBA in brackets. No bueno.")
			}

			log.WithFields(log.Fields{"matching": s[i : i+4]}).Debug()
			return true, nil
		}
	}

	return false, nil
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	supportsTLSCount := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		addr := scanner.Text()

		log.WithFields(log.Fields{"addr": addr}).Debug()

		if supportsTLS(addr) {
			log.Debug("String supports TLS")
			supportsTLSCount++
		}
	}

	fmt.Println("Number of addresses supporting TLS:", supportsTLSCount)
}
