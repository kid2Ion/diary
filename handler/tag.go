package handler

import (
	"diary/domain/model"
	"diary/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TagHandler interface {
	View() echo.HandlerFunc
	Add() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Delete() echo.HandlerFunc
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
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
		return c.JSON(http.StatusOK, tags)
	}
}

func (th *tagHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tag model.Tag
		c.Bind(&tag)
		id, err := th.TagUsecase.Add(&tag)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
		return c.JSON(http.StatusOK, id)
	}
}

func (th *tagHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
		var tag model.Tag
		c.Bind(&tag)
		err = th.TagUsecase.Edit(id, &tag)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, nil)
	}
}

func (th *tagHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
		err = th.TagUsecase.Delete(id)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, nil)
	}
}
