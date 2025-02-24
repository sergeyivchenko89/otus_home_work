package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {

	out := make(Out)
	tasksChannel := make(Bi, len(stages))
	resultChannel := make(Bi, len(stages))

	go func(stages []Stage) {
		for i, stage := range stages {
			if i == 0 {
				out = stage(tasksChannel)
			} else {
				out = stage(out)
			}
		}
	}(stages)

	go func() {
		for {
			select {
			case inValue, ok := <-in:
			case outValue, ok := <-out:

			}
		}
	}()

	return resultChannel
}
