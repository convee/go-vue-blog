package ding

import (
	"bytes"
	"github.com/convee/go-vue-blog/configs"
	"github.com/convee/go-vue-blog/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	dingTalkHost = "https://oapi.dingtalk.com"
	// todo 接入文档：https://open.dingtalk.com/document/group/custom-robot-access
	Token = "9a5a9ca03ff32415bd9884ac1dc1a92d5395b523c7ae6fcba1be25d36fec4cdf"
)

/**
 * 发送钉钉报警
 * token：报警机器人的token
 * content：报警内容
 * all：true：at所有人
 */

// SendAlert 发送钉钉报警
func SendAlert(title string, content interface{}, all bool) {
	cont, _ := jsoniter.Marshal(content)
	dingUrl := dingTalkHost + "/robot/send?access_token=" + Token
	data := make(map[string]interface{})

	data["msgtype"] = "text"
	data["text"] = map[string]string{
		"content": "时间：【" + time.Now().Format("2006-01-02 15:04:05") + "】\n事件：【" + configs.Conf.App.Env + "】" + title + "\n详情：" + string(cont),
	}
	data["at"] = map[string]interface{}{"atMobiles": [0]string{}, "isAtAll": all}

	bytePayload, err := jsoniter.Marshal(data)
	if err != nil {
		logger.GetLogger().Error("ding error", zap.Error(err))
		return
	}
	_, err = http.Post(dingUrl, "application/json", bytes.NewBuffer(bytePayload))
	if err != nil {
		logger.GetLogger().Error("ding error", zap.Error(err))
	}

}
