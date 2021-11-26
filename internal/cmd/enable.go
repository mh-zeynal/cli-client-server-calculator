package cmd

import (
	"csp/internal/types"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"csp/internal/calculations"
)

var(
	clientDial net.Conn //client socket
	clientError error	//connection error
)

/*
enableCmd is a predefined command to connect client to the server
and let it to send data once
*/
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "client-side starter command",
	Long:  `this command starts a connection between client and server`,
	Args:  cobra.MinimumNArgs(3),
	Run:   enableSocket,
}

//adds enableCmd to rootCmd
func init() {
	rootCmd.AddCommand(enableCmd)
}

//makes a connection between client and server and
func enableSocket(cmd *cobra.Command, args []string) {
	clientDial, clientError = net.Dial("tcp", "localhost:8080")
	if clientError != nil {
		fmt.Println("server is down :/\nfirstly, you must turn your server on")
		return
	}
	if calculations.IsInputValid(args) {
		operator := calculations.GetOperator(args)
		numbers := calculations.GetNumbers(args)
		message := types.Message{}
		message.Numbers =  numbers
		message.Operator = operator
		str, _ := json.Marshal(message)
		_, _ = clientDial.Write(str)
		data := make([]byte, 1024)
		_, _ = clientDial.Read(data)
		fmt.Print("result: " + string(data))
	}
}
