package main

import (
	"github.com/CarlKlagba/go-todo/notification"
	. "github.com/CarlKlagba/go-todo/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"sort"
	"strconv"
	"text/template"
	"time"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer(e *echo.Echo, paths ...string) {
	tmpl := &template.Template{}
	for i := range paths {
		template.Must(tmpl.ParseGlob(paths[i]))
	}
	t := newTemplate(tmpl)
	e.Renderer = t
}

var taskRepo TaskRepository

func newTemplate(templates *template.Template) echo.Renderer {
	return &Template{
		Templates: templates,
	}
}

func main() {
	//InitDatabase
	taskRepo = NewTaskRepository(InitDatabase())

	//Run notification server
	go notification.Run(taskRepo)

	e := echo.New()
	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	NewTemplateRenderer(e, "templates/*.html")

	// Routes
	e.GET("/", home)
	e.POST("/add-task", addTask)
	e.PUT("/complete-task/:id", completeTask)
	e.Logger.Fatal(e.Start(":8080"))
}

func completeTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var newTask = Task{}
	newTask, err := taskRepo.CompleteTask(id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "task-item", from(newTask))
}

func home(c echo.Context) error {
	tasks, _ := taskRepo.GetTasks()
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})

	var tasksMap = map[string][]TaskDto{
		"Tasks": fromList(tasks),
	}

	return c.Render(http.StatusOK, "index", tasksMap)
}

func addTask(c echo.Context) error {
	time.Sleep(1 * time.Second)

	name := c.FormValue("name")
	dueDateString := c.FormValue("dueDate")
	dueDate, _ := time.Parse("2006-01-02", dueDateString)

	priority := 999

	newTasks := Task{
		Task:      name,
		Priority:  priority,
		Completed: false,
		DueDate:   dueDate,
	}

	id, err := taskRepo.AddTask(newTasks)
	if err != nil {
		return err
	}

	task, _ := taskRepo.FindTask(id)

	return c.Render(http.StatusOK, "task-item", from(task))
}

type TaskDto struct {
	Id        int
	Task      string
	Priority  int
	Completed bool
	DueDate   time.Time
	DaysLeft  int
}

func from(task Task) TaskDto {
	return TaskDto{
		Id:        task.Id,
		Task:      task.Task,
		Priority:  task.Priority,
		Completed: task.Completed,
		DueDate:   task.DueDate,
		DaysLeft:  task.DaysLeft(),
	}
}

func fromList(tasks []Task) []TaskDto {
	var tasksDto []TaskDto
	for _, task := range tasks {
		tasksDto = append(tasksDto, from(task))
	}
	return tasksDto
}
