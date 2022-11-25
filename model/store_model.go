package model

const (
	Video = iota
	Audio
	Document
	Photo
)

type Stores struct {
	Id           int    `json:"id"`
	UserId       int64  `json:"userId"`
	FileId       string `json:"fileId"`
	FileUniqueId string `json:"fileUniqueId"`
	FileSize     int    `json:"fileSize"`
	FileName     string `json:"fileName"`
	MimeType     string `json:"mimeType"`
	Duration     int    `json:"duration"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Title        string `json:"title"`
	Performer    string `json:"performer"`
	FileType     int    `json:"fileType"`
	LocalPath    string `json:"localPath"`
	BakLocalPath string `json:"bakLocalPath"`
	CreateTime   string `json:"createTime"`
}
