package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"mecnave.com/mod/crud_manager/cmd/routes"
)

const port = ":8081"

func StartAPI() {

	app := fiber.New()
	routes.SetupRoutes(app)
	err := app.Listen(port)
	fmt.Println("App Listen")

	if err != nil {
		panic(err)
	}
}
