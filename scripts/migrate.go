package main

import (
	"database/sql"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/abdoroot/subscription/util"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}

	cmd := flag.String("cmd", "up", "migrate up or down")
	flag.Parse()
	logrus.Info("migration cmd=", *cmd)

	db, err := sql.Open("postgres", util.GetPostgresConnectionString())
	if err != nil {
		logrus.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	migrationsPath := filepath.Join("", "migrations")
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres", driver)
	if err != nil {
		logrus.Fatal("NewWithDatabaseInstance error : ", err)
	}

	switch *cmd {
	case "up":
		if err := m.Up(); err != nil {
			logrus.Fatal("migrate up err: ", err)
		}
	case "down":
		if err := m.Down(); err != nil {
			logrus.Fatal("migrate down err: ", err)
		}
	}
}
