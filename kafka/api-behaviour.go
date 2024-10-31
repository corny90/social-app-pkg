package pkg_kafka

import (
	"encoding/json"
)

type Payload struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// API-BEHAVIOUR - TRACKER

type TrackerPayloadRequest struct {
	UserID       string   `json:"userId"`
	PostID       string   `json:"postId"`
	TrackType    string   `json:"trackType"`
	ViewDuration int      `json:"viewDuration"`
	Keywords     []string `json:"keywords"`
	Clicks       Clicks   `json:"clicks"`
}

type ClickTrackerPayloadRequest struct {
	UserID      string `json:"userId"`
	PostID      string `json:"postId"`
	ClickAction string `json:"clickAction"`

	// AUX optional data based on ClickAction
	PeerUserID string `json:"peerUserId,omitempty"`
	CommentID  string `json:"commentId,omitempty"`
}

type Clicks struct {
	PostUser bool `json:"post_user"`
	Views    bool `json:"views"`
	Likes    bool `json:"likes"`
	Comments bool `json:"comments"`
}

type EventRequest struct {
	Name   string          `json:"event"`
	UserID string          `json:"userId"`
	Data   json.RawMessage `json:"data"`
}

type TrackerPayloadResponse struct {
	UserID string `json:"userId"`
	PostID string `json:"postId"`
	Status string `json:"status"`
}

// API-BEHAVIOUR - BEHAVIOUR

type BehaviourPayloadRequest struct {
	UserID string `json:"userId"`
}

type BehaviourPayloadResponse struct {
	UserID string          `json:"userId"`
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}
