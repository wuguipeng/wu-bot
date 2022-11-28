package bot

import (
	"os"
	"strings"
	"wu-bot/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func download(fileId string) (tgbotapi.File, error) {
	fileConfig := tgbotapi.FileConfig{
		FileID: fileId,
	}
	file, err := bot.GetFile(fileConfig)
	return file, err
}

// 重命名
func rename(FilePath, FileName string) string {
	split := strings.Split(FilePath, "/")
	split[len(split)-1] = FileName
	join := strings.Join(split, "/")
	err := os.Rename(FilePath, join)
	if err != nil {
		logger.Errorf("文件重命名失败！", err)
		return FilePath
	}
	return join
}
