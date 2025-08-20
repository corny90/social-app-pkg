package cronjob

import (
	"time"

	"github.com/gocql/gocql"
)

type PushNotificationRequest struct {
	UserID  []string `json:"user_id"`
	Message string   `json:"message"`
	Title   string   `json:"title"`
	Link    string   `json:"link"`
}

type PushNotificationResponse struct {
	Id string `json:"id"`
}

type Job struct {
	ID            gocql.UUID  `json:"id"`
	Type          string      `json:"type"`
	Payload       interface{} `json:"payload"`
	Status        string      `json:"status"`
	ExecutionTime time.Time   `json:"execution_time"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	FailedReason  string      `json:"failed_reason"`
}

type NewsResponse struct {
	Id                  gocql.UUID `json:"id"`
	Title               string     `json:"title"`
	Summary             string     `json:"summary"`
	ArticleText         string     `json:"article_text"`
	ImageUrl            string     `json:"image_url"`
	Keywords            []string   `json:"keywords"`
	DefaultLikeCount    int        `json:"default_like_count"`
	DefaultCommentCount int        `json:"default_comment_count"`
	RealLikeCount       int        `json:"real_like_count"`
	WritingStyle        string     `json:"writing_style"`
	CreatedAt           string     `json:"created_at"`
}
