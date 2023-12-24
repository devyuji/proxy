package main

import (
	"fmt"
	"os"

	"github.com/devyuji/proxy/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("unable to find .env file")
	}
}

func main() {
	app := fiber.New()

	// middleware
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		},
	))

	// routes
	app.Get("/", routes.Proxy)
	app.Get("/healthcheck", routes.Healthcheck)

	// app running on port
	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%s", port))
}
