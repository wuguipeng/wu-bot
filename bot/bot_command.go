package bot

import (
	"fmt"
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
		msg.Text = "更新按钮"
	case "rename":
		renameCommand(update)
	case "withArgument":
		msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
	default:
		msg.Text = "I don't know that command"
	}
	if len(stores) > 0 {
		msg.Text = "查询结果，共" + strconv.Itoa(len(stores)) + "条数据"
		msg.ReplyMarkup = inline(stores)
	}
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

func renameCommand(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	replace := strings.Split(update.Message.Text, " ")
	if len(replace) != 3 {
		msg.Text = "重命名格式错误! 格式为 /rename id newName"
		bot.Send(msg)
		return
	}
	fmt.Println(replace)
	var store model.Stores
	db.DB.Find(&store, "id = ? and user_id = ?", replace[1], update.CallbackQuery.From.ID)
	if (model.Stores{}) == store {
		return
	}
	newFilePath := rename(store.LocalPath, replace[2])
	store.LocalPath = newFilePath
	store.FileName = replace[2]
	db.DB.Save(&store)
	msg.Text = "更新成功，新文件名为： " + replace[2]
	bot.Send(msg)
}
