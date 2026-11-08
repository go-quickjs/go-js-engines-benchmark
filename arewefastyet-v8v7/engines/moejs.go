package engines

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MoeForge/moejs"
)

type MoeJS struct {
	rt      *moejs.Runtime
	output  [][]string
	baseDir string
}

func (m *MoeJS) Name() string {
	return "MoeJS"
}

func (m *MoeJS) Init() error {
	baseDir, err := filepath.Abs("v8-v7")
	if err != nil {
		return fmt.Errorf("could not resolve v8-v7 directory: %w", err)
	}
	m.baseDir = baseDir
	m.output = nil
	// V8v7 replaces Math.random with a deterministic generator.
	m.rt = moejs.NewRuntime(moejs.Options{MutableIntrinsics: true})
	if err := m.rt.SetGlobal("load", moejs.NativeFunc(func(r *moejs.Realm, _ moejs.Value, args []moejs.Value) (moejs.Value, error) {
		if len(args) == 0 {
			return moejs.Undefined(), fmt.Errorf("load: missing filename argument")
		}
		name, err := r.ToString(args[0])
		if err != nil {
			return moejs.Undefined(), err
		}
		return moejs.Undefined(), m.runScript(filepath.Join(m.baseDir, name.GoString()))
	})); err != nil {
		return err
	}
	return m.rt.SetGlobal("print", moejs.NativeFunc(func(r *moejs.Realm, _ moejs.Value, args []moejs.Value) (moejs.Value, error) {
		var line []string
		for i, arg := range args {
			text, err := r.ToString(arg)
			if err != nil {
				return moejs.Undefined(), err
			}
			if i > 0 {
				line = append(line, " ")
			}
			line = append(line, text.GoString())
		}
		for _, text := range line {
			fmt.Print(text)
		}
		fmt.Println()
		m.output = append(m.output, line)
		return moejs.Undefined(), nil
	}))
}

func (m *MoeJS) Run(inputFile string) ([][]string, error) {
	err := m.runScript(inputFile)
	return m.output, err
}

func (m *MoeJS) runScript(inputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", inputFile, err)
	}
	script, err := moejs.CompileScript(inputFile, string(content))
	if err != nil {
		return fmt.Errorf("could not compile %s: %w", inputFile, err)
	}
	// Loaded scripts share the runtime's global environment.
	_, err = m.rt.RunScript(script)
	return err
}

func (m *MoeJS) Close() error {
	m.rt = nil
	m.output = nil
	m.baseDir = ""
	return nil
}

var _ JSEngine = (*MoeJS)(nil)
