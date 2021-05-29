package reply

import "github.com/traPyojobot/gateway/pkg/model"

type ReplyService interface {
	GenerateReply(m model.Message) model.Message
}
