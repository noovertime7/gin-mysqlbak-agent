package dingproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type dingSendMessage struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
	Message      string `json:"message"`
	Title        string `json:"title"`
	Url          string `json:"url"`
}

func NewDingSender(token, secret, message, title, url string) *dingSendMessage {
	return &dingSendMessage{
		AccessToken:  token,
		AccessSecret: secret,
		Title:        title,
		Message:      message,
		Url:          url,
	}
}

func (d *dingSendMessage) SendMessage() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return "", err
	}
	res, err := http.Post(d.Url+"/ding/sendmsg", "application/json", buf)
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
	fmt.Println(d)
	res, err := http.Post(d.Url+"/ding/sendmd", "application/json", buf)
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
