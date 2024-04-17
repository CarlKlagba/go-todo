package main

import (
	. "github.com/CarlKlagba/go-todo/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
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
	/*var tasks = []Task{
		{1, "Sort papers", 1, false, time.Date(2024, 5, 15, 14, 30, 45, 100, time.Local)},
		{2, "Pay URSSAF", 3, false, time.Date(2024, 4, 15, 14, 30, 45, 100, time.Local)},
		{3, "Cancel subscription", 2, false, time.Date(2024, 5, 3, 14, 30, 45, 100, time.Local)},
	}
	for _, task := range tasks {
		taskRepo.AddTask(task)
	}*/

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
	e.Logger.Fatal(e.Start(":1323"))
}

func completeTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var newTask = Task{}
	newTask, err := taskRepo.CompleteTask(id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "task-item", newTask)
}

func home(c echo.Context) error {
	tasks, _ := taskRepo.GetTasks()
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})

	var tasksMap = map[string][]Task{
		"Tasks": tasks,
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

	log.Println("task: ", task)
	log.Println("days left: ", task.DaysLeft())

	return c.Render(http.StatusOK, "task-item", task)
}
