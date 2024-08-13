package utils

import (
	"github.com/gocql/gocql"
)

func StringIDtoUUID(paramName string) (gocql.UUID, error) {
	if paramName == "" {
		return gocql.UUID{}, nil
	}

	parsedUUID, err := gocql.ParseUUID(paramName)
	if err != nil {
		return gocql.UUID{}, err
	}

	return parsedUUID, nil
}
