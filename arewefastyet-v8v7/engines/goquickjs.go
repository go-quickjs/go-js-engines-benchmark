package engines

import (
	"fmt"
	"os"

	"github.com/go-quickjs/go-quickjs"
)

type GoQuickJS struct {
	rt     *quickjs.Runtime
	output [][]string
}

func (q *GoQuickJS) Name() string {
	return "GoQuickJS"
}

func (q *GoQuickJS) Init() error {
	q.rt = quickjs.New()
	q.output = nil
	if err := q.rt.Set("load", func(input string) error {
		content, err := os.ReadFile("v8-v7/" + input)
		if err != nil {
			return fmt.Errorf("could not read %s: %w", input, err)
		}
		if _, err := q.rt.EvalFile(input, string(content)); err != nil {
			return fmt.Errorf("could not run script %s: %w", input, err)
		}
		return nil
	}); err != nil {
		q.Close()
		return err
	}

	if err := q.rt.Set("print", func(args ...quickjs.Value) {
		var line []string
		for i, arg := range args {
			if i > 0 {
				fmt.Print(" ")
				line = append(line, " ")
			}
			text := arg.String()
			fmt.Print(text)
			line = append(line, text)
		}
		fmt.Println()
		q.output = append(q.output, line)
	}); err != nil {
		q.Close()
		return err
	}
	return nil
}

func (q *GoQuickJS) Run(inputFile string) ([][]string, error) {
	script, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", inputFile, err)
	}
	_, err = q.rt.EvalFile(inputFile, string(script))
	return q.output, err
}

func (q *GoQuickJS) Close() error {
	q.output = nil
	if q.rt == nil {
		return nil
	}
	err := q.rt.Close()
	q.rt = nil
	return err
}

var _ JSEngine = (*GoQuickJS)(nil)
