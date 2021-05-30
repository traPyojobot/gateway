package standerd

import (
	"encoding/json"
	"time"

	"github.com/traPyojobot/gateway/pkg/model"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

type replyRequest struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Text     string `json:"text"`
}
type replyResponce struct { //TODO ちゃんと決める
	Text string `json:"message"`
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
	text := replyResponce{}
	err = json.Unmarshal(b, &text)
	if err != nil {
		return nil, err
	}
	res := &model.Message{
		Text:      text.Text,
		CreatedAt: time.Now(),
	}
	return res, nil
}
