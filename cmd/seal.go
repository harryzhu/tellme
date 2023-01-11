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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	sealCmd.Flags().StringVar(&Auth, "auth", "plain", "smtp Auth: plain/login")
	sealCmd.Flags().StringVar(&User, "user", "", " smtp User")
	sealCmd.Flags().StringVar(&Password, "password", "", " smtp Password")
}
