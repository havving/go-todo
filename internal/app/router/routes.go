package router

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/app/handlers"
	"go-todo/internal/generated"
	"net/http"
)

func withoutContent(code int) func(echo.Context) error {
	return func(ctx echo.Context) error {
		return ctx.NoContent(code)
	}
}

var Routes = []Route{
	{
		Name:        "FavIcon",
		Method:      http.MethodGet,
		Pattern:     "/favicon.ico",
		HandlerFunc: withoutContent(http.StatusNoContent),
	},
}

func Register(e *echo.Echo) {
	for _, route := range Routes {
		e.Add(route.Method, route.Pattern, route.HandlerFunc)
	}
	h := new(handlers.APIHandlerBlock)

	generated.RegisterHandlers(e, h)
}
