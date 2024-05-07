package chat

import (
	"fmt"
	"regexp"
)

type ParseFunction func(string) (string, error)

// parseInt extracts the first sequence of digits from a string and converts it to an integer.
func parseInt(s string) (string, error) {
	// Compile a regular expression to find numbers.
	re := regexp.MustCompile(`\d+`)
	// Find the first match in the string.
	match := re.FindString(s)
	if match == "" {
		return "", fmt.Errorf("no numbers found in string")
	}

	return match, nil
}
