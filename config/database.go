package config

import (
	"database/sql"
	"fmt"
	"gobasic/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "1531"
	dbName   = "book_golang"
)

func DatabaseConnection() *sql.DB {
	// sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)
	fmt.Println(sqlInfo)

	db, err := sql.Open("mysql", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to database !!!")

	return db
}
