package utils

import (
	"errors"
	"regexp"
	"time"
)

func TimeFormatForDatabase(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date string is empty")
	}

	// Handle or clean double timezone offsets
	re := regexp.MustCompile(`(\+\d{4}) \+\d{4}`)
	dateStr = re.ReplaceAllString(dateStr, "$1")

	// First, try parsing with RFC 3339 Nano which expects the 'T' separator.
	parsedTime, err := time.Parse(time.RFC3339Nano, dateStr)
	if err == nil {
		return parsedTime, nil
	}

	// If RFC 3339 Nano fails, try the custom format
	const customFormat = "2006-01-02 15:04:05.999 -0700 MST"
	parsedTime, err = time.Parse(customFormat, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
