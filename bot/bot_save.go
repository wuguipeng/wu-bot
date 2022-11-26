package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
	"time"
	"wu-bot/db"
	"wu-bot/model"
)

func saveVideo(update tgbotapi.Update) {
	video := update.Message.Video
	fileInfo, err := download(video.FileID)
	if err != nil {
		saveFailMsg(update)
		return
	}
	rename(fileInfo.FilePath, video.FileName)
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
		LocalPath:    fileInfo.FilePath,
	}
	db.DB.Create(&stores)
	saveSuccessMsg(update, stores)
}

func saveAudio(update tgbotapi.Update) {
	audio := update.Message.Audio
	fileInfo, err := download(audio.FileID)
	if err != nil {
		saveFailMsg(update)
		return
	}
	rename(fileInfo.FilePath, audio.FileName)
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
		LocalPath:    fileInfo.FilePath,
		Performer:    audio.Performer,
	}
	db.DB.Create(&stores)
	saveSuccessMsg(update, stores)
}

func saveDocument(update tgbotapi.Update) {
	document := update.Message.Document
	fileInfo, err := download(document.FileID)
	if err != nil {
		saveFailMsg(update)
		return
	}
	rename(fileInfo.FilePath, document.FileName)
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		FileName:     document.FileName,
		MimeType:     document.MimeType,
		FileId:       document.FileID,
		FileUniqueId: document.FileUniqueID,
		FileSize:     document.FileSize,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Document,
		LocalPath:    fileInfo.FilePath,
	}
	db.DB.Create(&stores)
	saveSuccessMsg(update, stores)
}

func savePhoto(update tgbotapi.Update) {
	photo := update.Message.Photo
	photoSize := photo[len(photo)-1]
	fileInfo, err := download(photoSize.FileID)
	if err != nil {
		saveFailMsg(update)
		return
	}
	fileName := time.Now().Format("2006-01-02 15:04:05") + ".png"
	rename(fileInfo.FilePath, fileName)
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
		LocalPath:    fileInfo.FilePath,
	}
	db.DB.Create(&stores)
	saveSuccessMsg(update, stores)
}

func saveSuccessMsg(update tgbotapi.Update, store model.Stores) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "保存成功, 文件信息\r\n")
	out, _ := yaml.Marshal(store)
	msg.Text += string(out)
	bot.Send(msg)
}

func saveFailMsg(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "保存失败")
	bot.Send(msg)
}
