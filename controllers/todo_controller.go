package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Digisata/todolist_app/helpers"
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
	body := new(models.CreateTodoRequest)

	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    map[string]interface{}{},
		})
		return
	}

	errors := helpers.ValidateStruct(*body)
	if errors != nil {
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: fmt.Sprintf("%s cannot be null", errors[0].FailedField),
			Data:    map[string]interface{}{},
		})
		return
	}

	todo := models.Todo{ActivityGroupID: body.ActivityGroupID, Title: body.Title, IsActive: body.IsActive, Priority: body.Priority}
	controller.DB.Create(&todo)

	c.Status(http.StatusCreated).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
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
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

func (controller *todoController) FindById(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	if err := controller.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Todo with ID %s Not Found", id),
				Data:    map[string]interface{}{},
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
		Data:    todo,
	})
}

func (controller *todoController) Update(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	if err := controller.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Todo with ID %s Not Found", id),
				Data:    map[string]interface{}{},
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	body := new(models.UpdateTodoRequest)

	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: err.Error(),
			Data:    map[string]interface{}{},
		})
		return
	}

	if body.Title != "" {
		todo.Title = body.Title
	}

	if body.Priority != "" {
		todo.Priority = body.Priority
	}

	todo.IsActive = body.IsActive

	if err := controller.DB.Model(&todo).Updates(todo).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	})
}

func (controller *todoController) Delete(c *fiber.Ctx) {
	id := c.Params("id")

	var todo models.Todo
	if err := controller.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(models.BaseResponse{
				Status:  http.StatusText(http.StatusNotFound),
				Message: fmt.Sprintf("Todo with ID %s Not Found", id),
				Data:    map[string]interface{}{},
			})
			return
		}

		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if err := controller.DB.Delete(&todo).Error; err != nil {
		c.Status(http.StatusInternalServerError).JSON(models.BaseResponse{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK).JSON(models.BaseResponse{
		Status:  "Success",
		Message: "Success",
		Data:    map[string]interface{}{},
	})
}
