// ひとりごとを生成する機能の本体
package monologue

import (
	monologueService "github.com/traPyojobot/gateway/cmd/monologue"
	"github.com/traPyojobot/gateway/pkg/model"
)

// 生成(といってもモデル抽象化レイヤーに横流しするだけ)
func Create(m *model.Message) (*model.Message, error) {
	res, err := monologueService.GenerateMonologue(m)
	if err != nil {
		return nil, err
	}
	return res, nil
}
