package test

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-todo/internal/models"
	"log"
	"testing"
	// postgres driver
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	DATABASE = "board"
	USER     = "heybin"
	PASSWORD = "1234"
	SSL      = "disable"
)

func TestConnection(t *testing.T) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", HOST, PORT, USER, PASSWORD, DATABASE, SSL)

	db, err := sql.Open("postgres", conn)
	defer db.Close()

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
	log.Println("Success connection to DB")
}

func TestSqlx(t *testing.T) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=heybin password=1234 dbname=board sslmode=disable")
	if err != nil { // sqlx.Connect안에 이미 Ping() 검사를 함
		log.Fatalln(err)
		return
	}

	rows, err := db.Queryx("SELECT * FROM todo")
	if err != nil {
		log.Fatalln(err)
		return
	}

	for rows.Next() {
		entity := new(models.TodoBlock)
		err = rows.StructScan(&entity)
		if err != nil {
			return
		}
		fmt.Printf("%#v\n", entity)
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO todo VALUES ($1, $2, $3)", 2, "writing", true)
	tx.Commit()
}
