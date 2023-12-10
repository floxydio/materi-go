package main

import (
	"belajarbe/controller"
	"belajarbe/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/todo", controller.TodoGet)
	app.Post("/create-todo", controller.TodoPost)
	app.Put("/edit-todo/:id", controller.EditTodo)
	app.Delete("/delete-todo/:id", controller.DeleteTodo)

	app.Listen(":3000")
}
