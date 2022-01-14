/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"
	"log"

	"strings"

	"github.com/harryzhu/util"
	"github.com/spf13/cobra"
)

var (
	MailTitle string
	MailFile  string
	MailFrom  string
	MailTo    string
	MailCc    string
	MailBcc   string
)

// gossipCmd represents the gossip command
var gossipCmd = &cobra.Command{
	Use:   "gossip",
	Short: "A brief description of your command",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("gossip")

		//ak, _ := util.GetAccessKey("email", "pass")
		//log.Println(ak)

		u, p, err := util.ParseAccessKey(AccessKey)
		if err != nil {
			log.Println(err)
		}

		//log.Println(u, p)

		mc := &util.MailConfig
		mc.SmtpHost = SmtpHost
		mc.SmtpPort = SmtpPort
		mc.SmtpUsername = u
		mc.SmtpPassword = p

		mc.MailTitle = MailTitle
		mc.MailFile = MailFile

		mc.MailFrom = u
		mc.MailTo = strings.Split(strings.ReplaceAll(MailTo, ",", ";"), ";")
		mc.MailCc = strings.Split(strings.ReplaceAll(MailCc, ",", ";"), ";")
		mc.MailBcc = strings.Split(strings.ReplaceAll(MailBcc, ",", ";"), ";")

		mc.SmtpSendMailStartTLS()

	},
}

func init() {
	rootCmd.AddCommand(gossipCmd)

	gossipCmd.Flags().StringVar(&MailTitle, "title", "", "mail title")
	gossipCmd.Flags().StringVar(&MailFile, "file", "", "mail content from the file")
	gossipCmd.Flags().StringVar(&MailTo, "to", "", "to address(es),split by semicolon(;), add quotation-marks if the address includes dash(-).")
	gossipCmd.Flags().StringVar(&MailCc, "cc", "", "cc address(es),split by semicolon(;), add quotation-marks if the address includes dash(-).")
	gossipCmd.Flags().StringVar(&MailBcc, "bcc", "", "bcc address(es),split by semicolon(;), add quotation-marks if the address includes dash(-).")

	gossipCmd.MarkFlagRequired("title")
	gossipCmd.MarkFlagRequired("file")
	gossipCmd.MarkFlagRequired("to")
}
