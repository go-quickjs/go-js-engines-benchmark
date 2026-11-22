package engines

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ekuma"
)

type Ekuma struct {
	vm       *ekuma.VM
	toString ekuma.Value
	output   [][]string
	baseDir  string
}

func (e *Ekuma) Name() string {
	return "Ekuma"
}

func (e *Ekuma) Init() error {
	baseDir, err := filepath.Abs("v8-v7")
	if err != nil {
		return fmt.Errorf("could not resolve v8-v7 directory: %w", err)
	}
	e.baseDir = baseDir
	e.output = nil
	e.vm = &ekuma.VM{}
	global := e.vm.Global()
	e.toString, err = global.Get("String")
	if err != nil {
		e.Close()
		return err
	}
	if err := global.Set("load", e.vm.FuncOf("load", func(_ ekuma.Value, args []ekuma.Value) (any, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("load: missing filename argument")
		}
		name, err := e.stringValue(args[0])
		if err != nil {
			return nil, err
		}
		return ekuma.Undefined, e.runScript(filepath.Join(e.baseDir, name))
	})); err != nil {
		e.Close()
		return err
	}
	if err := global.Set("print", e.vm.FuncOf("print", func(_ ekuma.Value, args []ekuma.Value) (any, error) {
		var line []string
		for i, arg := range args {
			text, err := e.stringValue(arg)
			if err != nil {
				return nil, err
			}
			if i > 0 {
				line = append(line, " ")
			}
			line = append(line, text)
		}
		for _, text := range line {
			fmt.Print(text)
		}
		fmt.Println()
		e.output = append(e.output, line)
		return ekuma.Undefined, nil
	})); err != nil {
		e.Close()
		return err
	}
	return nil
}

func (e *Ekuma) stringValue(value ekuma.Value) (string, error) {
	// Value.String describes non-strings; invoke JavaScript String for coercion.
	text, err := e.toString.Invoke(context.Background(), value)
	if err != nil {
		return "", err
	}
	return text.String(), nil
}

func (e *Ekuma) Run(inputFile string) ([][]string, error) {
	err := e.runScript(inputFile)
	return e.output, err
}

func (e *Ekuma) runScript(inputFile string) error {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("could not read %s: %w", inputFile, err)
	}
	script, err := e.vm.ParseScript(inputFile, string(content))
	if err != nil {
		return err
	}
	_, err = script.Evaluate(context.Background())
	return err
}

func (e *Ekuma) Close() error {
	e.vm = nil
	e.toString = ekuma.Value{}
	e.output = nil
	e.baseDir = ""
	return nil
}

var _ JSEngine = (*Ekuma)(nil)
