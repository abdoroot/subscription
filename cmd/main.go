package main

import (
	"net/http"
	"os"

	"github.com/abdoroot/subscription/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	listenAddr := os.Getenv("HTTP_LISTER_ADDRESS")
	//Connect to database
	db, err := util.ConnectToPq()
	if err != nil {
		logrus.Fatal("error connecting to db", err)
	}
	_ = db

	router := fiber.New(fiber.Config{
		Immutable: true,
	})

	//static file server
	router.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./static"),
	}))

	logrus.Info("server is running at port ", listenAddr)
	logrus.Error(router.Listen(listenAddr))

}
