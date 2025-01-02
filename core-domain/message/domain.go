package message

import (
	"time"

	"github.com/gocql/gocql"
)

// ---------------------------------------------------------
// CLIENT REQUESTS MODELS

type UserMessageRequest struct {
	RecipientUUID gocql.UUID `json:"recipient"`
	Content       string     `json:"content"`
}

// ---------------------------------------------------------
// USER PROFILE MODEL

type UserProfile struct {
	Username           string   `json:"username"`
	Language           string   `json:"language"`
	Location           string   `json:"location"`
	LocationDate       string   `json:"location_date"`
	LookingFor         []string `json:"looking_for"`
	SelfDescription    string   `json:"self_description"`
	MyCharacteristics  []string `json:"my_characteristics"`
	FavoriteActivities []string `json:"favorite_activities"`
	WhatTurnsMeOn      []string `json:"what_turns_me_on"`
	Interests          []string `json:"interests"`
	Religion           string   `json:"religion"`
	Profession         string   `json:"profession"`
	Education          string   `json:"education"`
	MaritalStatus      string   `json:"marital_status"`
	Children           string   `json:"children"`
	Gender             string   `json:"gender"`
	Height             string   `json:"height"`
	Race               string   `json:"race"`
	HairColor          string   `json:"hair_color"`
	EyeColor           string   `json:"eye_color"`
	BodyType           string   `json:"body_type"`
	Piercings          string   `json:"piercings"`
	Tattoos            string   `json:"tattoos"`
	Smoking            string   `json:"smoking"`
	Drinking           string   `json:"drinking"`
	Position           string   `json:"position"`
	Exercise           string   `json:"exercise"`
	Pets               string   `json:"pets"`
	SexSpot            string   `json:"sex_spot"`
}

// Response for frontend
type ConversationRender struct {
	ConversationID gocql.UUID    `json:"conversation_id"`
	Participants   []Participant `json:"participants"`
	CreatedAt      time.Time     `json:"created_at"`
	EditedAt       time.Time     `json:"edited_at"`
	Metadata       string        `json:"metadata"`
	State          string        `json:"state"`
	TotalUnread    int           `json:"total_unread"`
}
type Participant struct {
	UserID    gocql.UUID `json:"user_id"`
	UserType  string     `json:"user_type"` // r or v
	UserRole  string     `json:"user_role"` // advertise or business
	Username  string     `json:"username"`
	AvatarUrl string     `json:"avatar_url"`
}

type ConversationBase struct {
	ConversationID gocql.UUID   `json:"conversation_id"`
	Participants   []gocql.UUID `json:"participants"`
	CreatedAt      time.Time    `json:"created_at"`
	EditedAt       time.Time    `json:"edited_at"`
	Metadata       string       `json:"metadata"`
	State          string       `json:"state"`
}

type ConversationByEdited struct {
	Client         string     `json:"client"`
	EditedAt       time.Time  `json:"edited_at"`
	ConversationID gocql.UUID `json:"conversation_id"`
}

type ConversationBySender struct {
	Sender         gocql.UUID `json:"sender"`
	EditedAt       time.Time  `json:"edited_at"`
	ConversationID gocql.UUID `json:"conversation_id"`
}

type ConversationByParticipants struct {
	Participants   []gocql.UUID `json:"participants"`
	ConversationID gocql.UUID   `json:"conversation_id"`
}

type Message struct {
	MessageID      gocql.UUID `json:"message_id"`
	ConversationID gocql.UUID `json:"conversation_id"`
	Sender         gocql.UUID `json:"sender_id"`
	SenderType     string     `json:"sender_type"`
	Content        string     `json:"content"`
	State          string     `json:"state"`
	CreatedAt      time.Time  `json:"created_at"`
	Metadata       string     `json:"metadata"`
}

type MsgMetadata struct {
	Attachments []string `json:"attachments"`
	Reactions   []string `json:"reactions"`
}

// ---------------------------------------------------------
// FEEDBACK MODEL

type Feedback struct {
	MessageID       gocql.UUID `json:"message_id"`
	ConversationID  gocql.UUID `json:"conversation_id"`
	CreatedAt       time.Time  `json:"created_at"`
	SenderID        gocql.UUID `json:"sender_id"`
	SenderType      string     `json:"sender_type"`
	SenderMessage   string     `json:"sender_message"`
	BotMessage      string     `json:"bot_message"`
	FeedbackMessage string     `json:"feedback_message"`
	FeedbackNote    string     `json:"feedback_note"`
}

// ---------------------------------------------------------
// WEBSOCKET MODEL

type WsPayload struct {
	Action  string             `json:"action"`
	Payload UserMessageRequest `json:"payload"`
}

// ---------------------------------------------------------
// KAFKA MODELS

type KafkaLLMRequestPayload struct {
	UserRID        string     `json:"user_r_id"`
	UserVID        string     `json:"user_v_id"`
	UsernameV      string     `json:"username_v"`
	Location       string     `json:"location"`
	Places         []string   `json:"places"`
	Day            string     `json:"day"`
	DayMonth       string     `json:"day_month"`
	ChatHistory    []Message  `json:"chat_history"`
	ConversationID gocql.UUID `json:"conversation_id"`
}

type KafkaLLMResponsePayload struct {
	Action string               `json:"action"`
	Data   KafkaLLMResponseData `json:"data"`
}

type KafkaLLMResponseData struct {
	UserRId        interface{} `json:"user_r_id"`
	UserVId        interface{} `json:"user_v_id"`
	UserVMsg       interface{} `json:"user_v_msg"`
	ConversationID gocql.UUID  `json:"conversation_id"`
}
