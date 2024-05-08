package val

import (
	"fmt"
	"net/mail"
	"regexp"
	"slices"

	"github.com/matheuspolitano/GadgetHub/utils"
)

var (
	isValidName  = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidPhone = regexp.MustCompile(`^\+[0-9]{1,3}\s?[0-9]{4,14}$`).MatchString
	isValidRole  = func(val string) bool {
		return slices.Contains([]string{utils.Admin, utils.Client}, val)
	}
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePhone(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}
	if !isValidPhone(value) {
		return fmt.Errorf("must contain +country code following by number")
	}
	return nil
}

func ValidateRole(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}
	if !isValidRole(value) {
		return fmt.Errorf("must be a valid role")
	}
	return nil
}
func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
