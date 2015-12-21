package main

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

func TestFindMD5(t *testing.T) {
	seed := "abcdef"
	result, err := FindMD5(seed)

	if err != nil {
		t.Error("Got error ", err)
	}

	if result != 609043 {
		t.Error("Expected 609043 got ", result)
	}
}
