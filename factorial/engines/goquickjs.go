package engines

import (
	"fmt"

	"github.com/go-quickjs/go-quickjs"
)

type GoQuickJS struct {
	JIT bool
	rt  *quickjs.Runtime
}

func (q *GoQuickJS) Name() string {
	if q.JIT {
		return "GoQuickJSJIT"
	}
	return "GoQuickJS"
}

func (q *GoQuickJS) Init() error {
	var opts []quickjs.Option
	if q.JIT {
		if !goQuickJSJITSupported {
			return fmt.Errorf("GoQuickJSJIT requires -tags quickjs_jit on linux/amd64, windows/amd64, or darwin/arm64")
		}
		opts = append(opts, quickjs.WithJIT(), quickjs.WithJITMemoryLimit(32*1024*1024))
	}
	q.rt = quickjs.New(opts...)
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
