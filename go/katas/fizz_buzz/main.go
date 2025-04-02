package main

import (
	"strconv"
)

// TODO: We're currently starting the FOURTH CYCLE of https://www.codurance.com/katas/fizzbuzz
type FizzBuzz struct{}

func (fb *FizzBuzz) Convert(number int) (string, error) {
	return strconv.Itoa(number), nil
	// return 0, fmt.Errorf("implement me!")
}
