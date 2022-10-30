package model

import "time"

type Diary struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	TagId     int    `json:"tagId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
