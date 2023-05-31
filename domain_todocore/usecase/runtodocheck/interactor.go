package runtodocheck

import (
	"context"
)

type runTodoCheckInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCheckInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...
	//!

	return res, nil
}
