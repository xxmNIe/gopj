package cmd

import (
	"github.com/spf13/cobra"
	"hello-world/server"
	"log"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Run the gRPC hello-world server",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err :=recover();err != nil {
				log.Println("Recover error : %v", err)
			}
		}()
		server.Server()
	},
}

func init() {
	serverCmd.Flags().StringVarP(&server.ServerPort,"port","p","8002","server port")
	serverCmd.Flags().StringVarP(&server.CertPemPath, "cert-pem", "", "/home/xxm/gopj/hello-world/certs/server.pem", "cert pem path")
	serverCmd.Flags().StringVarP(&server.CertKeyPath, "cert-key", "", "/home/xxm/gopj/hello-world/certs/server.key", "cert key path")
	serverCmd.Flags().StringVarP(&server.CertName, "cert-name", "", "hello-world", "server's hostname")
	rootCmd.AddCommand(serverCmd)
}