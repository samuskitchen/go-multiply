package service

import (
	"go-multiply/domain/multiply/application/command"
)

type MultiplyService struct {}

func NewService() command.MultiplyCommand {
	return &MultiplyService{}
}

//MultiplyNumbers function that multiplies numbers with each other
func (ms *MultiplyService) MultiplyNumbers(numbers ...float64) float64 {
	return numbers[0] * numbers[1]
}