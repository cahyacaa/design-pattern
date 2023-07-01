package concurrency_pattern

type Executor func(any) (any, error)
type Pipeline interface {
	Pipe(executor Executor) Pipeline
	Merge() <-chan any
}

type pipeline struct {
	dataIn    <-chan any
	errC      chan error
	executors []Executor
}

func New(f func(chan<- any)) Pipeline {
	inC := make(chan any, 3)
	go f(inC)
	return &pipeline{
		dataIn:    inC,
		errC:      make(chan error),
		executors: []Executor{},
	}
}
func (p *pipeline) Pipe(executor Executor) Pipeline {
	p.executors = append(p.executors, executor)
	return p
}
func (p *pipeline) Merge() <-chan any {
	for i := 0; i < len(p.executors); i++ {
		p.dataIn, p.errC = run(p.dataIn, p.executors[i])
	}
	return p.dataIn
}
func run(inC <-chan any, f Executor) (<-chan any, chan error) {
	outC := make(chan any)
	errC := make(chan error)

	go func() {
		defer close(outC)
		for v := range inC {
			res, err := f(v)
			if err != nil {
				errC <- err
				continue
			}
			outC <- res
		}
	}()

	return outC, errC
}
