package validators

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func ValidateGoodTitle(field validator.FieldLevel) bool {
	title := field.Field().String()
	firstRune := rune(title[0])
	return unicode.IsLetter(firstRune) && (len(title) != 0)
}
