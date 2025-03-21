package exercise1

import "testing"

type caseTest struct {
	name             string
	number, expected int
}

type errorCaseTests struct {
	name    string
	number  int
	doError bool
}

//Test normal cases, asserting positive numbers
func TestNormal(t *testing.T) {
	c := Calculations{}
	var normalTests = []caseTest{
		{"2 should be 2", 2, 2},
		{"3 should be 6", 3, 6},
		{"4 should be 24", 4, 24},
		{"10 should be 3628800", 10, 3628800},
	}

	for _, normal := range normalTests {
		t.Run(normal.name, func(t *testing.T) {
			result, err := c.CalculateFactorial(normal.number)
			if err != nil {
				t.Fatalf(`TestNormal(%d) error: %s`, normal.number, err)
			}

			if result != normal.expected {
				t.Errorf(`TestNormal(%d) = %d, want match for %d, nil`, normal.number, result, normal.expected)
			}
		})
	}
}

//Fuzz Test normal cases, asserting random positive numbers
func FuzzNormal(f *testing.F) {
	c := Calculations{}
	f.Add(10)
	f.Fuzz(func(t *testing.T, n int) {
		c.CalculateFactorial(n)
	})
}

//Test edge cases, asserting 0 and 1 where they are edge cases
func TestInitials(t *testing.T) {
	c := Calculations{}
	var caseTests = []caseTest{
		{"0 should be 1", 0, 1},
		{"1 should be 1", 1, 1},
	}
	for _, edge := range caseTests {
		t.Run(edge.name, func(t *testing.T) {
			result, err := c.CalculateFactorial(edge.number)
			if err != nil {
				t.Fatalf(`TestInitials(%d) error: %s`, edge.number, err)
			}

			if result != edge.expected {
				t.Errorf(`TestInitials(%d) = %d, want match for %d, nil`, edge.number, result, edge.expected)
			}
		})
	}
}

//Test error cases, asserting negative numbers, we want to return the error and handle it
func TestError(t *testing.T) {
	c := Calculations{}
	var errorTests = []errorCaseTests{
		{"0 shouldnt return error", 0, false},
		{"100 shouldnt return error", 100, false},
		{"-1 should return error", -1, true},
		{"-1000 should return error", -1000, true},
	}
	for _, e := range errorTests {
		t.Run(e.name, func(t *testing.T) {
			_, err := c.CalculateFactorial(e.number)
			if e.doError && err == nil {
				t.Errorf(`TestError(%d) = %t, want match for %t`, e.number, err == nil, e.doError)
			}
		})
	}
}
