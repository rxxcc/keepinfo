package utils

import (
	"fmt"
	"net/mail"
	"regexp"
)

var isValidName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if value == "" {
		return fmt.Errorf("field cannot be empty")
	}
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateName(value ...string) error {
	for _, val := range value {
		if err := ValidateString(val, 2, 100); err != nil {
			return fmt.Errorf("name %w", err)
		}
		if !isValidName(val) {
			return fmt.Errorf("name must contain only letters or spaces")
		}
	}

	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return fmt.Errorf("password %w", err)
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return fmt.Errorf("email %w", err)
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("%s is not a valid email address", value)
	}
	return nil
}
