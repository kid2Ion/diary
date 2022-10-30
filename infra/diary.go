package infra

import (
	"database/sql"
	"diary/domain/model"
	"diary/domain/repository"
	"fmt"
	"time"
)

var layout = "2006-01-02 15:04:05"

type DiaryRepository struct {
	SqlHandler
}

func NewDiaryRepository(sqHandler SqlHandler) repository.DiaryRepository {
	diaryRepository := DiaryRepository{sqHandler}
	return &diaryRepository
}

func (dr *DiaryRepository) FindAll() ([]*model.Diary, error) {
	rows, err := dr.SqlHandler.Conn.Query("SELECT * FROM diaries")
	if err != nil {
		return nil, fmt.Errorf("failed to get all diaries, %v", err)
	}

	diaries := []*model.Diary{}
	return rowsScan(diaries, rows), nil
}

func (dr *DiaryRepository) Find(word string) ([]*model.Diary, error) {
	rows, err := dr.SqlHandler.Conn.Query("SELECT * FROM diaries WHERE title LIKE ?", "%"+word+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to get diaries, %v", err)
	}

	diaries := []*model.Diary{}
	return rowsScan(diaries, rows), nil
}

func (dr *DiaryRepository) FindByTag(tagId int) ([]*model.Diary, error) {
	rows, err := dr.SqlHandler.Conn.Query("SELECT * FROM diaries WHERE tag_id = ?", tagId)
	diaries := []*model.Diary{}
	if err != nil {
		return nil, fmt.Errorf("failed to get diaries by tag_id: %v", err)
	}

	return rowsScan(diaries, rows), nil
}

func (dr *DiaryRepository) Create(diary *model.Diary) (int64, error) {
	res, err := dr.SqlHandler.Conn.Exec("INSERT INTO diaries (title,content,tag_id) VALUES (?, ?, ?)", diary.Title, diary.Content, diary.TagId)
	if err != nil {
		return 0, fmt.Errorf("failed to create diary: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dr *DiaryRepository) Update(id int, diary *model.Diary) error {
	_, err := dr.SqlHandler.Conn.Exec("UPDATE diaries SET title = ?, content = ?, tag_id = ? WHERE id = ?", diary.Title, diary.Content, diary.TagId, id)
	if err != nil {
		return fmt.Errorf("failed to update diaries: %v", err)
	}
	if err != nil {
		return err
	}
	return nil
}

func (dr *DiaryRepository) Delete(id int) error {
	_, err := dr.SqlHandler.Conn.Exec("DELETE FROM diaries WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete diary: %v", err)
	}
	if err != nil {
		return err
	}
	return nil
}

func rowsScan(diaries []*model.Diary, rows *sql.Rows) []*model.Diary {
	defer rows.Close()
	for rows.Next() {
		diary := model.Diary{}
		// time.Timeが直接scanできなかった
		var createdAt string
		var updatedAt string
		rows.Scan(&diary.Id, &diary.Title, &diary.Content, &diary.TagId, &createdAt, &updatedAt)
		diary.CreatedAt = stringToTime(createdAt)
		diary.UpdatedAt = stringToTime(updatedAt)
		diaries = append(diaries, &diary)
	}
	return diaries
}

func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}
