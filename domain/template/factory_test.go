package template

import (
	"fmt"
	"testing"
)

var F = NewTemplateFactory("test")

var sa = &DingSA{
	DingAccessToken: "77f579efbefeefc316b55d3caea1ba1963db2f1319aa7520cbfd9626de073fdc",
	DingSecret:      "SEC72586e3f7ff6db4b2ad24eac905f308a9ddb0b1b9809af31e5623a14abb424b2",
}

var info = &SendInfo{
	Host:             "test",
	ServiceName:      "test",
	DBName:           "test",
	BakStatus:        "test",
	OssStatus:        "test",
	EncryptionStatus: "test",
	BakTime:          "test",
	FileName:         "test",
}

func TestDing_SendBySelf(t *testing.T) {
	ding := F.Ding(sa, info, func(factory *shardTemplateFactory) *shardTemplateFactory {
		factory.Title = "OpTest"
		return factory
	})
	ding.SetTemplateFilePath("G:\\backupAgent\\domain\\template\\ding.md")
	message, err := ding.ParseMessage()
	if err != nil {
		t.Error(err)
		return
	}
	err = ding.SendBySelf(message)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDing_SendByProxy(t *testing.T) {
	ding := F.Ding(sa, info, func(factory *shardTemplateFactory) *shardTemplateFactory {
		factory.Title = "OpTest"
		return factory
	})
	ding.SetTemplateFilePath("G:\\backupAgent\\domain\\template\\ding.md")
	message, err := ding.ParseMessage()
	if err != nil {
		t.Error(err)
		return
	}
	addr := "10.244.188.123:39999"
	if err := ding.SendByProxy(message, addr); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("success")
}
