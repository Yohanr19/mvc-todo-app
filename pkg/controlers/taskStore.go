package controlers

import (
	"github.com/yohanr19/mvc-todo-app/models/psqlStore"
)

type TaskControler struct {
	 Store psqlStore.TaskStore
}

func (tc *TaskControler)Init() error{
	error := tc.Store.InitDB()
	return error
}

func (tc *TaskControler)Insert(text string) {
	tc.Store.InsertTask(text)
}
func (tc *TaskControler)GetAll()[]psqlStore.Task {
	tasks := tc.Store.GetTasks()
	return tasks
}
func (tc *TaskControler)SetState(id string, isActive bool) {
		tc.Store.SetIsActive(id,isActive)
}