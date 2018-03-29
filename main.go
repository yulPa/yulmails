package main

import (
	"github.com/spf13/cobra"

	"github.com/yulPa/yulmails/api"
)

var certFile string
var keyFile string

func main() {
	var cmdApi = &cobra.Command{
		Use:   "api ",
		Short: "Start the API configuration server",
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			api.Start(certFile, keyFile)
		},
	}
	cmdApi.Flags().StringVarP(&certFile, "tls-crt-file", "", "domain.tld.crt", "A certificate file")
	cmdApi.Flags().StringVarP(&keyFile, "tls-key-file", "", "domain.tld.key", "A key file")


	var rootCmd = &cobra.Command{Use: "yulmails"}
	rootCmd.AddCommand(cmdApi)
	rootCmd.Execute()
}
