package error

import (
	"fmt"
)

/*
 * type: WError(Wrapper error)
 * function : to return more detail message
*/
type WError struct {
	eCode int
	eDetail  string
}

func (err *WError) Code() int {
	return err.eCode
}

func (err *WError) Error() string {
	//return fmt.Sprintf("Error: %s, %s", fetchErrString(err.eCode), err.eDetail)
	return fmt.Sprintf("Code: %d, Error: %s, Detail: %s", err.eCode, fetchErrString(err.eCode), err.eDetail)
}

func (err *WError) String() string {
	if err.eCode == ERR_CODE_OK {
		return "ok"
	} else {
		return err.Error()
	}
}

func NewWError(code int, format string, args ...interface{}) *WError {
	detail := format
	if len(format) > 0 && len(args) > 0 {
		detail = fmt.Sprintf(format, args...)
	}
	err := &WError{
		eCode: code,
		eDetail: detail,
	}

	return err
}

func WrapError(code int, e error) *WError {
	err := &WError{
		eCode: code,
		eDetail: e.Error(),
	}

	return err
}
