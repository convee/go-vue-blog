package mail

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost     string
	MailPort     int
	MailUser     string   // 发件人
	MailUserName string   // 发件人名称
	MailPass     string   // 发件人密码
	MailTo       string   // 收件人 多个用,分割
	MailCc       string   // 抄送人 多个用,分割
	Subject      string   // 邮件主题
	Body         string   // 邮件内容
	Attachment   []string // 邮件附件
}

func Send(o *Options) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", m.FormatAddress(o.MailUser, o.MailUserName))
	//设置发送给多个用户
	mailArrTo := strings.Split(o.MailTo, ",")
	mailArrCc := strings.Split(o.MailCc, ",")
	m.SetHeader("To", mailArrTo...)
	m.SetHeader("Cc", mailArrCc...)

	//设置邮件主题
	m.SetHeader("Subject", o.Subject)

	if len(o.Attachment) > 0 {
		for _, file := range o.Attachment {
			name := filepath.Base(file)
			m.Attach(file, gomail.Rename(name),
				gomail.SetHeader(map[string][]string{
					"Content-Disposition": {
						fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", name)),
					},
				}),
			)

		}
	}

	//设置邮件正文
	m.SetBody("text/html", o.Body)

	d := gomail.NewDialer(o.MailHost, o.MailPort, o.MailUser, o.MailPass)

	return d.DialAndSend(m)
}
