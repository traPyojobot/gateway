package reply

type ReplyMessage interface {
	GetUserName() string
	GetUserID() string
	GetMessage() string
	GetProperties() map[string]interface{}
}
type ReplyService interface {
	GenerateReply(m ReplyMessage) ReplyMessage
}
