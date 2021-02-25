package command

type MultiplyCommand interface {
	MultiplyNumbers(numbers ...float64) float64
}
