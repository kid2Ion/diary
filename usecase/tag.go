package usecase

import (
	"diary/domain/model"
	"diary/domain/repository"
)

type TagUsecase interface {
	View() ([]*model.Tag, error)
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
