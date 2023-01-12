/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
)

var (
	Name     string
	Host     string
	Port     string
	Auth     string
	User     string
	Password string
)

// sealCmd represents the seal command
var sealCmd = &cobra.Command{
	Use:   "seal",
	Short: "for generating --accesskey",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		smtpaccess = NewSmtpAccess(Name, Host, Port, Auth, User, Password)
		smtpaccess.Seal()
	},
}

func init() {
	rootCmd.AddCommand(sealCmd)
	sealCmd.Flags().StringVar(&Name, "name", "smtp_conf_v1", " smtp description")
	sealCmd.Flags().StringVar(&Host, "host", "", " smtp Host")
	sealCmd.Flags().StringVar(&Port, "port", "25", " smtp Port")
	sealCmd.Flags().StringVar(&Auth, "auth", "", "smtp Auth: \"\"/plain/login")
	sealCmd.Flags().StringVar(&User, "user", "", " smtp User")
	sealCmd.Flags().StringVar(&Password, "password", "", " smtp Password")
}
