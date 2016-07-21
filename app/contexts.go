//************************************************************************//
// API "TaskList": Application Contexts
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
	uuid "github.com/goadesign/goa/uuid"
	"golang.org/x/net/context"
)

// CreateTaskContext provides the Task create action context.
type CreateTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Content *string
}

// NewCreateTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the Task controller create action.
func NewCreateTaskContext(ctx context.Context, service *goa.Service) (*CreateTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramContent := req.Params["content"]
	if len(paramContent) > 0 {
		rawContent := paramContent[0]
		rctx.Content = &rawContent
	}
	return &rctx, err
}

// OKContent sends a HTTP response with status code 200.
func (ctx *CreateTaskContext) OKContent(r *TasklistTaskContent) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateTaskContext) OK(r *TasklistTask) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Teapot sends a HTTP response with status code 418.
func (ctx *CreateTaskContext) Teapot() error {
	ctx.ResponseData.WriteHeader(418)
	return nil
}

// DeleteTaskContext provides the Task delete action context.
type DeleteTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TaskID uuid.UUID
}

// NewDeleteTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the Task controller delete action.
func NewDeleteTaskContext(ctx context.Context, service *goa.Service) (*DeleteTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTaskID := req.Params["taskID"]
	if len(paramTaskID) > 0 {
		rawTaskID := paramTaskID[0]
		if taskID, err2 := uuid.FromString(rawTaskID); err2 == nil {
			rctx.TaskID = taskID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("taskID", rawTaskID, "uuid"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteTaskContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Teapot sends a HTTP response with status code 418.
func (ctx *DeleteTaskContext) Teapot() error {
	ctx.ResponseData.WriteHeader(418)
	return nil
}

// ShowTaskContext provides the Task show action context.
type ShowTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TaskID uuid.UUID
}

// NewShowTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the Task controller show action.
func NewShowTaskContext(ctx context.Context, service *goa.Service) (*ShowTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTaskID := req.Params["taskID"]
	if len(paramTaskID) > 0 {
		rawTaskID := paramTaskID[0]
		if taskID, err2 := uuid.FromString(rawTaskID); err2 == nil {
			rctx.TaskID = taskID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("taskID", rawTaskID, "uuid"))
		}
	}
	return &rctx, err
}

// OKContent sends a HTTP response with status code 200.
func (ctx *ShowTaskContext) OKContent(r *TasklistTaskContent) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTaskContext) OK(r *TasklistTask) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Teapot sends a HTTP response with status code 418.
func (ctx *ShowTaskContext) Teapot() error {
	ctx.ResponseData.WriteHeader(418)
	return nil
}

// UpdateTaskContext provides the Task update action context.
type UpdateTaskContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Content *string
	TaskID  uuid.UUID
}

// NewUpdateTaskContext parses the incoming request URL and body, performs validations and creates the
// context used by the Task controller update action.
func NewUpdateTaskContext(ctx context.Context, service *goa.Service) (*UpdateTaskContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateTaskContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramContent := req.Params["content"]
	if len(paramContent) > 0 {
		rawContent := paramContent[0]
		rctx.Content = &rawContent
	}
	paramTaskID := req.Params["taskID"]
	if len(paramTaskID) > 0 {
		rawTaskID := paramTaskID[0]
		if taskID, err2 := uuid.FromString(rawTaskID); err2 == nil {
			rctx.TaskID = taskID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("taskID", rawTaskID, "uuid"))
		}
	}
	return &rctx, err
}

// OKContent sends a HTTP response with status code 200.
func (ctx *UpdateTaskContext) OKContent(r *TasklistTaskContent) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateTaskContext) OK(r *TasklistTask) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.task+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Teapot sends a HTTP response with status code 418.
func (ctx *UpdateTaskContext) Teapot() error {
	ctx.ResponseData.WriteHeader(418)
	return nil
}

// ShowTasklistContext provides the Tasklist show action context.
type ShowTasklistContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowTasklistContext parses the incoming request URL and body, performs validations and creates the
// context used by the Tasklist controller show action.
func NewShowTasklistContext(ctx context.Context, service *goa.Service) (*ShowTasklistContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowTasklistContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTasklistContext) OK(r *TasklistTasklist) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/tasklist.tasklist+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
