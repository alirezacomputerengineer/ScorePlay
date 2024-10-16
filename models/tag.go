package models

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

var Tags = []Tag{}
