package reply

import (
	standerd "github.com/traPyojobot/gateway/cmd/standard_reply"
	"github.com/traPyojobot/gateway/pkg/model"
)

type ReplyService interface {
	GenerateReply(*model.Message) (*model.Message, error)
}

func Create(m *model.Message) (*model.Message, error) {
	// TODO:サービス選択処理
	res, err := standerd.GenerateReply(m)
	if err != nil {
		return nil, err
	}
	return res, err
}
