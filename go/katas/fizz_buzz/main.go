package main

import (
	"fmt"
	"math"
	"strconv"
)

// TODO: We're currently implementing the last logic
// https://www.codurance.com/katas/fizzbuzz
type FizzBuzz struct{}

func (fb *FizzBuzz) Convert(number int) (string, error) {
	if number <= 0 {
		return "0", fmt.Errorf("number must be positive")
	}
	if math.MaxInt64 < number {
		return "0", fmt.Errorf("number must be lower than {%v}", number)
	}
	result := strconv.Itoa(number)
	isMultipleThree := number % 3 == 0
	isMultipleFive := number % 5 == 0
	isMultipleBoth := isMultipleThree && isMultipleFive
	if isMultipleBoth {
		result = "FizzBuzz"
	} else if isMultipleFive {
		result = "Buzz"
	} else if isMultipleThree {
		result = "Fizz"
	}
	return result, nil
}
