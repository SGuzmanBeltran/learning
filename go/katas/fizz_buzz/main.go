package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
	fizz := ""
	containsFive := strings.Contains(result, "5")
	containsThree := strings.Contains(result, "3")
	if containsThree {
		fizz += "Fizz"
	}
	isMultipleThree := number % 3 == 0
	isMultipleFive := number % 5 == 0
	isMultipleBoth := isMultipleThree && isMultipleFive
	if isMultipleBoth {
		fizz += "FizzBuzz"
	} else if isMultipleFive {
		fizz += "Buzz"
	} else if isMultipleThree {
		fizz += "Fizz"
	}

	if containsFive {
		fizz += "Buzz"
	}

	if len(fizz) > 0 {
		return fizz, nil
	}
	return result, nil
}
