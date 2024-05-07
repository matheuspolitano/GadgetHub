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
	Name             string            `mapstructure:"name"`
	Primary          bool              `mapstructure:"primary,omitempty"`
	NextAction       string            `mapstructure:"next_action,omitempty"`
	Description      string            `mapstructure:"description"`
	SavePayload      bool              `mapstructure:"save_payload,omitempty"`
	PayloadKey       string            `mapstructure:"payload_key,omitempty"`
	ParseFunc        string            `mapstructure:"parse_func,omitempty"`
	Regex            string            `mapstructure:"regex,omitempty"`
	Response         string            `mapstructure:"response,omitempty"`
	IfElseNextAction *IfElseNextAction `mapstructure:"if_else_next_action,omitempty"`
	CallFunction     string            `mapstructure:"callFunction,omitempty"`
}

type IfElseNextAction struct {
	If   *Action `mapstructure:"if"`
	Else *Action `mapstructure:"else"`
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

	if a.IfElseNextAction != nil {
		if matched, err := CheckPatternExists(message, a.Regex); err != nil {
			return nil, ErrorInRegex
		} else if matched {
			currentAction = a.IfElseNextAction.If
		} else {
			currentAction = a.IfElseNextAction.Else
		}
	} else {
		if !matched {
			return nil, ErrorMessageUnexpected
		}
	}

	if currentAction.ParseFunc != "" {
		if parseFunction, ok := MapFunctions[a.ParseFunc]; ok {
			message, err = parseFunction(message)
			if err != nil {
				return nil, ErrorInParseMessage
			}
		}

	}

	if currentAction.SavePayload {
		payload[currentAction.PayloadKey] = message
	}
	return currentAction, nil
}
