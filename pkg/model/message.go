package model

import "time"

type Message interface {
	// メッセージの内容を取得します
	GetContent() string
	// メッセージ作成者のユーザーIDを取得します
	GetUserID() string
	// メッセージ作成者のユーザーネームを取得します
	GetUserName() string
	// メッセージの背区政日時を取得します
	GetCreatedAt() time.Time
}
