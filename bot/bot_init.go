package bot

import (
	"fmt"
	"wu-bot/setting"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitBot() {
	//Telegram bot basic info
	token := setting.Setting.Bot.Token

	botInit, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, "http://www.xyxdbp.xyz:8081/bot%s/%s")
	if err != nil {
		fmt.Println("get tgbot error:", err)
	}
	botInit.Debug = true
	bot = botInit
}
