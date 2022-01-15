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
	Short: "send mail",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		u, p, err := util.ParseAccessKey(AccessKey)
		if err != nil {
			log.Println(err)
		}

		mc := util.NewMailSetup()
		mc.SMTPHost = SMTPHost
		mc.SMTPPort = SMTPPort
		mc.SMTPUsername = u
		mc.SMTPPassword = p

		mc.MailTitle = MailTitle
		mc.MailFile = MailFile

		mc.MailFrom = u
		mc.MailTo = strings.Split(strings.ReplaceAll(MailTo, ",", ";"), ";")
		mc.MailCc = strings.Split(strings.ReplaceAll(MailCc, ",", ";"), ";")
		mc.MailBcc = strings.Split(strings.ReplaceAll(MailBcc, ",", ";"), ";")

		mc.SMTPSendMailStartTLS()

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
