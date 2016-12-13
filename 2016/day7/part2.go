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

type aba []string

func supportsSSL(addr string) bool {
	re := regexp.MustCompile("\\[?[a-z]+\\]?")
	substrings := re.FindAllString(addr, -1)

	var abas aba
	var babs aba

	for _, sub := range substrings {
		found, err := findABAs(sub)
		if err != nil {
			// We found a BAB inside brackets. Add it to the list
			babs = append(babs, found...)
		} else if len(found) > 0 {
			abas = append(abas, found...)
		}
	}
	log.WithFields(log.Fields{"abas": abas, "babs": babs}).Debug()

	// Check if the any ABAs have a corresponding BAB
	for _, a := range abas {
		for _, b := range babs {
			log.WithFields(log.Fields{"aba": a, "bab": b}).Debug()
			if a[0] == b[1] && a[1] == b[0] {
				return true
			}
		}
	}
	return false
}

func findABAs(s string) ([]string, error) {
	var abas []string
	var brackets error

	for i, _ := range s {
		end := i + 2
		if end > len(s)-1 {
			return abas, brackets
		}

		if s[i] == s[end] && s[i] != s[i+1] {
			abas = append(abas, s[i:i+3])

			if s[0] == '[' {
				brackets = errors.New("is a BAB")
			}

			log.WithFields(log.Fields{"matching": s[i : i+3], "brackets": brackets}).Debug("Found matching")
		}
	}

	return abas, brackets
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	supportsSSLCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		addr := scanner.Text()
		log.WithFields(log.Fields{"addr": addr}).Debug()

		if supportsSSL(addr) {
			log.Debug("String supports SSL")
			supportsSSLCount++
		}
	}

	fmt.Println("Number of addresses supporting SSL:", supportsSSLCount)
}
