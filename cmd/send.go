package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"strconv"
)

type Message struct {
	Numbers []int   `json:"numbers"`
	Operator string `json:"operator"`
}

var(
	clientDial net.Conn
	clientError error
	isConnected bool = false
)

var sendCmd = &cobra.Command{
	Use: "send",
	Short: "sends message to server",
	Long: `this command sends message to server in a special format as show below:
	FIRST_NUMBER SECOND_NUMBER THIRD_NUMBER ... nTH_NUMBER OPERATOR`,
	Args: cobra.MinimumNArgs(3),
	Run:  sendMessage,
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func sendMessage(cmd *cobra.Command, args []string) {
	if isConnected {
		if isInputValid(args) {
			operator := getOperator(args)
			numbers := getNumbers(args)
			message := Message{}
			message.Numbers = numbers
			message.Operator = operator
			str, _ := json.Marshal(message)
			_, _ = clientDial.Write([]byte(str))
		}
		return
	}
	fmt.Println("you're not connected to the server :/")
}

func isInputValid(args []string) bool {
	if !doesContainTheOperator(args) {
		fmt.Println("there's a problem in operator validation. possible problems are:\n" +
			"1.you haven't entered any operator in your arguments\n" +
			"2.the argument you've entered isn't among valid operators(+, -, *, /)\n" +
			"3.you have entered more than one operator")
		return false
	}else if !doesContainInvalidElements(args) {
		fmt.Println("you have entered invalid input(s)")
		return false
	}
	return true
}

func doesContainTheOperator(args []string) bool {
	operatorCounter := 0
	for i := 0; i < len(args); i++ {
		if args[i] == "+" || args[i] == "-" ||
			args[i] == "*" || args[i] == "/"{
			operatorCounter += 1
		}
	}
	return operatorCounter == 1
}

func doesContainInvalidElements(args []string) bool {
	numCounter := 0
	for i := 0; i < len(args); i++ {
		if _, err := strconv.Atoi(args[i]); err != nil {
			numCounter += 1
		}
	}
	return numCounter == (len(args) - 1)
}

func getOperator(args []string) string {
	for i := 0; i < len(args); i++ {
		if args[i] == "+" || args[i] == "-" ||
			args[i] == "*" || args[i] == "/"{
			return args[i]
		}
	}
	return ""
}

func getNumbers(args []string) []int {
	nums := make([]int, 0)
	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			temp, _ := strconv.Atoi(arg)
			nums = append(nums, temp)
		}
	}
	return nums
}
