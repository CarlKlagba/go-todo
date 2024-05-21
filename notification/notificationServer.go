package notification

import (
	"fmt"
	. "github.com/CarlKlagba/go-todo/repository"
	"time"
)

func Run(repo TaskRepository) {
	fmt.Println("Run notification server")
	for {
		tasks, err := repo.GetTasks()
		if err == nil {
			fmt.Println("Sending notifications...")
			tasksToNotify := filterTaskNotify(tasks)
			SendNotification(tasksToNotify)
		}
		time.Sleep(30 * time.Second)
	}
}

func filterTaskNotify(ss []Task) (ret []Task) {
	for _, s := range ss {
		if s.DaysLeft() <= 2 && !s.Completed {
			ret = append(ret, s)
		}
	}
	return
}
