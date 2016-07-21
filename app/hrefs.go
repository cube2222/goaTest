//************************************************************************//
// API "TaskList": Application Resource Href Factories
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

import "fmt"

// TaskHref returns the resource href.
func TaskHref(taskID interface{}) string {
	return fmt.Sprintf("/tasks/%v", taskID)
}

// TasklistHref returns the resource href.
func TasklistHref() string {
	return fmt.Sprintf("/list/list")
}
