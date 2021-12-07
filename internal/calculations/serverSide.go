package calculations

import (
	"csp/internal/types"
)

//calls a function to do an operation based on operator
func DoOperation(nums []int, operator string) string {
	var temp types.CalculateResult
	switch operator {
	case "+":
		temp = &types.Sum{Nums: nums};      return getFinalValue(temp)
	case "*":
		temp = &types.Product{Nums: nums};  return getFinalValue(temp)
	case "-":
		temp = &types.Subtract{Nums: nums}; return getFinalValue(temp)
	case "/":
		temp = &types.Divide{Nums: nums};   return getFinalValue(temp)
	}
	return ""
}

//calculates the result depending on struct type
func getFinalValue(c types.CalculateResult) string {
	return c.GetResult()
}



