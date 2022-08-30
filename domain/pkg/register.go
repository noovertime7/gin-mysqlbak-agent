package pkg

import (
	"backupAgent/domain/config"
	"backupAgent/domain/pkg/log"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AgentRegister interface {
	Register() (string, error)
	DeRegister() (string, error)
}

type registry struct {
	ServiceName string `json:"service_name"`
	Address     string `json:"address"`
	Content     string `json:"content"`
	TaskNum     int    `json:"task_num"`
	FinishNum   int    `json:"finish_num"`
}

var (
	registerUrl   = config.GetStringConf("register", "registerUrl")
	deregisterUrl = config.GetStringConf("register", "deregisterUrl")
)

func NewAgentRegister(serviceName, address, content string, taskNUm, finishNum int) AgentRegister {
	return &registry{
		ServiceName: serviceName,
		Address:     address,
		Content:     content,
		TaskNum:     taskNUm,
		FinishNum:   finishNum,
	}
}

func (r *registry) Register() (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(r); err != nil {
		return "", err
	}
	log.Logger.Info("register向server注册服务", registerUrl)
	res, err := http.Post(registerUrl, "application/json", buf)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	log.Logger.Infof("register注册服务响应：%s", string(msg))
	return string(msg), err
}

func (r *registry) DeRegister() (string, error) {
	payload := strings.NewReader(fmt.Sprintf("{\"service_name\":\"%s\"}", r.ServiceName))
	req, _ := http.NewRequest("PUT", deregisterUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	log.Logger.Info("register向server注销服务", deregisterUrl)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	msg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	log.Logger.Infof("register注销服务响应：%s", string(msg))
	return string(msg), err
}
