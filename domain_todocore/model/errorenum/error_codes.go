package errorenum

import "github.com/taufiqkba/go-todolist-app/shared/model/apperror"

const (
	SomethingError      apperror.ErrorType = "ER0000 something error"
	MessageMustNotEmpty apperror.ErrorType = "ER0001 message must not empty"
)
