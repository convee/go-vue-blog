package middleware

import (
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/gin-gonic/gin"
)

type AuthBind struct {
	Appid     string `header:"appid" validate:"required"`
	Secret    string `header:"secret" validate:"required"`
	Sign      string `header:"sign" validate:"required"`
	Timestamp int    `header:"timestamp" validate:"required"`
}

func ApiAuth() (g gin.HandlerFunc) {
	return func(c *gin.Context) {
		var (
			ag       = app.Gin{C: c}
			authBind AuthBind
		)
		validateErr := app.BindHeader(c, &authBind)
		if len(validateErr) > 0 {
			ag.Error(e.INVALID_PARAMS, "", validateErr)
			c.Abort()
			return
		}

		// todo
		c.Next()
		return
	}
}
