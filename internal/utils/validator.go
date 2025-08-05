package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequest(c *gin.Context, req interface{}) (bool, []ValidationError) {
	if err := c.ShouldBindJSON(req); err != nil {
		var validationErrors []ValidationError

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrs {
				validationErrors = append(validationErrors, ValidationError{
					Field:   e.Field(),
					Message: getErrorMessage(e),
				})
			}
		} else {
			validationErrors = append(validationErrors, ValidationError{
				Field:   "request",
				Message: err.Error(),
			})
		}

		return false, validationErrors
	}

	return true, nil
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return "Nilai terlalu pendek"
	case "max":
		return "Nilai terlalu panjang"
	default:
		return "Validasi gagal pada field ini"
	}
}