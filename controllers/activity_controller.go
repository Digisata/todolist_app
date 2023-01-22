package controllers

import (
	"errors"
	"fmt"

	"github.com/Digisata/todolist_app/models"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type activityController struct {
	DB *gorm.DB
}

func NewActivityController(DB *gorm.DB) *activityController {
	return &activityController{
		DB: DB,
	}
}

func (controller *activityController) Create(c *fiber.Ctx) {
	var body struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}

	c.BodyParser(&body)

	if body.Title == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	}

	activity := models.Activity{Title: body.Title, Email: body.Email}
	result := controller.DB.Create(&activity)

	if result.Error != nil {
		c.Status(fiber.StatusBadRequest)
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func (controller *activityController) FindAll(c *fiber.Ctx) {
	var activities []models.Activity
	controller.DB.Find(&activities)

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activities,
	})
}

func (controller *activityController) FindById(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func (controller *activityController) Update(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	var body struct {
		Title string `json:"title"`
		Email string `json:"email"`
	}

	c.BodyParser(&body)

	result := controller.DB.Model(&activity).Updates(models.Activity{
		Title: body.Title,
		Email: body.Email,
	})

	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": result.Error.Error(),
		})
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    activity,
	})
}

func (controller *activityController) Delete(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %s Not Found", id),
		})
		return
	}

	controller.DB.Delete(&activity)

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success",
	})
}
