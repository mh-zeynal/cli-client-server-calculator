package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "client-side starter command",
	Long:  `this command starts a connection between client and server`,
	Run:   enableSocket,
}

func init() {
	rootCmd.AddCommand(enableCmd)
}

func enableSocket(cmd *cobra.Command, args []string) {
	clientDial, clientError = net.Dial("tcp", "localhost:8080")
	if clientError != nil {
		fmt.Println("server is down :/\nfirstly, you must turn your server on")
		return
	}
	fmt.Println("now you're connected :)")
	isConnected = true
}
