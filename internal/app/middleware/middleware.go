package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Register(e *echo.Echo) {
	// 디버깅 모드
	e.Debug = true

	// 로그 출력 미들웨어
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// 에러가 발생해도 서버가 바로 죽지 않고, 다시 살아나도록 하는 미들웨어
	e.Use(middleware.Recover())
	// CORS 에러 방지를 위한 미들웨어
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// OpenAPI 사용 미들웨어
	e.Use(oapiRequestValidate)
}
