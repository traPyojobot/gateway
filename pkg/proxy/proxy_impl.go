package proxy

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	GroupReply     = "reply"
	GroupMonologue = "monologue" //TODO 読み込み処理を書く
)

var allMLModel map[string][]MLModel //TODO 読み込み処理を書く

func (m *MLModel) sendRequest(json []byte) ([]byte, error) {
	resp, err := http.Post("http://postman.hijiki51.trap.show/", "application/json", bytes.NewBuffer(json)) //仮配置
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

func CreateMessageByGroup(group string, json []byte) ([]byte, error) {
	// ms := allMLModel[group]
	// m := ms[rand.Intn(len(ms))]//一時的にコメントアウト
	m := MLModel{}
	resp, err := m.sendRequest(json)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
