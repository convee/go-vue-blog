package tests

import (
	"github.com/convee/go-blog-api/pkg/utils"
	"testing"
)

func TestPassword(t *testing.T) {
	password := utils.GenPassword("yyds", "yyds")
	t.Log(password)
}
