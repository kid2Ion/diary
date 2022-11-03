package handler

import (
	"diary/domain/model"
	"diary/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type DiaryHandler interface {
	View() echo.HandlerFunc
	Search() echo.HandlerFunc
	SearchByTag() echo.HandlerFunc
	Add() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type diaryHandler struct {
	diaryUsecase usecase.DiaryUsecase
}

func NewDiaryHandler(diaryUsecase usecase.DiaryUsecase) DiaryHandler {
	diaryHandler := diaryHandler{diaryUsecase}
	return &diaryHandler
}

func (dh *diaryHandler) View() echo.HandlerFunc {
	return func(c echo.Context) error {
		diaries, err := dh.diaryUsecase.View()
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, diaries)
	}
}

func (dh *diaryHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.QueryParam("word")
		diaries, err := dh.diaryUsecase.Search(word)
		fmt.Println(err)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, diaries)
	}
}

func (dh *diaryHandler) SearchByTag() echo.HandlerFunc {
	return func(c echo.Context) error {
		tag := c.QueryParam("tagId")
		tagInt, err := strconv.Atoi(tag)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "failed to convert atoi",
			}
		}

		diaries, err := dh.diaryUsecase.SearchByTag(tagInt)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, diaries)
	}
}

func (dh *diaryHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var diary model.Diary
		c.Bind(&diary)
		id, err := dh.diaryUsecase.Add(&diary)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, id)
	}
}

func (dh *diaryHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("failed to convert atoi"))
		}
		var diary model.Diary
		c.Bind(&diary)
		err = dh.diaryUsecase.Edit(id, &diary)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, id)
	}
}

func (dh *diaryHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "failed to convert atoi",
			}
		}
		err = dh.diaryUsecase.Delete(id)
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		return c.JSON(http.StatusOK, id)
	}
}
