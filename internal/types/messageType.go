package types
/*
this struct is defined to send intended data to server and get the result
Numbers: a slice that contains all numbers that user enters
Operator: the type of operator that user enters(+, -, *, /)
*/
type Message struct {
	Numbers []int   `json:"numbers"`
	Operator string `json:"operator"`
}
