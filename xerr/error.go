package xerr

import (
	"fmt"
)

/*
 * type: xerr(x error)
 * function : to return more detail message
 */
type Xerr struct {
	eCode   int
	eMsg    string
}

func (err *Xerr) Code() int {
	return err.eCode
}

func (err *Xerr) Error() string {
	if err.eCode == ERR_CODE_OK {
		return "ok"
	} else {
		return fmt.Sprintf("Code: %d, Error: %s", err.eCode, err.eMsg)
	}
	return fmt.Sprintf("%s", err.eMsg)
}

func New(code int, format string, args ...interface{}) error {
	detail := format
	if len(format) > 0 && len(args) > 0 {
		detail = fmt.Sprintf(format, args...)
	}

	err := &Xerr{
		eCode: code,
		eMsg: detail,
	}

	return err
}