package reply

import "github.com/traPyojobot/gateway/pkg/gateway"

type ReplyService interface {
	GenerateReply(m gateway.Message) gateway.Message
}
