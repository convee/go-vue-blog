package tests

import (
	"github.com/convee/go-vue-blog/pkg/utils"
	"testing"
)

func TestPassword(t *testing.T) {
	password := utils.GenPassword("yyds", "yyds")
	t.Log(password)
}
