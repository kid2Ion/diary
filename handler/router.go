package handler

import "github.com/labstack/echo"

func InitRouting(e *echo.Echo, diaryHandler DiaryHandler, tagHandler TagHandler) {
	e.GET("/", diaryHandler.View())
	e.GET("/search", diaryHandler.Search())
	e.GET("/filter", diaryHandler.SearchByTag())
	e.POST("/new", diaryHandler.Add())
	e.PUT("/edit/:id", diaryHandler.Edit())
	e.DELETE("/delete/:id", diaryHandler.Delete())

	e.GET("tags", tagHandler.View())
}
