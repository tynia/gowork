package error

import (
	"fmt"
)

/*
 * type: WError(Wrapper error)
 * function : to return more detail message
 */
type WError struct {
	eCode   int
	eMsg    string
}

func (err *WError) Code() int {
	return err.eCode
}

func (err *WError) Error() string {
	return fmt.Sprintf("%s", err.eMsg)
}

func (err *WError) Detail() string {
	if err.eCode == ERR_CODE_OK {
		return "ok"
	} else {
		return fmt.Sprintf("Code: %d, Error: %s", err.eCode, err.eMsg)
	}
}

func NewWError(code int, format string, args ...interface{}) *WError {
	detail := format
	if len(format) > 0 && len(args) > 0 {
		detail = fmt.Sprintf(format, args...)
	}
	err := &WError{
		eCode:   code,
		eMsg: detail,
	}

	return err
}

func WrapError(code int, e error) *WError {
	err := &WError{
		eCode: code,
		eMsg:  e.Error(),
	}

	return err
}
