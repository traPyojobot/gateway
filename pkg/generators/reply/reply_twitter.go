package reply

type TwitterReply struct {
	Message    string
	UserID     string
	UserName   string
	Properties map[string]interface{}
}

func (t *TwitterReply) GetUserName() string {
	return t.UserName
}

func (t *TwitterReply) GetUserID() string {
	return t.UserID
}

func (t *TwitterReply) GetMessage() string {
	return t.Message
}
func (t *TwitterReply) SetMessage(m string) {
	t.Message = m
	return
}
func (t *TwitterReply) GetProperties() map[string]interface{} {
	return t.Properties
}
