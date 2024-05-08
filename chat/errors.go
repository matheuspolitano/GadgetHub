package chat

import "errors"

var (
	ErrorMessageUnexpected     = errors.New("error: message receive is unexpected to this action")
	ErrorParseFunctionNotExist = errors.New("error: message receive not exist")
	ErrorInRegex               = errors.New("error: regex expression is invalid")
	ErrorInParseMessage        = errors.New("error: regex expression in parse value")
)
