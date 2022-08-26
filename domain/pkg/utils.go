package pkg

import (
	"backupAgent/domain/pkg/log"
	"github.com/gorhill/cronexpr"
	"os"
	"time"
)

func Cronexpr(exprstr string) (string, error) {
	expr, err := cronexpr.Parse(exprstr) // 如果表达式解析错误将返回一个错误
	if err != nil {
		return "unknown", err
	}
	nextTime := expr.Next(time.Now())
	return nextTime.Format("2006年01月02日15:04:01"), nil
}

func IntToBool(a int64) bool {
	if a == 0 {
		return false
	}
	return true
}

func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		log.Logger.Errorf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		return
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Logger.Errorf("创建目录异常 -> %v\n", err)
		} else {
			log.Logger.Infof("创建文件夹%s成功!", path)
		}
	}
}

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func GetFileSize(fileanme string) int {
	fileInfo, err := os.Stat(fileanme)
	if err != nil {
		return 0
	}
	tmp := int(fileInfo.Size()) / 1024
	return tmp
}

func StatusConversion(a int64) string {
	switch a {
	case 0:
		return "失败"
	case 1:
		return "成功"
	case 2:
		return "未启用"
	}
	return "unknown"
}
