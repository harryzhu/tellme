package cmd

import (
	"errors"
	"net/smtp"
)

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
