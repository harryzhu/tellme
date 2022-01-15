package cmd

import (
	"os"

	"github.com/harryzhu/util"
	"github.com/spf13/cobra"
)

var (
	SmtpHost  string
	SmtpPort  string
	AccessKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tellops",
	Short: "general send-mail tool",
	Long:  `-`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	SmtpHost = util.GetEnv("TELLOPSSMTPHOST", "smtp.office365.com")
	SmtpPort = util.GetEnv("TELLOPSSMTPPORT", "587")
	AccessKey = util.GetEnv("TELLOPSACCESSKEY", "yiKSLz4ujLzPmJQsLf2kCTaI2HXlz61GBLkJZN2GDRM/xvXQIrCV4oMKDYweKfhj")
}
