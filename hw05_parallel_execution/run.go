package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	tasksChannel := make(chan Task, n)
	resultChannel := make(chan error)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(runnerNumber int) {
			fmt.Println("Начало работы воркера ", runnerNumber)
			defer func() {
				wg.Done()
				fmt.Println("Заврешение работы воркера ", runnerNumber)
			}()
			for task := range tasksChannel {
				resultChannel <- task()
			}
		}(i + 1)
	}

	go func() {
		wg.Wait()
		fmt.Println("Закрываем канал результатов")
		close(resultChannel)
	}()

	tasksLoaded := 0
	errorsCnt := 0
	err error

	defer func() {
		fmt.Println("Закрываем канал задач")
		close(tasksChannel)
	}()

	for {

		if tasksLoaded == 0 {
			fmt.Println("Первая итерация цикла. Загружаем задачи в канал задач")
			minValue := 0
			if n < len(tasks) {
				fmt.Println("Максимальное количество первых загруженных задач равно количеству воркеров.")
				minValue = n
			} else {
				fmt.Println("Максимальное количество первых загруженных задач равно общему количеству задач.")
				minValue = len(tasks)
			}

			for i := 0; i < minValue; i++ {
				tasksChannel <- tasks[tasksLoaded]
				tasksLoaded++
			}
		}

		if tasksLoaded == 0 {
			fmt.Println("Если количество загруженных задач равно нулю, то считаем что ошибок нет. Возвращаем nil")
			return nil
		}

		select {
		case result := <-resultChannel:
			if result != nil {
				errorsCnt++
				fmt.Println("Если воркер вернул ошибку, то увеличиваем счетчик ошибок на 1. Общее количество ошибок равно ", errorsCnt)
			}

			if errorsCnt == m {
				fmt.Println("Общее количество ошибок равно ", errorsCnt)
				tasksLoaded = len(tasks)
			}

			if tasksLoaded >= len(tasks) {
				fmt.Println("Завершение программы")
				return nil
			}

			tasksChannel <- tasks[tasksLoaded]
			tasksLoaded++
		}

	}

	if success {
		return nil
	} else {
		return ErrErrorsLimitExceeded
	}
}
