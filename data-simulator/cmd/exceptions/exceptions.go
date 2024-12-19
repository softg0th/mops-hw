package exceptions

import "fmt"

type CMDError struct {
	Field   string
	Message string
}

func (e *CMDError) Error() string {
	return fmt.Sprintf("Field: %s, Message: %s", e.Field, e.Message)
}
