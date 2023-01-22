package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Digisata/todolist_app/helpers"
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
	body := new(models.ActivityRequest)

	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	errors := helpers.ValidateStruct(*body)
	if errors != nil {
		msg, _ := json.Marshal(errors)
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: string(msg),
		})
		return
	}

	activity := models.Activity{Title: body.Title, Email: body.Email}
	controller.DB.Create(&activity)

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (controller *activityController) FindAll(c *fiber.Ctx) {
	var activities []models.Activity
	if err := controller.DB.Find(&activities).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

func (controller *activityController) FindById(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	if err := controller.DB.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (controller *activityController) Update(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	if err := controller.DB.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	body := new(models.ActivityRequest)

	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
		})
		return
	}

	errors := helpers.ValidateStruct(*body)
	if errors != nil {
		msg, _ := json.Marshal(errors)
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: string(msg),
		})
		return
	}

	if err := controller.DB.Model(&activity).Updates(models.Activity{
		Title: body.Title,
		Email: body.Email,
	}).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	})
}

func (controller *activityController) Delete(c *fiber.Ctx) {
	id := c.Params("id")

	var activity models.Activity
	if err := controller.DB.First(&activity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if err := controller.DB.Delete(&activity).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
	})
}
