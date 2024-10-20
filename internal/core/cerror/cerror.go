package cerror

import (
	"encoding/json"
	"fmt"
)

type customError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *customError) Error() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("cerror: %s", err.Error())
	}

	return string(bytes)
}

func New(code, message string) error {
	return &customError{
		Code:    code,
		Message: message,
	}
}
