package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Digisata/todolist_app/controllers"
	"github.com/Digisata/todolist_app/models"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	db.AutoMigrate(&models.Activity{})
	db.AutoMigrate(&models.Todo{})

	activityController := controllers.NewActivityController(db)
	todoController := controllers.NewTodoController(db)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg": "devcode updated",
		})
	})

	app.Post("/activity-groups", activityController.Create)
	app.Get("/activity-groups", activityController.FindAll)
	app.Get("/activity-groups/:id", activityController.FindById)
	app.Patch("/activity-groups/:id", activityController.Update)
	app.Delete("/activity-groups/:id", activityController.Delete)

	app.Post("/todo-items", todoController.Create)
	app.Get("/todo-items", todoController.FindAll)
	app.Get("/todo-items/:id", todoController.FindById)
	app.Patch("/todo-items/:id", todoController.Update)
	app.Delete("/todo-items/:id", todoController.Delete)

	app.Listen(":3030")
}
