package runtodocreate

import (
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/entity"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
	Todo *entity.Todo
}
