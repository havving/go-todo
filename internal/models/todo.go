package models

import (
	"log"
)

func (t TodoBlock) ListTodo() (list []TodoBlock, err error) {
	rows, err := SQL.Queryx("SELECT * FROM todo ORDER BY id")
	if err != nil {
		log.Fatalln(err)
		return
	}

	list, err = entityList(rows)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (t TodoBlock) InsertTodo(paramMap map[string]string) (err error) {
	_, err = SQL.Exec("INSERT INTO todo(name,completed) VALUES($1, $2)", paramMap["name"], false)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}

func (t TodoBlock) DeleteTodo(id int) (err error) {
	_, err = SQL.Exec("DELETE FROM todo WHERE id=$1", id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}

func (t TodoBlock) SelectTodo(id int) (*TodoBlock, error) {
	var entity TodoBlock
	err := SQL.QueryRowx("SELECT completed FROM todo WHERE id=$1", id).StructScan(&entity)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &entity, nil
}

func (t TodoBlock) UpdateTodo(completed bool, id int) (err error) {
	var whether bool // false
	if completed == false {
		whether = true
	}

	_, err = SQL.Exec("UPDATE todo SET completed=$1 where id=$2", whether, id)
	if err != nil {
		log.Fatalln(err)
		return
	}

	return
}
