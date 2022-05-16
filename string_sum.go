package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands   = errors.New("expecting two operands, but received more or less")
	errorUsupportedSymbol = errors.New("custom error")

	allowedOperands = map[rune]int{
		'0': 0,
		'1': 0,
		'2': 0,
		'3': 0,
		'4': 0,
		'5': 0,
		'6': 0,
		'7': 0,
		'8': 0,
		'9': 0,
	}
)

const EMPTY = ""

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf
func removeSpaces(input string) string {
	resultList := make([]rune, 0)
	for _, ch := range input {
		if !unicode.IsSpace(ch) {
			resultList = append(resultList, ch)
		}
	}
	return string(resultList)
}

func addIntFromStrInput(targetSlise []int, input string) []int {
	intToAdd, err := strconv.Atoi(input)
	if err == nil {
		targetSlise = append(targetSlise, intToAdd)
	}
	return targetSlise
}

func parseString(input string) (result []int, err error) {
	tempVal := EMPTY
	for _, val := range input {
		if val == '-' || val == '+' {
			if tempVal != EMPTY {
				result = addIntFromStrInput(result, tempVal)
			}
			tempVal = EMPTY

		} else if _, ok := allowedOperands[val]; !ok {
			_, err := strconv.Atoi(string(val))
			return result, fmt.Errorf("%w", err)
		}
		tempVal = tempVal + string(val)
	}
	if tempVal != EMPTY {
		result = addIntFromStrInput(result, tempVal)
	}
	if len(result) != 2 {
		return result, fmt.Errorf("%w", errorNotTwoOperands)
	}
	return result, nil
}

func StringSum(input string) (output string, err error) {
	input = removeSpaces(input)
	intResult := 0

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	operandsList, err := parseString(input)
	if err != nil {
		return "", err
	}
	for _, v := range operandsList {
		intResult = intResult + v
	}
	return strconv.Itoa(intResult), nil
}
