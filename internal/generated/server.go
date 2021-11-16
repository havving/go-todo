package generated

import "github.com/labstack/echo/v4"

type ServerInterface interface {
	// GetTodoListHandler (GET /api/todos)
	GetTodoListHandler(ctx echo.Context) error

	// PostTodoHandler (POST /api/todos)
	PostTodoHandler(ctx echo.Context) error

	// RemoveTodoHandler (DELETE /api/todos/:id)
	RemoveTodoHandler(ctx echo.Context) error

	// UpdateTodoHandler (PUT /api/todos/:id)
	UpdateTodoHandler(ctx echo.Context) error
}

type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

func (w *ServerInterfaceWrapper) GetTodoListServer(ctx echo.Context) error {
	err := w.Handler.GetTodoListHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) PostTodoServer(ctx echo.Context) error {
	err := w.Handler.PostTodoHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) RemoveTodoServer(ctx echo.Context) error {
	err := w.Handler.RemoveTodoHandler(ctx)

	return err
}

func (w *ServerInterfaceWrapper) UpdateTodoServer(ctx echo.Context) error {
	err := w.Handler.UpdateTodoHandler(ctx)

	return err
}

type EchoRouter interface {
	Static(prefix, root string) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

func RegisterHandlers(r EchoRouter, si ServerInterface) { // 웹 서버 핸들러 생성
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Static("/", "public")

	r.GET("/api/todos", wrapper.GetTodoListServer)
	r.POST("/api/todos", wrapper.PostTodoServer)
	r.DELETE("/api/todos/:id", wrapper.RemoveTodoServer)
	r.PUT("/api/todos/:id", wrapper.UpdateTodoServer)
}
