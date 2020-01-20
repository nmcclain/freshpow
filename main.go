package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	Version      = "1.0.0"
	minDelayTime = 30 * time.Second
)

var usage = `freshpow: trail and lift status alerts
Usage:
  freshpow [options] <resort>
  freshpow --list
  freshpow --help
  freshpow --version

Options:
  -l, --list                           List supported resorts.
  -d, --delay=<delay>                  Delay between requests [default: 5m].
  --debug                              Display debugging messages.
  -h, --help                           Show this screen.
  --version                            Show version.
`

// TK: add weather & conditions
func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}
	trackChanges(cfg)
}
