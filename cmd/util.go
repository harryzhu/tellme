package cmd

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type SmtpAccess struct {
	Name     string
	Host     string
	Port     string
	Auth     string
	User     string
	Password string
}

func NewSmtpAccess(name, host, port, auth, user, password string) *SmtpAccess {
	return &SmtpAccess{
		Name:     name,
		Host:     host,
		Port:     port,
		Auth:     auth,
		User:     user,
		Password: password,
	}

}

func (sa *SmtpAccess) Seal() (string, error) {
	data, err := json.Marshal(sa)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("ERROR: cannot marshal smtpAccess")
	}
	fmt.Println(string(data))
	pwd, err := AESEncode(data)
	if err != nil {
		return "", errors.New("ERROR: cannot encode smtpAccess")
	}
	fmt.Printf(" --accesskey=\"%v\"\n", pwd)
	return string(data), nil
}

func (sa *SmtpAccess) Unseal(pwd string) (*SmtpAccess, error) {
	plain, err := AESDecode(pwd)
	if err != nil {
		fmt.Println(err)
		return sa, errors.New("ERROR: cannot Unseal smtpAccess")
	}
	//fmt.Println(string(plain))

	err = json.Unmarshal(plain, sa)
	if err != nil {
		fmt.Println(err)
		return sa, errors.New("ERROR: cannot Unmarshal smtpAccess")
	}

	return sa, nil
}

func (sa *SmtpAccess) Send(m *Mail) error {
	if sa.Auth == "" && sa.User == "" && sa.Password == "" {
		return sa.SendWithAnonymous(m)
	}

	if sa.Auth == "plain" {
		return sa.SendWithPlain(m)
	}

	if sa.Auth == "login" {
		return sa.SendWithStartTLS(m)
	}

	return nil
}

func (sa *SmtpAccess) SendWithAnonymous(m *Mail) error {
	hostPort := strings.Join([]string{sa.Host, sa.Port}, ":")
	Tos := strings.Split(m.To, ";")
	err := smtp.SendMail(hostPort, nil, m.From, Tos, []byte(m.Message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("SendWithAnonymous ...")
	return nil
}

func (sa *SmtpAccess) SendWithPlain(m *Mail) error {
	hostPort := strings.Join([]string{sa.Host, sa.Port}, ":")
	Tos := strings.Split(m.To, ";")
	auth := smtp.PlainAuth("", sa.User, sa.Password, sa.Host)
	err := smtp.SendMail(hostPort, auth, m.From, Tos, []byte(m.Message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("SendWithPlain ...")
	return nil
}

func (sa *SmtpAccess) SendWithStartTLS(m *Mail) error {
	hostPort := strings.Join([]string{sa.Host, sa.Port}, ":")
	Tos := strings.Split(m.To, ";")

	smtpClient, err := smtp.Dial(hostPort)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer smtpClient.Close()

	fmt.Println("SendWithStartTLS ...")

	if ok, _ := smtpClient.Extension("STARTTLS"); ok {
		cfg := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         sa.Host}
		if err := smtpClient.StartTLS(cfg); err != nil {
			fmt.Println(err)
			return err
		}
	}

	a := NewLoginAuth(sa.User, sa.Password)
	if ok, _ := smtpClient.Extension("AUTH"); ok {
		if err := smtpClient.Auth(a); err != nil {
			fmt.Println(err)
			return err
		}
	}

	if err := smtpClient.Mail(m.From); err != nil {
		fmt.Println(err)
		return err
	}

	for _, addr := range Tos {
		if strings.Index(addr, "@") < 0 {
			continue
		}
		if err := smtpClient.Rcpt(addr); err != nil {
			fmt.Println(err)
			return err
		}
	}

	w, err := smtpClient.Data()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = w.Write([]byte(m.Message))
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = smtpClient.Quit()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetNowUnix() int64 {
	return time.Now().Unix()
}

func GetFileContent(src string) (body []byte, err error) {
	_, err = os.Stat(src)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func LoadAccessKey() error {
	smtpaccess = NewSmtpAccess("", "", "", "", "", "")
	ak := ""

	if AccessKey != "" {
		ak = AccessKey
		fmt.Println("using accesskey inline")
	} else {
		AccessKeyEnv := GetEnv("TELLMEACCESSKEY", "")
		if AccessKeyEnv != "" {
			ak = AccessKeyEnv
			fmt.Printf("using accesskey from env variable --accesskey=%v\n", ak)
		}
	}

	sa, err := smtpaccess.Unseal(ak)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("using config: ", sa.Name)
	}
	return nil
}
