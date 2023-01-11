package cmd

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/smtp"

	//"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// MailSetup ...
type MailSetup struct {
	SMTPHost      string `toml:"smtp_host"`
	SMTPPort      string `toml:"smtp_port"`
	SMTPAccessKey string `toml:"smtp_access_key"`
	SMTPStartTLS  bool   `toml:"smtp_starttls"`
	SMTPUsername  string
	SMTPPassword  string

	MailFrom  string   `toml:"mail_from"`
	MailTo    []string `toml:"mail_to"`
	MailCc    []string `toml:"mail_cc"`
	MailBcc   []string `toml:"mail_bcc"`
	MailFile  string
	MailTitle string

	Message string
}

// NewMailSetup ...
func NewMailSetup() *MailSetup {
	return &MailSetup{}
}

// LoadMailConfig ...
func LoadMailConfig(f string) MailSetup {
	var mailSetup MailSetup
	if _, err := toml.DecodeFile(f, &mailSetup); err != nil {
		log.Fatal(err)
	}

	if mailSetup.SMTPAccessKey != "" {
		u, p, err := ParseAccessKey(mailSetup.SMTPAccessKey)

		if err != nil {
			log.Fatal(err)
		}

		mailSetup.SMTPUsername = u
		mailSetup.SMTPPassword = p

	}
	return mailSetup
}

// LoginAuth for starttls
type LoginAuth struct {
	username string
	password string
}

// NewLoginAuth required for starttls
func NewLoginAuth(username, password string) smtp.Auth {
	return &LoginAuth{username, password}
}

// Start required for starttls
func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

// Next required for starttls
func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
		}
	}
	return nil, nil
}

// SMTPSendMail for not starttls
func (ms *MailSetup) SMTPSendMail() error {
	if ms.SMTPHost == "" || ms.SMTPPort == "" {
		log.Fatal("SMTPHost and SMTPPort cannot be empty")
	}

	hostPort := strings.Join([]string{ms.SMTPHost, ms.SMTPPort}, ":")

	log.Println(hostPort)

	if err := ms.SetMailMessage(); err != nil {
		log.Println(err)
		return err
	}

	var err error
	if ms.SMTPUsername == "" && ms.SMTPPassword == "" {
		log.Println("using Anonymous ...")
		err = smtp.SendMail(hostPort, nil, ms.MailFrom, ms.MailTo, []byte(ms.Message))
	} else {
		log.Println("using PlainAuth ...")
		auth := smtp.PlainAuth("", ms.SMTPUsername, ms.SMTPPassword, ms.SMTPHost)
		err = smtp.SendMail(hostPort, auth, ms.MailFrom, ms.MailTo, []byte(ms.Message))
	}

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SetMailMessage ......
func (ms *MailSetup) SetMailMessage() error {
	headers := make(map[string]string)
	headers["From"] = ms.MailFrom
	headers["To"] = strings.Join(ms.MailTo, ";")
	if len(ms.MailCc) > 0 {
		headers["Cc"] = strings.Join(ms.MailCc, ";")
	}
	if len(ms.MailBcc) > 0 {
		headers["Bcc"] = strings.Join(ms.MailBcc, ";")
	}

	headers["Subject"] = ms.MailTitle
	headers["Content-Type"] = "text/html"

	msg := ""
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body, err := GetFileContent(ms.MailFile)
	if err != nil {
		return err
	}

	msg += "\r\n" + string(body)

	ms.Message = msg

	return nil
}

// SMTPSendMailStartTLS ...
func (ms *MailSetup) SMTPSendMailStartTLS() error {
	log.Println("using STARTTLS ...")
	if ms.SMTPHost == "" || ms.SMTPPort == "" {
		log.Fatal("SMTPHost and SMTPPort cannot be empty")
	}

	hostPort := strings.Join([]string{ms.SMTPHost, ms.SMTPPort}, ":")

	log.Println(hostPort)

	smtpClient, err := smtp.Dial(hostPort)
	if err != nil {
		log.Println(err)
		return err
	}
	defer smtpClient.Close()

	if ok, _ := smtpClient.Extension("STARTTLS"); ok {
		cfg := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         ms.SMTPHost}
		if err := smtpClient.StartTLS(cfg); err != nil {
			log.Println(err)
			return err
		}
	}

	a := NewLoginAuth(ms.SMTPUsername, ms.SMTPPassword)
	if ok, _ := smtpClient.Extension("AUTH"); ok {
		if err := smtpClient.Auth(a); err != nil {
			log.Println(err)
			return err
		}
	}

	if err := smtpClient.Mail(ms.MailFrom); err != nil {
		log.Println(err)
		return err
	}

	for _, addr := range ms.MailTo {
		if strings.Index(addr, "@") < 0 {
			continue
		}
		if err := smtpClient.Rcpt(addr); err != nil {
			log.Println(err)
			return err
		}
	}

	if err := ms.SetMailMessage(); err != nil {
		log.Println(err)
		return err
	}

	w, err := smtpClient.Data()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = w.Write([]byte(ms.Message))
	if err != nil {
		log.Println(err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	smtpClient.Quit()
	return nil
}
