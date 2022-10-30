package model

import "time"

type Tag struct {
	Id         int    `json:"id"`
	TagContent string `json:"tag_content"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
