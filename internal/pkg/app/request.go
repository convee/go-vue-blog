package app

import (
	"github.com/convee/go-vue-blog/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// MarkErrors logs error logs
func MarkErrors(errors validator.ValidationErrors) {
	for _, err := range errors {
		logger.Error(err.Error())
	}
}
