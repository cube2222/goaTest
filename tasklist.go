package main

import (
	"github.com/cube2222/goaTest/app"
	"github.com/goadesign/goa"
)

// TasklistController implements the Tasklist resource.
type TasklistController struct {
	*goa.Controller
}

// NewTasklistController creates a Tasklist controller.
func NewTasklistController(service *goa.Service) *TasklistController {
	return &TasklistController{Controller: service.NewController("TasklistController")}
}

// Show runs the show action.
func (c *TasklistController) Show(ctx *app.ShowTasklistContext) error {
	// TasklistController_Show: start_implement

	res := &app.TasklistTasklist{}

	TaskMutex.RLock()
	for _, value := range Tasks {
		res.List = append(res.List, value)
	}
	TaskMutex.RLock()

	// TasklistController_Show: end_implement
	return ctx.OK(res)
}
