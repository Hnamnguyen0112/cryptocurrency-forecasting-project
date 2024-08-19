package main

import (
	"log"
	"time"

	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/config"
	"github.com/Hnamnguyen0112/cryptocurrency-forecasting-project/collector/pkg/database"
	"github.com/gofiber/fiber/v2"
)

const idleTimeout = 5 * time.Second

func main() {
  dbUser := config.Config("BLOCKCHAIN_DB_USER")
	dbPassword := config.Config("BLOCKCHAIN_DB_PASSWORD")
	dbHost := config.Config("BLOCKCHAIN_DB_HOST")
	dbPort := config.Config("BLOCKCHAIN_DB_PORT")
	dbName := config.Config("BLOCKCHAIN_DB_NAME")

  connectParams := database.ConnectParams{
    Host:     dbHost,
    Port:     dbPort,
    User:     dbUser,
    Password: dbPassword,
    Name:   dbName,
  }

  database.Connect(connectParams)

	app := fiber.New(fiber.Config{
		ReadBufferSize: 4096 * 2,
		Prefork:        true,
		IdleTimeout:    idleTimeout,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3002"))
}
