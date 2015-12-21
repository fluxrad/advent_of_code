package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.WarnLevel)

	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Couldn't read file %s\n", err)
	}

	seed := strings.Trim(string(data), "\n")
	for _, zeroes := range []string{"00000", "000000"} {
		result, err := FindMD5(seed, zeroes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Found int: %d, for seed %s with %s padding\n", result, seed, zeroes)
	}
}

// FindMD5 takes a seed string and finds an int where the resulting MD5 hash
// starts with six zeroes. For example "abcdef" => 609043.
func FindMD5(seed string, prefix string) (int, error) {
	// Stop after 100 million tries
	for i := 0; i <= 100000000; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", seed, i)))
		hashString := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(hashString, prefix) {
			return i, nil
		}
	}

	return 0, errors.New("could not find number")
}
