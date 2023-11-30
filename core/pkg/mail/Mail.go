package mail

import (
	"core/driver"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/jonboulle/clockwork"
)

// SendMail Todo: 增加StartTLS认证
func SendMail(subject, body string) error {
	mailCfg := driver.Base.Mail

	// 初始化邮件格式
	contentType := "Content-Type: text/plain; charset=UTF-8"
	nickname := "FanTracer"
	header := fmt.Sprintf("To: %s\r\nFrom: %s<%s>\r\nSubject: %s\r\n%s\r\n\r\n",
		strings.Join([]string{mailCfg.To}, ","),
		nickname,
		mailCfg.Account,
		subject,
		contentType)

	var auth smtp.Auth
	if mailCfg.Key != "" {
		auth = smtp.PlainAuth("", mailCfg.Account, mailCfg.Key, mailCfg.Addr)
	}

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", mailCfg.Addr, mailCfg.Port),
		auth,
		mailCfg.Account,
		[]string{mailCfg.To},
		[]byte(header+body))

	return err
}

func InitNotify(c clockwork.Clock) error {
	var err error = nil

	// 未启用邮件服务
	if !(driver.Base.Mail.Interval > 0) {
		return err
	}

	ticker := c.NewTicker(time.Duration(driver.Base.Mail.Interval) * time.Hour)

	body := fmt.Sprintf("在" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "时依赖服务初始化完毕，可以正常开始工作")
	err = SendMail("FanTracer Info", body)
	if err != nil {
		return err
	}

	go func() {
		defer ticker.Stop()

		for range ticker.Chan() {
			body = "Fantracer正常运行。"
			err = SendMail("FanTracer Info", body)

			if err != nil {
				return
			}
		}
	}()
	return err
}

// starttls认证: https://gist.github.com/homme/22b457eb054a07e7b2fb

//type loginAuth struct {
//	username, password string
//}
//
//// LoginAuth 用于进行startTLS认证
//func LoginAuth(username, password string) smtp.Auth {
//	return &loginAuth{username, password}
//}
//
//func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
//	return "LOGIN", []byte(a.username), nil
//}
//
//func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
//	if more {
//		switch string(fromServer) {
//		case "Username:":
//			return []byte(a.username), nil
//		case "Password:":
//			return []byte(a.password), nil
//		default:
//			return nil, errors.New("unknown fromServer")
//		}
//	}
//	return nil, nil
//}
