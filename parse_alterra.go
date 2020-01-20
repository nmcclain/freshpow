package main

import (
	"time"

	"github.com/buger/jsonparser"
	log "github.com/sirupsen/logrus"
)

func parseAlterraStatus(data []byte, source Source) (ResortStatus, error) {
	report := ResortStatus{
		Source: source,
		When:   time.Now(),
		Status: map[string]Entity{},
	}
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Fatalf("Parse err: %v", err)
		}
		jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if err != nil {
				log.Fatalf("Parse err: %v", err)
			}
			entity := Entity{Kind: "trail"}
			entity.Name, err = jsonparser.GetString(value, "Name")
			if err != nil {
				log.Fatalf("Inner parse err for type: %v", err)
			}
			entity.Status, err = jsonparser.GetString(value, "Status")
			if err != nil {
				log.Fatalf("Inner parse err for status: %v", err)
			}
			entity.Difficulty, err = jsonparser.GetString(value, "Difficulty")
			if err != nil {
				log.Fatalf("Inner parse err for difficulty: %v", err)
			}
			report.Status[entity.Name] = entity
		}, "Trails")

		jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if err != nil {
				log.Fatalf("Parse err: %v", err)
			}
			entity := Entity{Kind: "lift"}
			entity.Name, err = jsonparser.GetString(value, "Name")
			if err != nil {
				log.Fatalf("Inner parse err for type: %v", err)
			}
			entity.Status, err = jsonparser.GetString(value, "Status")
			if err != nil {
				log.Fatalf("Inner parse err for status: %v", err)
			}
			report.Status[entity.Name] = entity
		}, "Lifts")

	}, "MountainAreas")
	return report, nil
}
