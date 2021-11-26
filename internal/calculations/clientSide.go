package calculations

import (
	"fmt"
	"strconv"
)

//checks if we have exactly one operator and no invalid input
func IsInputValid(args []string) bool {
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

//checks if we have exactly one operator
func doesContainTheOperator(args []string) bool {
	operatorCounter := 0
	for i := 0; i < len(args); i++ {
		if args[i] == "+" || args[i] == "-" ||
			args[i] == "*" || args[i] == "/" {
			operatorCounter += 1
		}
	}
	return operatorCounter == 1
}

//checks if we have invalid characters/inputs
func doesContainInvalidElements(args []string) bool {
	numCounter := 0
	for i := 0; i < len(args); i++ {
		if _, err := strconv.Atoi(args[i]); err == nil {
			numCounter += 1
		}
	}
	return numCounter == (len(args) - 1)
}

//finds operator among arguments and returns it
func GetOperator(args []string) string {
	for i := 0; i < len(args); i++ {
		if args[i] == "+" || args[i] == "-" ||
			args[i] == "*" || args[i] == "/"{
			return args[i]
		}
	}
	return ""
}

//extracts and returns entered numbers and returns them as a slice
func GetNumbers(args []string) []int {
	nums := make([]int, 0)
	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err == nil {
			temp, _ := strconv.Atoi(arg)
			nums = append(nums, temp)
		}
	}
	return nums
}
