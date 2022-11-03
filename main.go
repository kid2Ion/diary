package main

import (
	"diary/handler"
	"diary/injector"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("---start server---")
	// Middleware
	diaryHandler := injector.InjectDiaryHandler()
	tagHandler := injector.InjectTagHandler()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handler.InitRouting(e, diaryHandler, tagHandler)
	e.Logger.Fatal(e.Start(":8082"))
}
