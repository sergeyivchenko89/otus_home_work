package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {

	if len(tasks) == 0 {
		return nil
	}

	wg := &sync.WaitGroup{}
	tasksChannel := make(chan Task, n)
	resultChannel := make(chan error)
	tasksInProgress := 0
	tasksExecuted := 0
	tasksWithErrors := 0

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(runnerNumber int) {
			defer wg.Done()
			for task := range tasksChannel {
				resultChannel <- task()
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	defer close(tasksChannel)

	inProgress := true
	for inProgress {

		select {
		case result := <-resultChannel:
			tasksInProgress--
			if result != nil {
				tasksWithErrors++
			}
			if tasksInProgress == 0 {
				inProgress = false
			}
		default:
		}

		if m > 0 && tasksWithErrors >= m {
			continue
		}

		if tasksExecuted < len(tasks) && tasksInProgress < n {
			tasksChannel <- tasks[tasksExecuted]
			tasksInProgress++
			tasksExecuted++
		}
	}

	if m > 0 && tasksWithErrors >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
