package model

import "time"

type Message struct {
	UserID    string
	UserName  string
	Text      string
	CreatedAt time.Time
}
