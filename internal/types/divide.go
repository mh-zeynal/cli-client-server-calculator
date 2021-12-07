package types

import "fmt"

//a new type for divide operation
type Divide struct {
	Nums []int
}

//divides numbers and returns result as a string
func (d Divide) GetResult() string {
	var res float32
	for i := 0; i < len(d.Nums); i++ {
		if i != 0 {
			if d.Nums[i] == 0 {
				return "divide by zero"
			}
			res /= float32(d.Nums[i])
			continue
		}
		res = float32(d.Nums[i])
	}
	return fmt.Sprintf("%g", res)
}
