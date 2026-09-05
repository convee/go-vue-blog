package ding

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/convee/go-blog-api/configs"
	"github.com/convee/go-blog-api/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"time"
)

//  接入文档：https://open.dingtalk.com/document/group/custom-robot-access

// AI-LOCK: 机器人 Uri 与 Secret 是凭证，只能来自配置，禁止写回常量

/**
 * 发送钉钉报警
 * token：报警机器人的token
 * content：报警内容
 * all：true：at所有人
 */

// SendAlert 发送钉钉报警
func SendAlert(title string, content interface{}, all bool) {
	cont, _ := jsoniter.Marshal(content)
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
	if configs.Conf.Ding.Uri == "" {
		return
	}
	dingUrl := makeDingUrl()
	_, err = http.Post(dingUrl, "application/json", bytes.NewBuffer(bytePayload))
	if err != nil {
		logger.GetLogger().Error("ding error", zap.Error(err))
	}
}

func makeDingUrl() string {
	secret := configs.Conf.Ding.Secret
	timestamp := time.Now().UnixNano() / 1000000
	signStr := fmt.Sprintf("%d\n%s", timestamp, secret)

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(signStr))
	sum := hash.Sum(nil)

	encode := base64.StdEncoding.EncodeToString(sum)
	urlEncode := url.QueryEscape(encode)
	return fmt.Sprintf("%s&timestamp=%d&sign=%s", configs.Conf.Ding.Uri, timestamp, urlEncode)

}
