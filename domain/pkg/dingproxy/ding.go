package dingproxy

import (
	"backupAgent/domain/config"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type dingSendMessage struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}

var url = "http://" + config.GetStringConf("dingProxyAgent", "addr")

func NewDingSender(token, secret, message string) *dingSendMessage {
	return &dingSendMessage{
		AccessToken:  token,
		AccessSecret: secret,
		Message:      message + "\n- 备注:" + config.GetStringConf("base", "content"),
		Title:        config.GetStringConf("dingProxyAgent", "title"),
	}
}

func (d *dingSendMessage) SendMessage() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return "", err
	}
	res, err := http.Post(url+"/ding/sendmsg", "application/json", buf)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	dingmsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(dingmsg), err
}

func (d *dingSendMessage) SendMarkdown() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return "", err
	}
	res, err := http.Post(url+"/ding/sendmd", "application/json", buf)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	dingmsg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(dingmsg), err
}
