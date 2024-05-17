package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdoroot/subscription/util"
	"github.com/abdoroot/subscription/views/home"
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

	router := fiber.New(fiber.Config{
		Immutable: true,
	})

	//static file server
	router.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./static"),
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return util.RenderHtml(c, home.Index())
	})

	router.Get("/users", func(c *fiber.Ctx) error {
		//for test perpose
		if IHtmxRequest(c) {
			return util.RenderHtml(c, home.Users())
		}
		return util.RenderHtml(c, home.FullUsers())
	})

	//kill server signal
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := router.Listen(listenAddr); err != nil {
			logrus.Error("Server starting error ", err)
		}
	}()
	logrus.Info("server is running at port ", listenAddr)
	<-killSig //wait until get any kill signal
	logrus.Info("close database conn")
	if err := db.Close(); err != nil {
		logrus.Error("closing database conn err", err)
	}
	logrus.Info("Shutting down server")
	if err := router.Shutdown(); err != nil {
		//gracefully shutdown
		logrus.Error("Shutting database server err", err)
		os.Exit(1)
	}
	logrus.Info("Server shutdown complete")
}

func IHtmxRequest(c *fiber.Ctx) bool {
	key := c.Request().Header.Peek("Hx-Request")
	keyString := string(key)
	return len(keyString) > 0
}
