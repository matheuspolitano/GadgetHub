package chat

import (
	"fmt"
	"regexp"
)

var MapFunctions map[string]ParseFunction

func init() {
	MapFunctions = map[string]ParseFunction{
		"parseInt": parseInt,
	}

}

type Action struct {
	Name         string   `mapstructure:"name"`
	Primary      bool     `mapstructure:"primary,omitempty"`
	NextAction   string   `mapstructure:"next_action,omitempty"`
	Description  string   `mapstructure:"description"`
	Payload      *Payload `mapstructure:"payload,omitempty"`
	Regex        string   `mapstructure:"regex,omitempty"`
	Response     string   `mapstructure:"response,omitempty"`
	IfElse       *IfElse  `mapstructure:"if_else,omitempty"`
	CallFunction string   `mapstructure:"callFunction,omitempty"`
}

type Payload struct {
	Key       string `mapstructure:"payload_key"`
	ParseFunc string `mapstructure:"parse_func,omitempty"`
}

type IfElse struct {
	Regex string  `mapstructure:"regex"`
	If    *Action `mapstructure:"if"`
	Else  *Action `mapstructure:"else"`
}

// CheckPatternExists checks if a regex pattern exists in a given string.
// It returns true if the pattern exists, false otherwise, and an error if the regex is not valid.
func CheckPatternExists(text, pattern string) (bool, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("error compiling regex: %v", err)
	}

	// Return if the pattern exists in the text
	return re.MatchString(text), nil
}

func (a *Action) CheckMessage(message string, payload map[string]string) (*Action, error) {
	var currentAction = a

	matched, err := CheckPatternExists(message, a.Regex)
	if err != nil {
		return nil, ErrorInRegex
	}
	if !matched {
		return nil, ErrorMessageUnexpected
	}

	if a.IfElse != nil {
		if matched, err := CheckPatternExists(message, a.IfElse.Regex); err != nil {
			return nil, ErrorInRegex
		} else if matched {
			currentAction = a.IfElse.If
		} else {
			currentAction = a.IfElse.Else
		}
	}

	if currentAction.Payload != nil {
		if parseFunction, ok := MapFunctions[a.Payload.ParseFunc]; ok || a.Payload.ParseFunc != "" {
			message, err = parseFunction(message)
			if err != nil {
				return nil, ErrorInParseMessage
			}
		}
		payload[currentAction.Payload.Key] = message

	}

	return currentAction, nil
}
