package main

import (
	"wu-bot/bot"
	"wu-bot/db"
	"wu-bot/setting"
)

func main() {
	setting.InitSetting()
	db.InitDB()
	bot.InitBot()
	bot.Message()
}
