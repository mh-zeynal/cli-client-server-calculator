package types

//an interface created to generalize GetResult method of operation structs
type CalculateResult interface {
	GetResult() string
}
