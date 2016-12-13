package common

import (
	"flag"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	debug := flag.Bool("d", false, "debug")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}
