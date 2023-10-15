package app

import (
	"github.com/convee/go-blog-api/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// MarkErrors logs error logs
func MarkErrors(errors validator.ValidationErrors) {
	for _, err := range errors {
		logger.Error(err.Error())
	}
}
