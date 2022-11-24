package main

import (
	"fmt"
	"wu-bot/setting"
)

func main() {
	setting.InitSetting()
	tgBot := setting.Setting.Bot
	fmt.Printf("setting.Setting.TgBot: %v\n", tgBot)
}
