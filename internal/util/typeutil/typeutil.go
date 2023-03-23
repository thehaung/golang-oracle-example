package typeutil

import "github.com/go-playground/validator/v10"

// HttpType
const (
	ContentType         = "Content-Type"
	ContentTypeJson     = "application/json"
	ContentTypeJsonUTF8 = "application/json charset=utf8"
)

func Validator() *validator.Validate {
	validate := validator.New()
	return validate
}
