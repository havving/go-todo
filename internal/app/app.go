package app

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/app/middleware"
	"go-todo/internal/app/router"
	"go-todo/internal/models"
	"log"
)

func Start() {
	log.Println("App Started")
	defer log.Println("App Stopped")

	// *Echo 객체 리턴
	e := echo.New()

	middleware.Register(e) // 미들웨어 설정
	router.Register(e)     // 라우터 설정
	models.Setup()         // 데이터베이스 설정

	// start server
	err := e.Start(":3000")
	if err != nil {
		panic(err)
	}

}
