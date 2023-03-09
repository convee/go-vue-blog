package ding

import (
	"github.com/convee/go-vue-blog/configs"
	"testing"
)

func Test_Ding(t *testing.T) {
	configs.Conf.App.Env = "DEV"
	SendAlert("Mpos-go", "钉钉预警测试", false)
}
