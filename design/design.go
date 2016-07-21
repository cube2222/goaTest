package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("TaskList", func() {
	Title("TaskList")
	Scheme("http")
	Description("Simple tasklist using goa.")
	Host("localhost:8080")
	//Security(BasicAuthSecurity("basic"))
})

var _ = Resource("Task", func() {
	BasePath("/tasks")
	DefaultMedia(Task)

	Action("show", func() {
		Description("Get task by ID.")
		Routing(GET("/:taskID"))
		Params(func() {
			Param("taskID", UUID, "Task ID")
		})
		Response(OK)
		Response(Teapot)
	})

	Action("create", func() {
		Description("Create task with content.")
		Routing(POST("/"))
		Params(func() {
			Param("content", String, "New task content.", func() {
			})
		})
		Response(OK)
		Response(Teapot)
	})

	Action("update", func() {
		Description("Update task content.")
		Routing(PUT("/:taskID"))
		Params(func() {
			Param("taskID", UUID, "Task ID")
			Param("content", String, "Updated task content.", func() {
			})
		} )
		Response(OK)
		Response(Teapot)
	})

	Action("delete", func() {
		Description("Delete task by ID.")
		Routing(DELETE("/:taskID"))
		Params(func() {
			Param("taskID", UUID, "Task ID")
		})
		Response(NoContent)
		Response(Teapot)
	})
})

var _ = Resource("Tasklist", func() {
	BasePath("/list")
	DefaultMedia(TaskList)

	Action("show", func() {
		Description("List all tasks.")
		Routing(GET("/list"))
		Response(OK)
	})
})

var Task = MediaType("application/tasklist.task+json", func() {
	Description("A single task.")
	Attributes(func() {
		Attribute("id", UUID, "Unique task ID.")
		Attribute("content", String, "Task content.")
	})
	View("default", func() {
		Attribute("id")
		Attribute("content")
	})
	View("content", func() {
		Attribute("content")
	})
})

var TaskList = MediaType("application/tasklist.tasklist+json", func() {
	Description("A group of tasks.")
	Attributes(func() {
		Attribute("list", ArrayOf(Task), "All the tasks.")
	})
	View("default", func() {
		Attribute("list")
	})
})
