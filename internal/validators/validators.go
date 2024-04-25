package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateGoodTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Good")
}
