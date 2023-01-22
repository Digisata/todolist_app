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

	activity := models.Activity{Title: body.Title, Email: body.Email}
	err := controller.DB.Create(&activity).Error
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
		"data":    activity,
	})
}

func (controller *activityController) FindAll(c *fiber.Ctx) {
	var activities []models.Activity
	err := controller.DB.Find(&activities).Error
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
		"data":    activities,
	})
}

func (controller *activityController) FindById(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %s Not Found", id),
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
		"data":    activity,
	})
}

func (controller *activityController) Update(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %s Not Found", id),
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
		Title string `json:"title"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	err = controller.DB.Model(&activity).Updates(models.Activity{
		Title: body.Title,
		Email: body.Email,
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
		"data":    activity,
	})
}

func (controller *activityController) Delete(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	err := controller.DB.First(&activity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(fiber.StatusOK).JSON(fiber.Map{
				"status":  "Not Found",
				"message": fmt.Sprintf("Activity with ID %s Not Found", id),
			})
			return
		}

		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	err = controller.DB.Delete(&activity).Error
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
