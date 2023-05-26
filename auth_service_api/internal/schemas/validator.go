package schemas

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validatable interface {
	Validate() error
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("sorting", ValidateSorting)
}

func ValidateSorting(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\w+\[(asc|desc)\]$`)
	return re.MatchString(fl.Field().String())
}

func Validate(s Validatable) error {
	return validate.Struct(s)
}
