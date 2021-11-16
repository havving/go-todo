package todo

import (
	"github.com/labstack/echo/v4"
	"go-todo/internal/models"
	"net/http"
	"strconv"
)

// GetTodoList 할 일 목록을 가져오는 함수
func GetTodoList(ctx echo.Context) (err error) {
	list, err := models.Todo.ListTodo()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	err = ctx.JSON(http.StatusOK, list)

	return err
}

// PostTodo 할 일을 목록에 추가하는 함수
func PostTodo(ctx echo.Context) error {
	params := make(map[string]string)
	// swagger TEST용
	err := ctx.Bind(&params)
	//err := json.NewDecoder(ctx.Request().Body).Decode(&params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = models.Todo.InsertTodo(params)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	err = ctx.JSON(http.StatusCreated, models.TodoBlock{Name: params["name"]})

	return err
}

// RemoveTodo 할 일을 목록에서 삭제하는 함수
func RemoveTodo(ctx echo.Context) error {
	vars := ctx.Param("id")
	id, _ := strconv.Atoi(vars) // 숫자로 이루어진 문자열을 숫자로 변환

	err := models.Todo.DeleteTodo(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Success{Success: false})
	}

	err = ctx.JSON(http.StatusOK, models.Success{Success: true})

	return err
}

// UpdateTodo 할 일을 수정(완료)하는 함수
func UpdateTodo(ctx echo.Context) error {
	vars := ctx.Param("id")
	id, _ := strconv.Atoi(vars)
	entity, err := models.Todo.SelectTodo(id)

	err = models.Todo.UpdateTodo(entity.Completed, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Success{Success: false})
	}

	err = ctx.JSON(http.StatusOK, models.Success{Success: true})

	return err
}
