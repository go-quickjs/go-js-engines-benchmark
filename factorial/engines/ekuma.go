package engines

import (
	"context"

	"github.com/hajimehoshi/ekuma"
)

type Ekuma struct {
	vm *ekuma.VM
}

func (e *Ekuma) Name() string {
	return "Ekuma"
}

func (e *Ekuma) Init() error {
	e.vm = &ekuma.VM{}
	// Ekuma initializes lazily; create its realm before timing starts.
	e.vm.InitialRealm()
	return nil
}

func (e *Ekuma) Run(input string) error {
	script, err := e.vm.ParseScript("factorial.js", input)
	if err != nil {
		return err
	}
	_, err = script.Evaluate(context.Background())
	return err
}

func (e *Ekuma) Close() error {
	e.vm = nil
	return nil
}

var _ JSEngine = (*Ekuma)(nil)
