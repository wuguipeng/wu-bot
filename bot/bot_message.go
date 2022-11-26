package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"wu-bot/db"
)

const basePath = "/home/wuguipeng/app/telegram-bot-api/"

func Message() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackData() != "" {
			if strings.Contains(update.CallbackQuery.Data, "/delete") {
				DeleteCallback(update)
			} else {
				QueryCallback(update)
			}
		}
		if update.Message == nil {
			continue
		}
		result := db.DB.Find(&update.Message.From)
		affected := result.RowsAffected
		if affected == 0 {
			db.DB.Create(&update.Message.From)
		}

		if update.Message.IsCommand() {
			go command(update)
		}

		if update.Message.Video != nil {
			go saveVideo(update)
		}

		if update.Message.Photo != nil {
			go savePhoto(update)
		}

		if update.Message.Document != nil {
			go saveDocument(update)
		}

		if update.Message.Audio != nil {
			go saveAudio(update)
		}
	}
}
