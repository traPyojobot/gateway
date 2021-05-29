package standerd

import (
	"encoding/json"

	"github.com/traPyojobot/gateway/pkg/model"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

type replyRequest struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Text     string `json:"text"`
}

func GenerateReply(m *model.Message) (*model.Message, error) { //ひとまずプロパティは切り捨てる
	req := &replyRequest{
		UserID:   m.UserID,
		UserName: m.UserName,
		Text:     m.Text,
	}
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	b, err := proxy.CreateMessageByGroup(proxy.GroupReply, j)
	if err != nil {
		return nil, err
	}
	var text string
	err = json.Unmarshal(b, &text)
	if err != nil {
		return nil, err
	}
	res := &model.Message{
		Text: text,
	}
	return res, nil
}
