package engines

import "github.com/Calcium-Ion/moejs"

type MoeJS struct {
	rt *moejs.Runtime
}

func (m *MoeJS) Name() string {
	return "MoeJS"
}

func (m *MoeJS) Init() error {
	m.rt = moejs.NewRuntime(moejs.Options{})
	return nil
}

func (m *MoeJS) Run(input string) error {
	script, err := moejs.CompileScript("factorial.js", input)
	if err != nil {
		return err
	}
	_, err = m.rt.RunScript(script)
	return err
}

func (m *MoeJS) Close() error {
	m.rt = nil
	return nil
}

var _ JSEngine = (*MoeJS)(nil)
