package password

import (
	"fmt"
	"unicode"
)

type PasswordValidation func(string) (bool, string)

func IsPasswordValid(password string, validators ...PasswordValidation) (bool, string) {
	isPasswordValid := true
	finalMessage := ""

	for _, validator := range validators {
		isValid, msg := validator(password)

		isPasswordValid = isPasswordValid && isValid
		finalMessage += msg
	}

	return isPasswordValid, finalMessage
}

func HaveNumber(password string) (bool, string) {
	minAmt := 1
	validAmt := 0

	for _, v := range password {
		if unicode.IsNumber(v) {
			validAmt++
		}
	}

	if validAmt >= minAmt {
		return true, ""
	}

	return false, fmt.Sprintf("Password must contain at least %d numbers", minAmt)
}

func HaveUpper(password string) (bool, string) {
	minAmt := 1
	validAmt := 0

	for _, v := range password {
		if unicode.IsUpper(v) {
			validAmt++
		}
	}

	if validAmt >= minAmt {
		return true, ""
	}

	return false, fmt.Sprintf("Password must contain at least %d uppercase letters", minAmt)
}

func HaveSymbol(password string) (bool, string) {
	minAmt := 1
	symbolAmt := 0

	for _, v := range password {
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			symbolAmt++
		}
	}

	if symbolAmt >= minAmt {
		return true, ""
	}

	return false, fmt.Sprintf("Password must contain at least %d symbol", minAmt)
}

func PasswordMaxLen(password string, maxLen int) PasswordValidation {
	return func(s string) (bool, string) {
		isValid := len(s) <= maxLen

		if isValid {
			return isValid, ""
		}

		return isValid, fmt.Sprintf("Password must be at maximum %d characters long;", maxLen)
	}
}

func PasswordMinLen(password string, minLen int) PasswordValidation {
	return func(s string) (bool, string) {
		isValid := len(s) >= minLen

		if isValid {
			return isValid, ""
		}

		return isValid, fmt.Sprintf("Password must be at least %d characters long;", minLen)
	}
}
