package models

type Media struct {
	ID     string   `json:"id"`
	Name   string   `json:"name" binding:"required"`
	Tags   []string `json:"tags" binding:"required"`
	FileUrl string  `json:"fileUrl"`
}

var Medias = []Media{}
