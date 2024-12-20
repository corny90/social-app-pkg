package media

import (
	"time"

	"github.com/gocql/gocql"
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
	// Metadata  string
}

type BackgroundMedia struct {
	UserID    gocql.UUID
	CreatedAt time.Time
	Filename  string
}
