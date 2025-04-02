package main

import (
	"strconv"
)

// TODO: We're currently implementing the last logic
// https://www.codurance.com/katas/fizzbuzz
type FizzBuzz struct{}

func (fb *FizzBuzz) Convert(number int) (string, error) {
	if number == 3 {
		return "Fizz", nil
	}
	return strconv.Itoa(number), nil
}
