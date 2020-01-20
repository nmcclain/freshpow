package main

import (
	"time"

	"github.com/buger/jsonparser"
	log "github.com/sirupsen/logrus"
)

func parsePowderStatus(data []byte, source Source) (ResortStatus, error) {
	report := ResortStatus{
		Source: source,
		When:   time.Now(),
		Status: map[string]Entity{},
	}
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			log.Fatalf("Parse err: %v", err)
		}
		kind, err := jsonparser.GetString(value, "type")
		if err != nil {
			log.Fatalf("Inner parse err for type: %v", err)
		}
		if kind != "lift" && kind != "trail" {
			// log.Printf("skipping = %+v\n", kind)
			return
		}
		entity := Entity{Kind: kind}
		if err != nil {
			log.Fatalf("Inner parse err for id: %v", err)
		}
		entity.Name, err = jsonparser.GetString(value, "properties", "name")
		if err != nil {
			log.Fatalf("Inner parse err for name: %v", err)
		}
		if entity.Kind == "trail" {
			entity.Status, err = jsonparser.GetString(value, "properties", "global_status")
			if err != nil {
				log.Printf("value = %+s\n", value)
				log.Fatalf("Inner parse err for name: %v", err)
			}
			entity.Difficulty, err = jsonparser.GetString(value, "properties", "subtype")
			if err != nil {
				log.Fatalf("Inner parse err for difficulty: %v", err)
			}

			switch entity.Difficulty {
			case "easier":
				entity.Difficulty = "Green"
			case "more_difficult":
				entity.Difficulty = "Blue"
			case "most_difficult":
				entity.Difficulty = "Black"
			case "extremely_difficult":
				entity.Difficulty = "Double Black"
			}

		} else { // assume "lift"
			jsonparser.ArrayEach(value, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				if err != nil {
					log.Fatalf("Inner parse err")
				}
				statusName, err := jsonparser.GetString(value, "status_name")
				if err != nil {
					log.Fatalf("Inner parse err for name: %v", err)
				}
				if statusName == "opening" {
					statusValue, err := jsonparser.GetString(value, "status_value")
					if err != nil {
						log.Fatalf("Inner parse err for name: %v", err)
					}
					entity.Status = statusValue
				}

			}, "status")
		}

		// cleanup
		if entity.Status == "opening" {
			entity.Status = "opened"
		}
		if entity.Status == "true" {
			entity.Status = "opened"
		}
		if entity.Status == "false" {
			entity.Status = "closed"
		}
		report.Status[entity.Name] = entity
	})
	return report, nil
}
