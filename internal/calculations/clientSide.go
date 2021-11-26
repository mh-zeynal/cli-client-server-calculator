package calculations

import (
	"fmt"
	"strconv"
)

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

func doesContainInvalidElements(args []string) bool {
	numCounter := 0
	for i := 0; i < len(args); i++ {
		if _, err := strconv.Atoi(args[i]); err == nil {
			numCounter += 1
		}
	}
	return numCounter == (len(args) - 1)
}

func GetOperator(args []string) string {
	for i := 0; i < len(args); i++ {
		if args[i] == "+" || args[i] == "-" ||
			args[i] == "*" || args[i] == "/"{
			return args[i]
		}
	}
	return ""
}

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
