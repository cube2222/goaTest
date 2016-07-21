package main

import (
	"github.com/cube2222/goaTest/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/uuid"
)

// TaskController implements the Task resource.
type TaskController struct {
	*goa.Controller
}

// NewTaskController creates a Task controller.
func NewTaskController(service *goa.Service) *TaskController {
	return &TaskController{Controller: service.NewController("TaskController")}
}

// Create runs the create action.
func (c *TaskController) Create(ctx *app.CreateTaskContext) error {
	// TaskController_Create: start_implement

	TaskMutex.Lock()
	newID := uuid.NewV4()
	Tasks[newID] = &app.TasklistTask{ID:&newID, Content:ctx.Content}
	TaskMutex.Unlock()

	// TaskController_Create: end_implement
	return ctx.OK(Tasks[newID])
}

// Delete runs the delete action.
func (c *TaskController) Delete(ctx *app.DeleteTaskContext) error {
	// TaskController_Delete: start_implement
	TaskMutex.Lock()
	delete(Tasks, ctx.TaskID)
	TaskMutex.Unlock()

	// TaskController_Delete: end_implement
	return ctx.NoContent()
}

// Show runs the show action.
func (c *TaskController) Show(ctx *app.ShowTaskContext) error {
	// TaskController_Show: start_implement

	TaskMutex.RLock()
	res, ok := Tasks[ctx.TaskID]
	TaskMutex.RUnlock()
	if ok != true {
		return ctx.Teapot()
	}

	// TaskController_Show: end_implement
	return ctx.OK(res)
}

// Update runs the update action.
func (c *TaskController) Update(ctx *app.UpdateTaskContext) error {
	// TaskController_Update: start_implement

	TaskMutex.RLock()
	_, ok := Tasks[ctx.TaskID]
	TaskMutex.RUnlock()
	if ok != true {
		ctx.Teapot()
	}
	res := &app.TasklistTask{ID:&ctx.TaskID, Content:ctx.Content}
	TaskMutex.Lock()
	Tasks[ctx.TaskID] = res
	TaskMutex.Unlock()

	// TaskController_Update: end_implement
	return ctx.OK(res)
}

// List runs the list action.
func (c *app.TasklistController) Show(ctx *app.ShowTasklistContext) error {
	// TaskController_List: start_implement
	res := &app.TasklistTasklist{}

	TaskMutex.RLock()
	for _, value := range Tasks {
		res.List = append(res.List, value)
	}
	TaskMutex.RLock()

	// TaskController_List: end_implement

	return ctx.OK(res)
}
