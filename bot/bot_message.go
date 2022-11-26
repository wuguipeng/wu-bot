package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"strings"
	"wu-bot/db"
	"wu-bot/model"
)

const basePath = "/home/wuguipeng/app/telegram-bot-api/"

func Message() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackData() != "" {
			if strings.Contains(update.CallbackQuery.Data, "/delete") {
				deleteCallback(update)
			} else {
				callback(update)
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

func deleteCallback(update tgbotapi.Update) {
	replace := strings.Replace(update.CallbackQuery.Data, "/delete ", "", 1)
	var store model.Stores
	db.DB.Find(&store, "id = ? and user_id = ?", replace, update.CallbackQuery.From.ID)
	err := os.Remove(store.LocalPath)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "删除失败")
	if err == nil {
		db.DB.Delete(&store)
		msg.Text = "删除成功"
	}
	bot.Send(msg)
}

func callback(update tgbotapi.Update) {
	var store model.Stores
	data := update.CallbackQuery.Data
	db.DB.Find(&store, "id = ? and user_id = ?", data, update.CallbackQuery.From.ID)

	id := tgbotapi.FileID(store.FileId)
	id.SendData()
	var msg tgbotapi.Chattable

	if store.FileType == model.Video {
		msg1 := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, id)
		msg1.ReplyMarkup = deleteStore(store)
		msg = msg1
	}
	if store.FileType == model.Document {
		msg2 := tgbotapi.NewDocument(update.CallbackQuery.Message.Chat.ID, id)
		msg2.ReplyMarkup = deleteStore(store)
		msg = msg2
	}
	if store.FileType == model.Audio {
		msg3 := tgbotapi.NewAudio(update.CallbackQuery.Message.Chat.ID, id)
		msg3.ReplyMarkup = deleteStore(store)
		msg = msg3
	}
	if store.FileType == model.Photo {
		msg4 := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, id)
		msg4.ReplyMarkup = deleteStore(store)
		msg = msg4
	}
	if msg != nil {
		bot.Send(msg)
	}
}

func deleteStore(store model.Stores) tgbotapi.InlineKeyboardMarkup {
	callbackDelete := "/delete " + strconv.Itoa(store.Id)
	inlineKeyboardButton := [][]tgbotapi.InlineKeyboardButton{
		{tgbotapi.InlineKeyboardButton{Text: "删除文件", CallbackData: &callbackDelete}},
	}
	reply := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardButton,
	}
	return reply
}
