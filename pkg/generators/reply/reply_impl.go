package reply

type TwitterReply struct {
	Message    string
	User       ReplyUser
	Properties map[string]interface{}
}

func (t *TwitterReply) GetUser() ReplyUser {
	return t.User
}

func (t *TwitterReply) GetMessage() string {
	return t.Message
}
func (t *TwitterReply) GetProperties() map[string]interface{} {
	return t.Properties
}
