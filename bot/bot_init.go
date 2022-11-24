package bot

import (
	"fmt"
	"wu-bot/setting"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitBot() {
	//Telegram bot basic info
	tgBottoken := setting.Setting.Bot.Token
	// tgBotid, err := j.settingService.GetTgBotChatId()
	// if err != nil {
	// 	logger.Warning("sendMsgToTgbot failed,GetTgBotChatId fail:", err)
	// 	return
	// }

	botInit, err := tgbotapi.NewBotAPI(tgBottoken)
	if err != nil {
		fmt.Println("get tgbot error:", err)
	}
	botInit.Debug = true
	bot = botInit
}
