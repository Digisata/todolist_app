package helpers

import (
	"github.com/Digisata/todolist_app/models"
	"github.com/go-playground/validator"
)

func ValidateStruct(data interface{}) []*models.ErrorResponse {
    var errors []*models.ErrorResponse
    err := validator.New().Struct(data)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var element models.ErrorResponse
            element.FailedField = err.StructNamespace()
            element.Tag = err.Tag()
            element.Value = err.Param()
            errors = append(errors, &element)
        }
    }
    return errors
}