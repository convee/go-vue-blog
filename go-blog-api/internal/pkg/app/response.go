package app

import (
	"github.com/convee/go-blog-api/internal/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
	return
}

// Error setting gin.JSON
func (g *Gin) Error(errCode int, msg string, data interface{}) {
	if msg == "" {
		msg = e.GetMsg(errCode)
	}
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
