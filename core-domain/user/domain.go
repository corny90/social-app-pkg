package user

import (
	"errors"
	"github.com/corny90/social-app-pkg/core-domain/social"
	"time"

	"github.com/gocql/gocql"
)

var (
	ErrInfoNotFound       = errors.New("user info not found")
	ErrorCreatingInfo     = errors.New("error creating user info")
	ErrorInfoAlreadyExist = errors.New("user info already exists")
	ErrorUpdatingInfo     = errors.New("error updating user info")

	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrorRegisteringUser  = errors.New("error registering user")
	ErrorUserAlreadyExist = errors.New("user already exists")
	ErrorUpdatingUser     = errors.New("error updating user")
)

// User struct ---------------------------------------------------------------------------------------------------------
type User struct {
	ID           gocql.UUID `json:"id,omitempty"`
	Username     string     `json:"username" validate:"required,min=4,max=20"`
	Email        string     `json:"email" validate:"required,email"`
	Password     string     `json:"password" validate:"required,min=4,max=20"`
	Jwt          string     `json:"jwt"`
	Phone        string     `json:"phone"`
	Birthday     string     `json:"birthday"`
	Avatar       string     `json:"avatar"`
	DateJoined   time.Time  `json:"date_joined"`
	DateEdited   time.Time  `json:"date_edited"`
	AccountType  string     `json:"account_type"`
	LocationJSON string     `json:"location"` // JSON strings for database storage
	AccountJSON  string     `json:"account"`  // JSON strings for database storage
	Location     Location   `json:"-"`        // structs for easier handling in Go
	Account      Account    `json:"-"`        // structs for easier handling in Go

	Info *Info `json:"info,omitempty"`
	//Media []*proto.Media `json:"media,omitempty"`
	Posts []social.Post `json:"posts,omitempty"`
}

// Location struct for user's location information
type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Geo     string `json:"geo"`
}

// Account struct for user's account information
type Account struct {
	SiteDomainPresence string   `json:"site_domain_presence"`
	PaymentProcess     string   `json:"payment_process"`
	Membership         string   `json:"membership"`
	Active             string   `json:"active"`
	Matches            []string `json:"matches"`
	Winks              []string `json:"winks"`
	MyWinks            []string `json:"my_winks"`
	Visitors           []string `json:"visitors"`
}

// Users struct --------------------------------------------------------------------------------------------------------
type Users struct {
	Users []User `json:"users"`
}

// Login struct --------------------------------------------------------------------------------------------------------
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=20"`
	Jwt      string `json:"jwt"`
}

// Update struct -------------------------------------------------------------------------------------------------------
type Update struct {
	Username     string   `json:"username" validate:"omitempty,min=4,max=20"`
	Email        string   `json:"email" validate:"omitempty,email"`
	Password     string   `json:"password"`
	Jwt          string   `json:"jwt"`
	Phone        string   `json:"phone"`
	Birthday     string   `json:"birthday"`
	Avatar       string   `json:"avatar"`
	DateEdited   string   `json:"date_edited"`
	AccountType  string   `json:"account_type"`
	LocationJSON string   `json:"location"` // JSON strings for database storage
	AccountJSON  string   `json:"account"`  // JSON strings for database storage
	Location     Location `json:"-"`        // structs for easier handling in Go
	Account      Account  `json:"-"`        // structs for easier handling in Go
}

type Info struct {
	UserID           gocql.UUID   `json:"user_id"`
	About            AboutInfo    `json:"about"`
	Social           SocialInfo   `json:"social"`
	Physical         PhysicalInfo `json:"physical"`
	Others           OtherInfo    `json:"others"`
	Behavior         BehaviorInfo `json:"behavior"`
	PreferedKeywords []string     `json:"prefered_keywords"`
}

type AboutInfo struct {
	LookingFor         []string `json:"looking_for"`
	SelfDescription    string   `json:"self_description"`
	MyCharacteristics  []string `json:"my_characteristics"`
	FavoriteActivities []string `json:"favorite_activities"`
	WhatTurnsMeOn      []string `json:"what_turns_me_on"`
	Interests          []string `json:"interests"`
}

type SocialInfo struct {
	Religion      string `json:"religion"`
	Profession    string `json:"profession"`
	Education     string `json:"education"`
	MaritalStatus string `json:"marital_status"`
	Children      string `json:"children"`
}

type PhysicalInfo struct {
	Gender    string `json:"gender"`
	Height    string `json:"height"`
	Race      string `json:"race"`
	HairColor string `json:"hair_color"`
	EyeColor  string `json:"eye_color"`
	BodyType  string `json:"body_type"`
	Piercings string `json:"piercings"`
	Tattoos   string `json:"tattoos"`
}

type OtherInfo struct {
	Smoking  string `json:"smoking"`
	Drinking string `json:"drinking"`
	Position string `json:"position"`
	Exercise string `json:"exercise"`
	Pets     string `json:"pets"`
	SexSpot  string `json:"sex_spot"`
}

type BehaviorInfo struct {
	Interests            []string `json:"interests"`
	Trends               []string `json:"trends"`
	Engagement           []string `json:"engagement"`
	AvrDailyTimeActive   int      `json:"avr_daily_time_active"`
	AvrDailyClicks       int      `json:"avr_daily_clicks"`
	AvrDailySentMessages int      `json:"avr_daily_sent_messages"`
	AvrDailySpends       int      `json:"avr_daily_spends"`
}
