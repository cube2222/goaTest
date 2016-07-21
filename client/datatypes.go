//************************************************************************//
// User Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/cube2222/goaTest/design
// --out=$(GOPATH)\src\github.com\cube2222\goaTest
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
)

// A single task.
type TasklistTask struct {
	// Task content.
	Content *string `json:"content,omitempty" xml:"content,omitempty" form:"content,omitempty"`
	// Unique task ID.
	ID *uuid.UUID `json:"id,omitempty" xml:"id,omitempty" form:"id,omitempty"`
}

// DecodeTasklistTask decodes the TasklistTask instance encoded in resp body.
func (c *Client) DecodeTasklistTask(resp *http.Response) (*TasklistTask, error) {
	var decoded TasklistTask
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
