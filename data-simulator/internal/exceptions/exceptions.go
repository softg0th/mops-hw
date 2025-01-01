package exceptions

import (
	"fmt"
	"os"
	"time"
)

type CMDError struct {
	Field   string
	Message string
}

func (e *CMDError) Error() string {
	return fmt.Sprintf("Field: %s, Message: %s", e.Field, e.Message)
}

func HandleError(err error) {
	fmt.Println(err)
	time.Sleep(time.Second)
	os.Exit(1)
}
