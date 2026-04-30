package handler

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		tag := fld.Tag.Get("json")
		if tag == "" || tag == "-" {
			return fld.Name
		}

		name, _, _ := strings.Cut(tag, ",")
		if name == "" {
			return fld.Name
		}
		return name
	})
}

func validateStruct(s interface{}) map[string]string {
	errs := make(map[string]string)
	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = humanMessage(e)
		}
	}
	return errs
}

func humanMessage(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "numeric":
		return fmt.Sprintf("%s must be a number", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", field)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, e.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, e.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, e.Param())
	case "gtfield":
		return fmt.Sprintf("%s must be after %s", field, e.Param())
	case "ltfield":
		return fmt.Sprintf("%s must be before %s", field, e.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}
