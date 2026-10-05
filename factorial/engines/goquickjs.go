package engines

import "github.com/go-quickjs/go-quickjs"

type GoQuickJS struct {
	rt *quickjs.Runtime
}

func (q *GoQuickJS) Name() string {
	return "GoQuickJS"
}

func (q *GoQuickJS) Init() error {
	q.rt = quickjs.New()
	return nil
}

func (q *GoQuickJS) Run(input string) error {
	_, err := q.rt.Eval(input)
	return err
}

func (q *GoQuickJS) Close() error {
	if q.rt == nil {
		return nil
	}
	err := q.rt.Close()
	q.rt = nil
	return err
}

var _ JSEngine = (*GoQuickJS)(nil)
