//************************************************************************//
// API "TaskList": Application Controllers
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/cube2222/goaTest/design
// --out=$(GOPATH)\src\github.com\cube2222\goaTest
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// TaskController is the controller interface for the Task actions.
type TaskController interface {
	goa.Muxer
	Create(*CreateTaskContext) error
	Delete(*DeleteTaskContext) error
	Show(*ShowTaskContext) error
	Update(*UpdateTaskContext) error
}

// MountTaskController "mounts" a Task resource controller on the given service.
func MountTaskController(service *goa.Service, ctrl TaskController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateTaskContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/tasks", ctrl.MuxHandler("Create", h, nil))
	service.LogInfo("mount", "ctrl", "Task", "action", "Create", "route", "POST /tasks")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteTaskContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/tasks/:taskID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Task", "action", "Delete", "route", "DELETE /tasks/:taskID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTaskContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/tasks/:taskID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Task", "action", "Show", "route", "GET /tasks/:taskID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateTaskContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/tasks/:taskID", ctrl.MuxHandler("Update", h, nil))
	service.LogInfo("mount", "ctrl", "Task", "action", "Update", "route", "PUT /tasks/:taskID")
}

// TasklistController is the controller interface for the Tasklist actions.
type TasklistController interface {
	goa.Muxer
	Show(*ShowTasklistContext) error
}

// MountTasklistController "mounts" a Tasklist resource controller on the given service.
func MountTasklistController(service *goa.Service, ctrl TasklistController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowTasklistContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/list/list", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Tasklist", "action", "Show", "route", "GET /list/list")
}
