/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	AccessKey  string
	smtpaccess *SmtpAccess
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tellme",
	Short: "",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		fmt.Println("pre")

		smtpaccess = NewSmtpAccess("", "", "", "", "", "")
		if AccessKey != "" {
			sa, err := smtpaccess.Unseal(AccessKey)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("using config: ", sa.Name)
			}

		}

	},
	Run: func(cmd *cobra.Command, args []string) {

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("post")
	},
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

	rootCmd.PersistentFlags().StringVar(&AccessKey, "accesskey", "", "accesskey")

}
