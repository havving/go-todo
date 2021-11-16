package middleware

import (
	"github.com/deepmap/oapi-codegen/examples/petstore-expanded/chi/api"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/labstack/echo/v4"
)

func oapiRequestValidate(next echo.HandlerFunc) echo.HandlerFunc {
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil

	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(err)
	}

	return func(ctx echo.Context) error {
		req := ctx.Request()
		_, _, errRoute := router.FindRoute(req)
		if errRoute == routers.ErrPathNotFound {
			return next(ctx)
		}

		return next(ctx)
	}
}
