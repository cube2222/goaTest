package test

import (
	"bytes"
	"fmt"
	"github.com/cube2222/goaTest/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// CreateTaskOKContent Create runs the method Create of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateTaskOKContent(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, content *string) (http.ResponseWriter, *app.TasklistTaskContent) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	createCtx, err := app.NewCreateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTaskContent
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTaskContent)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTaskContent", resp)
		}
	}

	// Return results
	return rw, mt
}

// CreateTaskOK Create runs the method Create of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateTaskOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, content *string) (http.ResponseWriter, *app.TasklistTask) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	createCtx, err := app.NewCreateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTask
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTask)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTask", resp)
		}
	}

	// Return results
	return rw, mt
}

// CreateTaskTeapot Create runs the method Create of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateTaskTeapot(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, content *string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	createCtx, err := app.NewCreateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 418 {
		t.Errorf("invalid response status code: got %+v, expected 418", rw.Code)
	}

	// Return results
	return rw
}

// DeleteTaskNoContent Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteTaskNoContent(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tasks/%v", taskID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// DeleteTaskTeapot Delete runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteTaskTeapot(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tasks/%v", taskID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 418 {
		t.Errorf("invalid response status code: got %+v, expected 418", rw.Code)
	}

	// Return results
	return rw
}

// ShowTaskOKContent Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowTaskOKContent(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID) (http.ResponseWriter, *app.TasklistTaskContent) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tasks/%v", taskID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	showCtx, err := app.NewShowTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTaskContent
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTaskContent)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTaskContent", resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowTaskOK Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowTaskOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID) (http.ResponseWriter, *app.TasklistTask) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tasks/%v", taskID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	showCtx, err := app.NewShowTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTask
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTask)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTask", resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowTaskTeapot Show runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowTaskTeapot(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/tasks/%v", taskID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	showCtx, err := app.NewShowTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 418 {
		t.Errorf("invalid response status code: got %+v, expected 418", rw.Code)
	}

	// Return results
	return rw
}

// UpdateTaskOKContent Update runs the method Update of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateTaskOKContent(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID, content *string) (http.ResponseWriter, *app.TasklistTaskContent) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks/%v", taskID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTaskContent
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTaskContent)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTaskContent", resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdateTaskOK Update runs the method Update of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateTaskOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID, content *string) (http.ResponseWriter, *app.TasklistTask) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks/%v", taskID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.TasklistTask
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.TasklistTask)
		if !ok {
			t.Errorf("invalid response media: got %+v, expected instance of app.TasklistTask", resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdateTaskTeapot Update runs the method Update of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateTaskTeapot(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.TaskController, taskID uuid.UUID, content *string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if content != nil {
		sliceVal := []string{*content}
		query["content"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/tasks/%v", taskID),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["taskID"] = []string{fmt.Sprintf("%v", taskID)}
	if content != nil {
		sliceVal := []string{*content}
		prms["content"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "TaskTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateTaskContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 418 {
		t.Errorf("invalid response status code: got %+v, expected 418", rw.Code)
	}

	// Return results
	return rw
}
