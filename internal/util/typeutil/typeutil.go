package typeutil

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

// HttpType
const (
	ContentType         = "Content-Type"
	ContentTypeJson     = "application/json"
	ContentTypeJsonUTF8 = "application/json charset=utf8"
)

var (
	validateUtil *validator.Validate
	once         sync.Once
)

func Validator() *validator.Validate {
	once.Do(func() {
		validateUtil = validator.New()
	})
	return validateUtil
}
