package types

import "strconv"

//a new type for sum operation
type Sum struct {
	Nums []int
}

//sums numbers and returns result as a string
func (s *Sum) GetResult() string {
	var res int
	for _, arg := range s.Nums{
		res += arg
	}
	return strconv.Itoa(res)
}
