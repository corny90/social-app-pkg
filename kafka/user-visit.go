package pkg_kafka

import (
	"time"
)

type UserVisitPayload struct {
	UserId         string    `json:"userId"`
	Timestamp      time.Time `json:"timestamp"`
	RemoteAddress  string    `json:"remoteAddr"`
	City           string    `json:"city"`
	Country        string    `json:"country"`
	Latitude       float32   `json:"lat"`
	Longitude      float32   `json:"log"`
	Device         string    `json:"device"`
	BrowserName    string    `json:"browserName"`
	BrowserVersion string    `json:"browserVersion"`
	OsName         string    `json:"osName"`
	OsVersion      string    `json:"osVersion"`
}
