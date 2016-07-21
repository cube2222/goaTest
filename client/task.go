package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CreateTaskPath computes a request path to the create action of Task.
func CreateTaskPath() string {
	return fmt.Sprintf("/tasks")
}

// Create task with content.
func (c *Client) CreateTask(ctx context.Context, path string, content *string) (*http.Response, error) {
	req, err := c.NewCreateTaskRequest(ctx, path, content)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateTaskRequest create the request corresponding to the create action endpoint of the Task resource.
func (c *Client) NewCreateTaskRequest(ctx context.Context, path string, content *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if content != nil {
		values.Set("content", *content)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteTaskPath computes a request path to the delete action of Task.
func DeleteTaskPath(taskID string) string {
	return fmt.Sprintf("/tasks/%v", taskID)
}

// Delete task by ID.
func (c *Client) DeleteTask(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteTaskRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteTaskRequest create the request corresponding to the delete action endpoint of the Task resource.
func (c *Client) NewDeleteTaskRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowTaskPath computes a request path to the show action of Task.
func ShowTaskPath(taskID string) string {
	return fmt.Sprintf("/tasks/%v", taskID)
}

// Get task by ID.
func (c *Client) ShowTask(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTaskRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTaskRequest create the request corresponding to the show action endpoint of the Task resource.
func (c *Client) NewShowTaskRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateTaskPath computes a request path to the update action of Task.
func UpdateTaskPath(taskID string) string {
	return fmt.Sprintf("/tasks/%v", taskID)
}

// Update task content.
func (c *Client) UpdateTask(ctx context.Context, path string, content *string) (*http.Response, error) {
	req, err := c.NewUpdateTaskRequest(ctx, path, content)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateTaskRequest create the request corresponding to the update action endpoint of the Task resource.
func (c *Client) NewUpdateTaskRequest(ctx context.Context, path string, content *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if content != nil {
		values.Set("content", *content)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
