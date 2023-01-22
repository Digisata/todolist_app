package controllers

import (
	"errors"
	"fmt"

	"github.com/Digisata/todolist_app/models"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type todoController struct {
	DB *gorm.DB
}

func NewTodoController(DB *gorm.DB) *todoController {
	return &todoController{
		DB: DB,
	}
}

func (controller *todoController) Create(c *fiber.Ctx) {
	var body struct {
		ActivityGroupId uint   `json:"activity_group_id"`
		Title           string `json:"title"`
		IsActive        bool   `json:"is_active"`
		Priority        string `json:"priority"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if body.Title == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	}

	if body.ActivityGroupId == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "activity_group_id cannot be null",
		})
		return
	}

	todo := models.Todo{ActivityGroupId: body.ActivityGroupId, Title: body.Title, IsActive: body.IsActive, Priority: body.Priority}
	err := controller.DB.Create(&todo).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func (controller *todoController) FindAll(c *fiber.Ctx) {
	var todos []models.Todo
	activityGroupId := c.Query("activity_group_id")

	var err error
	if activityGroupId == "" {
		err = controller.DB.Find(&todos).Error
	} else {
		err = controller.DB.Find(&todos, "activity_group_id = ?", activityGroupId).Error
	}

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    todos,
	})
}

func (controller *todoController) FindById(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	err := controller.DB.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %s Not Found", id),
			})
			return
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func (controller *todoController) Update(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	err := controller.DB.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %s Not Found", id),
			})
			return
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	var body struct {
		ActivityGroupId uint   `json:"activity_group_id"`
		Title           string `json:"title"`
		IsActive        bool   `json:"is_active"`
		Priority        string `json:"priority"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	err = controller.DB.Model(&todo).Updates(map[string]interface{}{
		"activity_group_id": body.ActivityGroupId,
		"title":             body.Title,
		"is_active":         body.IsActive,
		"priority":          body.Priority,
	}).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
}

func (controller *todoController) Delete(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	err := controller.DB.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Todo with ID %s Not Found", id),
			})
			return
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	err = controller.DB.Delete(&todo).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
	})
}
