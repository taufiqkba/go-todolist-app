package runtodocheck

import "github.com/taufiqkba/go-todolist-app/domain_todocore/model/repository"

type Outport interface {
	repository.FindOneTodoByIDRepo
	repository.SaveTodoRepo
}
