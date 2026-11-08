package engines

// JSEngine interface that any JS engine must implement
type JSEngine interface {
	Name() string
	Init() error
	Run(input string) ([][]string, error)
	Close() error
}

func Engines() []JSEngine {
	engines := []JSEngine{
		&GOJA{},
		&ModerncQuickJS{},
		&QJS{},
		&GoQuickJS{},
	}
	if goQuickJSJITSupported {
		engines = append(engines, &GoQuickJS{JIT: true})
	}
	return append(engines,
		&Paserati{},
		&MoeJS{},
		&Node{},
		&Node{Jitless: true},
	)
}
