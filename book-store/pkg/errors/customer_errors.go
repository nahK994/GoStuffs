package errors

import "fmt"

type NotFound struct {
	Msg string
}

func (e *NotFound) Error() string {
	return fmt.Sprint(e.Msg)
}

type BadRequest struct {
	Msg string
}

func (e *BadRequest) Error() string {
	return fmt.Sprint(e.Msg)
}
