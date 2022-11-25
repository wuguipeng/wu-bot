package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
	"wu-bot/db"
	"wu-bot/model"
)

func saveVideo(update tgbotapi.Update) error {
	video := update.Message.Video
	fileInfo, err := download(video.FileID)
	if err != nil {
		return err
	}
	//rename(file.FilePath, newPath)
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
	return nil
}

func saveAudio(update tgbotapi.Update) error {
	audio := update.Message.Audio
	fileInfo, err := download(audio.FileID)
	if err != nil {
		return err
	}
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
	return nil
}

func saveDocument(update tgbotapi.Update) error {
	document := update.Message.Document
	fileInfo, err := download(document.FileID)
	if err != nil {
		return err
	}
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
	return nil
}

func savePhoto(update tgbotapi.Update) error {
	photo := update.Message.Photo
	photoSize := photo[len(photo)-1]
	fileInfo, err := download(photoSize.FileID)
	if err != nil {
		return err
	}
	stores := model.Stores{
		UserId:       update.Message.From.ID,
		Width:        photoSize.Width,
		Height:       photoSize.Height,
		FileId:       photoSize.FileID,
		FileUniqueId: photoSize.FileUniqueID,
		FileSize:     photoSize.FileSize,
		CreateTime:   time.Now().Format("2006-01-02 15:04:05"),
		FileType:     model.Photo,
		LocalPath:    fileInfo.FilePath,
	}
	db.DB.Create(&stores)
	return nil
}
