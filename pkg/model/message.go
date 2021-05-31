package model

import "time"

type Message struct {
	UserID    string    `json:"user_id,omitempty" query:"id"`
	UserName  string    `json:"user_name,omitempty" query:"name"`
	Text      string    `json:"text,omitempty" query:"text"`
	CreatedAt time.Time `json:"created_at,omitempty" query:"created_at"`
}
