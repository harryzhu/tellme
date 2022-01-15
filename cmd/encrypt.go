/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/harryzhu/util"
	"github.com/spf13/cobra"
)

var (
	MailUsername string
	MailPassword string
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "generate accesskey for sub-command: gossip.",
	Long:  `export TELLOPSACCESSKEY to specify your accesskey.`,
	Run: func(cmd *cobra.Command, args []string) {
		ak, err := util.GetAccessKey(MailUsername, MailPassword)
		if err != nil {
			log.Println(err)
		}
		log.Printf("TELLOPSACCESSKEY=\"%s\"", ak)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().StringVar(&MailUsername, "username", "", "--username=\"smtp-email-address\"")
	encryptCmd.Flags().StringVar(&MailPassword, "password", "", "--username=\"smtp-password\"")

	encryptCmd.MarkFlagRequired("username")
	encryptCmd.MarkFlagRequired("password")
}
