package gateway

import "time"

type Message interface {
	GetContent() string
	GetUserID() string
	GetUserName() string
	GetCreatedAt() time.Time
}
