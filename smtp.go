package smtp

import (
	"fmt"
	"net/smtp"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/smtp", new(SMTP))
}

type SMTP struct{}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func plainAuth(host string, password string, fromMail string) smtp.Auth {
	return smtp.PlainAuth("", fromMail, password, host)
}

func (*SMTP) SendMail(host string, port string, fromMail string, password string, udwList []string, message string) {
	body := []byte(message)
	auth := plainAuth(host, password, fromMail)
	err := smtp.SendMail(host+":"+port, auth, fromMail, udwList, body)
	check(err)
}
