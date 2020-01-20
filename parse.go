package main

import (
	"fmt"
)

func parseStatus(data []byte, source Source) (ResortStatus, error) {
	switch source.Provider {
	case "powder":
		return parsePowderStatus(data, source)
	case "alterra":
		return parseAlterraStatus(data, source)
	case "snowcom":
		return parseSnowComStatus(data, source)
	}
	return ResortStatus{}, fmt.Errorf("Unsupported provider")
}
