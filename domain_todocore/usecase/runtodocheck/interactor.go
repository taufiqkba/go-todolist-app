package runtodocheck

import (
	"context"
	"fmt"
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

	//get data by id
	todoObj, err := r.outport.FindOneTodoByID(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	//change state to = true
	err = todoObj.Check()
	if err != nil {
		return nil, err
	}

	//save to persistent storage
	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	res.Todo = todoObj
	return res, nil
}
