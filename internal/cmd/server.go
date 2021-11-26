package cmd

import (
	"csp/internal/calculations"
	"csp/internal/types"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

var (
	server net.Listener //server socket
	err error			//connection error
	)
//serverCmd is a predefined command to launch server a let clients send message to it
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "runs server",
	Long:  `this command runs server and it won't turn off till you manually do it'`,
	Run:   runServer,
}

//adds serverCmd to rootCmd
func init() {
	rootCmd.AddCommand(serverCmd)
}

//runs the server and also does the calculations and send result to client
func runServer(cmd *cobra.Command, args []string) {
	server, err = net.Listen("tcp", ":8080")
	fmt.Println("now server is listening...")
	if err != nil {
		fmt.Println("server terminated")
		return
	}
	message := types.Message{}
	for  {
		listener, _ := server.Accept()
		_ = json.NewDecoder(listener).Decode(&message)
		_, writeError := listener.Write([]byte(calculations.DoOperation(message.Numbers, message.Operator)))
		if writeError != nil {
			fmt.Println("something went wrong with sending message")
			panic(writeError)
		}
	}
}


