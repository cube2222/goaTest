package cli

import (
	"fmt"
	"github.com/cube2222/goaTest/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

type (
	// CreateTaskCommand is the command line data structure for the create action of Task
	CreateTaskCommand struct {
		// New task content.
		Content     string
		PrettyPrint bool
	}

	// DeleteTaskCommand is the command line data structure for the delete action of Task
	DeleteTaskCommand struct {
		// Task ID
		TaskID      string
		PrettyPrint bool
	}

	// ListTaskCommand is the command line data structure for the list action of Task
	ListTaskCommand struct {
		PrettyPrint bool
	}

	// ShowTaskCommand is the command line data structure for the show action of Task
	ShowTaskCommand struct {
		// Task ID
		TaskID      string
		PrettyPrint bool
	}

	// UpdateTaskCommand is the command line data structure for the update action of Task
	UpdateTaskCommand struct {
		// Task ID
		TaskID string
		// Updated task content.
		Content     string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "create",
		Short: `Create task with content.`,
	}
	tmp1 := new(CreateTaskCommand)
	sub = &cobra.Command{
		Use:   `Task ["/tasks"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `Delete task by ID.`,
	}
	tmp2 := new(DeleteTaskCommand)
	sub = &cobra.Command{
		Use:   `Task ["/tasks/TASKID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `List all tasks.`,
	}
	tmp3 := new(ListTaskCommand)
	sub = &cobra.Command{
		Use:   `Task ["/tasks/list"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `Get task by ID.`,
	}
	tmp4 := new(ShowTaskCommand)
	sub = &cobra.Command{
		Use:   `Task ["/tasks/TASKID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `Update task content.`,
	}
	tmp5 := new(UpdateTaskCommand)
	sub = &cobra.Command{
		Use:   `Task ["/tasks/TASKID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

// Run makes the HTTP request corresponding to the CreateTaskCommand command.
func (cmd *CreateTaskCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/tasks"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateTask(ctx, path, stringFlagVal("content", cmd.Content))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateTaskCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var content string
	cc.Flags().StringVar(&cmd.Content, "content", content, `New task content.`)
}

// Run makes the HTTP request corresponding to the DeleteTaskCommand command.
func (cmd *DeleteTaskCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/tasks/%v", cmd.TaskID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteTask(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteTaskCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var taskID string
	cc.Flags().StringVar(&cmd.TaskID, "taskID", taskID, `Task ID`)
}

// Run makes the HTTP request corresponding to the ListTaskCommand command.
func (cmd *ListTaskCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/tasks/list"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListTask(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListTaskCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowTaskCommand command.
func (cmd *ShowTaskCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/tasks/%v", cmd.TaskID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowTask(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowTaskCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var taskID string
	cc.Flags().StringVar(&cmd.TaskID, "taskID", taskID, `Task ID`)
}

// Run makes the HTTP request corresponding to the UpdateTaskCommand command.
func (cmd *UpdateTaskCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/tasks/%v", cmd.TaskID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateTask(ctx, path, stringFlagVal("content", cmd.Content))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateTaskCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var taskID string
	cc.Flags().StringVar(&cmd.TaskID, "taskID", taskID, `Task ID`)
	var content string
	cc.Flags().StringVar(&cmd.Content, "content", content, `Updated task content.`)
}
