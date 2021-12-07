package types

import "strconv"

//a new type for product operation
type Product struct {
	Nums []int
}

//products numbers and returns result as a string
func (p Product) GetResult() string {
	var res int
	for i := 0; i < len(p.Nums); i++ {
		if i != 0 {
			res *= p.Nums[i]
			continue
		}
		res = p.Nums[i]
	}
	return strconv.Itoa(res)
}
