package reply

type TwitterReply struct {
	Message    string
	User       ReplyUser
	Properties map[string]interface{}
}

func (t *TwitterReply) GetUser() ReplyUser {
	return t.User
}
