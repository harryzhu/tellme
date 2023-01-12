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
	Short: "for generating --accesskey=, encrypt config details for security.",
	Long: `--name="" --host="" --port="" --auth="" --user="" --password=""
	--port default is 25, --auth default is None, alse support "plain" and "login",
	i.e. use: seal --host="127.0.0.1" --name="smtp-mypc-localhost-25"
	you will get:
	--accesskey="Zt5BmhsR9C8C029xblTUkhR0JJazlduCVdKZIIX2aPqN7HQzd9/Bq4HPp6qU0DkBQ7H243gI2akWljyv1Jpsh6sAg/4vTE+IpMbSTV6LkC+IVR2zS1B6z+XOWqkBrSNOj/4hm0DedQPxGcZ434LVbQ=="
	`,
	Run: func(cmd *cobra.Command, args []string) {
		smtpaccess = NewSmtpAccess(Name, Host, Port, Auth, User, Password)
		smtpaccess.Seal()
	},
}

func init() {
	rootCmd.AddCommand(sealCmd)
	sealCmd.Flags().StringVar(&Name, "name", "smtp_conf_v1", "config description")
	sealCmd.Flags().StringVar(&Host, "host", "", "smtp host address")
	sealCmd.Flags().StringVar(&Port, "port", "25", "smtp port")
	sealCmd.Flags().StringVar(&Auth, "auth", "", "smtp auth: \"\"/\"plain\"/\"login\"")
	sealCmd.Flags().StringVar(&User, "user", "", "smtp username")
	sealCmd.Flags().StringVar(&Password, "password", "", "smtp password")

	sealCmd.MarkFlagRequired("host")
}
