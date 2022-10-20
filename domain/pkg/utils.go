package pkg

import (
	"backupAgent/domain/pkg/log"
	"bytes"
	"github.com/gorhill/cronexpr"
	"math/big"
	"os"
	"path"
	"strings"
	"time"
)

func CornExprToTime(exprstr string) string {
	expr, err := cronexpr.Parse(exprstr) // 如果表达式解析错误将返回一个错误
	if err != nil {
		log.Logger.Error("cron表达式转换失败", err)
		return "unknown"
	}
	nextTime := expr.Next(time.Now())
	return nextTime.Format("2006年01月02日15:04:01")
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

// GetDateByKeepNumber 根据保留周期生成日期
func GetDateByKeepNumber(k int) string {
	return time.Now().AddDate(0, 0, -k).Format("2006-01-02")
}

func CleanLocalFile(FilePath string) error {
	log.Logger.Infof("删除本地文件:%s", FilePath)
	return os.Remove(FilePath)
}

//GetFilePath 去除后缀
func GetFilePath(filePath string) string {
	ext := path.Ext(filePath)
	return strings.TrimSuffix(filePath, ext)
}

// Base58Encoding base58编码
func Base58Encoding(src string) string {
	srcByte := []byte(src)
	// todo 转成十进制
	i := big.NewInt(0).SetBytes(srcByte)
	//  循环取余
	var modSlice []byte
	for i.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0)
		i58 := big.NewInt(58)
		i.DivMod(i, i58, mod)                         // 取余
		modSlice = append(modSlice, b58[mod.Int64()]) // 将余数添加到数组中
	}
	//  把0使用字节'1'代替
	for _, s := range srcByte {
		if s != 0 {
			break
		}
		modSlice = append(modSlice, byte('1'))
	}
	//  反转byte数组
	retModSlice := ReverseByteArr(modSlice)
	log.Logger.Info("Base58加密成功")
	return string(retModSlice)
}

// Base58Decoding base58解码
func Base58Decoding(src string) string {
	// 转成byte数组
	srcByte := []byte(src)
	// 这里得到的是十进制
	ret := big.NewInt(0)
	for _, b := range srcByte {
		i := bytes.IndexByte(b58, b)
		ret.Mul(ret, big.NewInt(58))       // 乘回去
		ret.Add(ret, big.NewInt(int64(i))) // 相加
	}
	log.Logger.Info("Base58解密成功")
	return string(ret.Bytes())
}

// ReverseByteArr byte数组进行反转方式2
func ReverseByteArr(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
