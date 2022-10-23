package config

import (
	"backupAgent/domain/pkg/log"
	"flag"
	"gopkg.in/ini.v1"
	"os"
	"testing"
)

var Config *ini.File

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "set config file")
	path := getConfigPath()
	InitConfig(path)
}

func getConfigPath() string {
	testing.Init()
	flag.Parse()
	if configFile != "" {
		return configFile
	}
	p, _ := os.Getwd()
	return p + "/domain/config/config.ini"
}

func InitConfig(path string) {
	if Config != nil {
		return
	}
	cfg, err := ini.Load(path)
	if err != nil {
		log.Logger.Error("加载配置文件失败,Fail to read file: ", err)
		return
	}
	log.Logger.Infof("加载配置文件成功，当前配置文件路径:%s", path)
	Config = cfg
}

func GetStringConf(section, key string) string {
	return Config.Section(section).Key(key).MustString("获取string失败")
}

func GetBoolConf(section, key string) bool {
	return Config.Section(section).Key(key).MustBool(false)
}

func GetIntConf(section, key string) int {
	return Config.Section(section).Key(key).MustInt(0)
}
