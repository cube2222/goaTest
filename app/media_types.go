//************************************************************************//
// API "TaskList": Application Media Types
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

import uuid "github.com/goadesign/goa/uuid"

// A single task. (content view)
//
// Identifier: application/tasklist.task+json
type TasklistTaskContent struct {
	// Task content.
	Content *string `json:"content,omitempty" xml:"content,omitempty" form:"content,omitempty"`
}

// A single task. (default view)
//
// Identifier: application/tasklist.task+json
type TasklistTask struct {
	// Task content.
	Content *string `json:"content,omitempty" xml:"content,omitempty" form:"content,omitempty"`
	// Unique task ID.
	ID *uuid.UUID `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
}

// A group of tasks. (default view)
//
// Identifier: application/tasklist.tasklist+json
type TasklistTasklist struct {
	// All the tasks.
	List []*TasklistTask `json:"list,omitempty" xml:"list,omitempty" form:"list,omitempty"`
}
