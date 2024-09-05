package message_llm

import (
	"github.com/corny90/social-app-pkg/core-domain/message"
	"github.com/gocql/gocql"
)

// ---------------------------------------------------------
// LLM

type LLMRequest struct {
	ChatHistory        []message.Message      `json:"chat_history"`
	ChatSummary        LLMConversationSummary `json:"chat_summary"`
	ProfileReal        LLMParticipantProfile  `json:"profile_r"`
	ProfileVirtual     LLMParticipantProfile  `json:"profile_v"`
	ParticipantsInfo   []LLMParticipantInfo   `json:"participants_info"`
	FollowUp           bool                   `json:"follow_up"`
	MaxPremiumMessages int                    `json:"max_premium_messages"`
	GapMoreThan8H      bool                   `json:"gap_more_than_8h"`
}

type LLMResponse struct {
	BotResponse         string   `json:"response"`
	EngagementRate      int      `json:"engagement_rate"`
	BlockRecommendation bool     `json:"block_recommendation"`
	AgentPersonality    []string `json:"agent_personality"`
}

type LLMConversationSummary struct {
	ConversationID         gocql.UUID `json:"conversation_id"`
	Summary                string     `json:"summary"`
	AgentIntroducedHerself bool       `json:"agent_introduced_herself"`
	BlockRecommendation    bool       `json:"block_recommendation"`
	IsPremium              bool       `json:"is_premium"`
	Location               string     `json:"location"`
	LocationPlaces         []string   `json:"location_places"`
}

type LLMParticipantProfile struct {
	Language           string   `json:"language"`
	LookingFor         []string `json:"looking_for"`
	SelfDescription    string   `json:"self_description"`
	MyCharacteristics  []string `json:"my_characteristics"`
	FavoriteActivities []string `json:"favorite_activities"`
	WhatTurnsMeOn      []string `json:"what_turns_me_on"`
	Interests          []string `json:"interests"`
	Religion           string   `json:"religion"`
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

type LLMParticipantInfo struct {
	UserID             gocql.UUID `json:"user_id"`
	UserType           string     `json:"user_type"`
	GivenName          string     `json:"user_name"`
	Location           string     `json:"location"`
	Profession         string     `json:"profession"`
	Personality        string     `json:"personality"`
	ImgProfileKeywords []string   `json:"img_profile_keywords"`
}
