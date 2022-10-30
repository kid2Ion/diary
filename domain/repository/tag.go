package repository

import "diary/domain/model"

type TagRepository interface {
	FindAll() ([]*model.Tag, error)
	// Create(tag *model.Tag) (int64, error)
	// Update(id int, tag *model.Tag) error
	// Delete(id int) error
}
