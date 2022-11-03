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

func (tr *TagRepository) Create(tag *model.Tag) (int64, error) {
	res, err := tr.SqlHandler.Conn.Exec("INSERT INTO tags (tag_content) VALUES (?)", tag.TagContent)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TagRepository) Update(id int, tag *model.Tag) error {
	_, err := tr.SqlHandler.Conn.Exec("UPDATE tags SET tag_content = ? WHERE id = ?", tag.TagContent, id)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TagRepository) Delete(id int) error {
	_, err := tr.SqlHandler.Conn.Exec("DELETE FROM tags WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
