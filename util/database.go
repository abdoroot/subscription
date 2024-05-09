package util

import (
	"fmt"
	"os"
	"reflect"
	"slices"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToPq() (*sqlx.DB, error) {

	psqlInfo := GetPqUrl(true)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetPqUrl(withDB bool) string {
	//postqres connect url
	loadEnvFile()
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

func loadEnvFile() error {
	return godotenv.Load("../.env")
}

func SqlxStructUpdateQueryBuilder(s any, tableName string, exclude ...string) (string, error) {
	//s struct type
	if reflect.TypeOf(s).Kind().String() != "struct" {
		//allow only struct types
		return "", fmt.Errorf("please use struct type only")
	}
	var (
		query   string
		tagName string = "db"
	)
	query = fmt.Sprintf("update %v set ", tableName)
	t := reflect.TypeOf(s)
	n := t.NumField()
	for i := 0; i < n; i++ {
		tag := t.Field(i).Tag.Get(tagName)
		if tag == "-" || slices.Contains(exclude, tag) {
			continue
		}
		query += fmt.Sprintf("%v=:%v%v", tag, tag, ",")
	}
	query = query[:len(query)-1] // remove the trailing comma
	v, ok := s.(types.User)
	if ok && v.Id != 0 {
		query += fmt.Sprintf(" where id=%v", v.Id)
	}

	return query, nil
}
