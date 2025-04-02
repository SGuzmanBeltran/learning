package main

import (
	"strconv"
)

type FizzBuzz struct {}

func (fb *FizzBuzz) Convert(number int) (string, error) {
	return strconv.Itoa(number), nil
	// return 0, fmt.Errorf("implement me!")
}