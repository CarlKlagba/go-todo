package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math"
	"time"
)

type Task struct {
	Id        int
	Task      string
	Priority  int
	Completed bool
	DueDate   time.Time
}

func (t *Task) DaysLeft() int {
	return int(math.Abs(time.Now().Sub(t.DueDate).Hours() / 24))
}

const createTaskSchema = `CREATE TABLE IF NOT EXISTS tasks (
   id INTEGER PRIMARY KEY,
   task TEXT,
   priority NUMBER,
   completed BOOLEAN,
   dueDate DATE
);`

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "db/planner.db")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(createTaskSchema); err != nil {
		log.Println(err)
	}
	return db
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return TaskRepository{db: db}
}

func (repo *TaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	rows, err := repo.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Task, &task.Priority, &task.Completed, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (repo *TaskRepository) AddTask(task Task) (int, error) {
	res, err := repo.db.Exec("INSERT INTO tasks(task, priority, completed, dueDate) VALUES (?,?,?,?);",
		task.Task, task.Priority, task.Completed, task.DueDate)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (repo *TaskRepository) CompleteTask(id int) (Task, error) {
	task, err2 := repo.FindTask(id)
	if err2 != nil {
		return Task{}, err2
	}
	_, err := repo.db.Exec("UPDATE tasks SET completed = ? WHERE id = ?", !task.Completed, id)
	if err != nil {
		return Task{}, err
	}

	newTask, err3 := repo.FindTask(id)
	if err2 != nil {
		return Task{}, err3
	}
	return newTask, nil
}

func (repo *TaskRepository) FindTask(id int) (Task, error) {
	var task Task
	err := repo.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id).Scan(&task.Id, &task.Task, &task.Priority, &task.Completed, &task.DueDate)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
