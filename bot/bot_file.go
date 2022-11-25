package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"wu-bot/logger"
)

func download(fileId string) (tgbotapi.File, error) {
	fileConfig := tgbotapi.FileConfig{
		FileID: fileId,
	}
	file, err := bot.GetFile(fileConfig)
	return file, err
}

func rename(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		logger.Errorf("文件重命名失败！", err)
	}
}
