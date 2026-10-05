package engines

import "github.com/nooga/paserati/pkg/driver"

type Paserati struct {
	p *driver.Paserati
}

func (p *Paserati) Name() string {
	return "Paserati"
}

func (p *Paserati) Init() error {
	p.p = driver.NewPaserati()
	p.p.SetIgnoreTypeErrors(true)
	p.p.SetSkipTypeCheck(true)
	return nil
}

func (p *Paserati) Run(input string) error {
	result, errs := p.p.EvalCode(input, false)
	if len(errs) > 0 {
		return errs[0]
	}
	if p.p.GetVM().IsUnwinding() {
		return p.p.GetVM().NewExceptionError(result)
	}
	return nil
}

func (p *Paserati) Close() error {
	if p.p != nil {
		p.p.Cleanup()
		p.p = nil
	}
	return nil
}

var _ JSEngine = (*Paserati)(nil)
