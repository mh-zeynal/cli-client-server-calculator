package types

import "strconv"

//a new type for subtract operation
type Subtract struct {
	Nums []int
}

//subtracts numbers and returns result as a string
func (s Subtract) GetResult() string {
	var res int
	for i := 0; i < len(s.Nums); i++ {
		if i != 0 {
			res -= s.Nums[i]
			continue
		}
		res = s.Nums[i]
	}
	return strconv.Itoa(res)
}
