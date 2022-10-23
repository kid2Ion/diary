package infra

import (
	"database/sql"
	"diary/domain/model"
	"diary/domain/repository"
	"fmt"
)

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

func (dr *DiaryRepository) FindByTag(tag int) ([]*model.Diary, error) {
	rows, err := dr.SqlHandler.Conn.Query("SELECT * FROM diaries WHERE tag = ?", tag)
	diaries := []*model.Diary{}
	if err != nil {
		return nil, fmt.Errorf("failed to get diaries bytag: %v", err)
	}

	return rowsScan(diaries, rows), nil
}

func (dr *DiaryRepository) Create(diary *model.Diary) (int64, error) {
	res, err := dr.SqlHandler.Conn.Exec("INSERT INTO diaries (title,content,tag) VALUES (?, ?, ?)", diary.Title, diary.Content, diary.Tag)
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
	_, err := dr.SqlHandler.Conn.Exec("UPDATE diaries SET title = ?, content = ?, tag = ? WHERE id = ?", diary.Title, diary.Content, diary.Tag, id)
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
		rows.Scan(&diary.Id, &diary.Title, &diary.Content, &diary.Tag, &diary.CreatedAt, &diary.UpdatedAt)
		diaries = append(diaries, &diary)
	}
	return diaries
}
