package repository

import "diary/domain/model"

type DiaryRepository interface {
	FindAll() ([]*model.Diary, error)
	Find(word string) ([]*model.Diary, error)
	FindByTag(tag int) ([]*model.Diary, error)
	Create(diary *model.Diary) (int64, error)
	Update(id int, diary *model.Diary) error
	Delete(id int) error
}
