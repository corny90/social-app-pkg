package initialcontent

import "github.com/gocql/gocql"

type InitialContent struct {
	ID       gocql.UUID `json:"id"`
	Kind     string     `json:"kind"`
	Location string     `json:"location"`
}
