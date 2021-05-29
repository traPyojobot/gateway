package proxy

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
)

const (
	GroupReply     = "reply"
	GroupMonologue = "monologue" //TODO 読み込み処理を書く
)

var allMLModel map[string][]mlModel //TODO 読み込み処理を書く

func (m *mlModel) sendRequest(json []byte) ([]byte, error) {
	resp, err := http.Post(m.Url.String(), "application/json", bytes.NewBuffer(json))
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
	ms := allMLModel[group]
	m := ms[rand.Intn(len(ms))]
	resp, err := m.sendRequest(json)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
