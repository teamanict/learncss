package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := create()
	run(app)
}

func create() *fiber.App {
	// prepare server
	app := fiber.New()
	return app
}

func pages(app *fiber.App)  {
	app.Static("/", "./pages")
}

func run(app *fiber.App) {
	// run server
	pages(app)
	log.Fatal(app.Listen(":3000"))
}