package utils

import "fmt"

type AppError struct {
	StatusCode int    `json:"-"`
	Err        string `json:"error,omitempty"`
}

func (m AppError) Error() string {
	return m.Err
}

func (m *AppError) AppendErrorToDesc(desc string) {
	m.Err = fmt.Sprintf("%s: %s", desc, m.Err)
}

func GetErrorStatusCode(err error) (statusCode int) {
	var e AppError
	var ok bool

	if e, ok = err.(AppError); ok {
		return e.StatusCode
	}
	return statusCode
}
