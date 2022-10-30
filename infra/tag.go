package infra

import (
	"diary/domain/model"
	"diary/domain/repository"
	myutil "diary/utility"
)

type TagRepository struct {
	SqlHandler
}

func NewTagRepository(sqlHandler SqlHandler) repository.TagRepository {
	tagRepository := TagRepository{sqlHandler}
	return &tagRepository
}

func (tr *TagRepository) FindAll() ([]*model.Tag, error) {
	rows, err := tr.SqlHandler.Conn.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tags := []*model.Tag{}
	for rows.Next() {
		tag := model.Tag{}
		// time.Timeが直接scanできなかった
		var createdAt string
		var updatedAt string
		rows.Scan(&tag.Id, &tag.TagContent, &createdAt, &updatedAt)
		tag.CreatedAt = myutil.StringToTime(createdAt)
		tag.UpdatedAt = myutil.StringToTime(updatedAt)
		tags = append(tags, &tag)
	}
	return tags, nil
}
