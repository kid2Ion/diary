package injector

import (
	"diary/domain/repository"
	"diary/handler"
	"diary/infra"
	"diary/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func InjectDiaryRepository() repository.DiaryRepository {
	sqlHandler := InjectDB()
	return infra.NewDiaryRepository(sqlHandler)
}

func InjectTagRepository() repository.TagRepository {
	sqlHandler := InjectDB()
	return infra.NewTagRepository(sqlHandler)
}

func InjectDiaryUsecase() usecase.DiaryUsecase {
	diaryRepository := InjectDiaryRepository()
	return usecase.NewDiaryUsecase(diaryRepository)
}

func InjectTagUsecase() usecase.TagUsecase {
	tagRepository := InjectTagRepository()
	return usecase.NewTagUsecase(tagRepository)
}

func InjectDiaryHandler() handler.DiaryHandler {
	diaryHandler := InjectDiaryUsecase()
	return handler.NewDiaryHandler(diaryHandler)
}

func InjectTagHandler() handler.TagHandler {
	tagUsecase := InjectTagUsecase()
	return handler.NewTagHandler(tagUsecase)
}
