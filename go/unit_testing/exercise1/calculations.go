package exercise1

import "errors"

type Calculations struct{}

func (c *Calculations) CalculateFactorial(number int) (int, error) {
	result := 1

	if number < 0 {
		return 0, errors.New("can't calculate factorial of a negative number")
	}

	if number == 0 || number == 1 {
		return 1, nil
	}

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result, nil
}
