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
		go func() {
			defer wg.Done()
			for task := range tasksChannel {
				resultChannel <- task()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	defer close(tasksChannel)

	for tasksExecuted == 0 || tasksInProgress != 0 {
		select {
		case result := <-resultChannel:
			tasksInProgress--
			if result != nil {
				tasksWithErrors++
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
