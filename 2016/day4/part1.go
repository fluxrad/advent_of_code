package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

// To sort by value rather than key, we have to create a sortable object type
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value < p[j].Value {
		return true
	} else if p[i].Value > p[j].Value {
		return false
	} else {
		// if counts are equal, sort by letter (ascending)
		return p[i].Key > p[j].Key
	}
}

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	log.WithFields(log.Fields{
		"func": "sortMapByValue",
		"m":    m,
	}).Debug()

	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(p))
	return p
}

func validateRoom(room, checksum string) bool {
	freq := make(map[string]int)
	// get rid of dashes and populate frequency
	re := regexp.MustCompile(`-`)
	room = re.ReplaceAllLiteralString(room, "")

	for _, v := range room {
		freq[string(v)] += 1
	}
	log.WithFields(log.Fields{
		"frequencies": freq,
	}).Debug()

	letterList := sortMapByValue(freq)
	log.Debug(letterList)

	var ck bytes.Buffer
	for i := 0; i < 5; i++ {
		ck.WriteString(letterList[i].Key)
	}

	log.WithFields(log.Fields{
		"calculated_checksum": ck.String(),
		"checksum":            checksum,
		"room":                room,
	}).Debug("validating room")

	if ck.String() == checksum {
		return true
	}
	return false
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	r := regexp.MustCompile(`(.+)-([0-9]+)\[([a-z]+)\]\s*`)
	sectorIDSum := 0
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		roomDetail := r.FindStringSubmatch(scanner.Text())

		log.WithFields(log.Fields{
			"line":     scanner.Text(),
			"room":     roomDetail[1],
			"sectorID": roomDetail[2],
			"checksum": roomDetail[3],
		}).Debug()

		if validateRoom(roomDetail[1], roomDetail[3]) {
			c, err := strconv.Atoi(roomDetail[2])
			if err != nil {
				log.Fatal(err)
			}
			sectorIDSum += c
		}
	}

	fmt.Printf("Sector ID sum of valid rooms is: %d\n", sectorIDSum)
}
