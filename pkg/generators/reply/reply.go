package reply

type ReplyUser struct {
	UserName string
	UserID   string
}
type ReplyMessage interface {
	GetUser() ReplyUser
	GetMessage() string
	GetProperties() map[string]interface{}
}
type ReplyService interface {
	GenerateReply(m ReplyMessage) ReplyMessage
}
