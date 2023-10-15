package middleware

import (
	"fmt"
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/pkg/ding"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// Recovery panic异常捕获，钉钉告警
func Recovery() func(c *gin.Context) {
	return gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		var appG = app.Gin{C: c}
		appG.Error(e.ERROR, "服务器繁忙，请重试", nil)
		ding.SendAlert("panic...", fmt.Sprintf("err:%v;stack:%s", err, string(debug.Stack())), false)
	})
}
