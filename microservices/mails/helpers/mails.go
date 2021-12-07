package helpers

import (
	gomail "gopkg.in/gomail.v2"
)


func SendMail(email, body string, attach *string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From",Env.MAIL_FROM)
	m.SetHeader("To",email)
	m.SetHeader("Cc",Env.MAIL_CC)
	m.SetBody("text/html", body)

	if attach != nil {
			m.Attach(*attach)
	}

	d := gomail.NewDialer(Env.SMTP_HOST, Env.SMTP_PORT, Env.SMTP_USER, Env.SMTP_PASSWORD)

	err = d.DialAndSend(m)

	return
}


