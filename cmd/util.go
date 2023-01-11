package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
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

func GetNowUnix() int64 {
	return time.Now().Unix()
}
