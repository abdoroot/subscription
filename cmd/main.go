package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listenAddr := os.Getenv("LISTEN_ADDR")
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
