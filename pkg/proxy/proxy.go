package proxy

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/traPyojobot/gateway/pkg/loader"
	"github.com/traPyojobot/gateway/pkg/model"
)

const (
	GroupReply     = "reply"
	GroupMonologue = "monologue" //TODO 読み込み処理を書く
)

var mlModel map[string][]model.MLModel //TODO 読み込み処理を書く

func LoadModel() error {
	_model, err := loader.LoadMLModel()
	mlModel = _model
	if err != nil {
		return err
	}
	return nil
}

func CreateMessageByGroup(group string, json []byte) ([]byte, error) {
	ms := mlModel[group]
	m := ms[rand.Intn(len(ms))]
	resp, err := sendRequest(m, json)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func sendRequest(m model.MLModel, json []byte) ([]byte, error) {
	resp, err := http.Post(m.Url.String(), "application/json", bytes.NewBuffer(json)) //仮配置
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
