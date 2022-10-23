package usecase

import (
	"diary/domain/model"
	"diary/domain/repository"
)

type DiaryUsecase interface {
	View() ([]*model.Diary, error)
	Search(word string) ([]*model.Diary, error)
	SearchByTag(tag int) ([]*model.Diary, error)
	Add(diary *model.Diary) (int64, error)
	Edit(id int, diary *model.Diary) error
	Delete(id int) error
}

type diaryUsecase struct {
	diaryRepository repository.DiaryRepository
}

func NewDiaryUsecase(diaryRepository repository.DiaryRepository) DiaryUsecase {
	diaryUsecase := diaryUsecase{diaryRepository: diaryRepository}
	return &diaryUsecase
}

func (du *diaryUsecase) View() ([]*model.Diary, error) {
	diaries, err := du.diaryRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return diaries, nil
}

func (du *diaryUsecase) Search(word string) ([]*model.Diary, error) {
	diaries, err := du.diaryRepository.Find(word)
	if err != nil {
		return nil, err
	}

	return diaries, nil
}

func (du *diaryUsecase) SearchByTag(tag int) ([]*model.Diary, error) {
	diaries, err := du.diaryRepository.FindByTag(tag)
	if err != nil {
		return nil, err
	}

	return diaries, err
}

func (du *diaryUsecase) Add(diary *model.Diary) (int64, error) {
	id, err := du.diaryRepository.Create(diary)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (du *diaryUsecase) Edit(id int, diary *model.Diary) error {
	err := du.diaryRepository.Update(id, diary)
	if err != nil {
		return err
	}

	return nil
}

func (du *diaryUsecase) Delete(id int) error {
	err := du.diaryRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
