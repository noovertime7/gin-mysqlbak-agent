package config

import (
	"backupAgent/domain/pkg/log"
	"flag"
	"gopkg.in/ini.v1"
	"os"
)

var Config *ini.File

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/domain/config/config.ini", "set config file")
	flag.Parse()
	if Config != nil {
		return
	}
	path, err := os.Getwd()
	cfg, err := ini.Load(path + configFile)
	if err != nil {
		log.Logger.Fatal("加载配置文件失败,Fail to read file: ", err)
	}
	log.Logger.Infof("加载配置文件成功，当前配置文件路径:%s", path+configFile)
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
