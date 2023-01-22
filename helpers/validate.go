package helpers

import (
	"github.com/Digisata/todolist_app/models"
	"github.com/go-playground/validator"
)

func ValidateStruct(data interface{}) []*models.ErrorResponse {
	Fields := map[string]string{
		"Title":           "title",
		"ActivityGroupID": "activity_group_id",
	}
	var errors []*models.ErrorResponse
	err := validator.New().Struct(data)
	if err != nil {
		for _, errValidate := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = Fields[errValidate.Field()]
			element.Tag = errValidate.Tag()
			element.Value = errValidate.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
