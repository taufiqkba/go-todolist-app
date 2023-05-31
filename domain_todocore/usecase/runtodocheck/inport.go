package runtodocheck

import (
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
}

type InportResponse struct {
}
