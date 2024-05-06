package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func ConnectToPq() (*sqlx.DB, error) {

	var (
		host         = os.Getenv("DB_HOST")
		port         = os.Getenv("DB_PORT")
		databaseName = os.Getenv("DB_DATABASE")
		dbUsername   = os.Getenv("DB_USERNAME")
		dbPassword   = os.Getenv("DB_PASSWORD")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, dbUsername, dbPassword, databaseName)
	logrus.Info(psqlInfo)
	db, err := sqlx.Connect("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}
