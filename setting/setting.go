package setting

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Bot Bot
}

type Bot struct {
	token string `yaml:"token"`
}

var Setting = Config{}

func InitSetting() {
	file, err := ioutil.ReadFile("./conf/default.yml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}
	err = yaml.Unmarshal(file, &Setting)
	fmt.Printf("asdf")
	if err != nil {
		log.Fatal("fail to yaml unmarshal:", err)
	}
}
