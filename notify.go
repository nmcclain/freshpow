package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func notifyChanges(cfg Config, diff Diff) error {
	if len(diff.Changes) < 1 {
		log.Debugf("No changes at %s...", cfg.Source.Name)
		return nil
	}
	log.Printf("Changes at %s: %+v", cfg.Source.Name, diff)
	message := fmt.Sprintf("Changes at %s:\n", cfg.Source.Name)
	for _, change := range diff.Changes {
		message += fmt.Sprintf("%+v\n", change)
	}
	if len(cfg.SlackWebhook) < 1 || len(cfg.SlackChannel) < 1 {
		return nil
	}

	client := http.Client{Timeout: time.Second * 5}
	pl := SlackPayload{
		Text:      message,
		Channel:   cfg.SlackChannel,
		Username:  "freshpow",
		IconEmoji: ":ski:",
	}
	args, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", cfg.SlackWebhook, bytes.NewBuffer(args))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending Slack message to %s: %s", cfg.SlackChannel, err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if string(body) != "ok" {
		return fmt.Errorf("Error sending Slack message to %s: %s", cfg.SlackChannel, body)
	}
	log.Debugf("Sent Slack message to %s", cfg.SlackChannel)
	return nil
}
