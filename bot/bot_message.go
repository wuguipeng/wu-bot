package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wu-bot/db"
	"wu-bot/model"
)

const basePath = "/home/wuguipeng/app/telegram-bot-api/"

func Message() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		isSave := false
		if update.CallbackData() != "" {
			callback(update)
		}
		if update.Message == nil {
			continue
		}
		result := db.DB.Find(&update.Message.From)
		affected := result.RowsAffected
		if affected == 0 {
			db.DB.Create(&update.Message.From)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.IsCommand() {
			command(update)
		}

		var err error
		if update.Message.Video != nil {
			err = saveVideo(update)
			isSave = true
		}

		if update.Message.Photo != nil {
			err = savePhoto(update)
			isSave = true
		}

		if update.Message.Document != nil {
			err = saveDocument(update)
			isSave = true
		}

		if update.Message.Audio != nil {
			err = saveAudio(update)
			isSave = true
		}

		if isSave {
			if err != nil {
				msg.Text = "保存失败！"
			} else {
				msg.Text = "保存成功！"
			}
			bot.Send(msg)
		}
	}
}

func callback(update tgbotapi.Update) {
	var store model.Stores
	data := update.CallbackQuery.Data
	db.DB.Find(&store, "id = ? and user_id = ?", data, update.CallbackQuery.From.ID)
	id := tgbotapi.FileID(store.FileId)
	id.SendData()
	var msg tgbotapi.Chattable
	if store.FileType == model.Video {
		msg = tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, id)
	}
	if store.FileType == model.Document {
		msg = tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, id)
	}
	if store.FileType == model.Audio {
		msg = tgbotapi.NewAudio(update.CallbackQuery.Message.Chat.ID, id)
	}
	if store.FileType == model.Photo {
		msg = tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, id)
	}
	if msg != nil {
		bot.Send(msg)
	}
}
