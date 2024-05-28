package main

import (
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/abdoroot/subscription/util"
	"github.com/abdoroot/subscription/views/customer"
	"github.com/abdoroot/subscription/views/home"
	"github.com/abdoroot/subscription/views/item"
	subitem "github.com/abdoroot/subscription/views/subscriptionitem"
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
		logrus.Fatal("error connecting to db ", err)
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

	c := router.Group("customer")
	{
		c.Get("/", func(c *fiber.Ctx) error {
			//for test perpose
			if IsHXRequest(c) {
				return util.RenderHtml(c, customer.Customer())
			}
			return util.RenderHtml(c, customer.FullCustomerTempl())
		})

		c.Get("/:id<int>", func(c *fiber.Ctx) error {
			//Show customer
			logrus.Info("Show customer route")
			if IsHXRequest(c) {
				return util.RenderHtml(c, customer.Show())
			}
			return util.RenderHtml(c, customer.FullShowTempl())
		})

		c.Get("/:id<int>/overview", func(c *fiber.Ctx) error {
			//Show customer
			//block html
			return util.RenderHtml(c, customer.OverviewTab())
		})

		c.Get("/:id<int>/transaction", func(c *fiber.Ctx) error {
			//Show customer
			//block html
			r := rand.Intn(math.MaxInt)
			rs := strconv.Itoa(r)
			return util.RenderHtml(c, customer.TransactionTab(rs))
		})

		c.Get("/:id<int>/mail", func(c *fiber.Ctx) error {
			//Show customer
			//block html
			r := rand.Intn(math.MaxInt)
			rs := strconv.Itoa(r)
			return util.RenderHtml(c, customer.MailTab(rs))
		})

		c.Get("/:id<int>/statement", func(c *fiber.Ctx) error {
			//Show customer
			//block html
			r := rand.Intn(math.MaxInt)
			rs := strconv.Itoa(r)
			return util.RenderHtml(c, customer.StatementTab(rs))
		})

		c.Get("/create", func(c *fiber.Ctx) error {
			//Create customer
			logrus.Info("create customer route")
			if IsHXRequest(c) {
				return util.RenderHtml(c, customer.Create())
			}
			return util.RenderHtml(c, customer.FullCreateTempl())
		})
	}

	i := router.Group("item")
	{
		i.Get("/", func(c *fiber.Ctx) error {
			if IsHXRequest(c) {
				return util.RenderHtml(c, item.Index())
			}
			return util.RenderHtml(c, item.FullTempl())
		})
		i.Get("/create", func(c *fiber.Ctx) error {
			if IsHXRequest(c) {
				return util.RenderHtml(c, item.Create())
			}
			return util.RenderHtml(c, item.FullCreate())
		})
	}

	si := router.Group("subscription-item")
	{
		//Subscriptions items
		si.Get("/", func(c *fiber.Ctx) error {
			if IsHXRequest(c) {
				return util.RenderHtml(c, subitem.Index())
			}
			return util.RenderHtml(c, subitem.FullTempl())
		})

		//Subscriptions items
		si.Get("/plan/create", func(c *fiber.Ctx) error {
			if IsHXRequest(c) {
				return util.RenderHtml(c, subitem.CreatePlan())
			}
			return util.RenderHtml(c, subitem.FullCreatePlan())
		})
	}

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

func IsHXRequest(c *fiber.Ctx) bool {
	key := c.Request().Header.Peek("Hx-Request")
	keyString := string(key)
	return len(keyString) > 0
}
