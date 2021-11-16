package models

import (
	"github.com/jmoiron/sqlx"
	"go-todo/pkg/database"
)

type TodoBlock struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Name      string `json:"name" db:"name"`
	Completed bool   `json:"completed,omitempty" db:"completed"`
}

type Success struct {
	Success bool `json:"success"`
}

var SQL *database.SQL
var Todo TodoBlock

func entityList(rows *sqlx.Rows) (list []TodoBlock, err error) {
	for rows.Next() {
		//entity := new(TodoBlock) // new: 포인터 반환
		entity := TodoBlock{}
		err = rows.StructScan(&entity) // StructScan must pass a pointer
		if err != nil {
			list = nil
			return
		}
		list = append(list, entity)
	}

	return
}

func Setup() {
	SQL = new(database.SQL)
	SQL.Connect()
}
