package main

import (
	"diary/handler"
	"diary/injector"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("---start server---")
	diaryHandler := injector.InjectDiaryHandler()
	e := echo.New()
	handler.InitRouting(e, diaryHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
