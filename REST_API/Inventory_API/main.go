package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"Inventory_APIs/database"
	"Inventory_APIs/routes"
	"Inventory_APIs/utils"
)

const DEFAULT_PORT = "3000"

func NewFiberApp() *fiber.App {
	var app *fiber.App = fiber.New()

	routes.SetupRoutes(app)

	return app
}

func main() {
	var app *fiber.App = NewFiberApp()

	database.InitDatabase(utils.GetValue("DB_NAME"))

	var PORT string = os.Getenv("PORT")

	if PORT == "" {
		PORT = DEFAULT_PORT
	}

	app.Listen(fmt.Sprintf(":%s", PORT))
}