package main

import (
	"time"
)

type Source struct {
	Name     string
	Method   string
	Url      string
	Args     string
	Provider string
}

type ResortStatus struct {
	Source Source
	Status map[string]Entity
	When   time.Time
}

type Entity struct {
	Kind       string
	Name       string
	Status     string
	Difficulty string
}

type Diff struct {
	Changes []string
}

type SlackPayload struct {
	Text      string `json:"text"`
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
}
