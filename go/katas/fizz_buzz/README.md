# FizzBuzz Implementation

This implementation of FizzBuzz follows the standard rules plus an additional rule for digits '3' and '5'.

## Rules

1. **Standard FizzBuzz Rules:**
   - If number is multiple of 3: add "Fizz"
   - If number is multiple of 5: add "Buzz"
   - If number is multiple of both 3 and 5: add "FizzBuzz"

2. **Digit Rule:**
   - If number contains digit '3': add "Fizz"
   - If number contains digit '5': add "Buzz"

3. **Number Conversion Rule:**
   - If no rules apply, convert the number to string

## Test Cases

### Basic Cases
- 1 -> "1" (converted to string)
- 2 -> "2" (converted to string)
- 4 -> "4" (converted to string)

### Standard FizzBuzz Cases
- 3 -> "FizzFizz" (multiple of 3 + contains '3')
- 5 -> "BuzzBuzz" (multiple of 5 + contains '5')
- 6 -> "Fizz" (multiple of 3)
- 9 -> "Fizz" (multiple of 3)
- 10 -> "Buzz" (multiple of 5)
- 15 -> "FizzBuzzBuzz" (multiple of both + contains '5')

### Digit Rule Cases
- 13 -> "Fizz" (contains '3')
- 23 -> "Fizz" (contains '3')
- 51 -> "Buzz" (contains '5')
- 52 -> "Buzz" (contains '5')

### Combined Cases
- 33 -> "FizzFizz" (multiple of 3 + contains '3')
- 35 -> "FizzBuzzBuzz" (contains '3' + contains '5')
- 53 -> "FizzBuzz" (contains '3' + contains '5')
- 55 -> "BuzzBuzz" (multiple of 5 + contains '5')

### Edge Cases
- 0 -> Error (number must be positive)
- Negative numbers -> Error (number must be positive)
- Numbers exceeding MaxInt64 -> Error (number must be lower than MaxInt64)