package string_sum

import (
	"errors"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")

	allowedOperands = map[rune]struct{}{
		'0': struct{}{},
		'1': struct{}{},
		'2': struct{}{},
		'3': struct{}{},
		'4': struct{}{},
		'5': struct{}{},
		'6': struct{}{},
		'7': struct{}{},
		'8': struct{}{},
		'9': struct{}{},
	}
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func checkOperands(input string) (res bool) {
	limit := 2
	exists := 0
	for _, v := range input {
		_, ok := allowedOperands[v]
		if ok {
			exists++
			if exists > limit {
				return false
			}
		}
	}
	if exists < limit {
		return false
	}
	return true
}

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return output, errorEmptyInput
	}
	if !checkOperands(input) {
		return output, errorNotTwoOperands
	}

	return "", nil
}
