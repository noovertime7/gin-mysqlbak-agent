package template

import (
	"backupAgent/domain/pkg/ding"
	"backupAgent/domain/pkg/dingproxy"
	"backupAgent/domain/pkg/log"
	"bytes"
	"errors"
	"html/template"
	"os"
	"strings"
)

type DingSender interface {
	ParseMessage() (string, error)
	SetTemplateFilePath(path string)
	SendByProxy(msg string, proxyAddr string) error
	SendBySelf(msg string) error
}

var _ DingSender = &Ding{}

func NewDingSender(sa *DingSA, info *SendInfo, title string) *Ding {
	return &Ding{
		Title:    title,
		SendInfo: info,
		DingSA:   sa,
	}
}

type SendInfo struct {
	Host             string
	ServiceName      string
	DBName           string
	BakStatus        string
	OssStatus        string
	EncryptionStatus string
	BakTime          string
	FileName         string
}

type Ding struct {
	Title    string
	FilePath string
	*SendInfo
	*DingSA
}

type DingSA struct {
	DingAccessToken string
	DingSecret      string
}

func (d *Ding) SetTemplateFilePath(p string) {
	d.FilePath = p
}

func (d *Ding) ParseMessage() (string, error) {
	if d.FilePath == "" {
		tempPath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		d.FilePath = tempPath + "/domain/template/ding.md"
	}
	tmpl, err := template.ParseFiles(d.FilePath)
	if err != nil {
		return "", err
	}
	b := new(bytes.Buffer)
	if err = tmpl.Execute(b, d); err != nil {
		return "", err
	}
	return b.String(), nil
}

func (d *Ding) SendByProxy(msg string, proxyAddr string) error {
	if !checkAddr(proxyAddr) {
		return errors.New("proxyAddr not legal")
	}
	url := "http://" + proxyAddr
	sender := dingproxy.NewDingSender(d.DingAccessToken, d.DingSecret, msg, d.Title, url)
	data, err := sender.SendMarkdown()
	if err != nil {
		return err
	}
	log.Logger.Info("发送结果:", data)
	return nil
}

func checkAddr(addr string) bool {
	return len(strings.Split(addr, ":")) > 0
}

func (d *Ding) SendBySelf(msg string) error {
	markContent := map[string]string{
		"title": "备份状态",
		"text":  msg,
	}
	webhook := ding.Webhook{AtAll: true, Secret: d.DingSecret, AccessToken: d.DingAccessToken}
	if err := webhook.SendMarkDown(markContent); err != nil {
		return err
	}
	return nil
}
