package app

import (
	"encoding/json"

	"github.com/traPyojobot/gateway/pkg/model"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

type standerdReply struct {
	MlService string
}

func (s *standerdReply) GenerateReply(m model.Message) (model.Message, error) { //ひとまずプロパティは切り捨てる
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
