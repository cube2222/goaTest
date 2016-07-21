//go:generate goagen bootstrap -d github.com/cube2222/goaTest/design

package main

import (
	"github.com/cube2222/goaTest/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/uuid"
	"sync"
)

var Tasks map[uuid.UUID]*app.TasklistTask
var TaskMutex sync.RWMutex

func main() {
	// Create service
	service := goa.New("TaskList")

	Tasks = make(map[uuid.UUID]*app.TasklistTask)
	TaskMutex = sync.RWMutex{}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Task" controller
	c := NewTaskController(service)
	app.MountTaskController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
