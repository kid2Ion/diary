package usecase

import (
	"diary/domain/model"
	"diary/domain/repository"
)

type TagUsecase interface {
	View() ([]*model.Tag, error)
	Add(tag *model.Tag) (int64, error)
	Edit(id int, tag *model.Tag) error
	Delete(id int) error
}

type tagUsecase struct {
	repository.TagRepository
}

func NewTagUsecase(tagRepository repository.TagRepository) TagUsecase {
	tagUsecase := tagUsecase{tagRepository}
	return &tagUsecase
}

func (tu *tagUsecase) View() ([]*model.Tag, error) {
	tags, err := tu.TagRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (tu tagUsecase) Add(tag *model.Tag) (int64, error) {
	id, err := tu.TagRepository.Create(tag)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tu tagUsecase) Edit(id int, tag *model.Tag) error {
	err := tu.TagRepository.Update(id, tag)
	if err != nil {
		return err
	}

	return nil
}

func (tu tagUsecase) Delete(id int) error {
	err := tu.TagRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
