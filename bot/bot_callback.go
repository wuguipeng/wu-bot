package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"strings"
	"wu-bot/db"
	"wu-bot/model"
)

func QueryCallback(update tgbotapi.Update) {
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

func DeleteCallback(update tgbotapi.Update) {
	replace := strings.Replace(update.CallbackQuery.Data, "/delete ", "", 1)
	var store model.Stores
	db.DB.Find(&store, "id = ? and user_id = ?", replace, update.CallbackQuery.From.ID)
	if (model.Stores{}) == store {
		return
	}

	_, err := os.Stat(store.LocalPath)
	msg := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
	if err != nil {
		db.DB.Delete(&store)
		msg.Text = "文件不存在，删除记录"
	} else {
		err = os.Remove(store.LocalPath)
		msg.Text = "删除失败"
		if err == nil {
			db.DB.Delete(&store)
			msg.Text = "删除成功"
		}
	}
	deleteMsg := tgbotapi.NewDeleteMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
	bot.Send(deleteMsg)
	bot.Send(msg)
}

func deleteStore(store model.Stores) tgbotapi.InlineKeyboardMarkup {
	callbackDelete := "/delete " + strconv.Itoa(store.Id)
	row := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("删除文件", callbackDelete),
	)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}
