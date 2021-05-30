// ひとりごと生成モデルの抽象化レイヤー
package monologue

import (
	"encoding/json"
	"time"

	"github.com/traPyojobot/gateway/pkg/model"
	"github.com/traPyojobot/gateway/pkg/proxy"
)

// TODO: データをちゃんと決める
// 生成サービスにぶん投げるデータ(モデル共通)
type monologueRequest struct{}

// 生成サービスから戻ってくるデータ(モデル共通)
type monologueResponse struct {
	Text string `json:"text"`
}

func GenerateMonologue(m *model.Message) (*model.Message, error) {
	req := &monologueRequest{}
	reqJson, err := json.Marshal(req)

	respJson, err := proxy.CreateMessageByGroup(proxy.GroupMonologue, reqJson)
	if err != nil {
		return nil, err
	}

	resp := &monologueResponse{}
	err = json.Unmarshal(respJson, resp)
	if err != nil {
		return nil, err
	}

	message := &model.Message{
		Text:      resp.Text,
		CreatedAt: time.Now(),
	}
	return message, nil
}
