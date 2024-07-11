package pkg_kafka

type UserLoginPayload struct {
	UserId     string `json:"userId"`
	Timestamp  string `json:"timestamp"`
	RemoteAddr string `json:"remoteAddr"`
}
