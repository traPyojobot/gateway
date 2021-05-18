package reply

type ReplyMessage interface {
	GetUser()
	GetMessage()
	GetProperties()
}

func GenerateReply() {}
func sendGenerator() {}
