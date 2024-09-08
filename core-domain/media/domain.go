package media

import (
	"github.com/gocql/gocql"
	"time"
)

type PostMedia struct {
	UserType      string
	UserID        gocql.UUID
	PostID        gocql.UUID
	CreatedAt     time.Time
	MediaID       gocql.UUID
	MediaType     string
	MediaFilename string
	CoverFilename string
	Metadata      string
}

type AvatarMedia struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Filename  string
	Metadata  string
}
