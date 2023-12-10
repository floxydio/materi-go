package controller

import (
	"belajarbe/database"
	"belajarbe/models"

	"github.com/gofiber/fiber/v2"
)

func TodoGet(c *fiber.Ctx) error {
	var todoModel []models.TodoModel

	// err := database.DBInit.Find(&todoModel).Error

	err := database.DBInit.Raw("SELECT * FROM todos").Scan(&todoModel).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Something went wrong",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   todoModel,
		"msg":    "Successfully GET Todo",
	})

}

func TodoPost(c *fiber.Ctx) error {
	var todoModel models.TodoModel

	err := c.BodyParser(&todoModel)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"err":     err.Error(),
			"message": "Something went wrong",
		})
	}

	errCreate := database.DBInit.Create(&todoModel).Error

	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"err":     errCreate.Error(),
			"message": "Something went wrong",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": 201,
		"msg":    "Sukses Input",
	})

}

func EditTodo(c *fiber.Ctx) error {
	var todoModel models.TodoModel
	id := c.Params("id")

	err := c.BodyParser(&todoModel)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"err":     err.Error(),
			"message": "Something went wrong",
		})
	}

	errEdit := database.DBInit.Where("id = ?", id).Updates(&todoModel).Error

	if errEdit != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"err":     errEdit.Error(),
			"message": "Something went wrong",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status": 201,
		"msg":    "Sukses Edit",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	var todoModel models.TodoModel

	id := c.Params("id")

	err := database.DBInit.Where("id = ?", id).Delete(&todoModel).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"err":     err.Error(),
			"message": "Something went wrong",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"msg":    "Sukses Hapus",
	})
}
