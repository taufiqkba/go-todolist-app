package runtodocheck

import (
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/entity"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/vo"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
)

type Inport = gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
	Todo *entity.Todo
}
