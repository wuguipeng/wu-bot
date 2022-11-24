package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot *Bot `yaml:"bot"`
}

type Bot struct {
	Token string `yaml:"token"`
}

var Setting = Config{}

func InitSetting() {
	file, err := ioutil.ReadFile("./conf/config.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, &Setting)
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
}
