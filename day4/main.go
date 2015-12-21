package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	data, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatalf("Couldn't read file %s\n", err)
	}

	seed := strings.Trim(string(data), "\n")

	result, err := FindMD5(seed)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The required nubmer for seed %s is %d\n", seed, result)
}

// FindMD5 takes a seed string and finds an int where the resulting MD5 hash
// starts with five zeroes. For example "abcdef" => 609043.
func FindMD5(seed string) (int, error) {
	// Stop after 100 million tries
	for i := 0; i <= 10000000; i++ {
		test := seed + strconv.Itoa(i)
		hash := md5.Sum([]byte(test))

		hashString := fmt.Sprintf("%x", hash)
		log.Debugf("test is: %s, hash is: %x, first five are: %s", test, hash, hashString[:5])

		if hashString[:5] == "00000" {
			return i, nil
		}
	}

	return 0, errors.New("could not find number")
}
