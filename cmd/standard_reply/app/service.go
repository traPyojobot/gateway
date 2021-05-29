package app

import (
	"encoding/json"

	"github.com/traPyojobot/gateway/pkg/generators/reply"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

type standerdReply struct {
	MlService string
}
type replyRequest struct {
	Message  string `json:"message"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func (s *standerdReply) GenerateReply(m reply.ReplyMessage) (reply.ReplyMessage, error) { //ひとまずプロパティは切り捨てる
	req := &replyRequest{
		Message:  m.GetMessage(),
		UserID:   m.GetUserID(),
		UserName: m.GetUserName(),
	}
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	b, err := proxy.CreateMessageByGroup(proxy.GroupReply, j)
	if err != nil {
		return nil, err
	}
	var message string
	err = json.Unmarshal(b, &message)
	if err != nil {
		return nil, err
	}
	m.SetMessage(message)
	return m, nil
}
