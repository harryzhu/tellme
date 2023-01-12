package cmd

import (
	//"crypto/tls"
	//"errors"
	"fmt"
	//"net/smtp"
	"strings"
)

type Mail struct {
	From      string
	To        string
	Cc        string
	Subject   string
	Body      string
	Signature string
	Message   string
}

func NewMail(from, to, cc string) *Mail {
	return &Mail{
		From: from,
		To:   to,
		Cc:   cc,
	}
}

func (m *Mail) WithSubject(s string) *Mail {
	m.Subject = s
	return m
}

func (m *Mail) WithBody(s string) *Mail {
	m.Body = s
	return m
}

func (m *Mail) WithSignature(s string) *Mail {
	m.Signature = s
	return m
}

func (m *Mail) Compose() *Mail {
	headers := make(map[string]string)
	headers["From"] = m.From
	headers["To"] = strings.ReplaceAll(m.To, ",", ";")
	if len(m.Cc) > 0 {
		headers["Cc"] = strings.ReplaceAll(m.Cc, ",", ";")
	}

	headers["Subject"] = m.Subject
	headers["Content-Type"] = "text/html"

	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	msg += "\r\n" + m.Body

	if m.Signature != "" {
		msg += "\r\n<hr/>" + m.Signature
	}

	m.Message = msg
	return m
}
