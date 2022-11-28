package bot

import (
	"fmt"
	"wu-bot/logger"
	"wu-bot/setting"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitBot() {
	logger.Info("初始化bot！")
	//Telegram bot basic info
	token := setting.Setting.Bot.Token

	api := setting.Setting.Bot.Api
	botInit, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, api)

	if err != nil {
		fmt.Println("get tgbot error:", err)
	}
	botInit.Debug = true
	bot = botInit
}
