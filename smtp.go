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

type options struct {
	Subject string   `js:"subject"`
	Message string   `js:"message"`
	UDW     []string `js:"udw"`
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func plainAuth(host string, password string, sender string) smtp.Auth {
	return smtp.PlainAuth("", sender, password, host)
}

func (*SMTP) SendMail(host string, port string, sender string, password string, recipient string, options options) {
	emailMessage := "From: " + sender + "\r\n" + "To: " + recipient + "\r\n"

	if options.Subject != "" {
		emailMessage += "Subject: " + options.Subject + "\r\n\r\n"
	}

	if options.Message != "" {
		emailMessage += options.Message
	}

	if len(options.UDW) == 0 {
		options.UDW = []string{recipient}
	}

	body := []byte(emailMessage)
	auth := plainAuth(host, password, sender)
	err := smtp.SendMail(host+":"+port, auth, sender, options.UDW, body)
	check(err)
}
