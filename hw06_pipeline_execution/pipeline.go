package hw06pipelineexecution

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
		isDone := false
		for {
			select {
			case outValue, ok := <-out:
				if !ok {
					close(returnChannel)
					return
				}
				if isDone {
					continue
				}
				returnChannel <- outValue
			case <-done:
				if isDone {
					continue
				}
				isDone = true
				close(returnChannel)
				continue
			}
		}
	}()

	return returnChannel
}
