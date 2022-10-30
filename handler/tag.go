package handler

import (
	"diary/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type TagHandler interface {
	View() echo.HandlerFunc
}

type tagHandler struct {
	usecase.TagUsecase
}

func NewTagHandler(tagUsecase usecase.TagUsecase) TagHandler {
	tagHandler := tagHandler{tagUsecase}
	return &tagHandler
}

func (th *tagHandler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		tags, err := th.TagUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, tags)
	}
}
