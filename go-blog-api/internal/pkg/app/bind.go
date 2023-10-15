package app

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"reflect"
	"strings"

	"github.com/convee/go-blog-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

func BindForm(c *gin.Context, form interface{}) []string {
	return Bind("form", c, form)
}

func BindJson(c *gin.Context, json interface{}) []string {
	return Bind("json", c, json)
}

func BindQuery(c *gin.Context, query interface{}) []string {
	return Bind("query", c, query)
}

func BindHeader(c *gin.Context, query interface{}) []string {
	return Bind("header", c, query)
}

func Bind(flag string, c *gin.Context, data interface{}) []string {
	var err error
	switch flag {
	case "json":
		err = c.ShouldBindJSON(data)
		break
	case "form":
		err = c.ShouldBind(data)
		break
	case "query":
		err = c.ShouldBindQuery(data)
		break
	case "header":
		err = c.ShouldBindHeader(data)
		break
	}
	if err != nil {
		logger.Error(err.Error())
		return []string{err.Error()}
	}

	// 创建翻译器
	zhTrans := zh.New() // 中文转换器
	enTrans := en.New() // 英文转换器

	uni := ut.New(zhTrans, zhTrans, enTrans) // 创建一个通用转换器

	curLocales := "zh"                        // 设置当前语言类型
	trans, _ := uni.GetTranslator(curLocales) // 获取对应语言的转换器

	validate := validator.New()
	switch curLocales {
	case "zh":
		// 内置tag注册 中文翻译器
		_ = zh_trans.RegisterDefaultTranslations(validate, trans)
		// 自定义tag注册 中文翻译器
		_ = validate.RegisterTranslation("unique_customer", trans, func(ut ut.Translator) error {
			return ut.Add("unique_customer", "客户已存在", false)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
	case "en":
		// 内置tag注册 英文翻译器
		_ = en_trans.RegisterDefaultTranslations(validate, trans)
	}

	// 注册 RegisterTagNameFunc, 标签显示中文
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err = validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var msg []string
		for _, e := range errs {
			msg = append(msg, fmt.Sprintf("%s", e.Translate(trans)))
		}
		return msg

	}
	return []string{}
}
