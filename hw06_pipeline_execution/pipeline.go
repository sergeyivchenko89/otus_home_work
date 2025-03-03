package hw06pipelineexecution

import "fmt"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {

	out := make(Out)
	returnChannel := make(Bi)

	for i, stage := range stages {
		if i == 0 {
			out = stage(in)
		} else {
			out = stage(out)
		}
	}

	go func() {
		isRunning := true
		isDone := false
		for isRunning {
			fmt.Println("Is Running: ", isRunning)
			select {
			case outValue, ok := <-out:
				if isDone {
					continue
				}
				if !ok {
					close(returnChannel)
					isRunning = false
					isDone = true
					continue
				}
				returnChannel <- outValue
			case <-done:
				fmt.Println("Close done channel")
				isDone = true
				isRunning = false
				close(returnChannel)
				return
			default:
			}
		}
	}()

	return returnChannel
}
