package util

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToPq() (*sqlx.DB, error) {

	psqlInfo := GetPqUrl(true)
	db, err := sqlx.Connect("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}

func GetPqUrl(withDB bool) string {
	//postqres connect url
	var (
		host         = os.Getenv("DB_HOST")
		port         = os.Getenv("DB_PORT")
		databaseName = os.Getenv("DB_DATABASE")
		dbUsername   = os.Getenv("DB_USERNAME")
		dbPassword   = os.Getenv("DB_PASSWORD")
	)
	if withDB {
		return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, dbUsername, dbPassword, databaseName)
	}
	return fmt.Sprintf("host=%s port=%s user=%s "+"password=%s sslmode=disable", host, port, dbUsername, dbPassword)
}

func GetPostgresConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUsername, dbPassword, host, port, databaseName)
}
