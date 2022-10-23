package model

import "time"

type Diary struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Tag       int    `json:"tag"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
