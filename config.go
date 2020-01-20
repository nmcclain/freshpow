package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/docopt/docopt-go"
)

type Config struct {
	Delay        time.Duration
	Source       Source
	SlackWebhook string
	SlackChannel string
	Debug        bool
}

func getConfig() (Config, error) {
	c := Config{}
	args, err := docopt.Parse(usage, nil, true, Version, false)
	if err != nil {
		return c, err
	}
	if args["--list"].(bool) {
		names := []string{}
		for name, _ := range getSources() {
			names = append(names, name)
		}
		fmt.Printf("Supported resort names: %s\n", strings.Join(names, ", "))
		os.Exit(0)
	}
	c.Debug = args["--debug"].(bool)
	if err != nil {
		return c, err
	}
	if c.Debug {
		log.SetLevel(log.DebugLevel)
	}
	c.Source, err = getSource(args["<resort>"].(string))
	if err != nil {
		return c, err
	}
	delaySpec := args["--delay"].(string)
	c.Delay, err = time.ParseDuration(delaySpec)
	if err != nil {
		return c, fmt.Errorf("Invalid --delay: %s", err)
	}
	if c.Delay < minDelayTime {
		log.Printf("Note: using minimum supported %s delay.", minDelayTime)
		c.Delay = minDelayTime
	}
	c.SlackWebhook = os.Getenv("SLACK_WEBHOOK")
	c.SlackChannel = os.Getenv("SLACK_CHANNEL")
	if len(c.SlackWebhook) < 1 || len(c.SlackChannel) < 1 {
		log.Debugf("Slack notifications disabled: SLACK_WEBHOOK and/or SLACK_CHANNEL environment variables not set")
	}
	return c, nil
}
