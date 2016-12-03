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
	log.SetLevel(log.InfoLevel)
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
	niceWordCount2 := 0
	for scanner.Scan() {
		if WordIsNice(scanner.Text()) {
			niceWordCount++
		}

		if WordIsNice2(scanner.Text()) {
			niceWordCount2++
		}
	}

	fmt.Printf("%d words are nice in part one\n", niceWordCount)
	fmt.Printf("%d words are nice in part two\n", niceWordCount2)
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

func WordIsNice2(str string) bool {
	log.Debugf("word is %s", str)
	repeats := false
	doubles := false

	for i, c := range strings.Split(str, "") {
		if i < len(str)-2 && c == string(str[i+2]) {
			repeats = true
		}

		if i >= len(str)-3 {
			continue
		}

		search := c + string(str[i+1])
		matched, err := regexp.MatchString(search, string(str[i+2:]))
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("search string is:%s, match result is: %t", search, matched)

		if matched {
			doubles = true
		}

	}

	if repeats && doubles {
		return true
	}
	return false
}
