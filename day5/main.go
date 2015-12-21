package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	niceWordCount := 0
	for scanner.Scan() {
		if WordIsNice(scanner.Text()) {
			niceWordCount++
		}
	}

	fmt.Printf("%d words are nice", niceWordCount)
}

func WordIsNice(str string) bool {
	vowels := 0
	double := false

	blacklist := []string{"ab", "cd", "pq", "xy"}
	for _, b := range blacklist {
		match, err := regexp.MatchString(b, str)
		if err != nil {
			log.Fatal(err)
		}

		if match == true {
			return false
		}
	}

	var prev string
	for _, c := range strings.Split(str, "") {
		if strings.ContainsAny(c, "aeiou") {
			vowels++
		}

		if prev == c {
			double = true
		}
		prev = c
	}

	if vowels >= 3 && double == true {
		return true
	}

	return false
}
