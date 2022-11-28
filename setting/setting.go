package setting

import (
	"io/ioutil"
	"wu-bot/logger"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot   *Bot   `yaml:"bot"`
	Mysql *Mysql `yaml:"mysql"`
}

type Bot struct {
	Token      string `yaml:"token"`
	IsLocalApi bool   `yaml:"isLocalApi"`
	Api        string `yaml:"api"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

var Setting = Config{}

func InitSetting() {
	logger.Info("初始化配置文件")
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		logger.Error("配置文件加载失败:", err)
	}
	err = yaml.Unmarshal(file, &Setting)
	if err != nil {
		logger.Error("配置文件反序列化失败:", err)
	}
}
