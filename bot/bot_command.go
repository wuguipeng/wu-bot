package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"wu-bot/db"
	"wu-bot/model"
)

func command(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	var stores []model.Stores
	switch update.Message.Command() {
	case "help":
		msg.Text = "type /status. /menu"
	case "status":
		msg.Text = "is Ok"
	case "Video":
		msg.Text = "视频查询结果"
		db.DB.Find(&stores, "file_type = ? and user_id = ?", model.Video, update.Message.From.ID)
	case "Audio":
		msg.Text = "音乐查询结果"
		db.DB.Find(&stores, "file_type = ? and user_id = ?", model.Audio, update.Message.From.ID)
	case "Document":
		msg.Text = "文档查询结果"
		db.DB.Find(&stores, "file_type = ? and user_id = ?", model.Document, update.Message.From.ID)
	case "Photo":
		msg.Text = "图片查询结果"
		db.DB.Find(&stores, "file_type = ? and user_id = ?", model.Photo, update.Message.From.ID)
	case "menu":
		// 按键
		msg.ReplyMarkup = keyboard()
	case "withArgument":
		msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
	default:
		msg.Text = "I don't know that command"
	}
	msg.ReplyMarkup = inline(stores)
	bot.Send(msg)
}

func inline(stores []model.Stores) tgbotapi.InlineKeyboardMarkup {

	// 构建动态数组
	inlineKeyboardButton := make([][]tgbotapi.InlineKeyboardButton, len(stores))
	for i := 0; i < len(stores); i++ {
		inlineKeyboardButton[i] = make([]tgbotapi.InlineKeyboardButton, 1)
	}
	// 赋值
	for i := 0; i < len(stores); i++ {
		callback := strconv.Itoa(stores[i].Id)
		var fileName = stores[i].FileName
		if fileName == "" {
			split := strings.Split(stores[i].LocalPath, "/")
			fileName = split[len(split)-1]
		}
		inlineKeyboardButton[i][0] = tgbotapi.InlineKeyboardButton{Text: fileName, CallbackData: &callback}
	}
	reply := tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboardButton,
	}
	return reply
}

func keyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboardButton := [][]tgbotapi.KeyboardButton{
		{tgbotapi.KeyboardButton{Text: "/Video"}, tgbotapi.KeyboardButton{Text: "/Audio"}},
		{tgbotapi.KeyboardButton{Text: "/Photo"}, tgbotapi.KeyboardButton{Text: "/Document"}}}

	keyboard := tgbotapi.ReplyKeyboardMarkup{
		Keyboard:              keyboardButton,
		ResizeKeyboard:        true,
		OneTimeKeyboard:       true,
		InputFieldPlaceholder: "/help",
		Selective:             true,
	}
	return keyboard
}
