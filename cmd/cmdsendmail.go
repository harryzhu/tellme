/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	From      string
	To        string
	Cc        string
	Subject   string
	File      string
	Signature string
)

var (
	m *Mail
)

// sendmailCmd represents the sendmail command
var sendmailCmd = &cobra.Command{
	Use:   "sendmail",
	Short: "send mail anonymously or with plain-auth or with startTLS.",
	Long: `--from=abc@def.com --to=u1@test.com;u2@beta.com -cc=u1@sand.com;u2@alpha.com
	--subject="mail-title" --file="mail-content: the file-path of the mail-body" 
	--signature="append someting after the mail-content"
	`,
	PreRun: func(cmd *cobra.Command, args []string) {
		mBody, err := GetFileContent(File)
		if err != nil {
			fmt.Println(err)
			//panic("cannot read file:" + File)
		}
		m = NewMail(From, To, Cc).WithSubject(Subject).WithBody(string(mBody)).WithSignature(Signature)
		m.Compose()
	},
	Run: func(cmd *cobra.Command, args []string) {
		smtpaccess.Send(m)
	},
}

func init() {
	rootCmd.AddCommand(sendmailCmd)

	sendmailCmd.Flags().StringVar(&From, "from", "", "")
	sendmailCmd.Flags().StringVar(&To, "to", "", "")
	sendmailCmd.Flags().StringVar(&Cc, "cc", "", "")
	sendmailCmd.Flags().StringVar(&Subject, "subject", "", "")
	sendmailCmd.Flags().StringVar(&File, "file", "", "")
	sendmailCmd.Flags().StringVar(&Signature, "signature", "", "")

	sendmailCmd.MarkFlagRequired("from")
	sendmailCmd.MarkFlagRequired("to")
	sendmailCmd.MarkFlagRequired("subject")
	sendmailCmd.MarkFlagRequired("file")
}
