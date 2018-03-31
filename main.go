package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/yulPa/yulmails/api"
	"github.com/yulPa/yulmails/pkg/sender"
)


func main() {

	var certFile string
	var keyFile string

	var cmdApi = &cobra.Command{
		Use:   "api ",
		Short: "Start the API configuration server",
		Long: "Before starting the API you will need to provider SSL certificate",
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Certs given: ", certFile)
			api.Start(certFile, keyFile)
		},
	}
	cmdApi.Flags().StringVarP(&certFile, "tls-crt-file", "", "", "A certificate file")
	cmdApi.Flags().StringVarP(&keyFile, "tls-key-file", "", "", "A key file")

	var cmdSender = &cobra.Command{
		Use:   "sender ",
		Short: "Start the sender node",
		Run: func(cmd *cobra.Command, args []string) {
			sender.Run()
		},
	}


	var rootCmd = &cobra.Command{Use: "yulmails"}
	rootCmd.AddCommand(cmdApi, cmdSender)
	rootCmd.Execute()
}
