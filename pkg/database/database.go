package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	// postgres driver
	_ "github.com/lib/pq"
)

type SQL struct {
	db        *sqlx.DB
	Exec      func(query string, args ...interface{}) (sql.Result, error)
	Queryx    func(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx func(query string, args ...interface{}) *sqlx.Row
}

func (x *SQL) Connect() (err error) {
	x.db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=heybin password=1234 dbname=board sslmode=disable")
	if err != nil {
		log.Warn(err, "")
		return
	}

	x.Exec = func(query string, args ...interface{}) (sql.Result, error) {
		log.Print("%s / %v", query, args)
		return x.db.Exec(query, args...)
	}

	x.Queryx = func(query string, args ...interface{}) (*sqlx.Rows, error) {
		log.Print("%s / %v", query, args)
		return x.db.Queryx(query, args...)
	}

	x.QueryRowx = func(query string, args ...interface{}) *sqlx.Row {
		log.Print("%s / %v", query, args)
		return x.db.QueryRowx(query, args...)
	}

	return
}
