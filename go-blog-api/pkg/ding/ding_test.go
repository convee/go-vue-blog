package ding

import (
	"github.com/convee/go-blog-api/configs"
	"testing"
)

func Test_Ding(t *testing.T) {
	configs.Conf.App.Env = "DEV"
	SendAlert("blog", "钉钉预警测试", false)
}
