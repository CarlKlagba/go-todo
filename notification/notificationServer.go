package notification

import (
	"fmt"
	. "github.com/CarlKlagba/go-todo/repository"
	"time"
)

func Run(repo TaskRepository) {
	fmt.Println("Run notification server")

	runCron(sendEmailNotification, repo)
}

func runCron(task func(TaskRepository), repo TaskRepository) {
	for {
		now := time.Now()
		isCron, err := IsTimeToCron(now, CronAny, 1, 1)
		if err != nil {
			fmt.Println("Error in cron", err)
		}
		if isCron {
			task(repo)
		}
		time.Sleep(55 * time.Second)
	}
}

func sendEmailNotification(repo TaskRepository) {
	tasks, err := repo.GetTasks()
	if err == nil {
		fmt.Println("Sending notifications...")
		tasksToNotify := filterTaskNotify(tasks)
		if len(tasksToNotify) > 0 {
			SendNotification(tasksToNotify)
		}
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
