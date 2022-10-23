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

func InjectDiaryUsecase() usecase.DiaryUsecase {
	diaryRepository := InjectDiaryRepository()
	return usecase.NewDiaryUsecase(diaryRepository)
}

func InjectDiaryHandler() handler.DiaryHandler {
	diaryHandler := InjectDiaryUsecase()
	return handler.NewDiaryHandler(diaryHandler)
}
