package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateStruct(data any) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.Field()
			element.Message = getErrorMsg(err)
			errors = append(errors, &element)
		}
	}

	return errors
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email format"
	case "min":
		return "value is too short"
	case "max":
		return "value is too long"
	}
	return "invalid value"
}
