package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func trackChanges(cfg Config) {
	var prev ResortStatus
	for {
		log.Debugf("Looking for changes at %s...", cfg.Source.Name)
		data, err := fetchStatus(cfg.Source.Method, cfg.Source.Url, cfg.Source.Args)
		if err != nil {
			log.Errorf("Error fetching resort status: %v", err)
			time.Sleep(cfg.Delay)
			continue
		}
		stat, err := parseStatus(data, cfg.Source)
		if err != nil {
			log.Errorf("Error parsing resort status: %v", err)
			time.Sleep(cfg.Delay)
			continue
		}
		if len(prev.Status) < 1 {
			log.Warnf("First %s status... nothing to compare yet.", cfg.Source.Name) // TK: store prev execution's status in a bolt.db?
			prev = stat
			time.Sleep(cfg.Delay)
			continue
		}
		diff, err := diffStatus(stat, prev)
		if err != nil {
			log.Errorf("Diff err: %v", err)
			time.Sleep(cfg.Delay)
			continue
		}
		prev = stat
		if err := notifyChanges(cfg, diff); err != nil {
			log.Errorf("Notify err: %v", err)
			time.Sleep(cfg.Delay)
			continue
		}
		time.Sleep(cfg.Delay)
	}

}

func diffStatus(cur, prev ResortStatus) (Diff, error) {
	diff := Diff{}
	for name, status := range cur.Status {
		if _, exist := prev.Status[name]; exist {
			if status.Status != prev.Status[name].Status {
				if status.Kind == "trail" {
					diff.Changes = append(diff.Changes, fmt.Sprintf("%s (%s) %s->%s", status.Name, status.Difficulty, prev.Status[name].Status, status.Status))
				} else {
					diff.Changes = append(diff.Changes, fmt.Sprintf("%s %s->%s", status.Name, prev.Status[name].Status, status.Status))
				}
			}
		}
	}
	return diff, nil
}

func fetchStatus(requestMethod, url, args string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 5}
	method := http.MethodGet
	if requestMethod == "POST" {
		method = http.MethodPost
	}
	req, err := http.NewRequest(method, url, strings.NewReader(args))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36 FreshPow/1.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
