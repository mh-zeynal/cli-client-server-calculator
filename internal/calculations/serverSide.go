package calculations

import (
	"fmt"
	"strconv"
)

func DoOperation(nums []int, operator string) string {
	switch operator {
	case "+":
		return doSum(nums)
	case "*":
		return doProduct(nums)
	case "-":
		return doSubmission(nums)
	case "/":
		return doDivision(nums)
	}
	return ""
}

func doSum(nums []int) string {
	var res int
	for _, arg := range nums{
		res += arg
	}
	return strconv.Itoa(res)
}

func doProduct(nums []int) string {
	var res int
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			res *= nums[i]
			continue
		}
		res = nums[i]
	}
	return strconv.Itoa(res)
}

func doSubmission(nums []int) string {
	var res int
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			res -= nums[i]
			continue
		}
		res = nums[i]
	}
	return strconv.Itoa(res)
}

func doDivision(nums []int) string {
	var res float32
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			if nums[i] == 0 {
				return "divide by zero"
			}
			res /= float32(nums[i])
			continue
		}
		res = float32(nums[i])
	}
	return fmt.Sprintf("%g", res)
}
