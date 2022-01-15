package cmd

import (
	"os"

	"github.com/harryzhu/util"
	"github.com/spf13/cobra"
)

const (
	SmtpHost string = "smtp.office365.com"
	SmtpPort string = "587"
)

var (
	AccessKey string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tellops",
	Short: "A brief description of your application",
	Long:  `-`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tellops.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	AccessKey = util.GetEnv("TELLOPSACCESSKEY", "yiKSLz4ujLzPmJQsLf2kCTaI2HXlz61GBLkJZN2GDRM/xvXQIrCV4oMKDYweKfhj")
}