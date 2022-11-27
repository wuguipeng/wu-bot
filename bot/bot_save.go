package bot

import (
	"strconv"
	"time"
	"wu-bot/db"
	"wu-bot/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func saveVideo(update tgbotapi.Update) {
	notify := downloadBeforeNotify(update)
	video := update.Message.Video
	fileInfo, err := download(video.FileID)
	if err != nil {
		saveFailMsg(update, notify)
		return
	}
	newFilePath := rename(fileInfo.FilePath, video.FileName)
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		Duration:     video.Duration,
		Width:        video.Width,
		Height:       video.Height,
		FileName:     video.FileName,
		MimeType:     video.MimeType,
		FileId:       video.FileID,
		FileUniqueId: video.FileUniqueID,
		FileSize:     video.FileSize,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Video,
		LocalPath:    newFilePath,
	}
	db.DB.Create(&stores)
	downloadAfterNotify(update, stores, notify)
}

func saveAudio(update tgbotapi.Update) {
	notify := downloadBeforeNotify(update)
	audio := update.Message.Audio
	fileInfo, err := download(audio.FileID)
	if err != nil {
		saveFailMsg(update, notify)
		return
	}
	newFilePath := rename(fileInfo.FilePath, audio.FileName)
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		Duration:     audio.Duration,
		FileName:     audio.FileName,
		MimeType:     audio.MimeType,
		Title:        audio.Title,
		FileId:       audio.FileID,
		FileUniqueId: audio.FileUniqueID,
		FileSize:     audio.FileSize,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Audio,
		LocalPath:    newFilePath,
		Performer:    audio.Performer,
	}
	db.DB.Create(&stores)
	downloadAfterNotify(update, stores, notify)
}

func saveDocument(update tgbotapi.Update) {
	notify := downloadBeforeNotify(update)
	document := update.Message.Document
	fileInfo, err := download(document.FileID)
	if err != nil {
		saveFailMsg(update, notify)
		return
	}
	newFilePath := rename(fileInfo.FilePath, document.FileName)
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		FileName:     document.FileName,
		MimeType:     document.MimeType,
		FileId:       document.FileID,
		FileUniqueId: document.FileUniqueID,
		FileSize:     document.FileSize,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Document,
		LocalPath:    newFilePath,
	}
	db.DB.Create(&stores)
	downloadAfterNotify(update, stores, notify)
}

func savePhoto(update tgbotapi.Update) {
	notify := downloadBeforeNotify(update)
	photo := update.Message.Photo
	photoSize := photo[len(photo)-1]
	fileInfo, err := download(photoSize.FileID)
	if err != nil {
		saveFailMsg(update, notify)
		return
	}
	fileName := time.Now().Format("2006-01-02 15:04:05") + strconv.Itoa(update.Message.MessageID) + ".png"
	newFilePath := rename(fileInfo.FilePath, fileName)
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		Width:        photoSize.Width,
		Height:       photoSize.Height,
		FileId:       photoSize.FileID,
		FileUniqueId: photoSize.FileUniqueID,
		FileSize:     photoSize.FileSize,
		FileName:     fileName,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Photo,
		LocalPath:    newFilePath,
	}
	db.DB.Create(&stores)
	downloadAfterNotify(update, stores, notify)
}

func downloadBeforeNotify(update tgbotapi.Update) tgbotapi.Message {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Downloading file ...")
	msg.ReplyToMessageID = update.Message.MessageID
	send, _ := bot.Send(msg)
	return send
}

func downloadAfterNotify(update tgbotapi.Update, store model.Stores, notify tgbotapi.Message) {
	msg := tgbotapi.NewEditMessageText(update.Message.Chat.ID, notify.MessageID, "Download completed \r\n")
	msg.Text += "FileName: " + store.FileName
	bot.Send(msg)
}

func saveFailMsg(update tgbotapi.Update, notify tgbotapi.Message) {
	msg := tgbotapi.NewEditMessageText(update.Message.Chat.ID, notify.MessageID, "Download Error ...")
	bot.Send(msg)
}
